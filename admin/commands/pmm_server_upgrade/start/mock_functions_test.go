// Code generated by mockery v1.0.0. DO NOT EDIT.

package start

import (
	context "context"

	types "github.com/docker/docker/api/types"
	client "github.com/docker/docker/client"
	mock "github.com/stretchr/testify/mock"
)

// mockFunctions is an autogenerated mock type for the functions type
type mockFunctions struct {
	mock.Mock
}

// ContainerInspect provides a mock function with given fields: ctx, containerID
func (_m *mockFunctions) ContainerInspect(ctx context.Context, containerID string) (types.ContainerJSON, error) {
	ret := _m.Called(ctx, containerID)

	var r0 types.ContainerJSON
	if rf, ok := ret.Get(0).(func(context.Context, string) types.ContainerJSON); ok {
		r0 = rf(ctx, containerID)
	} else {
		r0 = ret.Get(0).(types.ContainerJSON)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, containerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerList provides a mock function with given fields: ctx, options
func (_m *mockFunctions) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
	ret := _m.Called(ctx, options)

	var r0 []types.Container
	if rf, ok := ret.Get(0).(func(context.Context, types.ContainerListOptions) []types.Container); ok {
		r0 = rf(ctx, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Container)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.ContainerListOptions) error); ok {
		r1 = rf(ctx, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindServerContainers provides a mock function with given fields: ctx
func (_m *mockFunctions) FindServerContainers(ctx context.Context) ([]types.Container, error) {
	ret := _m.Called(ctx)

	var r0 []types.Container
	if rf, ok := ret.Get(0).(func(context.Context) []types.Container); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Container)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDockerClient provides a mock function with given fields:
func (_m *mockFunctions) GetDockerClient() *client.Client {
	ret := _m.Called()

	var r0 *client.Client
	if rf, ok := ret.Get(0).(func() *client.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.Client)
		}
	}

	return r0
}

// HaveDockerAccess provides a mock function with given fields: ctx
func (_m *mockFunctions) HaveDockerAccess(ctx context.Context) bool {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsErrNotFound provides a mock function with given fields: err
func (_m *mockFunctions) IsErrNotFound(err error) bool {
	ret := _m.Called(err)

	var r0 bool
	if rf, ok := ret.Get(0).(func(error) bool); ok {
		r0 = rf(err)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
