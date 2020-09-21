// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/core/charm (interfaces: StateCharm)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStateCharm is a mock of StateCharm interface
type MockStateCharm struct {
	ctrl     *gomock.Controller
	recorder *MockStateCharmMockRecorder
}

// MockStateCharmMockRecorder is the mock recorder for MockStateCharm
type MockStateCharmMockRecorder struct {
	mock *MockStateCharm
}

// NewMockStateCharm creates a new mock instance
func NewMockStateCharm(ctrl *gomock.Controller) *MockStateCharm {
	mock := &MockStateCharm{ctrl: ctrl}
	mock.recorder = &MockStateCharmMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStateCharm) EXPECT() *MockStateCharmMockRecorder {
	return m.recorder
}

// IsUploaded mocks base method
func (m *MockStateCharm) IsUploaded() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUploaded")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUploaded indicates an expected call of IsUploaded
func (mr *MockStateCharmMockRecorder) IsUploaded() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUploaded", reflect.TypeOf((*MockStateCharm)(nil).IsUploaded))
}
