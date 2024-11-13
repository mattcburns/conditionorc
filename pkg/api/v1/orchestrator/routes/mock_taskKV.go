// Code generated by mockery v2.42.1. DO NOT EDIT.

package routes

import (
	context "context"

	condition "github.com/metal-toolbox/rivets/v2/condition"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MocktaskKV is an autogenerated mock type for the taskKV type
type MocktaskKV struct {
	mock.Mock
}

type MocktaskKV_Expecter struct {
	mock *mock.Mock
}

func (_m *MocktaskKV) EXPECT() *MocktaskKV_Expecter {
	return &MocktaskKV_Expecter{mock: &_m.Mock}
}

// get provides a mock function with given fields: ctx, conditionKind, conditionID, serverID
func (_m *MocktaskKV) get(ctx context.Context, conditionKind condition.Kind, conditionID uuid.UUID, serverID uuid.UUID) (*condition.Task[interface{}, interface{}], error) {
	ret := _m.Called(ctx, conditionKind, conditionID, serverID)

	if len(ret) == 0 {
		panic("no return value specified for get")
	}

	var r0 *condition.Task[interface{}, interface{}]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, condition.Kind, uuid.UUID, uuid.UUID) (*condition.Task[interface{}, interface{}], error)); ok {
		return rf(ctx, conditionKind, conditionID, serverID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, condition.Kind, uuid.UUID, uuid.UUID) *condition.Task[interface{}, interface{}]); ok {
		r0 = rf(ctx, conditionKind, conditionID, serverID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*condition.Task[interface{}, interface{}])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, condition.Kind, uuid.UUID, uuid.UUID) error); ok {
		r1 = rf(ctx, conditionKind, conditionID, serverID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MocktaskKV_get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'get'
type MocktaskKV_get_Call struct {
	*mock.Call
}

// get is a helper method to define mock.On call
//   - ctx context.Context
//   - conditionKind condition.Kind
//   - conditionID uuid.UUID
//   - serverID uuid.UUID
func (_e *MocktaskKV_Expecter) get(ctx interface{}, conditionKind interface{}, conditionID interface{}, serverID interface{}) *MocktaskKV_get_Call {
	return &MocktaskKV_get_Call{Call: _e.mock.On("get", ctx, conditionKind, conditionID, serverID)}
}

func (_c *MocktaskKV_get_Call) Run(run func(ctx context.Context, conditionKind condition.Kind, conditionID uuid.UUID, serverID uuid.UUID)) *MocktaskKV_get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(condition.Kind), args[2].(uuid.UUID), args[3].(uuid.UUID))
	})
	return _c
}

func (_c *MocktaskKV_get_Call) Return(_a0 *condition.Task[interface{}, interface{}], _a1 error) *MocktaskKV_get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MocktaskKV_get_Call) RunAndReturn(run func(context.Context, condition.Kind, uuid.UUID, uuid.UUID) (*condition.Task[interface{}, interface{}], error)) *MocktaskKV_get_Call {
	_c.Call.Return(run)
	return _c
}

// publish provides a mock function with given fields: ctx, serverID, conditionID, conditionKind, task, onlyTimestamp
func (_m *MocktaskKV) publish(ctx context.Context, serverID string, conditionID string, conditionKind condition.Kind, task *condition.Task[interface{}, interface{}], onlyTimestamp bool) error {
	ret := _m.Called(ctx, serverID, conditionID, conditionKind, task, onlyTimestamp)

	if len(ret) == 0 {
		panic("no return value specified for publish")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, condition.Kind, *condition.Task[interface{}, interface{}], bool) error); ok {
		r0 = rf(ctx, serverID, conditionID, conditionKind, task, onlyTimestamp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MocktaskKV_publish_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'publish'
type MocktaskKV_publish_Call struct {
	*mock.Call
}

// publish is a helper method to define mock.On call
//   - ctx context.Context
//   - serverID string
//   - conditionID string
//   - conditionKind condition.Kind
//   - task *condition.Task[interface{},interface{}]
//   - onlyTimestamp bool
func (_e *MocktaskKV_Expecter) publish(ctx interface{}, serverID interface{}, conditionID interface{}, conditionKind interface{}, task interface{}, onlyTimestamp interface{}) *MocktaskKV_publish_Call {
	return &MocktaskKV_publish_Call{Call: _e.mock.On("publish", ctx, serverID, conditionID, conditionKind, task, onlyTimestamp)}
}

func (_c *MocktaskKV_publish_Call) Run(run func(ctx context.Context, serverID string, conditionID string, conditionKind condition.Kind, task *condition.Task[interface{}, interface{}], onlyTimestamp bool)) *MocktaskKV_publish_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(condition.Kind), args[4].(*condition.Task[interface{}, interface{}]), args[5].(bool))
	})
	return _c
}

func (_c *MocktaskKV_publish_Call) Return(_a0 error) *MocktaskKV_publish_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MocktaskKV_publish_Call) RunAndReturn(run func(context.Context, string, string, condition.Kind, *condition.Task[interface{}, interface{}], bool) error) *MocktaskKV_publish_Call {
	_c.Call.Return(run)
	return _c
}

// NewMocktaskKV creates a new instance of MocktaskKV. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMocktaskKV(t interface {
	mock.TestingT
	Cleanup(func())
}) *MocktaskKV {
	mock := &MocktaskKV{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
