// Code generated by MockGen. DO NOT EDIT.
// Source: range.go

// Package range_test is a generated GoMock package.
package range_test

import (
	reflect "reflect"
)

import (
	gomock "github.com/golang/mock/gomock"
)

// MockRange is a mock of Range interface.
type MockRange struct {
	ctrl     *gomock.Controller
	recorder *MockRangeMockRecorder
}

// MockRangeMockRecorder is the mock recorder for MockRange.
type MockRangeMockRecorder struct {
	mock *MockRange
}

// NewMockRange creates a new mock instance.
func NewMockRange(ctrl *gomock.Controller) *MockRange {
	mock := &MockRange{ctrl: ctrl}
	mock.recorder = &MockRangeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRange) EXPECT() *MockRangeMockRecorder {
	return m.recorder
}

// HasNext mocks base method.
func (m *MockRange) HasNext() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasNext")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasNext indicates an expected call of HasNext.
func (mr *MockRangeMockRecorder) HasNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasNext", reflect.TypeOf((*MockRange)(nil).HasNext))
}

// Next mocks base method.
func (m *MockRange) Next() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockRangeMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockRange)(nil).Next))
}
