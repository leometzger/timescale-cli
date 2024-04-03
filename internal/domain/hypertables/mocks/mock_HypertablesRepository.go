// Code generated by mockery v2.42.1. DO NOT EDIT.

package hypertables

import (
	hypertables "github.com/leometzger/timescale-cli/internal/domain/hypertables"
	mock "github.com/stretchr/testify/mock"
)

// MockHypertablesRepository is an autogenerated mock type for the HypertablesRepository type
type MockHypertablesRepository struct {
	mock.Mock
}

type MockHypertablesRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockHypertablesRepository) EXPECT() *MockHypertablesRepository_Expecter {
	return &MockHypertablesRepository_Expecter{mock: &_m.Mock}
}

// GetHypertables provides a mock function with given fields: filter
func (_m *MockHypertablesRepository) GetHypertables(filter *hypertables.HypertablesFilter) ([]hypertables.HypertableInfo, error) {
	ret := _m.Called(filter)

	if len(ret) == 0 {
		panic("no return value specified for GetHypertables")
	}

	var r0 []hypertables.HypertableInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(*hypertables.HypertablesFilter) ([]hypertables.HypertableInfo, error)); ok {
		return rf(filter)
	}
	if rf, ok := ret.Get(0).(func(*hypertables.HypertablesFilter) []hypertables.HypertableInfo); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]hypertables.HypertableInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(*hypertables.HypertablesFilter) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockHypertablesRepository_GetHypertables_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetHypertables'
type MockHypertablesRepository_GetHypertables_Call struct {
	*mock.Call
}

// GetHypertables is a helper method to define mock.On call
//   - filter *hypertables.HypertablesFilter
func (_e *MockHypertablesRepository_Expecter) GetHypertables(filter interface{}) *MockHypertablesRepository_GetHypertables_Call {
	return &MockHypertablesRepository_GetHypertables_Call{Call: _e.mock.On("GetHypertables", filter)}
}

func (_c *MockHypertablesRepository_GetHypertables_Call) Run(run func(filter *hypertables.HypertablesFilter)) *MockHypertablesRepository_GetHypertables_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*hypertables.HypertablesFilter))
	})
	return _c
}

func (_c *MockHypertablesRepository_GetHypertables_Call) Return(_a0 []hypertables.HypertableInfo, _a1 error) *MockHypertablesRepository_GetHypertables_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockHypertablesRepository_GetHypertables_Call) RunAndReturn(run func(*hypertables.HypertablesFilter) ([]hypertables.HypertableInfo, error)) *MockHypertablesRepository_GetHypertables_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockHypertablesRepository creates a new instance of MockHypertablesRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockHypertablesRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockHypertablesRepository {
	mock := &MockHypertablesRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
