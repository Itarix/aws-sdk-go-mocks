// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/iotdataplane/iotdataplaneiface (interfaces: IoTDataPlaneAPI)

// Package iotdataplanemock is a generated GoMock package.
package iotdataplanemock

import (
	aws "github.com/aws/aws-sdk-go/aws"
	request "github.com/aws/aws-sdk-go/aws/request"
	iotdataplane "github.com/aws/aws-sdk-go/service/iotdataplane"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIoTDataPlaneAPI is a mock of IoTDataPlaneAPI interface
type MockIoTDataPlaneAPI struct {
	ctrl     *gomock.Controller
	recorder *MockIoTDataPlaneAPIMockRecorder
}

// MockIoTDataPlaneAPIMockRecorder is the mock recorder for MockIoTDataPlaneAPI
type MockIoTDataPlaneAPIMockRecorder struct {
	mock *MockIoTDataPlaneAPI
}

// NewMockIoTDataPlaneAPI creates a new mock instance
func NewMockIoTDataPlaneAPI(ctrl *gomock.Controller) *MockIoTDataPlaneAPI {
	mock := &MockIoTDataPlaneAPI{ctrl: ctrl}
	mock.recorder = &MockIoTDataPlaneAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIoTDataPlaneAPI) EXPECT() *MockIoTDataPlaneAPIMockRecorder {
	return m.recorder
}

// DeleteThingShadow mocks base method
func (m *MockIoTDataPlaneAPI) DeleteThingShadow(arg0 *iotdataplane.DeleteThingShadowInput) (*iotdataplane.DeleteThingShadowOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteThingShadow", arg0)
	ret0, _ := ret[0].(*iotdataplane.DeleteThingShadowOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteThingShadow indicates an expected call of DeleteThingShadow
func (mr *MockIoTDataPlaneAPIMockRecorder) DeleteThingShadow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteThingShadow", reflect.TypeOf((*MockIoTDataPlaneAPI)(nil).DeleteThingShadow), arg0)
}

// DeleteThingShadowRequest mocks base method
func (m *MockIoTDataPlaneAPI) DeleteThingShadowRequest(arg0 *iotdataplane.DeleteThingShadowInput) (*request.Request, *iotdataplane.DeleteThingShadowOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteThingShadowRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iotdataplane.DeleteThingShadowOutput)
	return ret0, ret1
}

// DeleteThingShadowRequest indicates an expected call of DeleteThingShadowRequest
func (mr *MockIoTDataPlaneAPIMockRecorder) DeleteThingShadowRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteThingShadowRequest", reflect.TypeOf((*MockIoTDataPlaneAPI)(nil).DeleteThingShadowRequest), arg0)
}

// DeleteThingShadowWithContext mocks base method
func (m *MockIoTDataPlaneAPI) DeleteThingShadowWithContext(arg0 aws.Context, arg1 *iotdataplane.DeleteThingShadowInput, arg2 ...request.Option) (*iotdataplane.DeleteThingShadowOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteThingShadowWithContext", varargs...)
	ret0, _ := ret[0].(*iotdataplane.DeleteThingShadowOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteThingShadowWithContext indicates an expected call of DeleteThingShadowWithContext
func (mr *MockIoTDataPlaneAPIMockRecorder) DeleteThingShadowWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteThingShadowWithContext", reflect.TypeOf((*MockIoTDataPlaneAPI)(nil).DeleteThingShadowWithContext), varargs...)
}

// GetThingShadow mocks base method
func (m *MockIoTDataPlaneAPI) GetThingShadow(arg0 *iotdataplane.GetThingShadowInput) (*iotdataplane.GetThingShadowOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetThingShadow", arg0)
	ret0, _ := ret[0].(*iotdataplane.GetThingShadowOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetThingShadow indicates an expected call of GetThingShadow
func (mr *MockIoTDataPlaneAPIMockRecorder) GetThingShadow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThingShadow", reflect.TypeOf((*MockIoTDataPlaneAPI)(nil).GetThingShadow), arg0)
}

// GetThingShadowRequest mocks base method
func (m *MockIoTDataPlaneAPI) GetThingShadowRequest(arg0 *iotdataplane.GetThingShadowInput) (*request.Request, *iotdataplane.GetThingShadowOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetThingShadowRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iotdataplane.GetThingShadowOutput)
	return ret0, ret1
}

// GetThingShadowRequest indicates an expected call of GetThingShadowRequest
func (mr *MockIoTDataPlaneAPIMockRecorder) GetThingShadowRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThingShadowRequest", reflect.TypeOf((*MockIoTDataPlaneAPI)(nil).GetThingShadowRequest), arg0)
}

// GetThingShadowWithContext mocks base method
func (m *MockIoTDataPlaneAPI) GetThingShadowWithContext(arg0 aws.Context, arg1 *iotdataplane.GetThingShadowInput, arg2 ...request.Option) (*iotdataplane.GetThingShadowOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetThingShadowWithContext", varargs...)
	ret0, _ := ret[0].(*iotdataplane.GetThingShadowOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetThingShadowWithContext indicates an expected call of GetThingShadowWithContext
func (mr *MockIoTDataPlaneAPIMockRecorder) GetThingShadowWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThingShadowWithContext", reflect.TypeOf((*MockIoTDataPlaneAPI)(nil).GetThingShadowWithContext), varargs...)
}

// Publish mocks base method
func (m *MockIoTDataPlaneAPI) Publish(arg0 *iotdataplane.PublishInput) (*iotdataplane.PublishOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", arg0)
	ret0, _ := ret[0].(*iotdataplane.PublishOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Publish indicates an expected call of Publish
func (mr *MockIoTDataPlaneAPIMockRecorder) Publish(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockIoTDataPlaneAPI)(nil).Publish), arg0)
}

// PublishRequest mocks base method
func (m *MockIoTDataPlaneAPI) PublishRequest(arg0 *iotdataplane.PublishInput) (*request.Request, *iotdataplane.PublishOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iotdataplane.PublishOutput)
	return ret0, ret1
}

// PublishRequest indicates an expected call of PublishRequest
func (mr *MockIoTDataPlaneAPIMockRecorder) PublishRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishRequest", reflect.TypeOf((*MockIoTDataPlaneAPI)(nil).PublishRequest), arg0)
}

// PublishWithContext mocks base method
func (m *MockIoTDataPlaneAPI) PublishWithContext(arg0 aws.Context, arg1 *iotdataplane.PublishInput, arg2 ...request.Option) (*iotdataplane.PublishOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PublishWithContext", varargs...)
	ret0, _ := ret[0].(*iotdataplane.PublishOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublishWithContext indicates an expected call of PublishWithContext
func (mr *MockIoTDataPlaneAPIMockRecorder) PublishWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishWithContext", reflect.TypeOf((*MockIoTDataPlaneAPI)(nil).PublishWithContext), varargs...)
}

// UpdateThingShadow mocks base method
func (m *MockIoTDataPlaneAPI) UpdateThingShadow(arg0 *iotdataplane.UpdateThingShadowInput) (*iotdataplane.UpdateThingShadowOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateThingShadow", arg0)
	ret0, _ := ret[0].(*iotdataplane.UpdateThingShadowOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateThingShadow indicates an expected call of UpdateThingShadow
func (mr *MockIoTDataPlaneAPIMockRecorder) UpdateThingShadow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateThingShadow", reflect.TypeOf((*MockIoTDataPlaneAPI)(nil).UpdateThingShadow), arg0)
}

// UpdateThingShadowRequest mocks base method
func (m *MockIoTDataPlaneAPI) UpdateThingShadowRequest(arg0 *iotdataplane.UpdateThingShadowInput) (*request.Request, *iotdataplane.UpdateThingShadowOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateThingShadowRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iotdataplane.UpdateThingShadowOutput)
	return ret0, ret1
}

// UpdateThingShadowRequest indicates an expected call of UpdateThingShadowRequest
func (mr *MockIoTDataPlaneAPIMockRecorder) UpdateThingShadowRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateThingShadowRequest", reflect.TypeOf((*MockIoTDataPlaneAPI)(nil).UpdateThingShadowRequest), arg0)
}

// UpdateThingShadowWithContext mocks base method
func (m *MockIoTDataPlaneAPI) UpdateThingShadowWithContext(arg0 aws.Context, arg1 *iotdataplane.UpdateThingShadowInput, arg2 ...request.Option) (*iotdataplane.UpdateThingShadowOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateThingShadowWithContext", varargs...)
	ret0, _ := ret[0].(*iotdataplane.UpdateThingShadowOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateThingShadowWithContext indicates an expected call of UpdateThingShadowWithContext
func (mr *MockIoTDataPlaneAPIMockRecorder) UpdateThingShadowWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateThingShadowWithContext", reflect.TypeOf((*MockIoTDataPlaneAPI)(nil).UpdateThingShadowWithContext), varargs...)
}
