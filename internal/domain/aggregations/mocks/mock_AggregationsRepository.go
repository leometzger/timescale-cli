// Code generated by mockery v2.42.1. DO NOT EDIT.

package aggregations

import (
	aggregations "github.com/leometzger/timescale-cli/internal/domain/aggregations"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// MockAggregationsRepository is an autogenerated mock type for the AggregationsRepository type
type MockAggregationsRepository struct {
	mock.Mock
}

type MockAggregationsRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAggregationsRepository) EXPECT() *MockAggregationsRepository_Expecter {
	return &MockAggregationsRepository_Expecter{mock: &_m.Mock}
}

// Compress provides a mock function with given fields: viewName, olderThan, newerThan
func (_m *MockAggregationsRepository) Compress(viewName string, olderThan *time.Time, newerThan *time.Time) error {
	ret := _m.Called(viewName, olderThan, newerThan)

	if len(ret) == 0 {
		panic("no return value specified for Compress")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *time.Time, *time.Time) error); ok {
		r0 = rf(viewName, olderThan, newerThan)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAggregationsRepository_Compress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Compress'
type MockAggregationsRepository_Compress_Call struct {
	*mock.Call
}

// Compress is a helper method to define mock.On call
//   - viewName string
//   - olderThan *time.Time
//   - newerThan *time.Time
func (_e *MockAggregationsRepository_Expecter) Compress(viewName interface{}, olderThan interface{}, newerThan interface{}) *MockAggregationsRepository_Compress_Call {
	return &MockAggregationsRepository_Compress_Call{Call: _e.mock.On("Compress", viewName, olderThan, newerThan)}
}

func (_c *MockAggregationsRepository_Compress_Call) Run(run func(viewName string, olderThan *time.Time, newerThan *time.Time)) *MockAggregationsRepository_Compress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*time.Time), args[2].(*time.Time))
	})
	return _c
}

func (_c *MockAggregationsRepository_Compress_Call) Return(_a0 error) *MockAggregationsRepository_Compress_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAggregationsRepository_Compress_Call) RunAndReturn(run func(string, *time.Time, *time.Time) error) *MockAggregationsRepository_Compress_Call {
	_c.Call.Return(run)
	return _c
}

// GetAggregations provides a mock function with given fields: filter
func (_m *MockAggregationsRepository) GetAggregations(filter *aggregations.AggregationsFilter) ([]aggregations.ContinuousAggregation, error) {
	ret := _m.Called(filter)

	if len(ret) == 0 {
		panic("no return value specified for GetAggregations")
	}

	var r0 []aggregations.ContinuousAggregation
	var r1 error
	if rf, ok := ret.Get(0).(func(*aggregations.AggregationsFilter) ([]aggregations.ContinuousAggregation, error)); ok {
		return rf(filter)
	}
	if rf, ok := ret.Get(0).(func(*aggregations.AggregationsFilter) []aggregations.ContinuousAggregation); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]aggregations.ContinuousAggregation)
		}
	}

	if rf, ok := ret.Get(1).(func(*aggregations.AggregationsFilter) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAggregationsRepository_GetAggregations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAggregations'
type MockAggregationsRepository_GetAggregations_Call struct {
	*mock.Call
}

// GetAggregations is a helper method to define mock.On call
//   - filter *aggregations.AggregationsFilter
func (_e *MockAggregationsRepository_Expecter) GetAggregations(filter interface{}) *MockAggregationsRepository_GetAggregations_Call {
	return &MockAggregationsRepository_GetAggregations_Call{Call: _e.mock.On("GetAggregations", filter)}
}

func (_c *MockAggregationsRepository_GetAggregations_Call) Run(run func(filter *aggregations.AggregationsFilter)) *MockAggregationsRepository_GetAggregations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*aggregations.AggregationsFilter))
	})
	return _c
}

func (_c *MockAggregationsRepository_GetAggregations_Call) Return(_a0 []aggregations.ContinuousAggregation, _a1 error) *MockAggregationsRepository_GetAggregations_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAggregationsRepository_GetAggregations_Call) RunAndReturn(run func(*aggregations.AggregationsFilter) ([]aggregations.ContinuousAggregation, error)) *MockAggregationsRepository_GetAggregations_Call {
	_c.Call.Return(run)
	return _c
}

// Refresh provides a mock function with given fields: viewName, start, end
func (_m *MockAggregationsRepository) Refresh(viewName string, start time.Time, end time.Time) error {
	ret := _m.Called(viewName, start, end)

	if len(ret) == 0 {
		panic("no return value specified for Refresh")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, time.Time, time.Time) error); ok {
		r0 = rf(viewName, start, end)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAggregationsRepository_Refresh_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Refresh'
type MockAggregationsRepository_Refresh_Call struct {
	*mock.Call
}

// Refresh is a helper method to define mock.On call
//   - viewName string
//   - start time.Time
//   - end time.Time
func (_e *MockAggregationsRepository_Expecter) Refresh(viewName interface{}, start interface{}, end interface{}) *MockAggregationsRepository_Refresh_Call {
	return &MockAggregationsRepository_Refresh_Call{Call: _e.mock.On("Refresh", viewName, start, end)}
}

func (_c *MockAggregationsRepository_Refresh_Call) Run(run func(viewName string, start time.Time, end time.Time)) *MockAggregationsRepository_Refresh_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(time.Time), args[2].(time.Time))
	})
	return _c
}

func (_c *MockAggregationsRepository_Refresh_Call) Return(_a0 error) *MockAggregationsRepository_Refresh_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAggregationsRepository_Refresh_Call) RunAndReturn(run func(string, time.Time, time.Time) error) *MockAggregationsRepository_Refresh_Call {
	_c.Call.Return(run)
	return _c
}

// SetMaxTuplesDecompressedPerDmlTransaction provides a mock function with given fields: value
func (_m *MockAggregationsRepository) SetMaxTuplesDecompressedPerDmlTransaction(value int32) error {
	ret := _m.Called(value)

	if len(ret) == 0 {
		panic("no return value specified for SetMaxTuplesDecompressedPerDmlTransaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int32) error); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAggregationsRepository_SetMaxTuplesDecompressedPerDmlTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetMaxTuplesDecompressedPerDmlTransaction'
type MockAggregationsRepository_SetMaxTuplesDecompressedPerDmlTransaction_Call struct {
	*mock.Call
}

// SetMaxTuplesDecompressedPerDmlTransaction is a helper method to define mock.On call
//   - value int32
func (_e *MockAggregationsRepository_Expecter) SetMaxTuplesDecompressedPerDmlTransaction(value interface{}) *MockAggregationsRepository_SetMaxTuplesDecompressedPerDmlTransaction_Call {
	return &MockAggregationsRepository_SetMaxTuplesDecompressedPerDmlTransaction_Call{Call: _e.mock.On("SetMaxTuplesDecompressedPerDmlTransaction", value)}
}

func (_c *MockAggregationsRepository_SetMaxTuplesDecompressedPerDmlTransaction_Call) Run(run func(value int32)) *MockAggregationsRepository_SetMaxTuplesDecompressedPerDmlTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int32))
	})
	return _c
}

func (_c *MockAggregationsRepository_SetMaxTuplesDecompressedPerDmlTransaction_Call) Return(_a0 error) *MockAggregationsRepository_SetMaxTuplesDecompressedPerDmlTransaction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAggregationsRepository_SetMaxTuplesDecompressedPerDmlTransaction_Call) RunAndReturn(run func(int32) error) *MockAggregationsRepository_SetMaxTuplesDecompressedPerDmlTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAggregationsRepository creates a new instance of MockAggregationsRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAggregationsRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAggregationsRepository {
	mock := &MockAggregationsRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
