// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/cloudsearchdomain/cloudsearchdomainiface (interfaces: CloudSearchDomainAPI)

// Package cloudsearchdomainmock is a generated GoMock package.
package cloudsearchdomainmock

import (
	aws "github.com/aws/aws-sdk-go/aws"
	request "github.com/aws/aws-sdk-go/aws/request"
	cloudsearchdomain "github.com/aws/aws-sdk-go/service/cloudsearchdomain"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCloudSearchDomainAPI is a mock of CloudSearchDomainAPI interface
type MockCloudSearchDomainAPI struct {
	ctrl     *gomock.Controller
	recorder *MockCloudSearchDomainAPIMockRecorder
}

// MockCloudSearchDomainAPIMockRecorder is the mock recorder for MockCloudSearchDomainAPI
type MockCloudSearchDomainAPIMockRecorder struct {
	mock *MockCloudSearchDomainAPI
}

// NewMockCloudSearchDomainAPI creates a new mock instance
func NewMockCloudSearchDomainAPI(ctrl *gomock.Controller) *MockCloudSearchDomainAPI {
	mock := &MockCloudSearchDomainAPI{ctrl: ctrl}
	mock.recorder = &MockCloudSearchDomainAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCloudSearchDomainAPI) EXPECT() *MockCloudSearchDomainAPIMockRecorder {
	return m.recorder
}

// Search mocks base method
func (m *MockCloudSearchDomainAPI) Search(arg0 *cloudsearchdomain.SearchInput) (*cloudsearchdomain.SearchOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", arg0)
	ret0, _ := ret[0].(*cloudsearchdomain.SearchOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockCloudSearchDomainAPIMockRecorder) Search(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockCloudSearchDomainAPI)(nil).Search), arg0)
}

// SearchRequest mocks base method
func (m *MockCloudSearchDomainAPI) SearchRequest(arg0 *cloudsearchdomain.SearchInput) (*request.Request, *cloudsearchdomain.SearchOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudsearchdomain.SearchOutput)
	return ret0, ret1
}

// SearchRequest indicates an expected call of SearchRequest
func (mr *MockCloudSearchDomainAPIMockRecorder) SearchRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchRequest", reflect.TypeOf((*MockCloudSearchDomainAPI)(nil).SearchRequest), arg0)
}

// SearchWithContext mocks base method
func (m *MockCloudSearchDomainAPI) SearchWithContext(arg0 aws.Context, arg1 *cloudsearchdomain.SearchInput, arg2 ...request.Option) (*cloudsearchdomain.SearchOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchWithContext", varargs...)
	ret0, _ := ret[0].(*cloudsearchdomain.SearchOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchWithContext indicates an expected call of SearchWithContext
func (mr *MockCloudSearchDomainAPIMockRecorder) SearchWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchWithContext", reflect.TypeOf((*MockCloudSearchDomainAPI)(nil).SearchWithContext), varargs...)
}

// Suggest mocks base method
func (m *MockCloudSearchDomainAPI) Suggest(arg0 *cloudsearchdomain.SuggestInput) (*cloudsearchdomain.SuggestOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Suggest", arg0)
	ret0, _ := ret[0].(*cloudsearchdomain.SuggestOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Suggest indicates an expected call of Suggest
func (mr *MockCloudSearchDomainAPIMockRecorder) Suggest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Suggest", reflect.TypeOf((*MockCloudSearchDomainAPI)(nil).Suggest), arg0)
}

// SuggestRequest mocks base method
func (m *MockCloudSearchDomainAPI) SuggestRequest(arg0 *cloudsearchdomain.SuggestInput) (*request.Request, *cloudsearchdomain.SuggestOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SuggestRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudsearchdomain.SuggestOutput)
	return ret0, ret1
}

// SuggestRequest indicates an expected call of SuggestRequest
func (mr *MockCloudSearchDomainAPIMockRecorder) SuggestRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuggestRequest", reflect.TypeOf((*MockCloudSearchDomainAPI)(nil).SuggestRequest), arg0)
}

// SuggestWithContext mocks base method
func (m *MockCloudSearchDomainAPI) SuggestWithContext(arg0 aws.Context, arg1 *cloudsearchdomain.SuggestInput, arg2 ...request.Option) (*cloudsearchdomain.SuggestOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SuggestWithContext", varargs...)
	ret0, _ := ret[0].(*cloudsearchdomain.SuggestOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SuggestWithContext indicates an expected call of SuggestWithContext
func (mr *MockCloudSearchDomainAPIMockRecorder) SuggestWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuggestWithContext", reflect.TypeOf((*MockCloudSearchDomainAPI)(nil).SuggestWithContext), varargs...)
}

// UploadDocuments mocks base method
func (m *MockCloudSearchDomainAPI) UploadDocuments(arg0 *cloudsearchdomain.UploadDocumentsInput) (*cloudsearchdomain.UploadDocumentsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadDocuments", arg0)
	ret0, _ := ret[0].(*cloudsearchdomain.UploadDocumentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadDocuments indicates an expected call of UploadDocuments
func (mr *MockCloudSearchDomainAPIMockRecorder) UploadDocuments(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadDocuments", reflect.TypeOf((*MockCloudSearchDomainAPI)(nil).UploadDocuments), arg0)
}

// UploadDocumentsRequest mocks base method
func (m *MockCloudSearchDomainAPI) UploadDocumentsRequest(arg0 *cloudsearchdomain.UploadDocumentsInput) (*request.Request, *cloudsearchdomain.UploadDocumentsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadDocumentsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudsearchdomain.UploadDocumentsOutput)
	return ret0, ret1
}

// UploadDocumentsRequest indicates an expected call of UploadDocumentsRequest
func (mr *MockCloudSearchDomainAPIMockRecorder) UploadDocumentsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadDocumentsRequest", reflect.TypeOf((*MockCloudSearchDomainAPI)(nil).UploadDocumentsRequest), arg0)
}

// UploadDocumentsWithContext mocks base method
func (m *MockCloudSearchDomainAPI) UploadDocumentsWithContext(arg0 aws.Context, arg1 *cloudsearchdomain.UploadDocumentsInput, arg2 ...request.Option) (*cloudsearchdomain.UploadDocumentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UploadDocumentsWithContext", varargs...)
	ret0, _ := ret[0].(*cloudsearchdomain.UploadDocumentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadDocumentsWithContext indicates an expected call of UploadDocumentsWithContext
func (mr *MockCloudSearchDomainAPIMockRecorder) UploadDocumentsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadDocumentsWithContext", reflect.TypeOf((*MockCloudSearchDomainAPI)(nil).UploadDocumentsWithContext), varargs...)
}
