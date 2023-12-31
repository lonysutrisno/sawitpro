// Code generated by MockGen. DO NOT EDIT.
// Source: sawitpro/pkg (interfaces: JWTInterface)

// Package pkg is a generated GoMock package.
package mock

import (
	reflect "reflect"
	pkg "sawitpro/pkg"

	gomock "github.com/golang/mock/gomock"
)

// MockJWTInterface is a mock of JWTInterface interface.
type MockJWTInterface struct {
	ctrl     *gomock.Controller
	recorder *MockJWTInterfaceMockRecorder
}

// MockJWTInterfaceMockRecorder is the mock recorder for MockJWTInterface.
type MockJWTInterfaceMockRecorder struct {
	mock *MockJWTInterface
}

// NewMockJWTInterface creates a new mock instance.
func NewMockJWTInterface(ctrl *gomock.Controller) *MockJWTInterface {
	mock := &MockJWTInterface{ctrl: ctrl}
	mock.recorder = &MockJWTInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJWTInterface) EXPECT() *MockJWTInterfaceMockRecorder {
	return m.recorder
}

// GenerateJWTToken mocks base method.
func (m *MockJWTInterface) GenerateJWTToken(arg0 int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateJWTToken", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateJWTToken indicates an expected call of GenerateJWTToken.
func (mr *MockJWTInterfaceMockRecorder) GenerateJWTToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateJWTToken", reflect.TypeOf((*MockJWTInterface)(nil).GenerateJWTToken), arg0)
}

// Validate mocks base method.
func (m *MockJWTInterface) Validate(arg0 string) (pkg.JWTClaims, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0)
	ret0, _ := ret[0].(pkg.JWTClaims)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Validate indicates an expected call of Validate.
func (mr *MockJWTInterfaceMockRecorder) Validate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockJWTInterface)(nil).Validate), arg0)
}
