// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/costandusagereportservice/costandusagereportserviceiface (interfaces: CostandUsageReportServiceAPI)

// Package costandusagereportservicemock is a generated GoMock package.
package costandusagereportservicemock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	costandusagereportservice "github.com/aws/aws-sdk-go/service/costandusagereportservice"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCostandUsageReportServiceAPI is a mock of CostandUsageReportServiceAPI interface
type MockCostandUsageReportServiceAPI struct {
	ctrl     *gomock.Controller
	recorder *MockCostandUsageReportServiceAPIMockRecorder
}

// MockCostandUsageReportServiceAPIMockRecorder is the mock recorder for MockCostandUsageReportServiceAPI
type MockCostandUsageReportServiceAPIMockRecorder struct {
	mock *MockCostandUsageReportServiceAPI
}

// NewMockCostandUsageReportServiceAPI creates a new mock instance
func NewMockCostandUsageReportServiceAPI(ctrl *gomock.Controller) *MockCostandUsageReportServiceAPI {
	mock := &MockCostandUsageReportServiceAPI{ctrl: ctrl}
	mock.recorder = &MockCostandUsageReportServiceAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCostandUsageReportServiceAPI) EXPECT() *MockCostandUsageReportServiceAPIMockRecorder {
	return m.recorder
}

// DeleteReportDefinition mocks base method
func (m *MockCostandUsageReportServiceAPI) DeleteReportDefinition(arg0 *costandusagereportservice.DeleteReportDefinitionInput) (*costandusagereportservice.DeleteReportDefinitionOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteReportDefinition", arg0)
	ret0, _ := ret[0].(*costandusagereportservice.DeleteReportDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteReportDefinition indicates an expected call of DeleteReportDefinition
func (mr *MockCostandUsageReportServiceAPIMockRecorder) DeleteReportDefinition(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReportDefinition", reflect.TypeOf((*MockCostandUsageReportServiceAPI)(nil).DeleteReportDefinition), arg0)
}

// DeleteReportDefinitionRequest mocks base method
func (m *MockCostandUsageReportServiceAPI) DeleteReportDefinitionRequest(arg0 *costandusagereportservice.DeleteReportDefinitionInput) (*request.Request, *costandusagereportservice.DeleteReportDefinitionOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteReportDefinitionRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*costandusagereportservice.DeleteReportDefinitionOutput)
	return ret0, ret1
}

// DeleteReportDefinitionRequest indicates an expected call of DeleteReportDefinitionRequest
func (mr *MockCostandUsageReportServiceAPIMockRecorder) DeleteReportDefinitionRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReportDefinitionRequest", reflect.TypeOf((*MockCostandUsageReportServiceAPI)(nil).DeleteReportDefinitionRequest), arg0)
}

// DeleteReportDefinitionWithContext mocks base method
func (m *MockCostandUsageReportServiceAPI) DeleteReportDefinitionWithContext(arg0 context.Context, arg1 *costandusagereportservice.DeleteReportDefinitionInput, arg2 ...request.Option) (*costandusagereportservice.DeleteReportDefinitionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteReportDefinitionWithContext", varargs...)
	ret0, _ := ret[0].(*costandusagereportservice.DeleteReportDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteReportDefinitionWithContext indicates an expected call of DeleteReportDefinitionWithContext
func (mr *MockCostandUsageReportServiceAPIMockRecorder) DeleteReportDefinitionWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReportDefinitionWithContext", reflect.TypeOf((*MockCostandUsageReportServiceAPI)(nil).DeleteReportDefinitionWithContext), varargs...)
}

// DescribeReportDefinitions mocks base method
func (m *MockCostandUsageReportServiceAPI) DescribeReportDefinitions(arg0 *costandusagereportservice.DescribeReportDefinitionsInput) (*costandusagereportservice.DescribeReportDefinitionsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeReportDefinitions", arg0)
	ret0, _ := ret[0].(*costandusagereportservice.DescribeReportDefinitionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeReportDefinitions indicates an expected call of DescribeReportDefinitions
func (mr *MockCostandUsageReportServiceAPIMockRecorder) DescribeReportDefinitions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeReportDefinitions", reflect.TypeOf((*MockCostandUsageReportServiceAPI)(nil).DescribeReportDefinitions), arg0)
}

// DescribeReportDefinitionsPages mocks base method
func (m *MockCostandUsageReportServiceAPI) DescribeReportDefinitionsPages(arg0 *costandusagereportservice.DescribeReportDefinitionsInput, arg1 func(*costandusagereportservice.DescribeReportDefinitionsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeReportDefinitionsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DescribeReportDefinitionsPages indicates an expected call of DescribeReportDefinitionsPages
func (mr *MockCostandUsageReportServiceAPIMockRecorder) DescribeReportDefinitionsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeReportDefinitionsPages", reflect.TypeOf((*MockCostandUsageReportServiceAPI)(nil).DescribeReportDefinitionsPages), arg0, arg1)
}

// DescribeReportDefinitionsPagesWithContext mocks base method
func (m *MockCostandUsageReportServiceAPI) DescribeReportDefinitionsPagesWithContext(arg0 context.Context, arg1 *costandusagereportservice.DescribeReportDefinitionsInput, arg2 func(*costandusagereportservice.DescribeReportDefinitionsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeReportDefinitionsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DescribeReportDefinitionsPagesWithContext indicates an expected call of DescribeReportDefinitionsPagesWithContext
func (mr *MockCostandUsageReportServiceAPIMockRecorder) DescribeReportDefinitionsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeReportDefinitionsPagesWithContext", reflect.TypeOf((*MockCostandUsageReportServiceAPI)(nil).DescribeReportDefinitionsPagesWithContext), varargs...)
}

// DescribeReportDefinitionsRequest mocks base method
func (m *MockCostandUsageReportServiceAPI) DescribeReportDefinitionsRequest(arg0 *costandusagereportservice.DescribeReportDefinitionsInput) (*request.Request, *costandusagereportservice.DescribeReportDefinitionsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeReportDefinitionsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*costandusagereportservice.DescribeReportDefinitionsOutput)
	return ret0, ret1
}

// DescribeReportDefinitionsRequest indicates an expected call of DescribeReportDefinitionsRequest
func (mr *MockCostandUsageReportServiceAPIMockRecorder) DescribeReportDefinitionsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeReportDefinitionsRequest", reflect.TypeOf((*MockCostandUsageReportServiceAPI)(nil).DescribeReportDefinitionsRequest), arg0)
}

// DescribeReportDefinitionsWithContext mocks base method
func (m *MockCostandUsageReportServiceAPI) DescribeReportDefinitionsWithContext(arg0 context.Context, arg1 *costandusagereportservice.DescribeReportDefinitionsInput, arg2 ...request.Option) (*costandusagereportservice.DescribeReportDefinitionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeReportDefinitionsWithContext", varargs...)
	ret0, _ := ret[0].(*costandusagereportservice.DescribeReportDefinitionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeReportDefinitionsWithContext indicates an expected call of DescribeReportDefinitionsWithContext
func (mr *MockCostandUsageReportServiceAPIMockRecorder) DescribeReportDefinitionsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeReportDefinitionsWithContext", reflect.TypeOf((*MockCostandUsageReportServiceAPI)(nil).DescribeReportDefinitionsWithContext), varargs...)
}

// ListTagsForResource mocks base method
func (m *MockCostandUsageReportServiceAPI) ListTagsForResource(arg0 *costandusagereportservice.ListTagsForResourceInput) (*costandusagereportservice.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForResource", arg0)
	ret0, _ := ret[0].(*costandusagereportservice.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource
func (mr *MockCostandUsageReportServiceAPIMockRecorder) ListTagsForResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockCostandUsageReportServiceAPI)(nil).ListTagsForResource), arg0)
}

// ListTagsForResourceRequest mocks base method
func (m *MockCostandUsageReportServiceAPI) ListTagsForResourceRequest(arg0 *costandusagereportservice.ListTagsForResourceInput) (*request.Request, *costandusagereportservice.ListTagsForResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*costandusagereportservice.ListTagsForResourceOutput)
	return ret0, ret1
}

// ListTagsForResourceRequest indicates an expected call of ListTagsForResourceRequest
func (mr *MockCostandUsageReportServiceAPIMockRecorder) ListTagsForResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResourceRequest", reflect.TypeOf((*MockCostandUsageReportServiceAPI)(nil).ListTagsForResourceRequest), arg0)
}

// ListTagsForResourceWithContext mocks base method
func (m *MockCostandUsageReportServiceAPI) ListTagsForResourceWithContext(arg0 context.Context, arg1 *costandusagereportservice.ListTagsForResourceInput, arg2 ...request.Option) (*costandusagereportservice.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResourceWithContext", varargs...)
	ret0, _ := ret[0].(*costandusagereportservice.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResourceWithContext indicates an expected call of ListTagsForResourceWithContext
func (mr *MockCostandUsageReportServiceAPIMockRecorder) ListTagsForResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResourceWithContext", reflect.TypeOf((*MockCostandUsageReportServiceAPI)(nil).ListTagsForResourceWithContext), varargs...)
}

// ModifyReportDefinition mocks base method
func (m *MockCostandUsageReportServiceAPI) ModifyReportDefinition(arg0 *costandusagereportservice.ModifyReportDefinitionInput) (*costandusagereportservice.ModifyReportDefinitionOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyReportDefinition", arg0)
	ret0, _ := ret[0].(*costandusagereportservice.ModifyReportDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyReportDefinition indicates an expected call of ModifyReportDefinition
func (mr *MockCostandUsageReportServiceAPIMockRecorder) ModifyReportDefinition(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyReportDefinition", reflect.TypeOf((*MockCostandUsageReportServiceAPI)(nil).ModifyReportDefinition), arg0)
}

// ModifyReportDefinitionRequest mocks base method
func (m *MockCostandUsageReportServiceAPI) ModifyReportDefinitionRequest(arg0 *costandusagereportservice.ModifyReportDefinitionInput) (*request.Request, *costandusagereportservice.ModifyReportDefinitionOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyReportDefinitionRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*costandusagereportservice.ModifyReportDefinitionOutput)
	return ret0, ret1
}

// ModifyReportDefinitionRequest indicates an expected call of ModifyReportDefinitionRequest
func (mr *MockCostandUsageReportServiceAPIMockRecorder) ModifyReportDefinitionRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyReportDefinitionRequest", reflect.TypeOf((*MockCostandUsageReportServiceAPI)(nil).ModifyReportDefinitionRequest), arg0)
}

// ModifyReportDefinitionWithContext mocks base method
func (m *MockCostandUsageReportServiceAPI) ModifyReportDefinitionWithContext(arg0 context.Context, arg1 *costandusagereportservice.ModifyReportDefinitionInput, arg2 ...request.Option) (*costandusagereportservice.ModifyReportDefinitionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ModifyReportDefinitionWithContext", varargs...)
	ret0, _ := ret[0].(*costandusagereportservice.ModifyReportDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyReportDefinitionWithContext indicates an expected call of ModifyReportDefinitionWithContext
func (mr *MockCostandUsageReportServiceAPIMockRecorder) ModifyReportDefinitionWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyReportDefinitionWithContext", reflect.TypeOf((*MockCostandUsageReportServiceAPI)(nil).ModifyReportDefinitionWithContext), varargs...)
}

// PutReportDefinition mocks base method
func (m *MockCostandUsageReportServiceAPI) PutReportDefinition(arg0 *costandusagereportservice.PutReportDefinitionInput) (*costandusagereportservice.PutReportDefinitionOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutReportDefinition", arg0)
	ret0, _ := ret[0].(*costandusagereportservice.PutReportDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutReportDefinition indicates an expected call of PutReportDefinition
func (mr *MockCostandUsageReportServiceAPIMockRecorder) PutReportDefinition(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutReportDefinition", reflect.TypeOf((*MockCostandUsageReportServiceAPI)(nil).PutReportDefinition), arg0)
}

// PutReportDefinitionRequest mocks base method
func (m *MockCostandUsageReportServiceAPI) PutReportDefinitionRequest(arg0 *costandusagereportservice.PutReportDefinitionInput) (*request.Request, *costandusagereportservice.PutReportDefinitionOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutReportDefinitionRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*costandusagereportservice.PutReportDefinitionOutput)
	return ret0, ret1
}

// PutReportDefinitionRequest indicates an expected call of PutReportDefinitionRequest
func (mr *MockCostandUsageReportServiceAPIMockRecorder) PutReportDefinitionRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutReportDefinitionRequest", reflect.TypeOf((*MockCostandUsageReportServiceAPI)(nil).PutReportDefinitionRequest), arg0)
}

// PutReportDefinitionWithContext mocks base method
func (m *MockCostandUsageReportServiceAPI) PutReportDefinitionWithContext(arg0 context.Context, arg1 *costandusagereportservice.PutReportDefinitionInput, arg2 ...request.Option) (*costandusagereportservice.PutReportDefinitionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutReportDefinitionWithContext", varargs...)
	ret0, _ := ret[0].(*costandusagereportservice.PutReportDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutReportDefinitionWithContext indicates an expected call of PutReportDefinitionWithContext
func (mr *MockCostandUsageReportServiceAPIMockRecorder) PutReportDefinitionWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutReportDefinitionWithContext", reflect.TypeOf((*MockCostandUsageReportServiceAPI)(nil).PutReportDefinitionWithContext), varargs...)
}

// TagResource mocks base method
func (m *MockCostandUsageReportServiceAPI) TagResource(arg0 *costandusagereportservice.TagResourceInput) (*costandusagereportservice.TagResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagResource", arg0)
	ret0, _ := ret[0].(*costandusagereportservice.TagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagResource indicates an expected call of TagResource
func (mr *MockCostandUsageReportServiceAPIMockRecorder) TagResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResource", reflect.TypeOf((*MockCostandUsageReportServiceAPI)(nil).TagResource), arg0)
}

// TagResourceRequest mocks base method
func (m *MockCostandUsageReportServiceAPI) TagResourceRequest(arg0 *costandusagereportservice.TagResourceInput) (*request.Request, *costandusagereportservice.TagResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*costandusagereportservice.TagResourceOutput)
	return ret0, ret1
}

// TagResourceRequest indicates an expected call of TagResourceRequest
func (mr *MockCostandUsageReportServiceAPIMockRecorder) TagResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResourceRequest", reflect.TypeOf((*MockCostandUsageReportServiceAPI)(nil).TagResourceRequest), arg0)
}

// TagResourceWithContext mocks base method
func (m *MockCostandUsageReportServiceAPI) TagResourceWithContext(arg0 context.Context, arg1 *costandusagereportservice.TagResourceInput, arg2 ...request.Option) (*costandusagereportservice.TagResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TagResourceWithContext", varargs...)
	ret0, _ := ret[0].(*costandusagereportservice.TagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagResourceWithContext indicates an expected call of TagResourceWithContext
func (mr *MockCostandUsageReportServiceAPIMockRecorder) TagResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResourceWithContext", reflect.TypeOf((*MockCostandUsageReportServiceAPI)(nil).TagResourceWithContext), varargs...)
}

// UntagResource mocks base method
func (m *MockCostandUsageReportServiceAPI) UntagResource(arg0 *costandusagereportservice.UntagResourceInput) (*costandusagereportservice.UntagResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagResource", arg0)
	ret0, _ := ret[0].(*costandusagereportservice.UntagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagResource indicates an expected call of UntagResource
func (mr *MockCostandUsageReportServiceAPIMockRecorder) UntagResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResource", reflect.TypeOf((*MockCostandUsageReportServiceAPI)(nil).UntagResource), arg0)
}

// UntagResourceRequest mocks base method
func (m *MockCostandUsageReportServiceAPI) UntagResourceRequest(arg0 *costandusagereportservice.UntagResourceInput) (*request.Request, *costandusagereportservice.UntagResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*costandusagereportservice.UntagResourceOutput)
	return ret0, ret1
}

// UntagResourceRequest indicates an expected call of UntagResourceRequest
func (mr *MockCostandUsageReportServiceAPIMockRecorder) UntagResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResourceRequest", reflect.TypeOf((*MockCostandUsageReportServiceAPI)(nil).UntagResourceRequest), arg0)
}

// UntagResourceWithContext mocks base method
func (m *MockCostandUsageReportServiceAPI) UntagResourceWithContext(arg0 context.Context, arg1 *costandusagereportservice.UntagResourceInput, arg2 ...request.Option) (*costandusagereportservice.UntagResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UntagResourceWithContext", varargs...)
	ret0, _ := ret[0].(*costandusagereportservice.UntagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagResourceWithContext indicates an expected call of UntagResourceWithContext
func (mr *MockCostandUsageReportServiceAPIMockRecorder) UntagResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResourceWithContext", reflect.TypeOf((*MockCostandUsageReportServiceAPI)(nil).UntagResourceWithContext), varargs...)
}
