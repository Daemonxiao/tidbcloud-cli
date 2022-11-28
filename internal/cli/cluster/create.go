// Copyright 2022 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cluster

import (
	"fmt"
	"time"

	"tidbcloud-cli/internal"
	"tidbcloud-cli/internal/cli/project"
	"tidbcloud-cli/internal/config"
	"tidbcloud-cli/internal/flag"
	"tidbcloud-cli/internal/ui"
	"tidbcloud-cli/internal/util"

	clusterApi "github.com/c4pt0r/go-tidbcloud-sdk-v1/client/cluster"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/fatih/color"
	"github.com/juju/errors"
	"github.com/spf13/cobra"
)

type createClusterField int

const (
	clusterNameIdx createClusterField = iota
	passwordIdx
)

const (
	serverlessType = "SERVERLESS"
	deverloperType = "DEVELOPER"
)

type CreateServerlessOpts struct {
	serverlessProviders []*clusterApi.ListProviderRegionsOKBodyItemsItems0
}

func CreateCmd(h *internal.Helper) *cobra.Command {
	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create one cluster in the specified project.",
		Example: fmt.Sprintf(`  Create a cluster in interactive mode:
  $ %[1]s cluster create

  Create a cluster in non-interactive mode:
  $ %[1]s cluster create --project-id <project-id> --cluster-name <cluster-name> --cloud-provider <cloud-provider> -r <region> --root-password <password> --cluster-type <cluster-type>`,
			config.CliName),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// mark required flags in non-interactive mode
			if cmd.Flags().NFlag() != 0 {
				err := cmd.MarkFlagRequired(flag.ClusterName)
				if err != nil {
					return errors.Trace(err)
				}
				err = cmd.MarkFlagRequired(flag.ClusterType)
				if err != nil {
					return errors.Trace(err)
				}
				err = cmd.MarkFlagRequired(flag.CloudProvider)
				if err != nil {
					return errors.Trace(err)
				}
				err = cmd.MarkFlagRequired(flag.Region)
				if err != nil {
					return errors.Trace(err)
				}
				err = cmd.MarkFlagRequired(flag.RootPassword)
				if err != nil {
					return errors.Trace(err)
				}
				err = cmd.MarkFlagRequired(flag.ProjectID)
				if err != nil {
					return errors.Trace(err)
				}
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			d := h.Client()

			var clusterName string
			var clusterType string
			var cloudProvider string
			var region string
			var rootPassword string
			var projectID string
			if cmd.Flags().NFlag() == 0 {
				if !h.IOStreams.CanPrompt {
					return errors.New("The terminal doesn't support interactive mode, please use non-interactive mode")
				}

				// interactive mode
				regions, err := d.ListProviderRegions(clusterApi.NewListProviderRegionsParams())
				if err != nil {
					return errors.Trace(err)
				}

				opts := CreateServerlessOpts{}
				for i, item := range regions.Payload.Items {
					// Filter out non-serverless providers, currently only serverless is supported.
					// But currently serverless is called "DEVELOPER" in the API.
					if item.ClusterType == deverloperType {
						opts.serverlessProviders = append(opts.serverlessProviders, regions.Payload.Items[i])
					}
				}

				model, err := ui.InitialSelectModel([]interface{}{serverlessType}, "Choose the cluster type:")
				if err != nil {
					return err
				}
				p := tea.NewProgram(model)
				typeModel, err := p.StartReturningModel()
				if err != nil {
					return errors.Trace(err)
				}
				if m, _ := typeModel.(ui.SelectModel); m.Interrupted {
					return nil
				}
				clusterType = typeModel.(ui.SelectModel).Choices[typeModel.(ui.SelectModel).Selected].(string)

				// distinct cloud providers
				set := hashset.New()
				for _, provider := range opts.serverlessProviders {
					set.Add(provider.CloudProvider)
				}
				model, err = ui.InitialSelectModel(set.Values(), "Choose the cloud provider:")
				if err != nil {
					return err
				}
				p = tea.NewProgram(model)
				providerModel, err := p.StartReturningModel()
				if err != nil {
					return errors.Trace(err)
				}
				if m, _ := providerModel.(ui.SelectModel); m.Interrupted {
					return nil
				}
				cloudProvider = providerModel.(ui.SelectModel).Choices[providerModel.(ui.SelectModel).Selected].(string)

				// filter out regions for the selected cloud provider
				set = hashset.New()
				for _, provider := range opts.serverlessProviders {
					if provider.CloudProvider == providerModel.(ui.SelectModel).Choices[providerModel.(ui.SelectModel).Selected] {
						set.Add(provider.Region)
					}
				}
				model, err = ui.InitialSelectModel(set.Values(), "Choose the cloud region:")
				if err != nil {
					return err
				}
				p = tea.NewProgram(model)
				regionModel, err := p.StartReturningModel()
				if err != nil {
					return errors.Trace(err)
				}
				if m, _ := regionModel.(ui.SelectModel); m.Interrupted {
					return nil
				}
				region = regionModel.(ui.SelectModel).Choices[regionModel.(ui.SelectModel).Selected].(string)

				_, items, err := project.RetrieveProjects(h.QueryPageSize, h.Client())
				if err != nil {
					return err
				}
				set = hashset.New()
				for _, item := range items {
					set.Add(item.ID)
				}
				model, err = ui.InitialSelectModel(set.Values(), "Choose the project ID:")
				if err != nil {
					return err
				}
				p = tea.NewProgram(model)
				projectModel, err := p.StartReturningModel()
				if err != nil {
					return errors.Trace(err)
				}
				if m, _ := projectModel.(ui.SelectModel); m.Interrupted {
					return nil
				}
				projectID = projectModel.(ui.SelectModel).Choices[projectModel.(ui.SelectModel).Selected].(string)

				// variables for input
				p = tea.NewProgram(initialCreateInputModel())
				inputModel, err := p.StartReturningModel()
				if err != nil {
					return errors.Trace(err)
				}
				if inputModel.(ui.TextInputModel).Interrupted {
					return nil
				}

				clusterName = inputModel.(ui.TextInputModel).Inputs[clusterNameIdx].Value()
				rootPassword = inputModel.(ui.TextInputModel).Inputs[passwordIdx].Value()
			} else {
				// non-interactive mode, get values from flags
				var err error
				clusterName, err = cmd.Flags().GetString(flag.ClusterName)
				if err != nil {
					return errors.Trace(err)
				}
				clusterType, err = cmd.Flags().GetString(flag.ClusterType)
				if err != nil {
					return errors.Trace(err)
				}
				cloudProvider, err = cmd.Flags().GetString(flag.CloudProvider)
				if err != nil {
					return errors.Trace(err)
				}
				region, err = cmd.Flags().GetString(flag.Region)
				if err != nil {
					return errors.Trace(err)
				}
				rootPassword, err = cmd.Flags().GetString(flag.RootPassword)
				if err != nil {
					return errors.Trace(err)
				}
				projectID, err = cmd.Flags().GetString(flag.ProjectID)
				if err != nil {
					return errors.Trace(err)
				}
			}

			if clusterType != serverlessType {
				return errors.New("Currently only \"SERVERLESS\" cluster are supported to create in CLI")
			} else {
				// Currently serverless type is called \"DEVELOPER\" in API, but it will be changed to \"SERVERLESS\" soon.
				clusterType = deverloperType
			}

			clusterDefBody := &clusterApi.CreateClusterBody{}

			err := clusterDefBody.UnmarshalBinary([]byte(fmt.Sprintf(`{
			"name": "%s",
			"cluster_type": "%s",
			"cloud_provider": "%s",
			"region": "%s",
			"config" : {
				"root_password": "%s",
				"ip_access_list": [
					{
						"CIDR": "0.0.0.0/0",
						"description": "Allow All"
					}
				]
			}
			}`, clusterName, clusterType, cloudProvider, region, rootPassword)))
			if err != nil {
				return errors.Trace(err)
			}

			if h.IOStreams.CanPrompt {
				err2, done := CreateAndSpinnerWait(d, projectID, clusterDefBody, h)
				if done {
					return errors.Trace(err2)
				}
			} else {
				err2, done := CreateAndWaitReady(h, d, projectID, clusterDefBody)
				if done {
					return err2
				}
			}

			return nil
		},
	}

	createCmd.Flags().String(flag.ClusterName, "", "Name of the cluster to de created")
	createCmd.Flags().String(flag.ClusterType, "", "Cluster type, only support \"SERVERLESS\" now")
	createCmd.Flags().String(flag.CloudProvider, "", "Cloud provider, e.g. AWS")
	createCmd.Flags().StringP(flag.Region, flag.RegionShort, "", "Cloud region")
	createCmd.Flags().StringP(flag.ProjectID, flag.ProjectIDShort, "", "The ID of the project, in which the cluster will be created")
	createCmd.Flags().String(flag.RootPassword, "", "The root password of the cluster")
	return createCmd
}

func CreateAndWaitReady(h *internal.Helper, d util.CloudClient, projectID string, clusterDefBody *clusterApi.CreateClusterBody) (error, bool) {
	createClusterResult, err := d.CreateCluster(clusterApi.NewCreateClusterParams().WithProjectID(projectID).WithBody(*clusterDefBody))
	if err != nil {
		return errors.Trace(err), true
	}
	newClusterID := *createClusterResult.GetPayload().ID
	ticker := time.NewTicker(1 * time.Second)

	fmt.Fprintln(h.IOStreams.Out, "Waiting for cluster to be ready...")
	for {
		select {
		case <-time.After(2 * time.Minute):
			return errors.New("Timeout waiting for cluster to be ready, please check status on dashboard."), true
		case <-ticker.C:
			clusterResult, err := d.GetCluster(clusterApi.NewGetClusterParams().
				WithClusterID(newClusterID).
				WithProjectID(projectID))
			if err != nil {
				return errors.Trace(err), true
			}
			s := clusterResult.GetPayload().Status.ClusterStatus
			if s == "AVAILABLE" {
				fmt.Fprint(h.IOStreams.Out, color.GreenString("Cluster %s is ready.", newClusterID))
				return nil, true
			}
		}
	}
}

func CreateAndSpinnerWait(d util.CloudClient, projectID string, clusterDefBody *clusterApi.CreateClusterBody, h *internal.Helper) (error, bool) {
	// use spinner to indicate that the cluster is being created
	task := func() tea.Msg {
		createClusterResult, err := d.CreateCluster(clusterApi.NewCreateClusterParams().WithProjectID(projectID).WithBody(*clusterDefBody))
		if err != nil {
			return errors.Trace(err)
		}
		newClusterID := *createClusterResult.GetPayload().ID

		ticker := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-time.After(2 * time.Minute):
				return ui.Result("Timeout waiting for cluster to be ready, please check status on dashboard.")
			case <-ticker.C:
				clusterResult, err := d.GetCluster(clusterApi.NewGetClusterParams().
					WithClusterID(newClusterID).
					WithProjectID(projectID))
				if err != nil {
					return errors.Trace(err)
				}
				s := clusterResult.GetPayload().Status.ClusterStatus
				if s == "AVAILABLE" {
					return ui.Result(fmt.Sprintf("Cluster %s is ready.", newClusterID))
				}
			}
		}
	}

	p := tea.NewProgram(ui.InitialSpinnerModel(task, "Waiting for cluster to be ready..."))
	createModel, err := p.StartReturningModel()
	if err != nil {
		return errors.Trace(err), true
	}
	if m, _ := createModel.(ui.SpinnerModel); m.Err != nil {
		fmt.Fprintln(h.IOStreams.Err, color.RedString(m.Err.Error()))
	} else {
		fmt.Fprintln(h.IOStreams.Out, color.GreenString(m.Output))
	}
	return nil, false
}

func initialCreateInputModel() ui.TextInputModel {
	m := ui.TextInputModel{
		Inputs: make([]textinput.Model, 2),
	}

	var t textinput.Model
	for i := range m.Inputs {
		t = textinput.New()
		t.CursorStyle = cursorStyle
		t.CharLimit = 64
		f := createClusterField(i)

		switch f {
		case clusterNameIdx:
			t.Placeholder = "Cluster Name"
			t.Validate = func(s string) error {
				if len(s) == 0 {
					return errors.New("cluster Name is required")
				}
				return nil
			}
		case passwordIdx:
			t.Placeholder = "Password"
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'
			t.Validate = func(s string) error {
				if len(s) == 0 {
					return errors.New("password is required")
				}
				return nil
			}
		}

		m.Inputs[i] = t
	}

	return m
}
