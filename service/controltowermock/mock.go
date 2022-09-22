// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/controltower/controltoweriface (interfaces: ControlTowerAPI)

// Package controltowermock is a generated GoMock package.
package controltowermock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	controltower "github.com/aws/aws-sdk-go/service/controltower"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockControlTowerAPI is a mock of ControlTowerAPI interface
type MockControlTowerAPI struct {
	ctrl     *gomock.Controller
	recorder *MockControlTowerAPIMockRecorder
}

// MockControlTowerAPIMockRecorder is the mock recorder for MockControlTowerAPI
type MockControlTowerAPIMockRecorder struct {
	mock *MockControlTowerAPI
}

// NewMockControlTowerAPI creates a new mock instance
func NewMockControlTowerAPI(ctrl *gomock.Controller) *MockControlTowerAPI {
	mock := &MockControlTowerAPI{ctrl: ctrl}
	mock.recorder = &MockControlTowerAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockControlTowerAPI) EXPECT() *MockControlTowerAPIMockRecorder {
	return m.recorder
}

// DisableControl mocks base method
func (m *MockControlTowerAPI) DisableControl(arg0 *controltower.DisableControlInput) (*controltower.DisableControlOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableControl", arg0)
	ret0, _ := ret[0].(*controltower.DisableControlOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableControl indicates an expected call of DisableControl
func (mr *MockControlTowerAPIMockRecorder) DisableControl(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableControl", reflect.TypeOf((*MockControlTowerAPI)(nil).DisableControl), arg0)
}

// DisableControlRequest mocks base method
func (m *MockControlTowerAPI) DisableControlRequest(arg0 *controltower.DisableControlInput) (*request.Request, *controltower.DisableControlOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableControlRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*controltower.DisableControlOutput)
	return ret0, ret1
}

// DisableControlRequest indicates an expected call of DisableControlRequest
func (mr *MockControlTowerAPIMockRecorder) DisableControlRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableControlRequest", reflect.TypeOf((*MockControlTowerAPI)(nil).DisableControlRequest), arg0)
}

// DisableControlWithContext mocks base method
func (m *MockControlTowerAPI) DisableControlWithContext(arg0 context.Context, arg1 *controltower.DisableControlInput, arg2 ...request.Option) (*controltower.DisableControlOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DisableControlWithContext", varargs...)
	ret0, _ := ret[0].(*controltower.DisableControlOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableControlWithContext indicates an expected call of DisableControlWithContext
func (mr *MockControlTowerAPIMockRecorder) DisableControlWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableControlWithContext", reflect.TypeOf((*MockControlTowerAPI)(nil).DisableControlWithContext), varargs...)
}

// EnableControl mocks base method
func (m *MockControlTowerAPI) EnableControl(arg0 *controltower.EnableControlInput) (*controltower.EnableControlOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableControl", arg0)
	ret0, _ := ret[0].(*controltower.EnableControlOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableControl indicates an expected call of EnableControl
func (mr *MockControlTowerAPIMockRecorder) EnableControl(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableControl", reflect.TypeOf((*MockControlTowerAPI)(nil).EnableControl), arg0)
}

// EnableControlRequest mocks base method
func (m *MockControlTowerAPI) EnableControlRequest(arg0 *controltower.EnableControlInput) (*request.Request, *controltower.EnableControlOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableControlRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*controltower.EnableControlOutput)
	return ret0, ret1
}

// EnableControlRequest indicates an expected call of EnableControlRequest
func (mr *MockControlTowerAPIMockRecorder) EnableControlRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableControlRequest", reflect.TypeOf((*MockControlTowerAPI)(nil).EnableControlRequest), arg0)
}

// EnableControlWithContext mocks base method
func (m *MockControlTowerAPI) EnableControlWithContext(arg0 context.Context, arg1 *controltower.EnableControlInput, arg2 ...request.Option) (*controltower.EnableControlOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnableControlWithContext", varargs...)
	ret0, _ := ret[0].(*controltower.EnableControlOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableControlWithContext indicates an expected call of EnableControlWithContext
func (mr *MockControlTowerAPIMockRecorder) EnableControlWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableControlWithContext", reflect.TypeOf((*MockControlTowerAPI)(nil).EnableControlWithContext), varargs...)
}

// GetControlOperation mocks base method
func (m *MockControlTowerAPI) GetControlOperation(arg0 *controltower.GetControlOperationInput) (*controltower.GetControlOperationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetControlOperation", arg0)
	ret0, _ := ret[0].(*controltower.GetControlOperationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetControlOperation indicates an expected call of GetControlOperation
func (mr *MockControlTowerAPIMockRecorder) GetControlOperation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetControlOperation", reflect.TypeOf((*MockControlTowerAPI)(nil).GetControlOperation), arg0)
}

// GetControlOperationRequest mocks base method
func (m *MockControlTowerAPI) GetControlOperationRequest(arg0 *controltower.GetControlOperationInput) (*request.Request, *controltower.GetControlOperationOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetControlOperationRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*controltower.GetControlOperationOutput)
	return ret0, ret1
}

// GetControlOperationRequest indicates an expected call of GetControlOperationRequest
func (mr *MockControlTowerAPIMockRecorder) GetControlOperationRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetControlOperationRequest", reflect.TypeOf((*MockControlTowerAPI)(nil).GetControlOperationRequest), arg0)
}

// GetControlOperationWithContext mocks base method
func (m *MockControlTowerAPI) GetControlOperationWithContext(arg0 context.Context, arg1 *controltower.GetControlOperationInput, arg2 ...request.Option) (*controltower.GetControlOperationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetControlOperationWithContext", varargs...)
	ret0, _ := ret[0].(*controltower.GetControlOperationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetControlOperationWithContext indicates an expected call of GetControlOperationWithContext
func (mr *MockControlTowerAPIMockRecorder) GetControlOperationWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetControlOperationWithContext", reflect.TypeOf((*MockControlTowerAPI)(nil).GetControlOperationWithContext), varargs...)
}

// ListEnabledControls mocks base method
func (m *MockControlTowerAPI) ListEnabledControls(arg0 *controltower.ListEnabledControlsInput) (*controltower.ListEnabledControlsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEnabledControls", arg0)
	ret0, _ := ret[0].(*controltower.ListEnabledControlsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEnabledControls indicates an expected call of ListEnabledControls
func (mr *MockControlTowerAPIMockRecorder) ListEnabledControls(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEnabledControls", reflect.TypeOf((*MockControlTowerAPI)(nil).ListEnabledControls), arg0)
}

// ListEnabledControlsPages mocks base method
func (m *MockControlTowerAPI) ListEnabledControlsPages(arg0 *controltower.ListEnabledControlsInput, arg1 func(*controltower.ListEnabledControlsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEnabledControlsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListEnabledControlsPages indicates an expected call of ListEnabledControlsPages
func (mr *MockControlTowerAPIMockRecorder) ListEnabledControlsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEnabledControlsPages", reflect.TypeOf((*MockControlTowerAPI)(nil).ListEnabledControlsPages), arg0, arg1)
}

// ListEnabledControlsPagesWithContext mocks base method
func (m *MockControlTowerAPI) ListEnabledControlsPagesWithContext(arg0 context.Context, arg1 *controltower.ListEnabledControlsInput, arg2 func(*controltower.ListEnabledControlsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEnabledControlsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListEnabledControlsPagesWithContext indicates an expected call of ListEnabledControlsPagesWithContext
func (mr *MockControlTowerAPIMockRecorder) ListEnabledControlsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEnabledControlsPagesWithContext", reflect.TypeOf((*MockControlTowerAPI)(nil).ListEnabledControlsPagesWithContext), varargs...)
}

// ListEnabledControlsRequest mocks base method
func (m *MockControlTowerAPI) ListEnabledControlsRequest(arg0 *controltower.ListEnabledControlsInput) (*request.Request, *controltower.ListEnabledControlsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEnabledControlsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*controltower.ListEnabledControlsOutput)
	return ret0, ret1
}

// ListEnabledControlsRequest indicates an expected call of ListEnabledControlsRequest
func (mr *MockControlTowerAPIMockRecorder) ListEnabledControlsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEnabledControlsRequest", reflect.TypeOf((*MockControlTowerAPI)(nil).ListEnabledControlsRequest), arg0)
}

// ListEnabledControlsWithContext mocks base method
func (m *MockControlTowerAPI) ListEnabledControlsWithContext(arg0 context.Context, arg1 *controltower.ListEnabledControlsInput, arg2 ...request.Option) (*controltower.ListEnabledControlsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEnabledControlsWithContext", varargs...)
	ret0, _ := ret[0].(*controltower.ListEnabledControlsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEnabledControlsWithContext indicates an expected call of ListEnabledControlsWithContext
func (mr *MockControlTowerAPIMockRecorder) ListEnabledControlsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEnabledControlsWithContext", reflect.TypeOf((*MockControlTowerAPI)(nil).ListEnabledControlsWithContext), varargs...)
}