// Code generated by mockery v2.16.0. DO NOT EDIT.

package mock

import (
	connect_info_service "tidbcloud-cli/pkg/tidbcloud/connect_info/client/connect_info_service"
	import_service "tidbcloud-cli/pkg/tidbcloud/import/client/import_service"

	cluster "github.com/c4pt0r/go-tidbcloud-sdk-v1/client/cluster"

	mock "github.com/stretchr/testify/mock"

	os "os"

	project "github.com/c4pt0r/go-tidbcloud-sdk-v1/client/project"
)

// TiDBCloudClient is an autogenerated mock type for the TiDBCloudClient type
type TiDBCloudClient struct {
	mock.Mock
}

// CancelImport provides a mock function with given fields: params, opts
func (_m *TiDBCloudClient) CancelImport(params *import_service.CancelImportParams, opts ...import_service.ClientOption) (*import_service.CancelImportOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *import_service.CancelImportOK
	if rf, ok := ret.Get(0).(func(*import_service.CancelImportParams, ...import_service.ClientOption) *import_service.CancelImportOK); ok {
		r0 = rf(params, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*import_service.CancelImportOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*import_service.CancelImportParams, ...import_service.ClientOption) error); ok {
		r1 = rf(params, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCluster provides a mock function with given fields: params, opts
func (_m *TiDBCloudClient) CreateCluster(params *cluster.CreateClusterParams, opts ...cluster.ClientOption) (*cluster.CreateClusterOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *cluster.CreateClusterOK
	if rf, ok := ret.Get(0).(func(*cluster.CreateClusterParams, ...cluster.ClientOption) *cluster.CreateClusterOK); ok {
		r0 = rf(params, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cluster.CreateClusterOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*cluster.CreateClusterParams, ...cluster.ClientOption) error); ok {
		r1 = rf(params, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateImport provides a mock function with given fields: params, opts
func (_m *TiDBCloudClient) CreateImport(params *import_service.CreateImportParams, opts ...import_service.ClientOption) (*import_service.CreateImportOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *import_service.CreateImportOK
	if rf, ok := ret.Get(0).(func(*import_service.CreateImportParams, ...import_service.ClientOption) *import_service.CreateImportOK); ok {
		r0 = rf(params, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*import_service.CreateImportOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*import_service.CreateImportParams, ...import_service.ClientOption) error); ok {
		r1 = rf(params, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCluster provides a mock function with given fields: params, opts
func (_m *TiDBCloudClient) DeleteCluster(params *cluster.DeleteClusterParams, opts ...cluster.ClientOption) (*cluster.DeleteClusterOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *cluster.DeleteClusterOK
	if rf, ok := ret.Get(0).(func(*cluster.DeleteClusterParams, ...cluster.ClientOption) *cluster.DeleteClusterOK); ok {
		r0 = rf(params, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cluster.DeleteClusterOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*cluster.DeleteClusterParams, ...cluster.ClientOption) error); ok {
		r1 = rf(params, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateUploadURL provides a mock function with given fields: params, opts
func (_m *TiDBCloudClient) GenerateUploadURL(params *import_service.GenerateUploadURLParams, opts ...import_service.ClientOption) (*import_service.GenerateUploadURLOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *import_service.GenerateUploadURLOK
	if rf, ok := ret.Get(0).(func(*import_service.GenerateUploadURLParams, ...import_service.ClientOption) *import_service.GenerateUploadURLOK); ok {
		r0 = rf(params, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*import_service.GenerateUploadURLOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*import_service.GenerateUploadURLParams, ...import_service.ClientOption) error); ok {
		r1 = rf(params, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCluster provides a mock function with given fields: params, opts
func (_m *TiDBCloudClient) GetCluster(params *cluster.GetClusterParams, opts ...cluster.ClientOption) (*cluster.GetClusterOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *cluster.GetClusterOK
	if rf, ok := ret.Get(0).(func(*cluster.GetClusterParams, ...cluster.ClientOption) *cluster.GetClusterOK); ok {
		r0 = rf(params, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cluster.GetClusterOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*cluster.GetClusterParams, ...cluster.ClientOption) error); ok {
		r1 = rf(params, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConnectInfo provides a mock function with given fields: params, opts
func (_m *TiDBCloudClient) GetConnectInfo(params *connect_info_service.GetInfoParams, opts ...connect_info_service.ClientOption) (*connect_info_service.GetInfoOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *connect_info_service.GetInfoOK
	if rf, ok := ret.Get(0).(func(*connect_info_service.GetInfoParams, ...connect_info_service.ClientOption) *connect_info_service.GetInfoOK); ok {
		r0 = rf(params, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect_info_service.GetInfoOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*connect_info_service.GetInfoParams, ...connect_info_service.ClientOption) error); ok {
		r1 = rf(params, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetImport provides a mock function with given fields: params, opts
func (_m *TiDBCloudClient) GetImport(params *import_service.GetImportParams, opts ...import_service.ClientOption) (*import_service.GetImportOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *import_service.GetImportOK
	if rf, ok := ret.Get(0).(func(*import_service.GetImportParams, ...import_service.ClientOption) *import_service.GetImportOK); ok {
		r0 = rf(params, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*import_service.GetImportOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*import_service.GetImportParams, ...import_service.ClientOption) error); ok {
		r1 = rf(params, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListClustersOfProject provides a mock function with given fields: params, opts
func (_m *TiDBCloudClient) ListClustersOfProject(params *cluster.ListClustersOfProjectParams, opts ...cluster.ClientOption) (*cluster.ListClustersOfProjectOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *cluster.ListClustersOfProjectOK
	if rf, ok := ret.Get(0).(func(*cluster.ListClustersOfProjectParams, ...cluster.ClientOption) *cluster.ListClustersOfProjectOK); ok {
		r0 = rf(params, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cluster.ListClustersOfProjectOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*cluster.ListClustersOfProjectParams, ...cluster.ClientOption) error); ok {
		r1 = rf(params, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListImports provides a mock function with given fields: params, opts
func (_m *TiDBCloudClient) ListImports(params *import_service.ListImportsParams, opts ...import_service.ClientOption) (*import_service.ListImportsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *import_service.ListImportsOK
	if rf, ok := ret.Get(0).(func(*import_service.ListImportsParams, ...import_service.ClientOption) *import_service.ListImportsOK); ok {
		r0 = rf(params, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*import_service.ListImportsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*import_service.ListImportsParams, ...import_service.ClientOption) error); ok {
		r1 = rf(params, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProjects provides a mock function with given fields: params, opts
func (_m *TiDBCloudClient) ListProjects(params *project.ListProjectsParams, opts ...project.ClientOption) (*project.ListProjectsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *project.ListProjectsOK
	if rf, ok := ret.Get(0).(func(*project.ListProjectsParams, ...project.ClientOption) *project.ListProjectsOK); ok {
		r0 = rf(params, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*project.ListProjectsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*project.ListProjectsParams, ...project.ClientOption) error); ok {
		r1 = rf(params, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProviderRegions provides a mock function with given fields: params, opts
func (_m *TiDBCloudClient) ListProviderRegions(params *cluster.ListProviderRegionsParams, opts ...cluster.ClientOption) (*cluster.ListProviderRegionsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *cluster.ListProviderRegionsOK
	if rf, ok := ret.Get(0).(func(*cluster.ListProviderRegionsParams, ...cluster.ClientOption) *cluster.ListProviderRegionsOK); ok {
		r0 = rf(params, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cluster.ListProviderRegionsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*cluster.ListProviderRegionsParams, ...cluster.ClientOption) error); ok {
		r1 = rf(params, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PreSignedUrlUpload provides a mock function with given fields: url, uploadFile, size
func (_m *TiDBCloudClient) PreSignedUrlUpload(url *string, uploadFile *os.File, size int64) error {
	ret := _m.Called(url, uploadFile, size)

	var r0 error
	if rf, ok := ret.Get(0).(func(*string, *os.File, int64) error); ok {
		r0 = rf(url, uploadFile, size)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTiDBCloudClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewTiDBCloudClient creates a new instance of TiDBCloudClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTiDBCloudClient(t mockConstructorTestingTNewTiDBCloudClient) *TiDBCloudClient {
	mock := &TiDBCloudClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
