// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/configs/env.go

// Package mock_configs is a generated GoMock package.
package mock_configs

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEnvInterface is a mock of EnvInterface interface.
type MockEnvInterface struct {
	ctrl     *gomock.Controller
	recorder *MockEnvInterfaceMockRecorder
}

// MockEnvInterfaceMockRecorder is the mock recorder for MockEnvInterface.
type MockEnvInterfaceMockRecorder struct {
	mock *MockEnvInterface
}

// NewMockEnvInterface creates a new mock instance.
func NewMockEnvInterface(ctrl *gomock.Controller) *MockEnvInterface {
	mock := &MockEnvInterface{ctrl: ctrl}
	mock.recorder = &MockEnvInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnvInterface) EXPECT() *MockEnvInterfaceMockRecorder {
	return m.recorder
}

// GetDbName mocks base method.
func (m *MockEnvInterface) GetDbName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDbName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDbName indicates an expected call of GetDbName.
func (mr *MockEnvInterfaceMockRecorder) GetDbName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDbName", reflect.TypeOf((*MockEnvInterface)(nil).GetDbName))
}

// GetMongoUri mocks base method.
func (m *MockEnvInterface) GetMongoUri() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMongoUri")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetMongoUri indicates an expected call of GetMongoUri.
func (mr *MockEnvInterfaceMockRecorder) GetMongoUri() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMongoUri", reflect.TypeOf((*MockEnvInterface)(nil).GetMongoUri))
}
