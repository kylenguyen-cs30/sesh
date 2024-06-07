// Code generated by mockery v2.43.2. DO NOT EDIT.

package lister

import (
	model "github.com/joshmedeski/sesh/model"
	mock "github.com/stretchr/testify/mock"
)

// MockLister is an autogenerated mock type for the Lister type
type MockLister struct {
	mock.Mock
}

type MockLister_Expecter struct {
	mock *mock.Mock
}

func (_m *MockLister) EXPECT() *MockLister_Expecter {
	return &MockLister_Expecter{mock: &_m.Mock}
}

// FindConfigSession provides a mock function with given fields: name
func (_m *MockLister) FindConfigSession(name string) (model.SeshSession, bool) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for FindConfigSession")
	}

	var r0 model.SeshSession
	var r1 bool
	if rf, ok := ret.Get(0).(func(string) (model.SeshSession, bool)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) model.SeshSession); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(model.SeshSession)
	}

	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// MockLister_FindConfigSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindConfigSession'
type MockLister_FindConfigSession_Call struct {
	*mock.Call
}

// FindConfigSession is a helper method to define mock.On call
//   - name string
func (_e *MockLister_Expecter) FindConfigSession(name interface{}) *MockLister_FindConfigSession_Call {
	return &MockLister_FindConfigSession_Call{Call: _e.mock.On("FindConfigSession", name)}
}

func (_c *MockLister_FindConfigSession_Call) Run(run func(name string)) *MockLister_FindConfigSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockLister_FindConfigSession_Call) Return(_a0 model.SeshSession, _a1 bool) *MockLister_FindConfigSession_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLister_FindConfigSession_Call) RunAndReturn(run func(string) (model.SeshSession, bool)) *MockLister_FindConfigSession_Call {
	_c.Call.Return(run)
	return _c
}

// FindTmuxSession provides a mock function with given fields: name
func (_m *MockLister) FindTmuxSession(name string) (model.SeshSession, bool) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for FindTmuxSession")
	}

	var r0 model.SeshSession
	var r1 bool
	if rf, ok := ret.Get(0).(func(string) (model.SeshSession, bool)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) model.SeshSession); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(model.SeshSession)
	}

	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// MockLister_FindTmuxSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindTmuxSession'
type MockLister_FindTmuxSession_Call struct {
	*mock.Call
}

// FindTmuxSession is a helper method to define mock.On call
//   - name string
func (_e *MockLister_Expecter) FindTmuxSession(name interface{}) *MockLister_FindTmuxSession_Call {
	return &MockLister_FindTmuxSession_Call{Call: _e.mock.On("FindTmuxSession", name)}
}

func (_c *MockLister_FindTmuxSession_Call) Run(run func(name string)) *MockLister_FindTmuxSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockLister_FindTmuxSession_Call) Return(_a0 model.SeshSession, _a1 bool) *MockLister_FindTmuxSession_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLister_FindTmuxSession_Call) RunAndReturn(run func(string) (model.SeshSession, bool)) *MockLister_FindTmuxSession_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: opts
func (_m *MockLister) List(opts ListOptions) (model.SeshSessions, error) {
	ret := _m.Called(opts)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 model.SeshSessions
	var r1 error
	if rf, ok := ret.Get(0).(func(ListOptions) (model.SeshSessions, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(ListOptions) model.SeshSessions); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(model.SeshSessions)
	}

	if rf, ok := ret.Get(1).(func(ListOptions) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLister_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockLister_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - opts ListOptions
func (_e *MockLister_Expecter) List(opts interface{}) *MockLister_List_Call {
	return &MockLister_List_Call{Call: _e.mock.On("List", opts)}
}

func (_c *MockLister_List_Call) Run(run func(opts ListOptions)) *MockLister_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(ListOptions))
	})
	return _c
}

func (_c *MockLister_List_Call) Return(_a0 model.SeshSessions, _a1 error) *MockLister_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLister_List_Call) RunAndReturn(run func(ListOptions) (model.SeshSessions, error)) *MockLister_List_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockLister creates a new instance of MockLister. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLister(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLister {
	mock := &MockLister{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
