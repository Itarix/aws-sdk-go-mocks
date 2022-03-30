// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/finspacedata/finspacedataiface (interfaces: FinSpaceDataAPI)

// Package finspacedatamock is a generated GoMock package.
package finspacedatamock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	finspacedata "github.com/aws/aws-sdk-go/service/finspacedata"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFinSpaceDataAPI is a mock of FinSpaceDataAPI interface
type MockFinSpaceDataAPI struct {
	ctrl     *gomock.Controller
	recorder *MockFinSpaceDataAPIMockRecorder
}

// MockFinSpaceDataAPIMockRecorder is the mock recorder for MockFinSpaceDataAPI
type MockFinSpaceDataAPIMockRecorder struct {
	mock *MockFinSpaceDataAPI
}

// NewMockFinSpaceDataAPI creates a new mock instance
func NewMockFinSpaceDataAPI(ctrl *gomock.Controller) *MockFinSpaceDataAPI {
	mock := &MockFinSpaceDataAPI{ctrl: ctrl}
	mock.recorder = &MockFinSpaceDataAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFinSpaceDataAPI) EXPECT() *MockFinSpaceDataAPIMockRecorder {
	return m.recorder
}

// CreateChangeset mocks base method
func (m *MockFinSpaceDataAPI) CreateChangeset(arg0 *finspacedata.CreateChangesetInput) (*finspacedata.CreateChangesetOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChangeset", arg0)
	ret0, _ := ret[0].(*finspacedata.CreateChangesetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChangeset indicates an expected call of CreateChangeset
func (mr *MockFinSpaceDataAPIMockRecorder) CreateChangeset(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChangeset", reflect.TypeOf((*MockFinSpaceDataAPI)(nil).CreateChangeset), arg0)
}

// CreateChangesetRequest mocks base method
func (m *MockFinSpaceDataAPI) CreateChangesetRequest(arg0 *finspacedata.CreateChangesetInput) (*request.Request, *finspacedata.CreateChangesetOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChangesetRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*finspacedata.CreateChangesetOutput)
	return ret0, ret1
}

// CreateChangesetRequest indicates an expected call of CreateChangesetRequest
func (mr *MockFinSpaceDataAPIMockRecorder) CreateChangesetRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChangesetRequest", reflect.TypeOf((*MockFinSpaceDataAPI)(nil).CreateChangesetRequest), arg0)
}

// CreateChangesetWithContext mocks base method
func (m *MockFinSpaceDataAPI) CreateChangesetWithContext(arg0 context.Context, arg1 *finspacedata.CreateChangesetInput, arg2 ...request.Option) (*finspacedata.CreateChangesetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateChangesetWithContext", varargs...)
	ret0, _ := ret[0].(*finspacedata.CreateChangesetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChangesetWithContext indicates an expected call of CreateChangesetWithContext
func (mr *MockFinSpaceDataAPIMockRecorder) CreateChangesetWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChangesetWithContext", reflect.TypeOf((*MockFinSpaceDataAPI)(nil).CreateChangesetWithContext), varargs...)
}

// GetProgrammaticAccessCredentials mocks base method
func (m *MockFinSpaceDataAPI) GetProgrammaticAccessCredentials(arg0 *finspacedata.GetProgrammaticAccessCredentialsInput) (*finspacedata.GetProgrammaticAccessCredentialsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProgrammaticAccessCredentials", arg0)
	ret0, _ := ret[0].(*finspacedata.GetProgrammaticAccessCredentialsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProgrammaticAccessCredentials indicates an expected call of GetProgrammaticAccessCredentials
func (mr *MockFinSpaceDataAPIMockRecorder) GetProgrammaticAccessCredentials(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProgrammaticAccessCredentials", reflect.TypeOf((*MockFinSpaceDataAPI)(nil).GetProgrammaticAccessCredentials), arg0)
}

// GetProgrammaticAccessCredentialsRequest mocks base method
func (m *MockFinSpaceDataAPI) GetProgrammaticAccessCredentialsRequest(arg0 *finspacedata.GetProgrammaticAccessCredentialsInput) (*request.Request, *finspacedata.GetProgrammaticAccessCredentialsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProgrammaticAccessCredentialsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*finspacedata.GetProgrammaticAccessCredentialsOutput)
	return ret0, ret1
}

// GetProgrammaticAccessCredentialsRequest indicates an expected call of GetProgrammaticAccessCredentialsRequest
func (mr *MockFinSpaceDataAPIMockRecorder) GetProgrammaticAccessCredentialsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProgrammaticAccessCredentialsRequest", reflect.TypeOf((*MockFinSpaceDataAPI)(nil).GetProgrammaticAccessCredentialsRequest), arg0)
}

// GetProgrammaticAccessCredentialsWithContext mocks base method
func (m *MockFinSpaceDataAPI) GetProgrammaticAccessCredentialsWithContext(arg0 context.Context, arg1 *finspacedata.GetProgrammaticAccessCredentialsInput, arg2 ...request.Option) (*finspacedata.GetProgrammaticAccessCredentialsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProgrammaticAccessCredentialsWithContext", varargs...)
	ret0, _ := ret[0].(*finspacedata.GetProgrammaticAccessCredentialsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProgrammaticAccessCredentialsWithContext indicates an expected call of GetProgrammaticAccessCredentialsWithContext
func (mr *MockFinSpaceDataAPIMockRecorder) GetProgrammaticAccessCredentialsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProgrammaticAccessCredentialsWithContext", reflect.TypeOf((*MockFinSpaceDataAPI)(nil).GetProgrammaticAccessCredentialsWithContext), varargs...)
}

// GetWorkingLocation mocks base method
func (m *MockFinSpaceDataAPI) GetWorkingLocation(arg0 *finspacedata.GetWorkingLocationInput) (*finspacedata.GetWorkingLocationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkingLocation", arg0)
	ret0, _ := ret[0].(*finspacedata.GetWorkingLocationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkingLocation indicates an expected call of GetWorkingLocation
func (mr *MockFinSpaceDataAPIMockRecorder) GetWorkingLocation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkingLocation", reflect.TypeOf((*MockFinSpaceDataAPI)(nil).GetWorkingLocation), arg0)
}

// GetWorkingLocationRequest mocks base method
func (m *MockFinSpaceDataAPI) GetWorkingLocationRequest(arg0 *finspacedata.GetWorkingLocationInput) (*request.Request, *finspacedata.GetWorkingLocationOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkingLocationRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*finspacedata.GetWorkingLocationOutput)
	return ret0, ret1
}

// GetWorkingLocationRequest indicates an expected call of GetWorkingLocationRequest
func (mr *MockFinSpaceDataAPIMockRecorder) GetWorkingLocationRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkingLocationRequest", reflect.TypeOf((*MockFinSpaceDataAPI)(nil).GetWorkingLocationRequest), arg0)
}

// GetWorkingLocationWithContext mocks base method
func (m *MockFinSpaceDataAPI) GetWorkingLocationWithContext(arg0 context.Context, arg1 *finspacedata.GetWorkingLocationInput, arg2 ...request.Option) (*finspacedata.GetWorkingLocationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWorkingLocationWithContext", varargs...)
	ret0, _ := ret[0].(*finspacedata.GetWorkingLocationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkingLocationWithContext indicates an expected call of GetWorkingLocationWithContext
func (mr *MockFinSpaceDataAPIMockRecorder) GetWorkingLocationWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkingLocationWithContext", reflect.TypeOf((*MockFinSpaceDataAPI)(nil).GetWorkingLocationWithContext), varargs...)
}