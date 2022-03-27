// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/mobile/mobileiface (interfaces: MobileAPI)

// Package mobilemock is a generated GoMock package.
package mobilemock

import (
	context "context"
	reflect "reflect"

	request "github.com/aws/aws-sdk-go/aws/request"
	mobile "github.com/aws/aws-sdk-go/service/mobile"
	gomock "github.com/golang/mock/gomock"
)

// MockMobileAPI is a mock of MobileAPI interface.
type MockMobileAPI struct {
	ctrl     *gomock.Controller
	recorder *MockMobileAPIMockRecorder
}

// MockMobileAPIMockRecorder is the mock recorder for MockMobileAPI.
type MockMobileAPIMockRecorder struct {
	mock *MockMobileAPI
}

// NewMockMobileAPI creates a new mock instance.
func NewMockMobileAPI(ctrl *gomock.Controller) *MockMobileAPI {
	mock := &MockMobileAPI{ctrl: ctrl}
	mock.recorder = &MockMobileAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMobileAPI) EXPECT() *MockMobileAPIMockRecorder {
	return m.recorder
}

// CreateProject mocks base method.
func (m *MockMobileAPI) CreateProject(arg0 *mobile.CreateProjectInput) (*mobile.CreateProjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", arg0)
	ret0, _ := ret[0].(*mobile.CreateProjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockMobileAPIMockRecorder) CreateProject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockMobileAPI)(nil).CreateProject), arg0)
}

// CreateProjectRequest mocks base method.
func (m *MockMobileAPI) CreateProjectRequest(arg0 *mobile.CreateProjectInput) (*request.Request, *mobile.CreateProjectOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProjectRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mobile.CreateProjectOutput)
	return ret0, ret1
}

// CreateProjectRequest indicates an expected call of CreateProjectRequest.
func (mr *MockMobileAPIMockRecorder) CreateProjectRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProjectRequest", reflect.TypeOf((*MockMobileAPI)(nil).CreateProjectRequest), arg0)
}

// CreateProjectWithContext mocks base method.
func (m *MockMobileAPI) CreateProjectWithContext(arg0 context.Context, arg1 *mobile.CreateProjectInput, arg2 ...request.Option) (*mobile.CreateProjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateProjectWithContext", varargs...)
	ret0, _ := ret[0].(*mobile.CreateProjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProjectWithContext indicates an expected call of CreateProjectWithContext.
func (mr *MockMobileAPIMockRecorder) CreateProjectWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProjectWithContext", reflect.TypeOf((*MockMobileAPI)(nil).CreateProjectWithContext), varargs...)
}

// DeleteProject mocks base method.
func (m *MockMobileAPI) DeleteProject(arg0 *mobile.DeleteProjectInput) (*mobile.DeleteProjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProject", arg0)
	ret0, _ := ret[0].(*mobile.DeleteProjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProject indicates an expected call of DeleteProject.
func (mr *MockMobileAPIMockRecorder) DeleteProject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProject", reflect.TypeOf((*MockMobileAPI)(nil).DeleteProject), arg0)
}

// DeleteProjectRequest mocks base method.
func (m *MockMobileAPI) DeleteProjectRequest(arg0 *mobile.DeleteProjectInput) (*request.Request, *mobile.DeleteProjectOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProjectRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mobile.DeleteProjectOutput)
	return ret0, ret1
}

// DeleteProjectRequest indicates an expected call of DeleteProjectRequest.
func (mr *MockMobileAPIMockRecorder) DeleteProjectRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProjectRequest", reflect.TypeOf((*MockMobileAPI)(nil).DeleteProjectRequest), arg0)
}

// DeleteProjectWithContext mocks base method.
func (m *MockMobileAPI) DeleteProjectWithContext(arg0 context.Context, arg1 *mobile.DeleteProjectInput, arg2 ...request.Option) (*mobile.DeleteProjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteProjectWithContext", varargs...)
	ret0, _ := ret[0].(*mobile.DeleteProjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProjectWithContext indicates an expected call of DeleteProjectWithContext.
func (mr *MockMobileAPIMockRecorder) DeleteProjectWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProjectWithContext", reflect.TypeOf((*MockMobileAPI)(nil).DeleteProjectWithContext), varargs...)
}

// DescribeBundle mocks base method.
func (m *MockMobileAPI) DescribeBundle(arg0 *mobile.DescribeBundleInput) (*mobile.DescribeBundleOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeBundle", arg0)
	ret0, _ := ret[0].(*mobile.DescribeBundleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeBundle indicates an expected call of DescribeBundle.
func (mr *MockMobileAPIMockRecorder) DescribeBundle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeBundle", reflect.TypeOf((*MockMobileAPI)(nil).DescribeBundle), arg0)
}

// DescribeBundleRequest mocks base method.
func (m *MockMobileAPI) DescribeBundleRequest(arg0 *mobile.DescribeBundleInput) (*request.Request, *mobile.DescribeBundleOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeBundleRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mobile.DescribeBundleOutput)
	return ret0, ret1
}

// DescribeBundleRequest indicates an expected call of DescribeBundleRequest.
func (mr *MockMobileAPIMockRecorder) DescribeBundleRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeBundleRequest", reflect.TypeOf((*MockMobileAPI)(nil).DescribeBundleRequest), arg0)
}

// DescribeBundleWithContext mocks base method.
func (m *MockMobileAPI) DescribeBundleWithContext(arg0 context.Context, arg1 *mobile.DescribeBundleInput, arg2 ...request.Option) (*mobile.DescribeBundleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeBundleWithContext", varargs...)
	ret0, _ := ret[0].(*mobile.DescribeBundleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeBundleWithContext indicates an expected call of DescribeBundleWithContext.
func (mr *MockMobileAPIMockRecorder) DescribeBundleWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeBundleWithContext", reflect.TypeOf((*MockMobileAPI)(nil).DescribeBundleWithContext), varargs...)
}

// DescribeProject mocks base method.
func (m *MockMobileAPI) DescribeProject(arg0 *mobile.DescribeProjectInput) (*mobile.DescribeProjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeProject", arg0)
	ret0, _ := ret[0].(*mobile.DescribeProjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeProject indicates an expected call of DescribeProject.
func (mr *MockMobileAPIMockRecorder) DescribeProject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeProject", reflect.TypeOf((*MockMobileAPI)(nil).DescribeProject), arg0)
}

// DescribeProjectRequest mocks base method.
func (m *MockMobileAPI) DescribeProjectRequest(arg0 *mobile.DescribeProjectInput) (*request.Request, *mobile.DescribeProjectOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeProjectRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mobile.DescribeProjectOutput)
	return ret0, ret1
}

// DescribeProjectRequest indicates an expected call of DescribeProjectRequest.
func (mr *MockMobileAPIMockRecorder) DescribeProjectRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeProjectRequest", reflect.TypeOf((*MockMobileAPI)(nil).DescribeProjectRequest), arg0)
}

// DescribeProjectWithContext mocks base method.
func (m *MockMobileAPI) DescribeProjectWithContext(arg0 context.Context, arg1 *mobile.DescribeProjectInput, arg2 ...request.Option) (*mobile.DescribeProjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeProjectWithContext", varargs...)
	ret0, _ := ret[0].(*mobile.DescribeProjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeProjectWithContext indicates an expected call of DescribeProjectWithContext.
func (mr *MockMobileAPIMockRecorder) DescribeProjectWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeProjectWithContext", reflect.TypeOf((*MockMobileAPI)(nil).DescribeProjectWithContext), varargs...)
}

// ExportBundle mocks base method.
func (m *MockMobileAPI) ExportBundle(arg0 *mobile.ExportBundleInput) (*mobile.ExportBundleOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportBundle", arg0)
	ret0, _ := ret[0].(*mobile.ExportBundleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExportBundle indicates an expected call of ExportBundle.
func (mr *MockMobileAPIMockRecorder) ExportBundle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportBundle", reflect.TypeOf((*MockMobileAPI)(nil).ExportBundle), arg0)
}

// ExportBundleRequest mocks base method.
func (m *MockMobileAPI) ExportBundleRequest(arg0 *mobile.ExportBundleInput) (*request.Request, *mobile.ExportBundleOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportBundleRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mobile.ExportBundleOutput)
	return ret0, ret1
}

// ExportBundleRequest indicates an expected call of ExportBundleRequest.
func (mr *MockMobileAPIMockRecorder) ExportBundleRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportBundleRequest", reflect.TypeOf((*MockMobileAPI)(nil).ExportBundleRequest), arg0)
}

// ExportBundleWithContext mocks base method.
func (m *MockMobileAPI) ExportBundleWithContext(arg0 context.Context, arg1 *mobile.ExportBundleInput, arg2 ...request.Option) (*mobile.ExportBundleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExportBundleWithContext", varargs...)
	ret0, _ := ret[0].(*mobile.ExportBundleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExportBundleWithContext indicates an expected call of ExportBundleWithContext.
func (mr *MockMobileAPIMockRecorder) ExportBundleWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportBundleWithContext", reflect.TypeOf((*MockMobileAPI)(nil).ExportBundleWithContext), varargs...)
}

// ExportProject mocks base method.
func (m *MockMobileAPI) ExportProject(arg0 *mobile.ExportProjectInput) (*mobile.ExportProjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportProject", arg0)
	ret0, _ := ret[0].(*mobile.ExportProjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExportProject indicates an expected call of ExportProject.
func (mr *MockMobileAPIMockRecorder) ExportProject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportProject", reflect.TypeOf((*MockMobileAPI)(nil).ExportProject), arg0)
}

// ExportProjectRequest mocks base method.
func (m *MockMobileAPI) ExportProjectRequest(arg0 *mobile.ExportProjectInput) (*request.Request, *mobile.ExportProjectOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportProjectRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mobile.ExportProjectOutput)
	return ret0, ret1
}

// ExportProjectRequest indicates an expected call of ExportProjectRequest.
func (mr *MockMobileAPIMockRecorder) ExportProjectRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportProjectRequest", reflect.TypeOf((*MockMobileAPI)(nil).ExportProjectRequest), arg0)
}

// ExportProjectWithContext mocks base method.
func (m *MockMobileAPI) ExportProjectWithContext(arg0 context.Context, arg1 *mobile.ExportProjectInput, arg2 ...request.Option) (*mobile.ExportProjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExportProjectWithContext", varargs...)
	ret0, _ := ret[0].(*mobile.ExportProjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExportProjectWithContext indicates an expected call of ExportProjectWithContext.
func (mr *MockMobileAPIMockRecorder) ExportProjectWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportProjectWithContext", reflect.TypeOf((*MockMobileAPI)(nil).ExportProjectWithContext), varargs...)
}

// ListBundles mocks base method.
func (m *MockMobileAPI) ListBundles(arg0 *mobile.ListBundlesInput) (*mobile.ListBundlesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBundles", arg0)
	ret0, _ := ret[0].(*mobile.ListBundlesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBundles indicates an expected call of ListBundles.
func (mr *MockMobileAPIMockRecorder) ListBundles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBundles", reflect.TypeOf((*MockMobileAPI)(nil).ListBundles), arg0)
}

// ListBundlesPages mocks base method.
func (m *MockMobileAPI) ListBundlesPages(arg0 *mobile.ListBundlesInput, arg1 func(*mobile.ListBundlesOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBundlesPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListBundlesPages indicates an expected call of ListBundlesPages.
func (mr *MockMobileAPIMockRecorder) ListBundlesPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBundlesPages", reflect.TypeOf((*MockMobileAPI)(nil).ListBundlesPages), arg0, arg1)
}

// ListBundlesPagesWithContext mocks base method.
func (m *MockMobileAPI) ListBundlesPagesWithContext(arg0 context.Context, arg1 *mobile.ListBundlesInput, arg2 func(*mobile.ListBundlesOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBundlesPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListBundlesPagesWithContext indicates an expected call of ListBundlesPagesWithContext.
func (mr *MockMobileAPIMockRecorder) ListBundlesPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBundlesPagesWithContext", reflect.TypeOf((*MockMobileAPI)(nil).ListBundlesPagesWithContext), varargs...)
}

// ListBundlesRequest mocks base method.
func (m *MockMobileAPI) ListBundlesRequest(arg0 *mobile.ListBundlesInput) (*request.Request, *mobile.ListBundlesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBundlesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mobile.ListBundlesOutput)
	return ret0, ret1
}

// ListBundlesRequest indicates an expected call of ListBundlesRequest.
func (mr *MockMobileAPIMockRecorder) ListBundlesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBundlesRequest", reflect.TypeOf((*MockMobileAPI)(nil).ListBundlesRequest), arg0)
}

// ListBundlesWithContext mocks base method.
func (m *MockMobileAPI) ListBundlesWithContext(arg0 context.Context, arg1 *mobile.ListBundlesInput, arg2 ...request.Option) (*mobile.ListBundlesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBundlesWithContext", varargs...)
	ret0, _ := ret[0].(*mobile.ListBundlesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBundlesWithContext indicates an expected call of ListBundlesWithContext.
func (mr *MockMobileAPIMockRecorder) ListBundlesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBundlesWithContext", reflect.TypeOf((*MockMobileAPI)(nil).ListBundlesWithContext), varargs...)
}

// ListProjects mocks base method.
func (m *MockMobileAPI) ListProjects(arg0 *mobile.ListProjectsInput) (*mobile.ListProjectsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProjects", arg0)
	ret0, _ := ret[0].(*mobile.ListProjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjects indicates an expected call of ListProjects.
func (mr *MockMobileAPIMockRecorder) ListProjects(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjects", reflect.TypeOf((*MockMobileAPI)(nil).ListProjects), arg0)
}

// ListProjectsPages mocks base method.
func (m *MockMobileAPI) ListProjectsPages(arg0 *mobile.ListProjectsInput, arg1 func(*mobile.ListProjectsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProjectsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListProjectsPages indicates an expected call of ListProjectsPages.
func (mr *MockMobileAPIMockRecorder) ListProjectsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectsPages", reflect.TypeOf((*MockMobileAPI)(nil).ListProjectsPages), arg0, arg1)
}

// ListProjectsPagesWithContext mocks base method.
func (m *MockMobileAPI) ListProjectsPagesWithContext(arg0 context.Context, arg1 *mobile.ListProjectsInput, arg2 func(*mobile.ListProjectsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProjectsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListProjectsPagesWithContext indicates an expected call of ListProjectsPagesWithContext.
func (mr *MockMobileAPIMockRecorder) ListProjectsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectsPagesWithContext", reflect.TypeOf((*MockMobileAPI)(nil).ListProjectsPagesWithContext), varargs...)
}

// ListProjectsRequest mocks base method.
func (m *MockMobileAPI) ListProjectsRequest(arg0 *mobile.ListProjectsInput) (*request.Request, *mobile.ListProjectsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProjectsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mobile.ListProjectsOutput)
	return ret0, ret1
}

// ListProjectsRequest indicates an expected call of ListProjectsRequest.
func (mr *MockMobileAPIMockRecorder) ListProjectsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectsRequest", reflect.TypeOf((*MockMobileAPI)(nil).ListProjectsRequest), arg0)
}

// ListProjectsWithContext mocks base method.
func (m *MockMobileAPI) ListProjectsWithContext(arg0 context.Context, arg1 *mobile.ListProjectsInput, arg2 ...request.Option) (*mobile.ListProjectsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProjectsWithContext", varargs...)
	ret0, _ := ret[0].(*mobile.ListProjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjectsWithContext indicates an expected call of ListProjectsWithContext.
func (mr *MockMobileAPIMockRecorder) ListProjectsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectsWithContext", reflect.TypeOf((*MockMobileAPI)(nil).ListProjectsWithContext), varargs...)
}

// UpdateProject mocks base method.
func (m *MockMobileAPI) UpdateProject(arg0 *mobile.UpdateProjectInput) (*mobile.UpdateProjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProject", arg0)
	ret0, _ := ret[0].(*mobile.UpdateProjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProject indicates an expected call of UpdateProject.
func (mr *MockMobileAPIMockRecorder) UpdateProject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProject", reflect.TypeOf((*MockMobileAPI)(nil).UpdateProject), arg0)
}

// UpdateProjectRequest mocks base method.
func (m *MockMobileAPI) UpdateProjectRequest(arg0 *mobile.UpdateProjectInput) (*request.Request, *mobile.UpdateProjectOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProjectRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mobile.UpdateProjectOutput)
	return ret0, ret1
}

// UpdateProjectRequest indicates an expected call of UpdateProjectRequest.
func (mr *MockMobileAPIMockRecorder) UpdateProjectRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProjectRequest", reflect.TypeOf((*MockMobileAPI)(nil).UpdateProjectRequest), arg0)
}

// UpdateProjectWithContext mocks base method.
func (m *MockMobileAPI) UpdateProjectWithContext(arg0 context.Context, arg1 *mobile.UpdateProjectInput, arg2 ...request.Option) (*mobile.UpdateProjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateProjectWithContext", varargs...)
	ret0, _ := ret[0].(*mobile.UpdateProjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProjectWithContext indicates an expected call of UpdateProjectWithContext.
func (mr *MockMobileAPIMockRecorder) UpdateProjectWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProjectWithContext", reflect.TypeOf((*MockMobileAPI)(nil).UpdateProjectWithContext), varargs...)
}
