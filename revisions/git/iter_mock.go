// Code generated by MockGen. DO NOT EDIT.
// Source: gopkg.in/src-d/go-git.v4/plumbing/storer (interfaces: ReferenceIter)

// Package git is a generated GoMock package.
package git

import (
	gomock "github.com/golang/mock/gomock"
	plumbing "gopkg.in/src-d/go-git.v4/plumbing"
	reflect "reflect"
)

// MockReferenceIter is a mock of ReferenceIter interface
type MockReferenceIter struct {
	ctrl     *gomock.Controller
	recorder *MockReferenceIterMockRecorder
}

// MockReferenceIterMockRecorder is the mock recorder for MockReferenceIter
type MockReferenceIterMockRecorder struct {
	mock *MockReferenceIter
}

// NewMockReferenceIter creates a new mock instance
func NewMockReferenceIter(ctrl *gomock.Controller) *MockReferenceIter {
	mock := &MockReferenceIter{ctrl: ctrl}
	mock.recorder = &MockReferenceIterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReferenceIter) EXPECT() *MockReferenceIterMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockReferenceIter) Close() {
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockReferenceIterMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockReferenceIter)(nil).Close))
}

// ForEach mocks base method
func (m *MockReferenceIter) ForEach(arg0 func(*plumbing.Reference) error) error {
	ret := m.ctrl.Call(m, "ForEach", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForEach indicates an expected call of ForEach
func (mr *MockReferenceIterMockRecorder) ForEach(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForEach", reflect.TypeOf((*MockReferenceIter)(nil).ForEach), arg0)
}

// Next mocks base method
func (m *MockReferenceIter) Next() (*plumbing.Reference, error) {
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(*plumbing.Reference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Next indicates an expected call of Next
func (mr *MockReferenceIterMockRecorder) Next() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockReferenceIter)(nil).Next))
}