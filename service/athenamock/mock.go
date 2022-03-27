// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/athena/athenaiface (interfaces: AthenaAPI)

// Package athenamock is a generated GoMock package.
package athenamock

import (
	aws "github.com/aws/aws-sdk-go/aws"
	request "github.com/aws/aws-sdk-go/aws/request"
	athena "github.com/aws/aws-sdk-go/service/athena"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAthenaAPI is a mock of AthenaAPI interface
type MockAthenaAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAthenaAPIMockRecorder
}

// MockAthenaAPIMockRecorder is the mock recorder for MockAthenaAPI
type MockAthenaAPIMockRecorder struct {
	mock *MockAthenaAPI
}

// NewMockAthenaAPI creates a new mock instance
func NewMockAthenaAPI(ctrl *gomock.Controller) *MockAthenaAPI {
	mock := &MockAthenaAPI{ctrl: ctrl}
	mock.recorder = &MockAthenaAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAthenaAPI) EXPECT() *MockAthenaAPIMockRecorder {
	return m.recorder
}

// BatchGetNamedQuery mocks base method
func (m *MockAthenaAPI) BatchGetNamedQuery(arg0 *athena.BatchGetNamedQueryInput) (*athena.BatchGetNamedQueryOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchGetNamedQuery", arg0)
	ret0, _ := ret[0].(*athena.BatchGetNamedQueryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetNamedQuery indicates an expected call of BatchGetNamedQuery
func (mr *MockAthenaAPIMockRecorder) BatchGetNamedQuery(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetNamedQuery", reflect.TypeOf((*MockAthenaAPI)(nil).BatchGetNamedQuery), arg0)
}

// BatchGetNamedQueryRequest mocks base method
func (m *MockAthenaAPI) BatchGetNamedQueryRequest(arg0 *athena.BatchGetNamedQueryInput) (*request.Request, *athena.BatchGetNamedQueryOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchGetNamedQueryRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*athena.BatchGetNamedQueryOutput)
	return ret0, ret1
}

// BatchGetNamedQueryRequest indicates an expected call of BatchGetNamedQueryRequest
func (mr *MockAthenaAPIMockRecorder) BatchGetNamedQueryRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetNamedQueryRequest", reflect.TypeOf((*MockAthenaAPI)(nil).BatchGetNamedQueryRequest), arg0)
}

// BatchGetNamedQueryWithContext mocks base method
func (m *MockAthenaAPI) BatchGetNamedQueryWithContext(arg0 aws.Context, arg1 *athena.BatchGetNamedQueryInput, arg2 ...request.Option) (*athena.BatchGetNamedQueryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchGetNamedQueryWithContext", varargs...)
	ret0, _ := ret[0].(*athena.BatchGetNamedQueryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetNamedQueryWithContext indicates an expected call of BatchGetNamedQueryWithContext
func (mr *MockAthenaAPIMockRecorder) BatchGetNamedQueryWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetNamedQueryWithContext", reflect.TypeOf((*MockAthenaAPI)(nil).BatchGetNamedQueryWithContext), varargs...)
}

// BatchGetQueryExecution mocks base method
func (m *MockAthenaAPI) BatchGetQueryExecution(arg0 *athena.BatchGetQueryExecutionInput) (*athena.BatchGetQueryExecutionOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchGetQueryExecution", arg0)
	ret0, _ := ret[0].(*athena.BatchGetQueryExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetQueryExecution indicates an expected call of BatchGetQueryExecution
func (mr *MockAthenaAPIMockRecorder) BatchGetQueryExecution(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetQueryExecution", reflect.TypeOf((*MockAthenaAPI)(nil).BatchGetQueryExecution), arg0)
}

// BatchGetQueryExecutionRequest mocks base method
func (m *MockAthenaAPI) BatchGetQueryExecutionRequest(arg0 *athena.BatchGetQueryExecutionInput) (*request.Request, *athena.BatchGetQueryExecutionOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchGetQueryExecutionRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*athena.BatchGetQueryExecutionOutput)
	return ret0, ret1
}

// BatchGetQueryExecutionRequest indicates an expected call of BatchGetQueryExecutionRequest
func (mr *MockAthenaAPIMockRecorder) BatchGetQueryExecutionRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetQueryExecutionRequest", reflect.TypeOf((*MockAthenaAPI)(nil).BatchGetQueryExecutionRequest), arg0)
}

// BatchGetQueryExecutionWithContext mocks base method
func (m *MockAthenaAPI) BatchGetQueryExecutionWithContext(arg0 aws.Context, arg1 *athena.BatchGetQueryExecutionInput, arg2 ...request.Option) (*athena.BatchGetQueryExecutionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchGetQueryExecutionWithContext", varargs...)
	ret0, _ := ret[0].(*athena.BatchGetQueryExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetQueryExecutionWithContext indicates an expected call of BatchGetQueryExecutionWithContext
func (mr *MockAthenaAPIMockRecorder) BatchGetQueryExecutionWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetQueryExecutionWithContext", reflect.TypeOf((*MockAthenaAPI)(nil).BatchGetQueryExecutionWithContext), varargs...)
}

// CreateNamedQuery mocks base method
func (m *MockAthenaAPI) CreateNamedQuery(arg0 *athena.CreateNamedQueryInput) (*athena.CreateNamedQueryOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNamedQuery", arg0)
	ret0, _ := ret[0].(*athena.CreateNamedQueryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNamedQuery indicates an expected call of CreateNamedQuery
func (mr *MockAthenaAPIMockRecorder) CreateNamedQuery(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNamedQuery", reflect.TypeOf((*MockAthenaAPI)(nil).CreateNamedQuery), arg0)
}

// CreateNamedQueryRequest mocks base method
func (m *MockAthenaAPI) CreateNamedQueryRequest(arg0 *athena.CreateNamedQueryInput) (*request.Request, *athena.CreateNamedQueryOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNamedQueryRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*athena.CreateNamedQueryOutput)
	return ret0, ret1
}

// CreateNamedQueryRequest indicates an expected call of CreateNamedQueryRequest
func (mr *MockAthenaAPIMockRecorder) CreateNamedQueryRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNamedQueryRequest", reflect.TypeOf((*MockAthenaAPI)(nil).CreateNamedQueryRequest), arg0)
}

// CreateNamedQueryWithContext mocks base method
func (m *MockAthenaAPI) CreateNamedQueryWithContext(arg0 aws.Context, arg1 *athena.CreateNamedQueryInput, arg2 ...request.Option) (*athena.CreateNamedQueryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateNamedQueryWithContext", varargs...)
	ret0, _ := ret[0].(*athena.CreateNamedQueryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNamedQueryWithContext indicates an expected call of CreateNamedQueryWithContext
func (mr *MockAthenaAPIMockRecorder) CreateNamedQueryWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNamedQueryWithContext", reflect.TypeOf((*MockAthenaAPI)(nil).CreateNamedQueryWithContext), varargs...)
}

// DeleteNamedQuery mocks base method
func (m *MockAthenaAPI) DeleteNamedQuery(arg0 *athena.DeleteNamedQueryInput) (*athena.DeleteNamedQueryOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNamedQuery", arg0)
	ret0, _ := ret[0].(*athena.DeleteNamedQueryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteNamedQuery indicates an expected call of DeleteNamedQuery
func (mr *MockAthenaAPIMockRecorder) DeleteNamedQuery(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNamedQuery", reflect.TypeOf((*MockAthenaAPI)(nil).DeleteNamedQuery), arg0)
}

// DeleteNamedQueryRequest mocks base method
func (m *MockAthenaAPI) DeleteNamedQueryRequest(arg0 *athena.DeleteNamedQueryInput) (*request.Request, *athena.DeleteNamedQueryOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNamedQueryRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*athena.DeleteNamedQueryOutput)
	return ret0, ret1
}

// DeleteNamedQueryRequest indicates an expected call of DeleteNamedQueryRequest
func (mr *MockAthenaAPIMockRecorder) DeleteNamedQueryRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNamedQueryRequest", reflect.TypeOf((*MockAthenaAPI)(nil).DeleteNamedQueryRequest), arg0)
}

// DeleteNamedQueryWithContext mocks base method
func (m *MockAthenaAPI) DeleteNamedQueryWithContext(arg0 aws.Context, arg1 *athena.DeleteNamedQueryInput, arg2 ...request.Option) (*athena.DeleteNamedQueryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteNamedQueryWithContext", varargs...)
	ret0, _ := ret[0].(*athena.DeleteNamedQueryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteNamedQueryWithContext indicates an expected call of DeleteNamedQueryWithContext
func (mr *MockAthenaAPIMockRecorder) DeleteNamedQueryWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNamedQueryWithContext", reflect.TypeOf((*MockAthenaAPI)(nil).DeleteNamedQueryWithContext), varargs...)
}

// GetNamedQuery mocks base method
func (m *MockAthenaAPI) GetNamedQuery(arg0 *athena.GetNamedQueryInput) (*athena.GetNamedQueryOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamedQuery", arg0)
	ret0, _ := ret[0].(*athena.GetNamedQueryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNamedQuery indicates an expected call of GetNamedQuery
func (mr *MockAthenaAPIMockRecorder) GetNamedQuery(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamedQuery", reflect.TypeOf((*MockAthenaAPI)(nil).GetNamedQuery), arg0)
}

// GetNamedQueryRequest mocks base method
func (m *MockAthenaAPI) GetNamedQueryRequest(arg0 *athena.GetNamedQueryInput) (*request.Request, *athena.GetNamedQueryOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamedQueryRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*athena.GetNamedQueryOutput)
	return ret0, ret1
}

// GetNamedQueryRequest indicates an expected call of GetNamedQueryRequest
func (mr *MockAthenaAPIMockRecorder) GetNamedQueryRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamedQueryRequest", reflect.TypeOf((*MockAthenaAPI)(nil).GetNamedQueryRequest), arg0)
}

// GetNamedQueryWithContext mocks base method
func (m *MockAthenaAPI) GetNamedQueryWithContext(arg0 aws.Context, arg1 *athena.GetNamedQueryInput, arg2 ...request.Option) (*athena.GetNamedQueryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetNamedQueryWithContext", varargs...)
	ret0, _ := ret[0].(*athena.GetNamedQueryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNamedQueryWithContext indicates an expected call of GetNamedQueryWithContext
func (mr *MockAthenaAPIMockRecorder) GetNamedQueryWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamedQueryWithContext", reflect.TypeOf((*MockAthenaAPI)(nil).GetNamedQueryWithContext), varargs...)
}

// GetQueryExecution mocks base method
func (m *MockAthenaAPI) GetQueryExecution(arg0 *athena.GetQueryExecutionInput) (*athena.GetQueryExecutionOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueryExecution", arg0)
	ret0, _ := ret[0].(*athena.GetQueryExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQueryExecution indicates an expected call of GetQueryExecution
func (mr *MockAthenaAPIMockRecorder) GetQueryExecution(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryExecution", reflect.TypeOf((*MockAthenaAPI)(nil).GetQueryExecution), arg0)
}

// GetQueryExecutionRequest mocks base method
func (m *MockAthenaAPI) GetQueryExecutionRequest(arg0 *athena.GetQueryExecutionInput) (*request.Request, *athena.GetQueryExecutionOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueryExecutionRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*athena.GetQueryExecutionOutput)
	return ret0, ret1
}

// GetQueryExecutionRequest indicates an expected call of GetQueryExecutionRequest
func (mr *MockAthenaAPIMockRecorder) GetQueryExecutionRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryExecutionRequest", reflect.TypeOf((*MockAthenaAPI)(nil).GetQueryExecutionRequest), arg0)
}

// GetQueryExecutionWithContext mocks base method
func (m *MockAthenaAPI) GetQueryExecutionWithContext(arg0 aws.Context, arg1 *athena.GetQueryExecutionInput, arg2 ...request.Option) (*athena.GetQueryExecutionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetQueryExecutionWithContext", varargs...)
	ret0, _ := ret[0].(*athena.GetQueryExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQueryExecutionWithContext indicates an expected call of GetQueryExecutionWithContext
func (mr *MockAthenaAPIMockRecorder) GetQueryExecutionWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryExecutionWithContext", reflect.TypeOf((*MockAthenaAPI)(nil).GetQueryExecutionWithContext), varargs...)
}

// GetQueryResults mocks base method
func (m *MockAthenaAPI) GetQueryResults(arg0 *athena.GetQueryResultsInput) (*athena.GetQueryResultsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueryResults", arg0)
	ret0, _ := ret[0].(*athena.GetQueryResultsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQueryResults indicates an expected call of GetQueryResults
func (mr *MockAthenaAPIMockRecorder) GetQueryResults(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryResults", reflect.TypeOf((*MockAthenaAPI)(nil).GetQueryResults), arg0)
}

// GetQueryResultsPages mocks base method
func (m *MockAthenaAPI) GetQueryResultsPages(arg0 *athena.GetQueryResultsInput, arg1 func(*athena.GetQueryResultsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueryResultsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetQueryResultsPages indicates an expected call of GetQueryResultsPages
func (mr *MockAthenaAPIMockRecorder) GetQueryResultsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryResultsPages", reflect.TypeOf((*MockAthenaAPI)(nil).GetQueryResultsPages), arg0, arg1)
}

// GetQueryResultsPagesWithContext mocks base method
func (m *MockAthenaAPI) GetQueryResultsPagesWithContext(arg0 aws.Context, arg1 *athena.GetQueryResultsInput, arg2 func(*athena.GetQueryResultsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetQueryResultsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetQueryResultsPagesWithContext indicates an expected call of GetQueryResultsPagesWithContext
func (mr *MockAthenaAPIMockRecorder) GetQueryResultsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryResultsPagesWithContext", reflect.TypeOf((*MockAthenaAPI)(nil).GetQueryResultsPagesWithContext), varargs...)
}

// GetQueryResultsRequest mocks base method
func (m *MockAthenaAPI) GetQueryResultsRequest(arg0 *athena.GetQueryResultsInput) (*request.Request, *athena.GetQueryResultsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueryResultsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*athena.GetQueryResultsOutput)
	return ret0, ret1
}

// GetQueryResultsRequest indicates an expected call of GetQueryResultsRequest
func (mr *MockAthenaAPIMockRecorder) GetQueryResultsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryResultsRequest", reflect.TypeOf((*MockAthenaAPI)(nil).GetQueryResultsRequest), arg0)
}

// GetQueryResultsWithContext mocks base method
func (m *MockAthenaAPI) GetQueryResultsWithContext(arg0 aws.Context, arg1 *athena.GetQueryResultsInput, arg2 ...request.Option) (*athena.GetQueryResultsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetQueryResultsWithContext", varargs...)
	ret0, _ := ret[0].(*athena.GetQueryResultsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQueryResultsWithContext indicates an expected call of GetQueryResultsWithContext
func (mr *MockAthenaAPIMockRecorder) GetQueryResultsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryResultsWithContext", reflect.TypeOf((*MockAthenaAPI)(nil).GetQueryResultsWithContext), varargs...)
}

// ListNamedQueries mocks base method
func (m *MockAthenaAPI) ListNamedQueries(arg0 *athena.ListNamedQueriesInput) (*athena.ListNamedQueriesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNamedQueries", arg0)
	ret0, _ := ret[0].(*athena.ListNamedQueriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNamedQueries indicates an expected call of ListNamedQueries
func (mr *MockAthenaAPIMockRecorder) ListNamedQueries(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNamedQueries", reflect.TypeOf((*MockAthenaAPI)(nil).ListNamedQueries), arg0)
}

// ListNamedQueriesPages mocks base method
func (m *MockAthenaAPI) ListNamedQueriesPages(arg0 *athena.ListNamedQueriesInput, arg1 func(*athena.ListNamedQueriesOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNamedQueriesPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListNamedQueriesPages indicates an expected call of ListNamedQueriesPages
func (mr *MockAthenaAPIMockRecorder) ListNamedQueriesPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNamedQueriesPages", reflect.TypeOf((*MockAthenaAPI)(nil).ListNamedQueriesPages), arg0, arg1)
}

// ListNamedQueriesPagesWithContext mocks base method
func (m *MockAthenaAPI) ListNamedQueriesPagesWithContext(arg0 aws.Context, arg1 *athena.ListNamedQueriesInput, arg2 func(*athena.ListNamedQueriesOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListNamedQueriesPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListNamedQueriesPagesWithContext indicates an expected call of ListNamedQueriesPagesWithContext
func (mr *MockAthenaAPIMockRecorder) ListNamedQueriesPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNamedQueriesPagesWithContext", reflect.TypeOf((*MockAthenaAPI)(nil).ListNamedQueriesPagesWithContext), varargs...)
}

// ListNamedQueriesRequest mocks base method
func (m *MockAthenaAPI) ListNamedQueriesRequest(arg0 *athena.ListNamedQueriesInput) (*request.Request, *athena.ListNamedQueriesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNamedQueriesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*athena.ListNamedQueriesOutput)
	return ret0, ret1
}

// ListNamedQueriesRequest indicates an expected call of ListNamedQueriesRequest
func (mr *MockAthenaAPIMockRecorder) ListNamedQueriesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNamedQueriesRequest", reflect.TypeOf((*MockAthenaAPI)(nil).ListNamedQueriesRequest), arg0)
}

// ListNamedQueriesWithContext mocks base method
func (m *MockAthenaAPI) ListNamedQueriesWithContext(arg0 aws.Context, arg1 *athena.ListNamedQueriesInput, arg2 ...request.Option) (*athena.ListNamedQueriesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListNamedQueriesWithContext", varargs...)
	ret0, _ := ret[0].(*athena.ListNamedQueriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNamedQueriesWithContext indicates an expected call of ListNamedQueriesWithContext
func (mr *MockAthenaAPIMockRecorder) ListNamedQueriesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNamedQueriesWithContext", reflect.TypeOf((*MockAthenaAPI)(nil).ListNamedQueriesWithContext), varargs...)
}

// ListQueryExecutions mocks base method
func (m *MockAthenaAPI) ListQueryExecutions(arg0 *athena.ListQueryExecutionsInput) (*athena.ListQueryExecutionsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListQueryExecutions", arg0)
	ret0, _ := ret[0].(*athena.ListQueryExecutionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListQueryExecutions indicates an expected call of ListQueryExecutions
func (mr *MockAthenaAPIMockRecorder) ListQueryExecutions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListQueryExecutions", reflect.TypeOf((*MockAthenaAPI)(nil).ListQueryExecutions), arg0)
}

// ListQueryExecutionsPages mocks base method
func (m *MockAthenaAPI) ListQueryExecutionsPages(arg0 *athena.ListQueryExecutionsInput, arg1 func(*athena.ListQueryExecutionsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListQueryExecutionsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListQueryExecutionsPages indicates an expected call of ListQueryExecutionsPages
func (mr *MockAthenaAPIMockRecorder) ListQueryExecutionsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListQueryExecutionsPages", reflect.TypeOf((*MockAthenaAPI)(nil).ListQueryExecutionsPages), arg0, arg1)
}

// ListQueryExecutionsPagesWithContext mocks base method
func (m *MockAthenaAPI) ListQueryExecutionsPagesWithContext(arg0 aws.Context, arg1 *athena.ListQueryExecutionsInput, arg2 func(*athena.ListQueryExecutionsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListQueryExecutionsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListQueryExecutionsPagesWithContext indicates an expected call of ListQueryExecutionsPagesWithContext
func (mr *MockAthenaAPIMockRecorder) ListQueryExecutionsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListQueryExecutionsPagesWithContext", reflect.TypeOf((*MockAthenaAPI)(nil).ListQueryExecutionsPagesWithContext), varargs...)
}

// ListQueryExecutionsRequest mocks base method
func (m *MockAthenaAPI) ListQueryExecutionsRequest(arg0 *athena.ListQueryExecutionsInput) (*request.Request, *athena.ListQueryExecutionsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListQueryExecutionsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*athena.ListQueryExecutionsOutput)
	return ret0, ret1
}

// ListQueryExecutionsRequest indicates an expected call of ListQueryExecutionsRequest
func (mr *MockAthenaAPIMockRecorder) ListQueryExecutionsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListQueryExecutionsRequest", reflect.TypeOf((*MockAthenaAPI)(nil).ListQueryExecutionsRequest), arg0)
}

// ListQueryExecutionsWithContext mocks base method
func (m *MockAthenaAPI) ListQueryExecutionsWithContext(arg0 aws.Context, arg1 *athena.ListQueryExecutionsInput, arg2 ...request.Option) (*athena.ListQueryExecutionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListQueryExecutionsWithContext", varargs...)
	ret0, _ := ret[0].(*athena.ListQueryExecutionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListQueryExecutionsWithContext indicates an expected call of ListQueryExecutionsWithContext
func (mr *MockAthenaAPIMockRecorder) ListQueryExecutionsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListQueryExecutionsWithContext", reflect.TypeOf((*MockAthenaAPI)(nil).ListQueryExecutionsWithContext), varargs...)
}

// StartQueryExecution mocks base method
func (m *MockAthenaAPI) StartQueryExecution(arg0 *athena.StartQueryExecutionInput) (*athena.StartQueryExecutionOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartQueryExecution", arg0)
	ret0, _ := ret[0].(*athena.StartQueryExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartQueryExecution indicates an expected call of StartQueryExecution
func (mr *MockAthenaAPIMockRecorder) StartQueryExecution(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartQueryExecution", reflect.TypeOf((*MockAthenaAPI)(nil).StartQueryExecution), arg0)
}

// StartQueryExecutionRequest mocks base method
func (m *MockAthenaAPI) StartQueryExecutionRequest(arg0 *athena.StartQueryExecutionInput) (*request.Request, *athena.StartQueryExecutionOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartQueryExecutionRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*athena.StartQueryExecutionOutput)
	return ret0, ret1
}

// StartQueryExecutionRequest indicates an expected call of StartQueryExecutionRequest
func (mr *MockAthenaAPIMockRecorder) StartQueryExecutionRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartQueryExecutionRequest", reflect.TypeOf((*MockAthenaAPI)(nil).StartQueryExecutionRequest), arg0)
}

// StartQueryExecutionWithContext mocks base method
func (m *MockAthenaAPI) StartQueryExecutionWithContext(arg0 aws.Context, arg1 *athena.StartQueryExecutionInput, arg2 ...request.Option) (*athena.StartQueryExecutionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StartQueryExecutionWithContext", varargs...)
	ret0, _ := ret[0].(*athena.StartQueryExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartQueryExecutionWithContext indicates an expected call of StartQueryExecutionWithContext
func (mr *MockAthenaAPIMockRecorder) StartQueryExecutionWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartQueryExecutionWithContext", reflect.TypeOf((*MockAthenaAPI)(nil).StartQueryExecutionWithContext), varargs...)
}

// StopQueryExecution mocks base method
func (m *MockAthenaAPI) StopQueryExecution(arg0 *athena.StopQueryExecutionInput) (*athena.StopQueryExecutionOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopQueryExecution", arg0)
	ret0, _ := ret[0].(*athena.StopQueryExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopQueryExecution indicates an expected call of StopQueryExecution
func (mr *MockAthenaAPIMockRecorder) StopQueryExecution(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopQueryExecution", reflect.TypeOf((*MockAthenaAPI)(nil).StopQueryExecution), arg0)
}

// StopQueryExecutionRequest mocks base method
func (m *MockAthenaAPI) StopQueryExecutionRequest(arg0 *athena.StopQueryExecutionInput) (*request.Request, *athena.StopQueryExecutionOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopQueryExecutionRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*athena.StopQueryExecutionOutput)
	return ret0, ret1
}

// StopQueryExecutionRequest indicates an expected call of StopQueryExecutionRequest
func (mr *MockAthenaAPIMockRecorder) StopQueryExecutionRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopQueryExecutionRequest", reflect.TypeOf((*MockAthenaAPI)(nil).StopQueryExecutionRequest), arg0)
}

// StopQueryExecutionWithContext mocks base method
func (m *MockAthenaAPI) StopQueryExecutionWithContext(arg0 aws.Context, arg1 *athena.StopQueryExecutionInput, arg2 ...request.Option) (*athena.StopQueryExecutionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StopQueryExecutionWithContext", varargs...)
	ret0, _ := ret[0].(*athena.StopQueryExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopQueryExecutionWithContext indicates an expected call of StopQueryExecutionWithContext
func (mr *MockAthenaAPIMockRecorder) StopQueryExecutionWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopQueryExecutionWithContext", reflect.TypeOf((*MockAthenaAPI)(nil).StopQueryExecutionWithContext), varargs...)
}
