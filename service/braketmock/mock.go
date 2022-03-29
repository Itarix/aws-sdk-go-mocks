// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/braket/braketiface (interfaces: BraketAPI)

// Package braketmock is a generated GoMock package.
package braketmock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	braket "github.com/aws/aws-sdk-go/service/braket"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBraketAPI is a mock of BraketAPI interface
type MockBraketAPI struct {
	ctrl     *gomock.Controller
	recorder *MockBraketAPIMockRecorder
}

// MockBraketAPIMockRecorder is the mock recorder for MockBraketAPI
type MockBraketAPIMockRecorder struct {
	mock *MockBraketAPI
}

// NewMockBraketAPI creates a new mock instance
func NewMockBraketAPI(ctrl *gomock.Controller) *MockBraketAPI {
	mock := &MockBraketAPI{ctrl: ctrl}
	mock.recorder = &MockBraketAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBraketAPI) EXPECT() *MockBraketAPIMockRecorder {
	return m.recorder
}

// CancelQuantumTask mocks base method
func (m *MockBraketAPI) CancelQuantumTask(arg0 *braket.CancelQuantumTaskInput) (*braket.CancelQuantumTaskOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelQuantumTask", arg0)
	ret0, _ := ret[0].(*braket.CancelQuantumTaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelQuantumTask indicates an expected call of CancelQuantumTask
func (mr *MockBraketAPIMockRecorder) CancelQuantumTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelQuantumTask", reflect.TypeOf((*MockBraketAPI)(nil).CancelQuantumTask), arg0)
}

// CancelQuantumTaskRequest mocks base method
func (m *MockBraketAPI) CancelQuantumTaskRequest(arg0 *braket.CancelQuantumTaskInput) (*request.Request, *braket.CancelQuantumTaskOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelQuantumTaskRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*braket.CancelQuantumTaskOutput)
	return ret0, ret1
}

// CancelQuantumTaskRequest indicates an expected call of CancelQuantumTaskRequest
func (mr *MockBraketAPIMockRecorder) CancelQuantumTaskRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelQuantumTaskRequest", reflect.TypeOf((*MockBraketAPI)(nil).CancelQuantumTaskRequest), arg0)
}

// CancelQuantumTaskWithContext mocks base method
func (m *MockBraketAPI) CancelQuantumTaskWithContext(arg0 context.Context, arg1 *braket.CancelQuantumTaskInput, arg2 ...request.Option) (*braket.CancelQuantumTaskOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CancelQuantumTaskWithContext", varargs...)
	ret0, _ := ret[0].(*braket.CancelQuantumTaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelQuantumTaskWithContext indicates an expected call of CancelQuantumTaskWithContext
func (mr *MockBraketAPIMockRecorder) CancelQuantumTaskWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelQuantumTaskWithContext", reflect.TypeOf((*MockBraketAPI)(nil).CancelQuantumTaskWithContext), varargs...)
}

// CreateQuantumTask mocks base method
func (m *MockBraketAPI) CreateQuantumTask(arg0 *braket.CreateQuantumTaskInput) (*braket.CreateQuantumTaskOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateQuantumTask", arg0)
	ret0, _ := ret[0].(*braket.CreateQuantumTaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateQuantumTask indicates an expected call of CreateQuantumTask
func (mr *MockBraketAPIMockRecorder) CreateQuantumTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateQuantumTask", reflect.TypeOf((*MockBraketAPI)(nil).CreateQuantumTask), arg0)
}

// CreateQuantumTaskRequest mocks base method
func (m *MockBraketAPI) CreateQuantumTaskRequest(arg0 *braket.CreateQuantumTaskInput) (*request.Request, *braket.CreateQuantumTaskOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateQuantumTaskRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*braket.CreateQuantumTaskOutput)
	return ret0, ret1
}

// CreateQuantumTaskRequest indicates an expected call of CreateQuantumTaskRequest
func (mr *MockBraketAPIMockRecorder) CreateQuantumTaskRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateQuantumTaskRequest", reflect.TypeOf((*MockBraketAPI)(nil).CreateQuantumTaskRequest), arg0)
}

// CreateQuantumTaskWithContext mocks base method
func (m *MockBraketAPI) CreateQuantumTaskWithContext(arg0 context.Context, arg1 *braket.CreateQuantumTaskInput, arg2 ...request.Option) (*braket.CreateQuantumTaskOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateQuantumTaskWithContext", varargs...)
	ret0, _ := ret[0].(*braket.CreateQuantumTaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateQuantumTaskWithContext indicates an expected call of CreateQuantumTaskWithContext
func (mr *MockBraketAPIMockRecorder) CreateQuantumTaskWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateQuantumTaskWithContext", reflect.TypeOf((*MockBraketAPI)(nil).CreateQuantumTaskWithContext), varargs...)
}

// GetDevice mocks base method
func (m *MockBraketAPI) GetDevice(arg0 *braket.GetDeviceInput) (*braket.GetDeviceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevice", arg0)
	ret0, _ := ret[0].(*braket.GetDeviceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevice indicates an expected call of GetDevice
func (mr *MockBraketAPIMockRecorder) GetDevice(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevice", reflect.TypeOf((*MockBraketAPI)(nil).GetDevice), arg0)
}

// GetDeviceRequest mocks base method
func (m *MockBraketAPI) GetDeviceRequest(arg0 *braket.GetDeviceInput) (*request.Request, *braket.GetDeviceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*braket.GetDeviceOutput)
	return ret0, ret1
}

// GetDeviceRequest indicates an expected call of GetDeviceRequest
func (mr *MockBraketAPIMockRecorder) GetDeviceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceRequest", reflect.TypeOf((*MockBraketAPI)(nil).GetDeviceRequest), arg0)
}

// GetDeviceWithContext mocks base method
func (m *MockBraketAPI) GetDeviceWithContext(arg0 context.Context, arg1 *braket.GetDeviceInput, arg2 ...request.Option) (*braket.GetDeviceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDeviceWithContext", varargs...)
	ret0, _ := ret[0].(*braket.GetDeviceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceWithContext indicates an expected call of GetDeviceWithContext
func (mr *MockBraketAPIMockRecorder) GetDeviceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceWithContext", reflect.TypeOf((*MockBraketAPI)(nil).GetDeviceWithContext), varargs...)
}

// GetQuantumTask mocks base method
func (m *MockBraketAPI) GetQuantumTask(arg0 *braket.GetQuantumTaskInput) (*braket.GetQuantumTaskOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuantumTask", arg0)
	ret0, _ := ret[0].(*braket.GetQuantumTaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQuantumTask indicates an expected call of GetQuantumTask
func (mr *MockBraketAPIMockRecorder) GetQuantumTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuantumTask", reflect.TypeOf((*MockBraketAPI)(nil).GetQuantumTask), arg0)
}

// GetQuantumTaskRequest mocks base method
func (m *MockBraketAPI) GetQuantumTaskRequest(arg0 *braket.GetQuantumTaskInput) (*request.Request, *braket.GetQuantumTaskOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuantumTaskRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*braket.GetQuantumTaskOutput)
	return ret0, ret1
}

// GetQuantumTaskRequest indicates an expected call of GetQuantumTaskRequest
func (mr *MockBraketAPIMockRecorder) GetQuantumTaskRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuantumTaskRequest", reflect.TypeOf((*MockBraketAPI)(nil).GetQuantumTaskRequest), arg0)
}

// GetQuantumTaskWithContext mocks base method
func (m *MockBraketAPI) GetQuantumTaskWithContext(arg0 context.Context, arg1 *braket.GetQuantumTaskInput, arg2 ...request.Option) (*braket.GetQuantumTaskOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetQuantumTaskWithContext", varargs...)
	ret0, _ := ret[0].(*braket.GetQuantumTaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQuantumTaskWithContext indicates an expected call of GetQuantumTaskWithContext
func (mr *MockBraketAPIMockRecorder) GetQuantumTaskWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuantumTaskWithContext", reflect.TypeOf((*MockBraketAPI)(nil).GetQuantumTaskWithContext), varargs...)
}

// SearchDevices mocks base method
func (m *MockBraketAPI) SearchDevices(arg0 *braket.SearchDevicesInput) (*braket.SearchDevicesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchDevices", arg0)
	ret0, _ := ret[0].(*braket.SearchDevicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchDevices indicates an expected call of SearchDevices
func (mr *MockBraketAPIMockRecorder) SearchDevices(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchDevices", reflect.TypeOf((*MockBraketAPI)(nil).SearchDevices), arg0)
}

// SearchDevicesPages mocks base method
func (m *MockBraketAPI) SearchDevicesPages(arg0 *braket.SearchDevicesInput, arg1 func(*braket.SearchDevicesOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchDevicesPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SearchDevicesPages indicates an expected call of SearchDevicesPages
func (mr *MockBraketAPIMockRecorder) SearchDevicesPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchDevicesPages", reflect.TypeOf((*MockBraketAPI)(nil).SearchDevicesPages), arg0, arg1)
}

// SearchDevicesPagesWithContext mocks base method
func (m *MockBraketAPI) SearchDevicesPagesWithContext(arg0 context.Context, arg1 *braket.SearchDevicesInput, arg2 func(*braket.SearchDevicesOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchDevicesPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SearchDevicesPagesWithContext indicates an expected call of SearchDevicesPagesWithContext
func (mr *MockBraketAPIMockRecorder) SearchDevicesPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchDevicesPagesWithContext", reflect.TypeOf((*MockBraketAPI)(nil).SearchDevicesPagesWithContext), varargs...)
}

// SearchDevicesRequest mocks base method
func (m *MockBraketAPI) SearchDevicesRequest(arg0 *braket.SearchDevicesInput) (*request.Request, *braket.SearchDevicesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchDevicesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*braket.SearchDevicesOutput)
	return ret0, ret1
}

// SearchDevicesRequest indicates an expected call of SearchDevicesRequest
func (mr *MockBraketAPIMockRecorder) SearchDevicesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchDevicesRequest", reflect.TypeOf((*MockBraketAPI)(nil).SearchDevicesRequest), arg0)
}

// SearchDevicesWithContext mocks base method
func (m *MockBraketAPI) SearchDevicesWithContext(arg0 context.Context, arg1 *braket.SearchDevicesInput, arg2 ...request.Option) (*braket.SearchDevicesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchDevicesWithContext", varargs...)
	ret0, _ := ret[0].(*braket.SearchDevicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchDevicesWithContext indicates an expected call of SearchDevicesWithContext
func (mr *MockBraketAPIMockRecorder) SearchDevicesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchDevicesWithContext", reflect.TypeOf((*MockBraketAPI)(nil).SearchDevicesWithContext), varargs...)
}

// SearchQuantumTasks mocks base method
func (m *MockBraketAPI) SearchQuantumTasks(arg0 *braket.SearchQuantumTasksInput) (*braket.SearchQuantumTasksOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchQuantumTasks", arg0)
	ret0, _ := ret[0].(*braket.SearchQuantumTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchQuantumTasks indicates an expected call of SearchQuantumTasks
func (mr *MockBraketAPIMockRecorder) SearchQuantumTasks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchQuantumTasks", reflect.TypeOf((*MockBraketAPI)(nil).SearchQuantumTasks), arg0)
}

// SearchQuantumTasksPages mocks base method
func (m *MockBraketAPI) SearchQuantumTasksPages(arg0 *braket.SearchQuantumTasksInput, arg1 func(*braket.SearchQuantumTasksOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchQuantumTasksPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SearchQuantumTasksPages indicates an expected call of SearchQuantumTasksPages
func (mr *MockBraketAPIMockRecorder) SearchQuantumTasksPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchQuantumTasksPages", reflect.TypeOf((*MockBraketAPI)(nil).SearchQuantumTasksPages), arg0, arg1)
}

// SearchQuantumTasksPagesWithContext mocks base method
func (m *MockBraketAPI) SearchQuantumTasksPagesWithContext(arg0 context.Context, arg1 *braket.SearchQuantumTasksInput, arg2 func(*braket.SearchQuantumTasksOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchQuantumTasksPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SearchQuantumTasksPagesWithContext indicates an expected call of SearchQuantumTasksPagesWithContext
func (mr *MockBraketAPIMockRecorder) SearchQuantumTasksPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchQuantumTasksPagesWithContext", reflect.TypeOf((*MockBraketAPI)(nil).SearchQuantumTasksPagesWithContext), varargs...)
}

// SearchQuantumTasksRequest mocks base method
func (m *MockBraketAPI) SearchQuantumTasksRequest(arg0 *braket.SearchQuantumTasksInput) (*request.Request, *braket.SearchQuantumTasksOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchQuantumTasksRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*braket.SearchQuantumTasksOutput)
	return ret0, ret1
}

// SearchQuantumTasksRequest indicates an expected call of SearchQuantumTasksRequest
func (mr *MockBraketAPIMockRecorder) SearchQuantumTasksRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchQuantumTasksRequest", reflect.TypeOf((*MockBraketAPI)(nil).SearchQuantumTasksRequest), arg0)
}

// SearchQuantumTasksWithContext mocks base method
func (m *MockBraketAPI) SearchQuantumTasksWithContext(arg0 context.Context, arg1 *braket.SearchQuantumTasksInput, arg2 ...request.Option) (*braket.SearchQuantumTasksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchQuantumTasksWithContext", varargs...)
	ret0, _ := ret[0].(*braket.SearchQuantumTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchQuantumTasksWithContext indicates an expected call of SearchQuantumTasksWithContext
func (mr *MockBraketAPIMockRecorder) SearchQuantumTasksWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchQuantumTasksWithContext", reflect.TypeOf((*MockBraketAPI)(nil).SearchQuantumTasksWithContext), varargs...)
}