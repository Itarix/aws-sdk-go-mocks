// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/sagemakeredgemanager/sagemakeredgemanageriface (interfaces: SagemakerEdgeManagerAPI)

// Package sagemakeredgemanagermock is a generated GoMock package.
package sagemakeredgemanagermock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	sagemakeredgemanager "github.com/aws/aws-sdk-go/service/sagemakeredgemanager"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSagemakerEdgeManagerAPI is a mock of SagemakerEdgeManagerAPI interface
type MockSagemakerEdgeManagerAPI struct {
	ctrl     *gomock.Controller
	recorder *MockSagemakerEdgeManagerAPIMockRecorder
}

// MockSagemakerEdgeManagerAPIMockRecorder is the mock recorder for MockSagemakerEdgeManagerAPI
type MockSagemakerEdgeManagerAPIMockRecorder struct {
	mock *MockSagemakerEdgeManagerAPI
}

// NewMockSagemakerEdgeManagerAPI creates a new mock instance
func NewMockSagemakerEdgeManagerAPI(ctrl *gomock.Controller) *MockSagemakerEdgeManagerAPI {
	mock := &MockSagemakerEdgeManagerAPI{ctrl: ctrl}
	mock.recorder = &MockSagemakerEdgeManagerAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSagemakerEdgeManagerAPI) EXPECT() *MockSagemakerEdgeManagerAPIMockRecorder {
	return m.recorder
}

// GetDeviceRegistration mocks base method
func (m *MockSagemakerEdgeManagerAPI) GetDeviceRegistration(arg0 *sagemakeredgemanager.GetDeviceRegistrationInput) (*sagemakeredgemanager.GetDeviceRegistrationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceRegistration", arg0)
	ret0, _ := ret[0].(*sagemakeredgemanager.GetDeviceRegistrationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceRegistration indicates an expected call of GetDeviceRegistration
func (mr *MockSagemakerEdgeManagerAPIMockRecorder) GetDeviceRegistration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceRegistration", reflect.TypeOf((*MockSagemakerEdgeManagerAPI)(nil).GetDeviceRegistration), arg0)
}

// GetDeviceRegistrationRequest mocks base method
func (m *MockSagemakerEdgeManagerAPI) GetDeviceRegistrationRequest(arg0 *sagemakeredgemanager.GetDeviceRegistrationInput) (*request.Request, *sagemakeredgemanager.GetDeviceRegistrationOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceRegistrationRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sagemakeredgemanager.GetDeviceRegistrationOutput)
	return ret0, ret1
}

// GetDeviceRegistrationRequest indicates an expected call of GetDeviceRegistrationRequest
func (mr *MockSagemakerEdgeManagerAPIMockRecorder) GetDeviceRegistrationRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceRegistrationRequest", reflect.TypeOf((*MockSagemakerEdgeManagerAPI)(nil).GetDeviceRegistrationRequest), arg0)
}

// GetDeviceRegistrationWithContext mocks base method
func (m *MockSagemakerEdgeManagerAPI) GetDeviceRegistrationWithContext(arg0 context.Context, arg1 *sagemakeredgemanager.GetDeviceRegistrationInput, arg2 ...request.Option) (*sagemakeredgemanager.GetDeviceRegistrationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDeviceRegistrationWithContext", varargs...)
	ret0, _ := ret[0].(*sagemakeredgemanager.GetDeviceRegistrationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceRegistrationWithContext indicates an expected call of GetDeviceRegistrationWithContext
func (mr *MockSagemakerEdgeManagerAPIMockRecorder) GetDeviceRegistrationWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceRegistrationWithContext", reflect.TypeOf((*MockSagemakerEdgeManagerAPI)(nil).GetDeviceRegistrationWithContext), varargs...)
}

// SendHeartbeat mocks base method
func (m *MockSagemakerEdgeManagerAPI) SendHeartbeat(arg0 *sagemakeredgemanager.SendHeartbeatInput) (*sagemakeredgemanager.SendHeartbeatOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeartbeat", arg0)
	ret0, _ := ret[0].(*sagemakeredgemanager.SendHeartbeatOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendHeartbeat indicates an expected call of SendHeartbeat
func (mr *MockSagemakerEdgeManagerAPIMockRecorder) SendHeartbeat(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeartbeat", reflect.TypeOf((*MockSagemakerEdgeManagerAPI)(nil).SendHeartbeat), arg0)
}

// SendHeartbeatRequest mocks base method
func (m *MockSagemakerEdgeManagerAPI) SendHeartbeatRequest(arg0 *sagemakeredgemanager.SendHeartbeatInput) (*request.Request, *sagemakeredgemanager.SendHeartbeatOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeartbeatRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sagemakeredgemanager.SendHeartbeatOutput)
	return ret0, ret1
}

// SendHeartbeatRequest indicates an expected call of SendHeartbeatRequest
func (mr *MockSagemakerEdgeManagerAPIMockRecorder) SendHeartbeatRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeartbeatRequest", reflect.TypeOf((*MockSagemakerEdgeManagerAPI)(nil).SendHeartbeatRequest), arg0)
}

// SendHeartbeatWithContext mocks base method
func (m *MockSagemakerEdgeManagerAPI) SendHeartbeatWithContext(arg0 context.Context, arg1 *sagemakeredgemanager.SendHeartbeatInput, arg2 ...request.Option) (*sagemakeredgemanager.SendHeartbeatOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendHeartbeatWithContext", varargs...)
	ret0, _ := ret[0].(*sagemakeredgemanager.SendHeartbeatOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendHeartbeatWithContext indicates an expected call of SendHeartbeatWithContext
func (mr *MockSagemakerEdgeManagerAPIMockRecorder) SendHeartbeatWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeartbeatWithContext", reflect.TypeOf((*MockSagemakerEdgeManagerAPI)(nil).SendHeartbeatWithContext), varargs...)
}