// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/s3control/s3controliface (interfaces: S3ControlAPI)

// Package s3controlmock is a generated GoMock package.
package s3controlmock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	s3control "github.com/aws/aws-sdk-go/service/s3control"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockS3ControlAPI is a mock of S3ControlAPI interface
type MockS3ControlAPI struct {
	ctrl     *gomock.Controller
	recorder *MockS3ControlAPIMockRecorder
}

// MockS3ControlAPIMockRecorder is the mock recorder for MockS3ControlAPI
type MockS3ControlAPIMockRecorder struct {
	mock *MockS3ControlAPI
}

// NewMockS3ControlAPI creates a new mock instance
func NewMockS3ControlAPI(ctrl *gomock.Controller) *MockS3ControlAPI {
	mock := &MockS3ControlAPI{ctrl: ctrl}
	mock.recorder = &MockS3ControlAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockS3ControlAPI) EXPECT() *MockS3ControlAPIMockRecorder {
	return m.recorder
}

// CreateJob mocks base method
func (m *MockS3ControlAPI) CreateJob(arg0 *s3control.CreateJobInput) (*s3control.CreateJobOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJob", arg0)
	ret0, _ := ret[0].(*s3control.CreateJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJob indicates an expected call of CreateJob
func (mr *MockS3ControlAPIMockRecorder) CreateJob(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJob", reflect.TypeOf((*MockS3ControlAPI)(nil).CreateJob), arg0)
}

// CreateJobRequest mocks base method
func (m *MockS3ControlAPI) CreateJobRequest(arg0 *s3control.CreateJobInput) (*request.Request, *s3control.CreateJobOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJobRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*s3control.CreateJobOutput)
	return ret0, ret1
}

// CreateJobRequest indicates an expected call of CreateJobRequest
func (mr *MockS3ControlAPIMockRecorder) CreateJobRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJobRequest", reflect.TypeOf((*MockS3ControlAPI)(nil).CreateJobRequest), arg0)
}

// CreateJobWithContext mocks base method
func (m *MockS3ControlAPI) CreateJobWithContext(arg0 context.Context, arg1 *s3control.CreateJobInput, arg2 ...request.Option) (*s3control.CreateJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateJobWithContext", varargs...)
	ret0, _ := ret[0].(*s3control.CreateJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJobWithContext indicates an expected call of CreateJobWithContext
func (mr *MockS3ControlAPIMockRecorder) CreateJobWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJobWithContext", reflect.TypeOf((*MockS3ControlAPI)(nil).CreateJobWithContext), varargs...)
}

// DeletePublicAccessBlock mocks base method
func (m *MockS3ControlAPI) DeletePublicAccessBlock(arg0 *s3control.DeletePublicAccessBlockInput) (*s3control.DeletePublicAccessBlockOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePublicAccessBlock", arg0)
	ret0, _ := ret[0].(*s3control.DeletePublicAccessBlockOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePublicAccessBlock indicates an expected call of DeletePublicAccessBlock
func (mr *MockS3ControlAPIMockRecorder) DeletePublicAccessBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePublicAccessBlock", reflect.TypeOf((*MockS3ControlAPI)(nil).DeletePublicAccessBlock), arg0)
}

// DeletePublicAccessBlockRequest mocks base method
func (m *MockS3ControlAPI) DeletePublicAccessBlockRequest(arg0 *s3control.DeletePublicAccessBlockInput) (*request.Request, *s3control.DeletePublicAccessBlockOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePublicAccessBlockRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*s3control.DeletePublicAccessBlockOutput)
	return ret0, ret1
}

// DeletePublicAccessBlockRequest indicates an expected call of DeletePublicAccessBlockRequest
func (mr *MockS3ControlAPIMockRecorder) DeletePublicAccessBlockRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePublicAccessBlockRequest", reflect.TypeOf((*MockS3ControlAPI)(nil).DeletePublicAccessBlockRequest), arg0)
}

// DeletePublicAccessBlockWithContext mocks base method
func (m *MockS3ControlAPI) DeletePublicAccessBlockWithContext(arg0 context.Context, arg1 *s3control.DeletePublicAccessBlockInput, arg2 ...request.Option) (*s3control.DeletePublicAccessBlockOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeletePublicAccessBlockWithContext", varargs...)
	ret0, _ := ret[0].(*s3control.DeletePublicAccessBlockOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePublicAccessBlockWithContext indicates an expected call of DeletePublicAccessBlockWithContext
func (mr *MockS3ControlAPIMockRecorder) DeletePublicAccessBlockWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePublicAccessBlockWithContext", reflect.TypeOf((*MockS3ControlAPI)(nil).DeletePublicAccessBlockWithContext), varargs...)
}

// DescribeJob mocks base method
func (m *MockS3ControlAPI) DescribeJob(arg0 *s3control.DescribeJobInput) (*s3control.DescribeJobOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeJob", arg0)
	ret0, _ := ret[0].(*s3control.DescribeJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeJob indicates an expected call of DescribeJob
func (mr *MockS3ControlAPIMockRecorder) DescribeJob(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeJob", reflect.TypeOf((*MockS3ControlAPI)(nil).DescribeJob), arg0)
}

// DescribeJobRequest mocks base method
func (m *MockS3ControlAPI) DescribeJobRequest(arg0 *s3control.DescribeJobInput) (*request.Request, *s3control.DescribeJobOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeJobRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*s3control.DescribeJobOutput)
	return ret0, ret1
}

// DescribeJobRequest indicates an expected call of DescribeJobRequest
func (mr *MockS3ControlAPIMockRecorder) DescribeJobRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeJobRequest", reflect.TypeOf((*MockS3ControlAPI)(nil).DescribeJobRequest), arg0)
}

// DescribeJobWithContext mocks base method
func (m *MockS3ControlAPI) DescribeJobWithContext(arg0 context.Context, arg1 *s3control.DescribeJobInput, arg2 ...request.Option) (*s3control.DescribeJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeJobWithContext", varargs...)
	ret0, _ := ret[0].(*s3control.DescribeJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeJobWithContext indicates an expected call of DescribeJobWithContext
func (mr *MockS3ControlAPIMockRecorder) DescribeJobWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeJobWithContext", reflect.TypeOf((*MockS3ControlAPI)(nil).DescribeJobWithContext), varargs...)
}

// GetPublicAccessBlock mocks base method
func (m *MockS3ControlAPI) GetPublicAccessBlock(arg0 *s3control.GetPublicAccessBlockInput) (*s3control.GetPublicAccessBlockOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicAccessBlock", arg0)
	ret0, _ := ret[0].(*s3control.GetPublicAccessBlockOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicAccessBlock indicates an expected call of GetPublicAccessBlock
func (mr *MockS3ControlAPIMockRecorder) GetPublicAccessBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicAccessBlock", reflect.TypeOf((*MockS3ControlAPI)(nil).GetPublicAccessBlock), arg0)
}

// GetPublicAccessBlockRequest mocks base method
func (m *MockS3ControlAPI) GetPublicAccessBlockRequest(arg0 *s3control.GetPublicAccessBlockInput) (*request.Request, *s3control.GetPublicAccessBlockOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicAccessBlockRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*s3control.GetPublicAccessBlockOutput)
	return ret0, ret1
}

// GetPublicAccessBlockRequest indicates an expected call of GetPublicAccessBlockRequest
func (mr *MockS3ControlAPIMockRecorder) GetPublicAccessBlockRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicAccessBlockRequest", reflect.TypeOf((*MockS3ControlAPI)(nil).GetPublicAccessBlockRequest), arg0)
}

// GetPublicAccessBlockWithContext mocks base method
func (m *MockS3ControlAPI) GetPublicAccessBlockWithContext(arg0 context.Context, arg1 *s3control.GetPublicAccessBlockInput, arg2 ...request.Option) (*s3control.GetPublicAccessBlockOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPublicAccessBlockWithContext", varargs...)
	ret0, _ := ret[0].(*s3control.GetPublicAccessBlockOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicAccessBlockWithContext indicates an expected call of GetPublicAccessBlockWithContext
func (mr *MockS3ControlAPIMockRecorder) GetPublicAccessBlockWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicAccessBlockWithContext", reflect.TypeOf((*MockS3ControlAPI)(nil).GetPublicAccessBlockWithContext), varargs...)
}

// ListJobs mocks base method
func (m *MockS3ControlAPI) ListJobs(arg0 *s3control.ListJobsInput) (*s3control.ListJobsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListJobs", arg0)
	ret0, _ := ret[0].(*s3control.ListJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListJobs indicates an expected call of ListJobs
func (mr *MockS3ControlAPIMockRecorder) ListJobs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListJobs", reflect.TypeOf((*MockS3ControlAPI)(nil).ListJobs), arg0)
}

// ListJobsPages mocks base method
func (m *MockS3ControlAPI) ListJobsPages(arg0 *s3control.ListJobsInput, arg1 func(*s3control.ListJobsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListJobsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListJobsPages indicates an expected call of ListJobsPages
func (mr *MockS3ControlAPIMockRecorder) ListJobsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListJobsPages", reflect.TypeOf((*MockS3ControlAPI)(nil).ListJobsPages), arg0, arg1)
}

// ListJobsPagesWithContext mocks base method
func (m *MockS3ControlAPI) ListJobsPagesWithContext(arg0 context.Context, arg1 *s3control.ListJobsInput, arg2 func(*s3control.ListJobsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListJobsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListJobsPagesWithContext indicates an expected call of ListJobsPagesWithContext
func (mr *MockS3ControlAPIMockRecorder) ListJobsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListJobsPagesWithContext", reflect.TypeOf((*MockS3ControlAPI)(nil).ListJobsPagesWithContext), varargs...)
}

// ListJobsRequest mocks base method
func (m *MockS3ControlAPI) ListJobsRequest(arg0 *s3control.ListJobsInput) (*request.Request, *s3control.ListJobsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListJobsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*s3control.ListJobsOutput)
	return ret0, ret1
}

// ListJobsRequest indicates an expected call of ListJobsRequest
func (mr *MockS3ControlAPIMockRecorder) ListJobsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListJobsRequest", reflect.TypeOf((*MockS3ControlAPI)(nil).ListJobsRequest), arg0)
}

// ListJobsWithContext mocks base method
func (m *MockS3ControlAPI) ListJobsWithContext(arg0 context.Context, arg1 *s3control.ListJobsInput, arg2 ...request.Option) (*s3control.ListJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListJobsWithContext", varargs...)
	ret0, _ := ret[0].(*s3control.ListJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListJobsWithContext indicates an expected call of ListJobsWithContext
func (mr *MockS3ControlAPIMockRecorder) ListJobsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListJobsWithContext", reflect.TypeOf((*MockS3ControlAPI)(nil).ListJobsWithContext), varargs...)
}

// PutPublicAccessBlock mocks base method
func (m *MockS3ControlAPI) PutPublicAccessBlock(arg0 *s3control.PutPublicAccessBlockInput) (*s3control.PutPublicAccessBlockOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutPublicAccessBlock", arg0)
	ret0, _ := ret[0].(*s3control.PutPublicAccessBlockOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutPublicAccessBlock indicates an expected call of PutPublicAccessBlock
func (mr *MockS3ControlAPIMockRecorder) PutPublicAccessBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutPublicAccessBlock", reflect.TypeOf((*MockS3ControlAPI)(nil).PutPublicAccessBlock), arg0)
}

// PutPublicAccessBlockRequest mocks base method
func (m *MockS3ControlAPI) PutPublicAccessBlockRequest(arg0 *s3control.PutPublicAccessBlockInput) (*request.Request, *s3control.PutPublicAccessBlockOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutPublicAccessBlockRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*s3control.PutPublicAccessBlockOutput)
	return ret0, ret1
}

// PutPublicAccessBlockRequest indicates an expected call of PutPublicAccessBlockRequest
func (mr *MockS3ControlAPIMockRecorder) PutPublicAccessBlockRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutPublicAccessBlockRequest", reflect.TypeOf((*MockS3ControlAPI)(nil).PutPublicAccessBlockRequest), arg0)
}

// PutPublicAccessBlockWithContext mocks base method
func (m *MockS3ControlAPI) PutPublicAccessBlockWithContext(arg0 context.Context, arg1 *s3control.PutPublicAccessBlockInput, arg2 ...request.Option) (*s3control.PutPublicAccessBlockOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutPublicAccessBlockWithContext", varargs...)
	ret0, _ := ret[0].(*s3control.PutPublicAccessBlockOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutPublicAccessBlockWithContext indicates an expected call of PutPublicAccessBlockWithContext
func (mr *MockS3ControlAPIMockRecorder) PutPublicAccessBlockWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutPublicAccessBlockWithContext", reflect.TypeOf((*MockS3ControlAPI)(nil).PutPublicAccessBlockWithContext), varargs...)
}

// UpdateJobPriority mocks base method
func (m *MockS3ControlAPI) UpdateJobPriority(arg0 *s3control.UpdateJobPriorityInput) (*s3control.UpdateJobPriorityOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateJobPriority", arg0)
	ret0, _ := ret[0].(*s3control.UpdateJobPriorityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateJobPriority indicates an expected call of UpdateJobPriority
func (mr *MockS3ControlAPIMockRecorder) UpdateJobPriority(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateJobPriority", reflect.TypeOf((*MockS3ControlAPI)(nil).UpdateJobPriority), arg0)
}

// UpdateJobPriorityRequest mocks base method
func (m *MockS3ControlAPI) UpdateJobPriorityRequest(arg0 *s3control.UpdateJobPriorityInput) (*request.Request, *s3control.UpdateJobPriorityOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateJobPriorityRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*s3control.UpdateJobPriorityOutput)
	return ret0, ret1
}

// UpdateJobPriorityRequest indicates an expected call of UpdateJobPriorityRequest
func (mr *MockS3ControlAPIMockRecorder) UpdateJobPriorityRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateJobPriorityRequest", reflect.TypeOf((*MockS3ControlAPI)(nil).UpdateJobPriorityRequest), arg0)
}

// UpdateJobPriorityWithContext mocks base method
func (m *MockS3ControlAPI) UpdateJobPriorityWithContext(arg0 context.Context, arg1 *s3control.UpdateJobPriorityInput, arg2 ...request.Option) (*s3control.UpdateJobPriorityOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateJobPriorityWithContext", varargs...)
	ret0, _ := ret[0].(*s3control.UpdateJobPriorityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateJobPriorityWithContext indicates an expected call of UpdateJobPriorityWithContext
func (mr *MockS3ControlAPIMockRecorder) UpdateJobPriorityWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateJobPriorityWithContext", reflect.TypeOf((*MockS3ControlAPI)(nil).UpdateJobPriorityWithContext), varargs...)
}

// UpdateJobStatus mocks base method
func (m *MockS3ControlAPI) UpdateJobStatus(arg0 *s3control.UpdateJobStatusInput) (*s3control.UpdateJobStatusOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateJobStatus", arg0)
	ret0, _ := ret[0].(*s3control.UpdateJobStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateJobStatus indicates an expected call of UpdateJobStatus
func (mr *MockS3ControlAPIMockRecorder) UpdateJobStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateJobStatus", reflect.TypeOf((*MockS3ControlAPI)(nil).UpdateJobStatus), arg0)
}

// UpdateJobStatusRequest mocks base method
func (m *MockS3ControlAPI) UpdateJobStatusRequest(arg0 *s3control.UpdateJobStatusInput) (*request.Request, *s3control.UpdateJobStatusOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateJobStatusRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*s3control.UpdateJobStatusOutput)
	return ret0, ret1
}

// UpdateJobStatusRequest indicates an expected call of UpdateJobStatusRequest
func (mr *MockS3ControlAPIMockRecorder) UpdateJobStatusRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateJobStatusRequest", reflect.TypeOf((*MockS3ControlAPI)(nil).UpdateJobStatusRequest), arg0)
}

// UpdateJobStatusWithContext mocks base method
func (m *MockS3ControlAPI) UpdateJobStatusWithContext(arg0 context.Context, arg1 *s3control.UpdateJobStatusInput, arg2 ...request.Option) (*s3control.UpdateJobStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateJobStatusWithContext", varargs...)
	ret0, _ := ret[0].(*s3control.UpdateJobStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateJobStatusWithContext indicates an expected call of UpdateJobStatusWithContext
func (mr *MockS3ControlAPIMockRecorder) UpdateJobStatusWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateJobStatusWithContext", reflect.TypeOf((*MockS3ControlAPI)(nil).UpdateJobStatusWithContext), varargs...)
}