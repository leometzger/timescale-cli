// Code generated by mockery v2.42.1. DO NOT EDIT.

package printer

import (
	printer "github.com/leometzger/timescale-cli/internal/printer"
	mock "github.com/stretchr/testify/mock"
)

// MockPrinter is an autogenerated mock type for the Printer type
type MockPrinter struct {
	mock.Mock
}

type MockPrinter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPrinter) EXPECT() *MockPrinter_Expecter {
	return &MockPrinter_Expecter{mock: &_m.Mock}
}

// Print provides a mock function with given fields: ref, values
func (_m *MockPrinter) Print(ref printer.Printable, values []printer.Printable) error {
	ret := _m.Called(ref, values)

	if len(ret) == 0 {
		panic("no return value specified for Print")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(printer.Printable, []printer.Printable) error); ok {
		r0 = rf(ref, values)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPrinter_Print_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Print'
type MockPrinter_Print_Call struct {
	*mock.Call
}

// Print is a helper method to define mock.On call
//   - ref printer.Printable
//   - values []printer.Printable
func (_e *MockPrinter_Expecter) Print(ref interface{}, values interface{}) *MockPrinter_Print_Call {
	return &MockPrinter_Print_Call{Call: _e.mock.On("Print", ref, values)}
}

func (_c *MockPrinter_Print_Call) Run(run func(ref printer.Printable, values []printer.Printable)) *MockPrinter_Print_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(printer.Printable), args[1].([]printer.Printable))
	})
	return _c
}

func (_c *MockPrinter_Print_Call) Return(_a0 error) *MockPrinter_Print_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPrinter_Print_Call) RunAndReturn(run func(printer.Printable, []printer.Printable) error) *MockPrinter_Print_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPrinter creates a new instance of MockPrinter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPrinter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPrinter {
	mock := &MockPrinter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
