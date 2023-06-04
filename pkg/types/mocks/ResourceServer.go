// Code generated by mockery v2.30.16. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	v1beta1 "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

// ResourceServer is an autogenerated mock type for the ResourceServer type
type ResourceServer struct {
	mock.Mock
}

// Allocate provides a mock function with given fields: _a0, _a1
func (_m *ResourceServer) Allocate(_a0 context.Context, _a1 *v1beta1.AllocateRequest) (*v1beta1.AllocateResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1beta1.AllocateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.AllocateRequest) (*v1beta1.AllocateResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.AllocateRequest) *v1beta1.AllocateResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.AllocateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.AllocateRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDevicePluginOptions provides a mock function with given fields: _a0, _a1
func (_m *ResourceServer) GetDevicePluginOptions(_a0 context.Context, _a1 *v1beta1.Empty) (*v1beta1.DevicePluginOptions, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1beta1.DevicePluginOptions
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.Empty) (*v1beta1.DevicePluginOptions, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.Empty) *v1beta1.DevicePluginOptions); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.DevicePluginOptions)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.Empty) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPreferredAllocation provides a mock function with given fields: _a0, _a1
func (_m *ResourceServer) GetPreferredAllocation(_a0 context.Context, _a1 *v1beta1.PreferredAllocationRequest) (*v1beta1.PreferredAllocationResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1beta1.PreferredAllocationResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.PreferredAllocationRequest) (*v1beta1.PreferredAllocationResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.PreferredAllocationRequest) *v1beta1.PreferredAllocationResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.PreferredAllocationResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.PreferredAllocationRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Init provides a mock function with given fields:
func (_m *ResourceServer) Init() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListAndWatch provides a mock function with given fields: _a0, _a1
func (_m *ResourceServer) ListAndWatch(_a0 *v1beta1.Empty, _a1 v1beta1.DevicePlugin_ListAndWatchServer) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1beta1.Empty, v1beta1.DevicePlugin_ListAndWatchServer) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PreStartContainer provides a mock function with given fields: _a0, _a1
func (_m *ResourceServer) PreStartContainer(_a0 context.Context, _a1 *v1beta1.PreStartContainerRequest) (*v1beta1.PreStartContainerResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1beta1.PreStartContainerResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.PreStartContainerRequest) (*v1beta1.PreStartContainerResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.PreStartContainerRequest) *v1beta1.PreStartContainerResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.PreStartContainerResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.PreStartContainerRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Start provides a mock function with given fields:
func (_m *ResourceServer) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *ResourceServer) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Watch provides a mock function with given fields:
func (_m *ResourceServer) Watch() {
	_m.Called()
}

// NewResourceServer creates a new instance of ResourceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewResourceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *ResourceServer {
	mock := &ResourceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
