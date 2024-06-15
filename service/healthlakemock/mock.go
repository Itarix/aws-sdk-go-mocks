// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/healthlake/healthlakeiface (interfaces: HealthLakeAPI)

// Package healthlakemock is a generated GoMock package.
package healthlakemock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	healthlake "github.com/aws/aws-sdk-go/service/healthlake"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockHealthLakeAPI is a mock of HealthLakeAPI interface
type MockHealthLakeAPI struct {
	ctrl     *gomock.Controller
	recorder *MockHealthLakeAPIMockRecorder
}

// MockHealthLakeAPIMockRecorder is the mock recorder for MockHealthLakeAPI
type MockHealthLakeAPIMockRecorder struct {
	mock *MockHealthLakeAPI
}

// NewMockHealthLakeAPI creates a new mock instance
func NewMockHealthLakeAPI(ctrl *gomock.Controller) *MockHealthLakeAPI {
	mock := &MockHealthLakeAPI{ctrl: ctrl}
	mock.recorder = &MockHealthLakeAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHealthLakeAPI) EXPECT() *MockHealthLakeAPIMockRecorder {
	return m.recorder
}

// CreateFHIRDatastore mocks base method
func (m *MockHealthLakeAPI) CreateFHIRDatastore(arg0 *healthlake.CreateFHIRDatastoreInput) (*healthlake.CreateFHIRDatastoreOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFHIRDatastore", arg0)
	ret0, _ := ret[0].(*healthlake.CreateFHIRDatastoreOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFHIRDatastore indicates an expected call of CreateFHIRDatastore
func (mr *MockHealthLakeAPIMockRecorder) CreateFHIRDatastore(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFHIRDatastore", reflect.TypeOf((*MockHealthLakeAPI)(nil).CreateFHIRDatastore), arg0)
}

// CreateFHIRDatastoreRequest mocks base method
func (m *MockHealthLakeAPI) CreateFHIRDatastoreRequest(arg0 *healthlake.CreateFHIRDatastoreInput) (*request.Request, *healthlake.CreateFHIRDatastoreOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFHIRDatastoreRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*healthlake.CreateFHIRDatastoreOutput)
	return ret0, ret1
}

// CreateFHIRDatastoreRequest indicates an expected call of CreateFHIRDatastoreRequest
func (mr *MockHealthLakeAPIMockRecorder) CreateFHIRDatastoreRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFHIRDatastoreRequest", reflect.TypeOf((*MockHealthLakeAPI)(nil).CreateFHIRDatastoreRequest), arg0)
}

// CreateFHIRDatastoreWithContext mocks base method
func (m *MockHealthLakeAPI) CreateFHIRDatastoreWithContext(arg0 context.Context, arg1 *healthlake.CreateFHIRDatastoreInput, arg2 ...request.Option) (*healthlake.CreateFHIRDatastoreOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateFHIRDatastoreWithContext", varargs...)
	ret0, _ := ret[0].(*healthlake.CreateFHIRDatastoreOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFHIRDatastoreWithContext indicates an expected call of CreateFHIRDatastoreWithContext
func (mr *MockHealthLakeAPIMockRecorder) CreateFHIRDatastoreWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFHIRDatastoreWithContext", reflect.TypeOf((*MockHealthLakeAPI)(nil).CreateFHIRDatastoreWithContext), varargs...)
}

// DeleteFHIRDatastore mocks base method
func (m *MockHealthLakeAPI) DeleteFHIRDatastore(arg0 *healthlake.DeleteFHIRDatastoreInput) (*healthlake.DeleteFHIRDatastoreOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFHIRDatastore", arg0)
	ret0, _ := ret[0].(*healthlake.DeleteFHIRDatastoreOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFHIRDatastore indicates an expected call of DeleteFHIRDatastore
func (mr *MockHealthLakeAPIMockRecorder) DeleteFHIRDatastore(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFHIRDatastore", reflect.TypeOf((*MockHealthLakeAPI)(nil).DeleteFHIRDatastore), arg0)
}

// DeleteFHIRDatastoreRequest mocks base method
func (m *MockHealthLakeAPI) DeleteFHIRDatastoreRequest(arg0 *healthlake.DeleteFHIRDatastoreInput) (*request.Request, *healthlake.DeleteFHIRDatastoreOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFHIRDatastoreRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*healthlake.DeleteFHIRDatastoreOutput)
	return ret0, ret1
}

// DeleteFHIRDatastoreRequest indicates an expected call of DeleteFHIRDatastoreRequest
func (mr *MockHealthLakeAPIMockRecorder) DeleteFHIRDatastoreRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFHIRDatastoreRequest", reflect.TypeOf((*MockHealthLakeAPI)(nil).DeleteFHIRDatastoreRequest), arg0)
}

// DeleteFHIRDatastoreWithContext mocks base method
func (m *MockHealthLakeAPI) DeleteFHIRDatastoreWithContext(arg0 context.Context, arg1 *healthlake.DeleteFHIRDatastoreInput, arg2 ...request.Option) (*healthlake.DeleteFHIRDatastoreOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteFHIRDatastoreWithContext", varargs...)
	ret0, _ := ret[0].(*healthlake.DeleteFHIRDatastoreOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFHIRDatastoreWithContext indicates an expected call of DeleteFHIRDatastoreWithContext
func (mr *MockHealthLakeAPIMockRecorder) DeleteFHIRDatastoreWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFHIRDatastoreWithContext", reflect.TypeOf((*MockHealthLakeAPI)(nil).DeleteFHIRDatastoreWithContext), varargs...)
}

// DescribeFHIRDatastore mocks base method
func (m *MockHealthLakeAPI) DescribeFHIRDatastore(arg0 *healthlake.DescribeFHIRDatastoreInput) (*healthlake.DescribeFHIRDatastoreOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeFHIRDatastore", arg0)
	ret0, _ := ret[0].(*healthlake.DescribeFHIRDatastoreOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeFHIRDatastore indicates an expected call of DescribeFHIRDatastore
func (mr *MockHealthLakeAPIMockRecorder) DescribeFHIRDatastore(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFHIRDatastore", reflect.TypeOf((*MockHealthLakeAPI)(nil).DescribeFHIRDatastore), arg0)
}

// DescribeFHIRDatastoreRequest mocks base method
func (m *MockHealthLakeAPI) DescribeFHIRDatastoreRequest(arg0 *healthlake.DescribeFHIRDatastoreInput) (*request.Request, *healthlake.DescribeFHIRDatastoreOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeFHIRDatastoreRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*healthlake.DescribeFHIRDatastoreOutput)
	return ret0, ret1
}

// DescribeFHIRDatastoreRequest indicates an expected call of DescribeFHIRDatastoreRequest
func (mr *MockHealthLakeAPIMockRecorder) DescribeFHIRDatastoreRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFHIRDatastoreRequest", reflect.TypeOf((*MockHealthLakeAPI)(nil).DescribeFHIRDatastoreRequest), arg0)
}

// DescribeFHIRDatastoreWithContext mocks base method
func (m *MockHealthLakeAPI) DescribeFHIRDatastoreWithContext(arg0 context.Context, arg1 *healthlake.DescribeFHIRDatastoreInput, arg2 ...request.Option) (*healthlake.DescribeFHIRDatastoreOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeFHIRDatastoreWithContext", varargs...)
	ret0, _ := ret[0].(*healthlake.DescribeFHIRDatastoreOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeFHIRDatastoreWithContext indicates an expected call of DescribeFHIRDatastoreWithContext
func (mr *MockHealthLakeAPIMockRecorder) DescribeFHIRDatastoreWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFHIRDatastoreWithContext", reflect.TypeOf((*MockHealthLakeAPI)(nil).DescribeFHIRDatastoreWithContext), varargs...)
}

// DescribeFHIRExportJob mocks base method
func (m *MockHealthLakeAPI) DescribeFHIRExportJob(arg0 *healthlake.DescribeFHIRExportJobInput) (*healthlake.DescribeFHIRExportJobOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeFHIRExportJob", arg0)
	ret0, _ := ret[0].(*healthlake.DescribeFHIRExportJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeFHIRExportJob indicates an expected call of DescribeFHIRExportJob
func (mr *MockHealthLakeAPIMockRecorder) DescribeFHIRExportJob(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFHIRExportJob", reflect.TypeOf((*MockHealthLakeAPI)(nil).DescribeFHIRExportJob), arg0)
}

// DescribeFHIRExportJobRequest mocks base method
func (m *MockHealthLakeAPI) DescribeFHIRExportJobRequest(arg0 *healthlake.DescribeFHIRExportJobInput) (*request.Request, *healthlake.DescribeFHIRExportJobOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeFHIRExportJobRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*healthlake.DescribeFHIRExportJobOutput)
	return ret0, ret1
}

// DescribeFHIRExportJobRequest indicates an expected call of DescribeFHIRExportJobRequest
func (mr *MockHealthLakeAPIMockRecorder) DescribeFHIRExportJobRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFHIRExportJobRequest", reflect.TypeOf((*MockHealthLakeAPI)(nil).DescribeFHIRExportJobRequest), arg0)
}

// DescribeFHIRExportJobWithContext mocks base method
func (m *MockHealthLakeAPI) DescribeFHIRExportJobWithContext(arg0 context.Context, arg1 *healthlake.DescribeFHIRExportJobInput, arg2 ...request.Option) (*healthlake.DescribeFHIRExportJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeFHIRExportJobWithContext", varargs...)
	ret0, _ := ret[0].(*healthlake.DescribeFHIRExportJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeFHIRExportJobWithContext indicates an expected call of DescribeFHIRExportJobWithContext
func (mr *MockHealthLakeAPIMockRecorder) DescribeFHIRExportJobWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFHIRExportJobWithContext", reflect.TypeOf((*MockHealthLakeAPI)(nil).DescribeFHIRExportJobWithContext), varargs...)
}

// DescribeFHIRImportJob mocks base method
func (m *MockHealthLakeAPI) DescribeFHIRImportJob(arg0 *healthlake.DescribeFHIRImportJobInput) (*healthlake.DescribeFHIRImportJobOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeFHIRImportJob", arg0)
	ret0, _ := ret[0].(*healthlake.DescribeFHIRImportJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeFHIRImportJob indicates an expected call of DescribeFHIRImportJob
func (mr *MockHealthLakeAPIMockRecorder) DescribeFHIRImportJob(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFHIRImportJob", reflect.TypeOf((*MockHealthLakeAPI)(nil).DescribeFHIRImportJob), arg0)
}

// DescribeFHIRImportJobRequest mocks base method
func (m *MockHealthLakeAPI) DescribeFHIRImportJobRequest(arg0 *healthlake.DescribeFHIRImportJobInput) (*request.Request, *healthlake.DescribeFHIRImportJobOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeFHIRImportJobRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*healthlake.DescribeFHIRImportJobOutput)
	return ret0, ret1
}

// DescribeFHIRImportJobRequest indicates an expected call of DescribeFHIRImportJobRequest
func (mr *MockHealthLakeAPIMockRecorder) DescribeFHIRImportJobRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFHIRImportJobRequest", reflect.TypeOf((*MockHealthLakeAPI)(nil).DescribeFHIRImportJobRequest), arg0)
}

// DescribeFHIRImportJobWithContext mocks base method
func (m *MockHealthLakeAPI) DescribeFHIRImportJobWithContext(arg0 context.Context, arg1 *healthlake.DescribeFHIRImportJobInput, arg2 ...request.Option) (*healthlake.DescribeFHIRImportJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeFHIRImportJobWithContext", varargs...)
	ret0, _ := ret[0].(*healthlake.DescribeFHIRImportJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeFHIRImportJobWithContext indicates an expected call of DescribeFHIRImportJobWithContext
func (mr *MockHealthLakeAPIMockRecorder) DescribeFHIRImportJobWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFHIRImportJobWithContext", reflect.TypeOf((*MockHealthLakeAPI)(nil).DescribeFHIRImportJobWithContext), varargs...)
}

// ListFHIRDatastores mocks base method
func (m *MockHealthLakeAPI) ListFHIRDatastores(arg0 *healthlake.ListFHIRDatastoresInput) (*healthlake.ListFHIRDatastoresOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFHIRDatastores", arg0)
	ret0, _ := ret[0].(*healthlake.ListFHIRDatastoresOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFHIRDatastores indicates an expected call of ListFHIRDatastores
func (mr *MockHealthLakeAPIMockRecorder) ListFHIRDatastores(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFHIRDatastores", reflect.TypeOf((*MockHealthLakeAPI)(nil).ListFHIRDatastores), arg0)
}

// ListFHIRDatastoresPages mocks base method
func (m *MockHealthLakeAPI) ListFHIRDatastoresPages(arg0 *healthlake.ListFHIRDatastoresInput, arg1 func(*healthlake.ListFHIRDatastoresOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFHIRDatastoresPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListFHIRDatastoresPages indicates an expected call of ListFHIRDatastoresPages
func (mr *MockHealthLakeAPIMockRecorder) ListFHIRDatastoresPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFHIRDatastoresPages", reflect.TypeOf((*MockHealthLakeAPI)(nil).ListFHIRDatastoresPages), arg0, arg1)
}

// ListFHIRDatastoresPagesWithContext mocks base method
func (m *MockHealthLakeAPI) ListFHIRDatastoresPagesWithContext(arg0 context.Context, arg1 *healthlake.ListFHIRDatastoresInput, arg2 func(*healthlake.ListFHIRDatastoresOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFHIRDatastoresPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListFHIRDatastoresPagesWithContext indicates an expected call of ListFHIRDatastoresPagesWithContext
func (mr *MockHealthLakeAPIMockRecorder) ListFHIRDatastoresPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFHIRDatastoresPagesWithContext", reflect.TypeOf((*MockHealthLakeAPI)(nil).ListFHIRDatastoresPagesWithContext), varargs...)
}

// ListFHIRDatastoresRequest mocks base method
func (m *MockHealthLakeAPI) ListFHIRDatastoresRequest(arg0 *healthlake.ListFHIRDatastoresInput) (*request.Request, *healthlake.ListFHIRDatastoresOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFHIRDatastoresRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*healthlake.ListFHIRDatastoresOutput)
	return ret0, ret1
}

// ListFHIRDatastoresRequest indicates an expected call of ListFHIRDatastoresRequest
func (mr *MockHealthLakeAPIMockRecorder) ListFHIRDatastoresRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFHIRDatastoresRequest", reflect.TypeOf((*MockHealthLakeAPI)(nil).ListFHIRDatastoresRequest), arg0)
}

// ListFHIRDatastoresWithContext mocks base method
func (m *MockHealthLakeAPI) ListFHIRDatastoresWithContext(arg0 context.Context, arg1 *healthlake.ListFHIRDatastoresInput, arg2 ...request.Option) (*healthlake.ListFHIRDatastoresOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFHIRDatastoresWithContext", varargs...)
	ret0, _ := ret[0].(*healthlake.ListFHIRDatastoresOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFHIRDatastoresWithContext indicates an expected call of ListFHIRDatastoresWithContext
func (mr *MockHealthLakeAPIMockRecorder) ListFHIRDatastoresWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFHIRDatastoresWithContext", reflect.TypeOf((*MockHealthLakeAPI)(nil).ListFHIRDatastoresWithContext), varargs...)
}

// ListFHIRExportJobs mocks base method
func (m *MockHealthLakeAPI) ListFHIRExportJobs(arg0 *healthlake.ListFHIRExportJobsInput) (*healthlake.ListFHIRExportJobsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFHIRExportJobs", arg0)
	ret0, _ := ret[0].(*healthlake.ListFHIRExportJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFHIRExportJobs indicates an expected call of ListFHIRExportJobs
func (mr *MockHealthLakeAPIMockRecorder) ListFHIRExportJobs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFHIRExportJobs", reflect.TypeOf((*MockHealthLakeAPI)(nil).ListFHIRExportJobs), arg0)
}

// ListFHIRExportJobsPages mocks base method
func (m *MockHealthLakeAPI) ListFHIRExportJobsPages(arg0 *healthlake.ListFHIRExportJobsInput, arg1 func(*healthlake.ListFHIRExportJobsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFHIRExportJobsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListFHIRExportJobsPages indicates an expected call of ListFHIRExportJobsPages
func (mr *MockHealthLakeAPIMockRecorder) ListFHIRExportJobsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFHIRExportJobsPages", reflect.TypeOf((*MockHealthLakeAPI)(nil).ListFHIRExportJobsPages), arg0, arg1)
}

// ListFHIRExportJobsPagesWithContext mocks base method
func (m *MockHealthLakeAPI) ListFHIRExportJobsPagesWithContext(arg0 context.Context, arg1 *healthlake.ListFHIRExportJobsInput, arg2 func(*healthlake.ListFHIRExportJobsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFHIRExportJobsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListFHIRExportJobsPagesWithContext indicates an expected call of ListFHIRExportJobsPagesWithContext
func (mr *MockHealthLakeAPIMockRecorder) ListFHIRExportJobsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFHIRExportJobsPagesWithContext", reflect.TypeOf((*MockHealthLakeAPI)(nil).ListFHIRExportJobsPagesWithContext), varargs...)
}

// ListFHIRExportJobsRequest mocks base method
func (m *MockHealthLakeAPI) ListFHIRExportJobsRequest(arg0 *healthlake.ListFHIRExportJobsInput) (*request.Request, *healthlake.ListFHIRExportJobsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFHIRExportJobsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*healthlake.ListFHIRExportJobsOutput)
	return ret0, ret1
}

// ListFHIRExportJobsRequest indicates an expected call of ListFHIRExportJobsRequest
func (mr *MockHealthLakeAPIMockRecorder) ListFHIRExportJobsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFHIRExportJobsRequest", reflect.TypeOf((*MockHealthLakeAPI)(nil).ListFHIRExportJobsRequest), arg0)
}

// ListFHIRExportJobsWithContext mocks base method
func (m *MockHealthLakeAPI) ListFHIRExportJobsWithContext(arg0 context.Context, arg1 *healthlake.ListFHIRExportJobsInput, arg2 ...request.Option) (*healthlake.ListFHIRExportJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFHIRExportJobsWithContext", varargs...)
	ret0, _ := ret[0].(*healthlake.ListFHIRExportJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFHIRExportJobsWithContext indicates an expected call of ListFHIRExportJobsWithContext
func (mr *MockHealthLakeAPIMockRecorder) ListFHIRExportJobsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFHIRExportJobsWithContext", reflect.TypeOf((*MockHealthLakeAPI)(nil).ListFHIRExportJobsWithContext), varargs...)
}

// ListFHIRImportJobs mocks base method
func (m *MockHealthLakeAPI) ListFHIRImportJobs(arg0 *healthlake.ListFHIRImportJobsInput) (*healthlake.ListFHIRImportJobsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFHIRImportJobs", arg0)
	ret0, _ := ret[0].(*healthlake.ListFHIRImportJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFHIRImportJobs indicates an expected call of ListFHIRImportJobs
func (mr *MockHealthLakeAPIMockRecorder) ListFHIRImportJobs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFHIRImportJobs", reflect.TypeOf((*MockHealthLakeAPI)(nil).ListFHIRImportJobs), arg0)
}

// ListFHIRImportJobsPages mocks base method
func (m *MockHealthLakeAPI) ListFHIRImportJobsPages(arg0 *healthlake.ListFHIRImportJobsInput, arg1 func(*healthlake.ListFHIRImportJobsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFHIRImportJobsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListFHIRImportJobsPages indicates an expected call of ListFHIRImportJobsPages
func (mr *MockHealthLakeAPIMockRecorder) ListFHIRImportJobsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFHIRImportJobsPages", reflect.TypeOf((*MockHealthLakeAPI)(nil).ListFHIRImportJobsPages), arg0, arg1)
}

// ListFHIRImportJobsPagesWithContext mocks base method
func (m *MockHealthLakeAPI) ListFHIRImportJobsPagesWithContext(arg0 context.Context, arg1 *healthlake.ListFHIRImportJobsInput, arg2 func(*healthlake.ListFHIRImportJobsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFHIRImportJobsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListFHIRImportJobsPagesWithContext indicates an expected call of ListFHIRImportJobsPagesWithContext
func (mr *MockHealthLakeAPIMockRecorder) ListFHIRImportJobsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFHIRImportJobsPagesWithContext", reflect.TypeOf((*MockHealthLakeAPI)(nil).ListFHIRImportJobsPagesWithContext), varargs...)
}

// ListFHIRImportJobsRequest mocks base method
func (m *MockHealthLakeAPI) ListFHIRImportJobsRequest(arg0 *healthlake.ListFHIRImportJobsInput) (*request.Request, *healthlake.ListFHIRImportJobsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFHIRImportJobsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*healthlake.ListFHIRImportJobsOutput)
	return ret0, ret1
}

// ListFHIRImportJobsRequest indicates an expected call of ListFHIRImportJobsRequest
func (mr *MockHealthLakeAPIMockRecorder) ListFHIRImportJobsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFHIRImportJobsRequest", reflect.TypeOf((*MockHealthLakeAPI)(nil).ListFHIRImportJobsRequest), arg0)
}

// ListFHIRImportJobsWithContext mocks base method
func (m *MockHealthLakeAPI) ListFHIRImportJobsWithContext(arg0 context.Context, arg1 *healthlake.ListFHIRImportJobsInput, arg2 ...request.Option) (*healthlake.ListFHIRImportJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFHIRImportJobsWithContext", varargs...)
	ret0, _ := ret[0].(*healthlake.ListFHIRImportJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFHIRImportJobsWithContext indicates an expected call of ListFHIRImportJobsWithContext
func (mr *MockHealthLakeAPIMockRecorder) ListFHIRImportJobsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFHIRImportJobsWithContext", reflect.TypeOf((*MockHealthLakeAPI)(nil).ListFHIRImportJobsWithContext), varargs...)
}

// ListTagsForResource mocks base method
func (m *MockHealthLakeAPI) ListTagsForResource(arg0 *healthlake.ListTagsForResourceInput) (*healthlake.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForResource", arg0)
	ret0, _ := ret[0].(*healthlake.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource
func (mr *MockHealthLakeAPIMockRecorder) ListTagsForResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockHealthLakeAPI)(nil).ListTagsForResource), arg0)
}

// ListTagsForResourceRequest mocks base method
func (m *MockHealthLakeAPI) ListTagsForResourceRequest(arg0 *healthlake.ListTagsForResourceInput) (*request.Request, *healthlake.ListTagsForResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*healthlake.ListTagsForResourceOutput)
	return ret0, ret1
}

// ListTagsForResourceRequest indicates an expected call of ListTagsForResourceRequest
func (mr *MockHealthLakeAPIMockRecorder) ListTagsForResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResourceRequest", reflect.TypeOf((*MockHealthLakeAPI)(nil).ListTagsForResourceRequest), arg0)
}

// ListTagsForResourceWithContext mocks base method
func (m *MockHealthLakeAPI) ListTagsForResourceWithContext(arg0 context.Context, arg1 *healthlake.ListTagsForResourceInput, arg2 ...request.Option) (*healthlake.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResourceWithContext", varargs...)
	ret0, _ := ret[0].(*healthlake.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResourceWithContext indicates an expected call of ListTagsForResourceWithContext
func (mr *MockHealthLakeAPIMockRecorder) ListTagsForResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResourceWithContext", reflect.TypeOf((*MockHealthLakeAPI)(nil).ListTagsForResourceWithContext), varargs...)
}

// StartFHIRExportJob mocks base method
func (m *MockHealthLakeAPI) StartFHIRExportJob(arg0 *healthlake.StartFHIRExportJobInput) (*healthlake.StartFHIRExportJobOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartFHIRExportJob", arg0)
	ret0, _ := ret[0].(*healthlake.StartFHIRExportJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartFHIRExportJob indicates an expected call of StartFHIRExportJob
func (mr *MockHealthLakeAPIMockRecorder) StartFHIRExportJob(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartFHIRExportJob", reflect.TypeOf((*MockHealthLakeAPI)(nil).StartFHIRExportJob), arg0)
}

// StartFHIRExportJobRequest mocks base method
func (m *MockHealthLakeAPI) StartFHIRExportJobRequest(arg0 *healthlake.StartFHIRExportJobInput) (*request.Request, *healthlake.StartFHIRExportJobOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartFHIRExportJobRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*healthlake.StartFHIRExportJobOutput)
	return ret0, ret1
}

// StartFHIRExportJobRequest indicates an expected call of StartFHIRExportJobRequest
func (mr *MockHealthLakeAPIMockRecorder) StartFHIRExportJobRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartFHIRExportJobRequest", reflect.TypeOf((*MockHealthLakeAPI)(nil).StartFHIRExportJobRequest), arg0)
}

// StartFHIRExportJobWithContext mocks base method
func (m *MockHealthLakeAPI) StartFHIRExportJobWithContext(arg0 context.Context, arg1 *healthlake.StartFHIRExportJobInput, arg2 ...request.Option) (*healthlake.StartFHIRExportJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StartFHIRExportJobWithContext", varargs...)
	ret0, _ := ret[0].(*healthlake.StartFHIRExportJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartFHIRExportJobWithContext indicates an expected call of StartFHIRExportJobWithContext
func (mr *MockHealthLakeAPIMockRecorder) StartFHIRExportJobWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartFHIRExportJobWithContext", reflect.TypeOf((*MockHealthLakeAPI)(nil).StartFHIRExportJobWithContext), varargs...)
}

// StartFHIRImportJob mocks base method
func (m *MockHealthLakeAPI) StartFHIRImportJob(arg0 *healthlake.StartFHIRImportJobInput) (*healthlake.StartFHIRImportJobOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartFHIRImportJob", arg0)
	ret0, _ := ret[0].(*healthlake.StartFHIRImportJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartFHIRImportJob indicates an expected call of StartFHIRImportJob
func (mr *MockHealthLakeAPIMockRecorder) StartFHIRImportJob(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartFHIRImportJob", reflect.TypeOf((*MockHealthLakeAPI)(nil).StartFHIRImportJob), arg0)
}

// StartFHIRImportJobRequest mocks base method
func (m *MockHealthLakeAPI) StartFHIRImportJobRequest(arg0 *healthlake.StartFHIRImportJobInput) (*request.Request, *healthlake.StartFHIRImportJobOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartFHIRImportJobRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*healthlake.StartFHIRImportJobOutput)
	return ret0, ret1
}

// StartFHIRImportJobRequest indicates an expected call of StartFHIRImportJobRequest
func (mr *MockHealthLakeAPIMockRecorder) StartFHIRImportJobRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartFHIRImportJobRequest", reflect.TypeOf((*MockHealthLakeAPI)(nil).StartFHIRImportJobRequest), arg0)
}

// StartFHIRImportJobWithContext mocks base method
func (m *MockHealthLakeAPI) StartFHIRImportJobWithContext(arg0 context.Context, arg1 *healthlake.StartFHIRImportJobInput, arg2 ...request.Option) (*healthlake.StartFHIRImportJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StartFHIRImportJobWithContext", varargs...)
	ret0, _ := ret[0].(*healthlake.StartFHIRImportJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartFHIRImportJobWithContext indicates an expected call of StartFHIRImportJobWithContext
func (mr *MockHealthLakeAPIMockRecorder) StartFHIRImportJobWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartFHIRImportJobWithContext", reflect.TypeOf((*MockHealthLakeAPI)(nil).StartFHIRImportJobWithContext), varargs...)
}

// TagResource mocks base method
func (m *MockHealthLakeAPI) TagResource(arg0 *healthlake.TagResourceInput) (*healthlake.TagResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagResource", arg0)
	ret0, _ := ret[0].(*healthlake.TagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagResource indicates an expected call of TagResource
func (mr *MockHealthLakeAPIMockRecorder) TagResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResource", reflect.TypeOf((*MockHealthLakeAPI)(nil).TagResource), arg0)
}

// TagResourceRequest mocks base method
func (m *MockHealthLakeAPI) TagResourceRequest(arg0 *healthlake.TagResourceInput) (*request.Request, *healthlake.TagResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*healthlake.TagResourceOutput)
	return ret0, ret1
}

// TagResourceRequest indicates an expected call of TagResourceRequest
func (mr *MockHealthLakeAPIMockRecorder) TagResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResourceRequest", reflect.TypeOf((*MockHealthLakeAPI)(nil).TagResourceRequest), arg0)
}

// TagResourceWithContext mocks base method
func (m *MockHealthLakeAPI) TagResourceWithContext(arg0 context.Context, arg1 *healthlake.TagResourceInput, arg2 ...request.Option) (*healthlake.TagResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TagResourceWithContext", varargs...)
	ret0, _ := ret[0].(*healthlake.TagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagResourceWithContext indicates an expected call of TagResourceWithContext
func (mr *MockHealthLakeAPIMockRecorder) TagResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResourceWithContext", reflect.TypeOf((*MockHealthLakeAPI)(nil).TagResourceWithContext), varargs...)
}

// UntagResource mocks base method
func (m *MockHealthLakeAPI) UntagResource(arg0 *healthlake.UntagResourceInput) (*healthlake.UntagResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagResource", arg0)
	ret0, _ := ret[0].(*healthlake.UntagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagResource indicates an expected call of UntagResource
func (mr *MockHealthLakeAPIMockRecorder) UntagResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResource", reflect.TypeOf((*MockHealthLakeAPI)(nil).UntagResource), arg0)
}

// UntagResourceRequest mocks base method
func (m *MockHealthLakeAPI) UntagResourceRequest(arg0 *healthlake.UntagResourceInput) (*request.Request, *healthlake.UntagResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*healthlake.UntagResourceOutput)
	return ret0, ret1
}

// UntagResourceRequest indicates an expected call of UntagResourceRequest
func (mr *MockHealthLakeAPIMockRecorder) UntagResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResourceRequest", reflect.TypeOf((*MockHealthLakeAPI)(nil).UntagResourceRequest), arg0)
}

// UntagResourceWithContext mocks base method
func (m *MockHealthLakeAPI) UntagResourceWithContext(arg0 context.Context, arg1 *healthlake.UntagResourceInput, arg2 ...request.Option) (*healthlake.UntagResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UntagResourceWithContext", varargs...)
	ret0, _ := ret[0].(*healthlake.UntagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagResourceWithContext indicates an expected call of UntagResourceWithContext
func (mr *MockHealthLakeAPIMockRecorder) UntagResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResourceWithContext", reflect.TypeOf((*MockHealthLakeAPI)(nil).UntagResourceWithContext), varargs...)
}
