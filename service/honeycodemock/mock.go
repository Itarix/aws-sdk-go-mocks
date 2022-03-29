// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/honeycode/honeycodeiface (interfaces: HoneycodeAPI)

// Package honeycodemock is a generated GoMock package.
package honeycodemock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	honeycode "github.com/aws/aws-sdk-go/service/honeycode"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockHoneycodeAPI is a mock of HoneycodeAPI interface
type MockHoneycodeAPI struct {
	ctrl     *gomock.Controller
	recorder *MockHoneycodeAPIMockRecorder
}

// MockHoneycodeAPIMockRecorder is the mock recorder for MockHoneycodeAPI
type MockHoneycodeAPIMockRecorder struct {
	mock *MockHoneycodeAPI
}

// NewMockHoneycodeAPI creates a new mock instance
func NewMockHoneycodeAPI(ctrl *gomock.Controller) *MockHoneycodeAPI {
	mock := &MockHoneycodeAPI{ctrl: ctrl}
	mock.recorder = &MockHoneycodeAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHoneycodeAPI) EXPECT() *MockHoneycodeAPIMockRecorder {
	return m.recorder
}

// GetScreenData mocks base method
func (m *MockHoneycodeAPI) GetScreenData(arg0 *honeycode.GetScreenDataInput) (*honeycode.GetScreenDataOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScreenData", arg0)
	ret0, _ := ret[0].(*honeycode.GetScreenDataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScreenData indicates an expected call of GetScreenData
func (mr *MockHoneycodeAPIMockRecorder) GetScreenData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScreenData", reflect.TypeOf((*MockHoneycodeAPI)(nil).GetScreenData), arg0)
}

// GetScreenDataRequest mocks base method
func (m *MockHoneycodeAPI) GetScreenDataRequest(arg0 *honeycode.GetScreenDataInput) (*request.Request, *honeycode.GetScreenDataOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScreenDataRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*honeycode.GetScreenDataOutput)
	return ret0, ret1
}

// GetScreenDataRequest indicates an expected call of GetScreenDataRequest
func (mr *MockHoneycodeAPIMockRecorder) GetScreenDataRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScreenDataRequest", reflect.TypeOf((*MockHoneycodeAPI)(nil).GetScreenDataRequest), arg0)
}

// GetScreenDataWithContext mocks base method
func (m *MockHoneycodeAPI) GetScreenDataWithContext(arg0 context.Context, arg1 *honeycode.GetScreenDataInput, arg2 ...request.Option) (*honeycode.GetScreenDataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetScreenDataWithContext", varargs...)
	ret0, _ := ret[0].(*honeycode.GetScreenDataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScreenDataWithContext indicates an expected call of GetScreenDataWithContext
func (mr *MockHoneycodeAPIMockRecorder) GetScreenDataWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScreenDataWithContext", reflect.TypeOf((*MockHoneycodeAPI)(nil).GetScreenDataWithContext), varargs...)
}

// InvokeScreenAutomation mocks base method
func (m *MockHoneycodeAPI) InvokeScreenAutomation(arg0 *honeycode.InvokeScreenAutomationInput) (*honeycode.InvokeScreenAutomationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvokeScreenAutomation", arg0)
	ret0, _ := ret[0].(*honeycode.InvokeScreenAutomationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InvokeScreenAutomation indicates an expected call of InvokeScreenAutomation
func (mr *MockHoneycodeAPIMockRecorder) InvokeScreenAutomation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeScreenAutomation", reflect.TypeOf((*MockHoneycodeAPI)(nil).InvokeScreenAutomation), arg0)
}

// InvokeScreenAutomationRequest mocks base method
func (m *MockHoneycodeAPI) InvokeScreenAutomationRequest(arg0 *honeycode.InvokeScreenAutomationInput) (*request.Request, *honeycode.InvokeScreenAutomationOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvokeScreenAutomationRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*honeycode.InvokeScreenAutomationOutput)
	return ret0, ret1
}

// InvokeScreenAutomationRequest indicates an expected call of InvokeScreenAutomationRequest
func (mr *MockHoneycodeAPIMockRecorder) InvokeScreenAutomationRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeScreenAutomationRequest", reflect.TypeOf((*MockHoneycodeAPI)(nil).InvokeScreenAutomationRequest), arg0)
}

// InvokeScreenAutomationWithContext mocks base method
func (m *MockHoneycodeAPI) InvokeScreenAutomationWithContext(arg0 context.Context, arg1 *honeycode.InvokeScreenAutomationInput, arg2 ...request.Option) (*honeycode.InvokeScreenAutomationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InvokeScreenAutomationWithContext", varargs...)
	ret0, _ := ret[0].(*honeycode.InvokeScreenAutomationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InvokeScreenAutomationWithContext indicates an expected call of InvokeScreenAutomationWithContext
func (mr *MockHoneycodeAPIMockRecorder) InvokeScreenAutomationWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeScreenAutomationWithContext", reflect.TypeOf((*MockHoneycodeAPI)(nil).InvokeScreenAutomationWithContext), varargs...)
}