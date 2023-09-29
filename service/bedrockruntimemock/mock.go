// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/bedrockruntime/bedrockruntimeiface (interfaces: BedrockRuntimeAPI)

// Package bedrockruntimemock is a generated GoMock package.
package bedrockruntimemock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	bedrockruntime "github.com/aws/aws-sdk-go/service/bedrockruntime"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBedrockRuntimeAPI is a mock of BedrockRuntimeAPI interface
type MockBedrockRuntimeAPI struct {
	ctrl     *gomock.Controller
	recorder *MockBedrockRuntimeAPIMockRecorder
}

// MockBedrockRuntimeAPIMockRecorder is the mock recorder for MockBedrockRuntimeAPI
type MockBedrockRuntimeAPIMockRecorder struct {
	mock *MockBedrockRuntimeAPI
}

// NewMockBedrockRuntimeAPI creates a new mock instance
func NewMockBedrockRuntimeAPI(ctrl *gomock.Controller) *MockBedrockRuntimeAPI {
	mock := &MockBedrockRuntimeAPI{ctrl: ctrl}
	mock.recorder = &MockBedrockRuntimeAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBedrockRuntimeAPI) EXPECT() *MockBedrockRuntimeAPIMockRecorder {
	return m.recorder
}

// InvokeModel mocks base method
func (m *MockBedrockRuntimeAPI) InvokeModel(arg0 *bedrockruntime.InvokeModelInput) (*bedrockruntime.InvokeModelOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvokeModel", arg0)
	ret0, _ := ret[0].(*bedrockruntime.InvokeModelOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InvokeModel indicates an expected call of InvokeModel
func (mr *MockBedrockRuntimeAPIMockRecorder) InvokeModel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeModel", reflect.TypeOf((*MockBedrockRuntimeAPI)(nil).InvokeModel), arg0)
}

// InvokeModelRequest mocks base method
func (m *MockBedrockRuntimeAPI) InvokeModelRequest(arg0 *bedrockruntime.InvokeModelInput) (*request.Request, *bedrockruntime.InvokeModelOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvokeModelRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*bedrockruntime.InvokeModelOutput)
	return ret0, ret1
}

// InvokeModelRequest indicates an expected call of InvokeModelRequest
func (mr *MockBedrockRuntimeAPIMockRecorder) InvokeModelRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeModelRequest", reflect.TypeOf((*MockBedrockRuntimeAPI)(nil).InvokeModelRequest), arg0)
}

// InvokeModelWithContext mocks base method
func (m *MockBedrockRuntimeAPI) InvokeModelWithContext(arg0 context.Context, arg1 *bedrockruntime.InvokeModelInput, arg2 ...request.Option) (*bedrockruntime.InvokeModelOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InvokeModelWithContext", varargs...)
	ret0, _ := ret[0].(*bedrockruntime.InvokeModelOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InvokeModelWithContext indicates an expected call of InvokeModelWithContext
func (mr *MockBedrockRuntimeAPIMockRecorder) InvokeModelWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeModelWithContext", reflect.TypeOf((*MockBedrockRuntimeAPI)(nil).InvokeModelWithContext), varargs...)
}

// InvokeModelWithResponseStream mocks base method
func (m *MockBedrockRuntimeAPI) InvokeModelWithResponseStream(arg0 *bedrockruntime.InvokeModelWithResponseStreamInput) (*bedrockruntime.InvokeModelWithResponseStreamOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvokeModelWithResponseStream", arg0)
	ret0, _ := ret[0].(*bedrockruntime.InvokeModelWithResponseStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InvokeModelWithResponseStream indicates an expected call of InvokeModelWithResponseStream
func (mr *MockBedrockRuntimeAPIMockRecorder) InvokeModelWithResponseStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeModelWithResponseStream", reflect.TypeOf((*MockBedrockRuntimeAPI)(nil).InvokeModelWithResponseStream), arg0)
}

// InvokeModelWithResponseStreamRequest mocks base method
func (m *MockBedrockRuntimeAPI) InvokeModelWithResponseStreamRequest(arg0 *bedrockruntime.InvokeModelWithResponseStreamInput) (*request.Request, *bedrockruntime.InvokeModelWithResponseStreamOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvokeModelWithResponseStreamRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*bedrockruntime.InvokeModelWithResponseStreamOutput)
	return ret0, ret1
}

// InvokeModelWithResponseStreamRequest indicates an expected call of InvokeModelWithResponseStreamRequest
func (mr *MockBedrockRuntimeAPIMockRecorder) InvokeModelWithResponseStreamRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeModelWithResponseStreamRequest", reflect.TypeOf((*MockBedrockRuntimeAPI)(nil).InvokeModelWithResponseStreamRequest), arg0)
}

// InvokeModelWithResponseStreamWithContext mocks base method
func (m *MockBedrockRuntimeAPI) InvokeModelWithResponseStreamWithContext(arg0 context.Context, arg1 *bedrockruntime.InvokeModelWithResponseStreamInput, arg2 ...request.Option) (*bedrockruntime.InvokeModelWithResponseStreamOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InvokeModelWithResponseStreamWithContext", varargs...)
	ret0, _ := ret[0].(*bedrockruntime.InvokeModelWithResponseStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InvokeModelWithResponseStreamWithContext indicates an expected call of InvokeModelWithResponseStreamWithContext
func (mr *MockBedrockRuntimeAPIMockRecorder) InvokeModelWithResponseStreamWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeModelWithResponseStreamWithContext", reflect.TypeOf((*MockBedrockRuntimeAPI)(nil).InvokeModelWithResponseStreamWithContext), varargs...)
}