// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/arczonalshift/arczonalshiftiface (interfaces: ARCZonalShiftAPI)

// Package arczonalshiftmock is a generated GoMock package.
package arczonalshiftmock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	arczonalshift "github.com/aws/aws-sdk-go/service/arczonalshift"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockARCZonalShiftAPI is a mock of ARCZonalShiftAPI interface
type MockARCZonalShiftAPI struct {
	ctrl     *gomock.Controller
	recorder *MockARCZonalShiftAPIMockRecorder
}

// MockARCZonalShiftAPIMockRecorder is the mock recorder for MockARCZonalShiftAPI
type MockARCZonalShiftAPIMockRecorder struct {
	mock *MockARCZonalShiftAPI
}

// NewMockARCZonalShiftAPI creates a new mock instance
func NewMockARCZonalShiftAPI(ctrl *gomock.Controller) *MockARCZonalShiftAPI {
	mock := &MockARCZonalShiftAPI{ctrl: ctrl}
	mock.recorder = &MockARCZonalShiftAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockARCZonalShiftAPI) EXPECT() *MockARCZonalShiftAPIMockRecorder {
	return m.recorder
}

// CancelZonalShift mocks base method
func (m *MockARCZonalShiftAPI) CancelZonalShift(arg0 *arczonalshift.CancelZonalShiftInput) (*arczonalshift.CancelZonalShiftOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelZonalShift", arg0)
	ret0, _ := ret[0].(*arczonalshift.CancelZonalShiftOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelZonalShift indicates an expected call of CancelZonalShift
func (mr *MockARCZonalShiftAPIMockRecorder) CancelZonalShift(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelZonalShift", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).CancelZonalShift), arg0)
}

// CancelZonalShiftRequest mocks base method
func (m *MockARCZonalShiftAPI) CancelZonalShiftRequest(arg0 *arczonalshift.CancelZonalShiftInput) (*request.Request, *arczonalshift.CancelZonalShiftOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelZonalShiftRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*arczonalshift.CancelZonalShiftOutput)
	return ret0, ret1
}

// CancelZonalShiftRequest indicates an expected call of CancelZonalShiftRequest
func (mr *MockARCZonalShiftAPIMockRecorder) CancelZonalShiftRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelZonalShiftRequest", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).CancelZonalShiftRequest), arg0)
}

// CancelZonalShiftWithContext mocks base method
func (m *MockARCZonalShiftAPI) CancelZonalShiftWithContext(arg0 context.Context, arg1 *arczonalshift.CancelZonalShiftInput, arg2 ...request.Option) (*arczonalshift.CancelZonalShiftOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CancelZonalShiftWithContext", varargs...)
	ret0, _ := ret[0].(*arczonalshift.CancelZonalShiftOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelZonalShiftWithContext indicates an expected call of CancelZonalShiftWithContext
func (mr *MockARCZonalShiftAPIMockRecorder) CancelZonalShiftWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelZonalShiftWithContext", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).CancelZonalShiftWithContext), varargs...)
}

// CreatePracticeRunConfiguration mocks base method
func (m *MockARCZonalShiftAPI) CreatePracticeRunConfiguration(arg0 *arczonalshift.CreatePracticeRunConfigurationInput) (*arczonalshift.CreatePracticeRunConfigurationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePracticeRunConfiguration", arg0)
	ret0, _ := ret[0].(*arczonalshift.CreatePracticeRunConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePracticeRunConfiguration indicates an expected call of CreatePracticeRunConfiguration
func (mr *MockARCZonalShiftAPIMockRecorder) CreatePracticeRunConfiguration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePracticeRunConfiguration", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).CreatePracticeRunConfiguration), arg0)
}

// CreatePracticeRunConfigurationRequest mocks base method
func (m *MockARCZonalShiftAPI) CreatePracticeRunConfigurationRequest(arg0 *arczonalshift.CreatePracticeRunConfigurationInput) (*request.Request, *arczonalshift.CreatePracticeRunConfigurationOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePracticeRunConfigurationRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*arczonalshift.CreatePracticeRunConfigurationOutput)
	return ret0, ret1
}

// CreatePracticeRunConfigurationRequest indicates an expected call of CreatePracticeRunConfigurationRequest
func (mr *MockARCZonalShiftAPIMockRecorder) CreatePracticeRunConfigurationRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePracticeRunConfigurationRequest", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).CreatePracticeRunConfigurationRequest), arg0)
}

// CreatePracticeRunConfigurationWithContext mocks base method
func (m *MockARCZonalShiftAPI) CreatePracticeRunConfigurationWithContext(arg0 context.Context, arg1 *arczonalshift.CreatePracticeRunConfigurationInput, arg2 ...request.Option) (*arczonalshift.CreatePracticeRunConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreatePracticeRunConfigurationWithContext", varargs...)
	ret0, _ := ret[0].(*arczonalshift.CreatePracticeRunConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePracticeRunConfigurationWithContext indicates an expected call of CreatePracticeRunConfigurationWithContext
func (mr *MockARCZonalShiftAPIMockRecorder) CreatePracticeRunConfigurationWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePracticeRunConfigurationWithContext", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).CreatePracticeRunConfigurationWithContext), varargs...)
}

// DeletePracticeRunConfiguration mocks base method
func (m *MockARCZonalShiftAPI) DeletePracticeRunConfiguration(arg0 *arczonalshift.DeletePracticeRunConfigurationInput) (*arczonalshift.DeletePracticeRunConfigurationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePracticeRunConfiguration", arg0)
	ret0, _ := ret[0].(*arczonalshift.DeletePracticeRunConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePracticeRunConfiguration indicates an expected call of DeletePracticeRunConfiguration
func (mr *MockARCZonalShiftAPIMockRecorder) DeletePracticeRunConfiguration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePracticeRunConfiguration", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).DeletePracticeRunConfiguration), arg0)
}

// DeletePracticeRunConfigurationRequest mocks base method
func (m *MockARCZonalShiftAPI) DeletePracticeRunConfigurationRequest(arg0 *arczonalshift.DeletePracticeRunConfigurationInput) (*request.Request, *arczonalshift.DeletePracticeRunConfigurationOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePracticeRunConfigurationRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*arczonalshift.DeletePracticeRunConfigurationOutput)
	return ret0, ret1
}

// DeletePracticeRunConfigurationRequest indicates an expected call of DeletePracticeRunConfigurationRequest
func (mr *MockARCZonalShiftAPIMockRecorder) DeletePracticeRunConfigurationRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePracticeRunConfigurationRequest", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).DeletePracticeRunConfigurationRequest), arg0)
}

// DeletePracticeRunConfigurationWithContext mocks base method
func (m *MockARCZonalShiftAPI) DeletePracticeRunConfigurationWithContext(arg0 context.Context, arg1 *arczonalshift.DeletePracticeRunConfigurationInput, arg2 ...request.Option) (*arczonalshift.DeletePracticeRunConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeletePracticeRunConfigurationWithContext", varargs...)
	ret0, _ := ret[0].(*arczonalshift.DeletePracticeRunConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePracticeRunConfigurationWithContext indicates an expected call of DeletePracticeRunConfigurationWithContext
func (mr *MockARCZonalShiftAPIMockRecorder) DeletePracticeRunConfigurationWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePracticeRunConfigurationWithContext", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).DeletePracticeRunConfigurationWithContext), varargs...)
}

// GetManagedResource mocks base method
func (m *MockARCZonalShiftAPI) GetManagedResource(arg0 *arczonalshift.GetManagedResourceInput) (*arczonalshift.GetManagedResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManagedResource", arg0)
	ret0, _ := ret[0].(*arczonalshift.GetManagedResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetManagedResource indicates an expected call of GetManagedResource
func (mr *MockARCZonalShiftAPIMockRecorder) GetManagedResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManagedResource", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).GetManagedResource), arg0)
}

// GetManagedResourceRequest mocks base method
func (m *MockARCZonalShiftAPI) GetManagedResourceRequest(arg0 *arczonalshift.GetManagedResourceInput) (*request.Request, *arczonalshift.GetManagedResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManagedResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*arczonalshift.GetManagedResourceOutput)
	return ret0, ret1
}

// GetManagedResourceRequest indicates an expected call of GetManagedResourceRequest
func (mr *MockARCZonalShiftAPIMockRecorder) GetManagedResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManagedResourceRequest", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).GetManagedResourceRequest), arg0)
}

// GetManagedResourceWithContext mocks base method
func (m *MockARCZonalShiftAPI) GetManagedResourceWithContext(arg0 context.Context, arg1 *arczonalshift.GetManagedResourceInput, arg2 ...request.Option) (*arczonalshift.GetManagedResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetManagedResourceWithContext", varargs...)
	ret0, _ := ret[0].(*arczonalshift.GetManagedResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetManagedResourceWithContext indicates an expected call of GetManagedResourceWithContext
func (mr *MockARCZonalShiftAPIMockRecorder) GetManagedResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManagedResourceWithContext", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).GetManagedResourceWithContext), varargs...)
}

// ListAutoshifts mocks base method
func (m *MockARCZonalShiftAPI) ListAutoshifts(arg0 *arczonalshift.ListAutoshiftsInput) (*arczonalshift.ListAutoshiftsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAutoshifts", arg0)
	ret0, _ := ret[0].(*arczonalshift.ListAutoshiftsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAutoshifts indicates an expected call of ListAutoshifts
func (mr *MockARCZonalShiftAPIMockRecorder) ListAutoshifts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAutoshifts", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).ListAutoshifts), arg0)
}

// ListAutoshiftsPages mocks base method
func (m *MockARCZonalShiftAPI) ListAutoshiftsPages(arg0 *arczonalshift.ListAutoshiftsInput, arg1 func(*arczonalshift.ListAutoshiftsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAutoshiftsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListAutoshiftsPages indicates an expected call of ListAutoshiftsPages
func (mr *MockARCZonalShiftAPIMockRecorder) ListAutoshiftsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAutoshiftsPages", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).ListAutoshiftsPages), arg0, arg1)
}

// ListAutoshiftsPagesWithContext mocks base method
func (m *MockARCZonalShiftAPI) ListAutoshiftsPagesWithContext(arg0 context.Context, arg1 *arczonalshift.ListAutoshiftsInput, arg2 func(*arczonalshift.ListAutoshiftsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAutoshiftsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListAutoshiftsPagesWithContext indicates an expected call of ListAutoshiftsPagesWithContext
func (mr *MockARCZonalShiftAPIMockRecorder) ListAutoshiftsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAutoshiftsPagesWithContext", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).ListAutoshiftsPagesWithContext), varargs...)
}

// ListAutoshiftsRequest mocks base method
func (m *MockARCZonalShiftAPI) ListAutoshiftsRequest(arg0 *arczonalshift.ListAutoshiftsInput) (*request.Request, *arczonalshift.ListAutoshiftsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAutoshiftsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*arczonalshift.ListAutoshiftsOutput)
	return ret0, ret1
}

// ListAutoshiftsRequest indicates an expected call of ListAutoshiftsRequest
func (mr *MockARCZonalShiftAPIMockRecorder) ListAutoshiftsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAutoshiftsRequest", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).ListAutoshiftsRequest), arg0)
}

// ListAutoshiftsWithContext mocks base method
func (m *MockARCZonalShiftAPI) ListAutoshiftsWithContext(arg0 context.Context, arg1 *arczonalshift.ListAutoshiftsInput, arg2 ...request.Option) (*arczonalshift.ListAutoshiftsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAutoshiftsWithContext", varargs...)
	ret0, _ := ret[0].(*arczonalshift.ListAutoshiftsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAutoshiftsWithContext indicates an expected call of ListAutoshiftsWithContext
func (mr *MockARCZonalShiftAPIMockRecorder) ListAutoshiftsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAutoshiftsWithContext", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).ListAutoshiftsWithContext), varargs...)
}

// ListManagedResources mocks base method
func (m *MockARCZonalShiftAPI) ListManagedResources(arg0 *arczonalshift.ListManagedResourcesInput) (*arczonalshift.ListManagedResourcesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListManagedResources", arg0)
	ret0, _ := ret[0].(*arczonalshift.ListManagedResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListManagedResources indicates an expected call of ListManagedResources
func (mr *MockARCZonalShiftAPIMockRecorder) ListManagedResources(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListManagedResources", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).ListManagedResources), arg0)
}

// ListManagedResourcesPages mocks base method
func (m *MockARCZonalShiftAPI) ListManagedResourcesPages(arg0 *arczonalshift.ListManagedResourcesInput, arg1 func(*arczonalshift.ListManagedResourcesOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListManagedResourcesPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListManagedResourcesPages indicates an expected call of ListManagedResourcesPages
func (mr *MockARCZonalShiftAPIMockRecorder) ListManagedResourcesPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListManagedResourcesPages", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).ListManagedResourcesPages), arg0, arg1)
}

// ListManagedResourcesPagesWithContext mocks base method
func (m *MockARCZonalShiftAPI) ListManagedResourcesPagesWithContext(arg0 context.Context, arg1 *arczonalshift.ListManagedResourcesInput, arg2 func(*arczonalshift.ListManagedResourcesOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListManagedResourcesPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListManagedResourcesPagesWithContext indicates an expected call of ListManagedResourcesPagesWithContext
func (mr *MockARCZonalShiftAPIMockRecorder) ListManagedResourcesPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListManagedResourcesPagesWithContext", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).ListManagedResourcesPagesWithContext), varargs...)
}

// ListManagedResourcesRequest mocks base method
func (m *MockARCZonalShiftAPI) ListManagedResourcesRequest(arg0 *arczonalshift.ListManagedResourcesInput) (*request.Request, *arczonalshift.ListManagedResourcesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListManagedResourcesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*arczonalshift.ListManagedResourcesOutput)
	return ret0, ret1
}

// ListManagedResourcesRequest indicates an expected call of ListManagedResourcesRequest
func (mr *MockARCZonalShiftAPIMockRecorder) ListManagedResourcesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListManagedResourcesRequest", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).ListManagedResourcesRequest), arg0)
}

// ListManagedResourcesWithContext mocks base method
func (m *MockARCZonalShiftAPI) ListManagedResourcesWithContext(arg0 context.Context, arg1 *arczonalshift.ListManagedResourcesInput, arg2 ...request.Option) (*arczonalshift.ListManagedResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListManagedResourcesWithContext", varargs...)
	ret0, _ := ret[0].(*arczonalshift.ListManagedResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListManagedResourcesWithContext indicates an expected call of ListManagedResourcesWithContext
func (mr *MockARCZonalShiftAPIMockRecorder) ListManagedResourcesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListManagedResourcesWithContext", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).ListManagedResourcesWithContext), varargs...)
}

// ListZonalShifts mocks base method
func (m *MockARCZonalShiftAPI) ListZonalShifts(arg0 *arczonalshift.ListZonalShiftsInput) (*arczonalshift.ListZonalShiftsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListZonalShifts", arg0)
	ret0, _ := ret[0].(*arczonalshift.ListZonalShiftsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListZonalShifts indicates an expected call of ListZonalShifts
func (mr *MockARCZonalShiftAPIMockRecorder) ListZonalShifts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListZonalShifts", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).ListZonalShifts), arg0)
}

// ListZonalShiftsPages mocks base method
func (m *MockARCZonalShiftAPI) ListZonalShiftsPages(arg0 *arczonalshift.ListZonalShiftsInput, arg1 func(*arczonalshift.ListZonalShiftsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListZonalShiftsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListZonalShiftsPages indicates an expected call of ListZonalShiftsPages
func (mr *MockARCZonalShiftAPIMockRecorder) ListZonalShiftsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListZonalShiftsPages", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).ListZonalShiftsPages), arg0, arg1)
}

// ListZonalShiftsPagesWithContext mocks base method
func (m *MockARCZonalShiftAPI) ListZonalShiftsPagesWithContext(arg0 context.Context, arg1 *arczonalshift.ListZonalShiftsInput, arg2 func(*arczonalshift.ListZonalShiftsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListZonalShiftsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListZonalShiftsPagesWithContext indicates an expected call of ListZonalShiftsPagesWithContext
func (mr *MockARCZonalShiftAPIMockRecorder) ListZonalShiftsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListZonalShiftsPagesWithContext", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).ListZonalShiftsPagesWithContext), varargs...)
}

// ListZonalShiftsRequest mocks base method
func (m *MockARCZonalShiftAPI) ListZonalShiftsRequest(arg0 *arczonalshift.ListZonalShiftsInput) (*request.Request, *arczonalshift.ListZonalShiftsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListZonalShiftsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*arczonalshift.ListZonalShiftsOutput)
	return ret0, ret1
}

// ListZonalShiftsRequest indicates an expected call of ListZonalShiftsRequest
func (mr *MockARCZonalShiftAPIMockRecorder) ListZonalShiftsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListZonalShiftsRequest", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).ListZonalShiftsRequest), arg0)
}

// ListZonalShiftsWithContext mocks base method
func (m *MockARCZonalShiftAPI) ListZonalShiftsWithContext(arg0 context.Context, arg1 *arczonalshift.ListZonalShiftsInput, arg2 ...request.Option) (*arczonalshift.ListZonalShiftsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListZonalShiftsWithContext", varargs...)
	ret0, _ := ret[0].(*arczonalshift.ListZonalShiftsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListZonalShiftsWithContext indicates an expected call of ListZonalShiftsWithContext
func (mr *MockARCZonalShiftAPIMockRecorder) ListZonalShiftsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListZonalShiftsWithContext", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).ListZonalShiftsWithContext), varargs...)
}

// StartZonalShift mocks base method
func (m *MockARCZonalShiftAPI) StartZonalShift(arg0 *arczonalshift.StartZonalShiftInput) (*arczonalshift.StartZonalShiftOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartZonalShift", arg0)
	ret0, _ := ret[0].(*arczonalshift.StartZonalShiftOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartZonalShift indicates an expected call of StartZonalShift
func (mr *MockARCZonalShiftAPIMockRecorder) StartZonalShift(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartZonalShift", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).StartZonalShift), arg0)
}

// StartZonalShiftRequest mocks base method
func (m *MockARCZonalShiftAPI) StartZonalShiftRequest(arg0 *arczonalshift.StartZonalShiftInput) (*request.Request, *arczonalshift.StartZonalShiftOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartZonalShiftRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*arczonalshift.StartZonalShiftOutput)
	return ret0, ret1
}

// StartZonalShiftRequest indicates an expected call of StartZonalShiftRequest
func (mr *MockARCZonalShiftAPIMockRecorder) StartZonalShiftRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartZonalShiftRequest", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).StartZonalShiftRequest), arg0)
}

// StartZonalShiftWithContext mocks base method
func (m *MockARCZonalShiftAPI) StartZonalShiftWithContext(arg0 context.Context, arg1 *arczonalshift.StartZonalShiftInput, arg2 ...request.Option) (*arczonalshift.StartZonalShiftOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StartZonalShiftWithContext", varargs...)
	ret0, _ := ret[0].(*arczonalshift.StartZonalShiftOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartZonalShiftWithContext indicates an expected call of StartZonalShiftWithContext
func (mr *MockARCZonalShiftAPIMockRecorder) StartZonalShiftWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartZonalShiftWithContext", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).StartZonalShiftWithContext), varargs...)
}

// UpdatePracticeRunConfiguration mocks base method
func (m *MockARCZonalShiftAPI) UpdatePracticeRunConfiguration(arg0 *arczonalshift.UpdatePracticeRunConfigurationInput) (*arczonalshift.UpdatePracticeRunConfigurationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePracticeRunConfiguration", arg0)
	ret0, _ := ret[0].(*arczonalshift.UpdatePracticeRunConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePracticeRunConfiguration indicates an expected call of UpdatePracticeRunConfiguration
func (mr *MockARCZonalShiftAPIMockRecorder) UpdatePracticeRunConfiguration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePracticeRunConfiguration", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).UpdatePracticeRunConfiguration), arg0)
}

// UpdatePracticeRunConfigurationRequest mocks base method
func (m *MockARCZonalShiftAPI) UpdatePracticeRunConfigurationRequest(arg0 *arczonalshift.UpdatePracticeRunConfigurationInput) (*request.Request, *arczonalshift.UpdatePracticeRunConfigurationOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePracticeRunConfigurationRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*arczonalshift.UpdatePracticeRunConfigurationOutput)
	return ret0, ret1
}

// UpdatePracticeRunConfigurationRequest indicates an expected call of UpdatePracticeRunConfigurationRequest
func (mr *MockARCZonalShiftAPIMockRecorder) UpdatePracticeRunConfigurationRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePracticeRunConfigurationRequest", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).UpdatePracticeRunConfigurationRequest), arg0)
}

// UpdatePracticeRunConfigurationWithContext mocks base method
func (m *MockARCZonalShiftAPI) UpdatePracticeRunConfigurationWithContext(arg0 context.Context, arg1 *arczonalshift.UpdatePracticeRunConfigurationInput, arg2 ...request.Option) (*arczonalshift.UpdatePracticeRunConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdatePracticeRunConfigurationWithContext", varargs...)
	ret0, _ := ret[0].(*arczonalshift.UpdatePracticeRunConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePracticeRunConfigurationWithContext indicates an expected call of UpdatePracticeRunConfigurationWithContext
func (mr *MockARCZonalShiftAPIMockRecorder) UpdatePracticeRunConfigurationWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePracticeRunConfigurationWithContext", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).UpdatePracticeRunConfigurationWithContext), varargs...)
}

// UpdateZonalAutoshiftConfiguration mocks base method
func (m *MockARCZonalShiftAPI) UpdateZonalAutoshiftConfiguration(arg0 *arczonalshift.UpdateZonalAutoshiftConfigurationInput) (*arczonalshift.UpdateZonalAutoshiftConfigurationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateZonalAutoshiftConfiguration", arg0)
	ret0, _ := ret[0].(*arczonalshift.UpdateZonalAutoshiftConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateZonalAutoshiftConfiguration indicates an expected call of UpdateZonalAutoshiftConfiguration
func (mr *MockARCZonalShiftAPIMockRecorder) UpdateZonalAutoshiftConfiguration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateZonalAutoshiftConfiguration", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).UpdateZonalAutoshiftConfiguration), arg0)
}

// UpdateZonalAutoshiftConfigurationRequest mocks base method
func (m *MockARCZonalShiftAPI) UpdateZonalAutoshiftConfigurationRequest(arg0 *arczonalshift.UpdateZonalAutoshiftConfigurationInput) (*request.Request, *arczonalshift.UpdateZonalAutoshiftConfigurationOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateZonalAutoshiftConfigurationRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*arczonalshift.UpdateZonalAutoshiftConfigurationOutput)
	return ret0, ret1
}

// UpdateZonalAutoshiftConfigurationRequest indicates an expected call of UpdateZonalAutoshiftConfigurationRequest
func (mr *MockARCZonalShiftAPIMockRecorder) UpdateZonalAutoshiftConfigurationRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateZonalAutoshiftConfigurationRequest", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).UpdateZonalAutoshiftConfigurationRequest), arg0)
}

// UpdateZonalAutoshiftConfigurationWithContext mocks base method
func (m *MockARCZonalShiftAPI) UpdateZonalAutoshiftConfigurationWithContext(arg0 context.Context, arg1 *arczonalshift.UpdateZonalAutoshiftConfigurationInput, arg2 ...request.Option) (*arczonalshift.UpdateZonalAutoshiftConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateZonalAutoshiftConfigurationWithContext", varargs...)
	ret0, _ := ret[0].(*arczonalshift.UpdateZonalAutoshiftConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateZonalAutoshiftConfigurationWithContext indicates an expected call of UpdateZonalAutoshiftConfigurationWithContext
func (mr *MockARCZonalShiftAPIMockRecorder) UpdateZonalAutoshiftConfigurationWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateZonalAutoshiftConfigurationWithContext", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).UpdateZonalAutoshiftConfigurationWithContext), varargs...)
}

// UpdateZonalShift mocks base method
func (m *MockARCZonalShiftAPI) UpdateZonalShift(arg0 *arczonalshift.UpdateZonalShiftInput) (*arczonalshift.UpdateZonalShiftOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateZonalShift", arg0)
	ret0, _ := ret[0].(*arczonalshift.UpdateZonalShiftOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateZonalShift indicates an expected call of UpdateZonalShift
func (mr *MockARCZonalShiftAPIMockRecorder) UpdateZonalShift(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateZonalShift", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).UpdateZonalShift), arg0)
}

// UpdateZonalShiftRequest mocks base method
func (m *MockARCZonalShiftAPI) UpdateZonalShiftRequest(arg0 *arczonalshift.UpdateZonalShiftInput) (*request.Request, *arczonalshift.UpdateZonalShiftOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateZonalShiftRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*arczonalshift.UpdateZonalShiftOutput)
	return ret0, ret1
}

// UpdateZonalShiftRequest indicates an expected call of UpdateZonalShiftRequest
func (mr *MockARCZonalShiftAPIMockRecorder) UpdateZonalShiftRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateZonalShiftRequest", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).UpdateZonalShiftRequest), arg0)
}

// UpdateZonalShiftWithContext mocks base method
func (m *MockARCZonalShiftAPI) UpdateZonalShiftWithContext(arg0 context.Context, arg1 *arczonalshift.UpdateZonalShiftInput, arg2 ...request.Option) (*arczonalshift.UpdateZonalShiftOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateZonalShiftWithContext", varargs...)
	ret0, _ := ret[0].(*arczonalshift.UpdateZonalShiftOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateZonalShiftWithContext indicates an expected call of UpdateZonalShiftWithContext
func (mr *MockARCZonalShiftAPIMockRecorder) UpdateZonalShiftWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateZonalShiftWithContext", reflect.TypeOf((*MockARCZonalShiftAPI)(nil).UpdateZonalShiftWithContext), varargs...)
}