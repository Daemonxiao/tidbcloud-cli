// Code generated by mockery v2.14.1. DO NOT EDIT.

package mock

import (
	cluster "github.com/c4pt0r/go-tidbcloud-sdk-v1/client/cluster"
	mock "github.com/stretchr/testify/mock"

	project "github.com/c4pt0r/go-tidbcloud-sdk-v1/client/project"
)

// ApiClient is an autogenerated mock type for the ApiClient type
type ApiClient struct {
	mock.Mock
}

// CreateCluster provides a mock function with given fields: params, opts
func (_m *ApiClient) CreateCluster(params *cluster.CreateClusterParams, opts ...cluster.ClientOption) (*cluster.CreateClusterOK, error) {
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

// DeleteCluster provides a mock function with given fields: params, opts
func (_m *ApiClient) DeleteCluster(params *cluster.DeleteClusterParams, opts ...cluster.ClientOption) (*cluster.DeleteClusterOK, error) {
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

// GetCluster provides a mock function with given fields: params, opts
func (_m *ApiClient) GetCluster(params *cluster.GetClusterParams, opts ...cluster.ClientOption) (*cluster.GetClusterOK, error) {
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

// ListClustersOfProject provides a mock function with given fields: params, opts
func (_m *ApiClient) ListClustersOfProject(params *cluster.ListClustersOfProjectParams, opts ...cluster.ClientOption) (*cluster.ListClustersOfProjectOK, error) {
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

// ListProjects provides a mock function with given fields: params, opts
func (_m *ApiClient) ListProjects(params *project.ListProjectsParams, opts ...project.ClientOption) (*project.ListProjectsOK, error) {
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
func (_m *ApiClient) ListProviderRegions(params *cluster.ListProviderRegionsParams, opts ...cluster.ClientOption) (*cluster.ListProviderRegionsOK, error) {
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

type mockConstructorTestingTNewApiClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewApiClient creates a new instance of ApiClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewApiClient(t mockConstructorTestingTNewApiClient) *ApiClient {
	mock := &ApiClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
