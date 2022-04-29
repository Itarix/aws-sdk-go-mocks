// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/chimesdkmediapipelines/chimesdkmediapipelinesiface (interfaces: ChimeSDKMediaPipelinesAPI)

// Package chimesdkmediapipelinesmock is a generated GoMock package.
package chimesdkmediapipelinesmock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	chimesdkmediapipelines "github.com/aws/aws-sdk-go/service/chimesdkmediapipelines"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockChimeSDKMediaPipelinesAPI is a mock of ChimeSDKMediaPipelinesAPI interface
type MockChimeSDKMediaPipelinesAPI struct {
	ctrl     *gomock.Controller
	recorder *MockChimeSDKMediaPipelinesAPIMockRecorder
}

// MockChimeSDKMediaPipelinesAPIMockRecorder is the mock recorder for MockChimeSDKMediaPipelinesAPI
type MockChimeSDKMediaPipelinesAPIMockRecorder struct {
	mock *MockChimeSDKMediaPipelinesAPI
}

// NewMockChimeSDKMediaPipelinesAPI creates a new mock instance
func NewMockChimeSDKMediaPipelinesAPI(ctrl *gomock.Controller) *MockChimeSDKMediaPipelinesAPI {
	mock := &MockChimeSDKMediaPipelinesAPI{ctrl: ctrl}
	mock.recorder = &MockChimeSDKMediaPipelinesAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChimeSDKMediaPipelinesAPI) EXPECT() *MockChimeSDKMediaPipelinesAPIMockRecorder {
	return m.recorder
}

// CreateMediaCapturePipeline mocks base method
func (m *MockChimeSDKMediaPipelinesAPI) CreateMediaCapturePipeline(arg0 *chimesdkmediapipelines.CreateMediaCapturePipelineInput) (*chimesdkmediapipelines.CreateMediaCapturePipelineOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMediaCapturePipeline", arg0)
	ret0, _ := ret[0].(*chimesdkmediapipelines.CreateMediaCapturePipelineOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMediaCapturePipeline indicates an expected call of CreateMediaCapturePipeline
func (mr *MockChimeSDKMediaPipelinesAPIMockRecorder) CreateMediaCapturePipeline(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMediaCapturePipeline", reflect.TypeOf((*MockChimeSDKMediaPipelinesAPI)(nil).CreateMediaCapturePipeline), arg0)
}

// CreateMediaCapturePipelineRequest mocks base method
func (m *MockChimeSDKMediaPipelinesAPI) CreateMediaCapturePipelineRequest(arg0 *chimesdkmediapipelines.CreateMediaCapturePipelineInput) (*request.Request, *chimesdkmediapipelines.CreateMediaCapturePipelineOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMediaCapturePipelineRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*chimesdkmediapipelines.CreateMediaCapturePipelineOutput)
	return ret0, ret1
}

// CreateMediaCapturePipelineRequest indicates an expected call of CreateMediaCapturePipelineRequest
func (mr *MockChimeSDKMediaPipelinesAPIMockRecorder) CreateMediaCapturePipelineRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMediaCapturePipelineRequest", reflect.TypeOf((*MockChimeSDKMediaPipelinesAPI)(nil).CreateMediaCapturePipelineRequest), arg0)
}

// CreateMediaCapturePipelineWithContext mocks base method
func (m *MockChimeSDKMediaPipelinesAPI) CreateMediaCapturePipelineWithContext(arg0 context.Context, arg1 *chimesdkmediapipelines.CreateMediaCapturePipelineInput, arg2 ...request.Option) (*chimesdkmediapipelines.CreateMediaCapturePipelineOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMediaCapturePipelineWithContext", varargs...)
	ret0, _ := ret[0].(*chimesdkmediapipelines.CreateMediaCapturePipelineOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMediaCapturePipelineWithContext indicates an expected call of CreateMediaCapturePipelineWithContext
func (mr *MockChimeSDKMediaPipelinesAPIMockRecorder) CreateMediaCapturePipelineWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMediaCapturePipelineWithContext", reflect.TypeOf((*MockChimeSDKMediaPipelinesAPI)(nil).CreateMediaCapturePipelineWithContext), varargs...)
}

// DeleteMediaCapturePipeline mocks base method
func (m *MockChimeSDKMediaPipelinesAPI) DeleteMediaCapturePipeline(arg0 *chimesdkmediapipelines.DeleteMediaCapturePipelineInput) (*chimesdkmediapipelines.DeleteMediaCapturePipelineOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMediaCapturePipeline", arg0)
	ret0, _ := ret[0].(*chimesdkmediapipelines.DeleteMediaCapturePipelineOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMediaCapturePipeline indicates an expected call of DeleteMediaCapturePipeline
func (mr *MockChimeSDKMediaPipelinesAPIMockRecorder) DeleteMediaCapturePipeline(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMediaCapturePipeline", reflect.TypeOf((*MockChimeSDKMediaPipelinesAPI)(nil).DeleteMediaCapturePipeline), arg0)
}

// DeleteMediaCapturePipelineRequest mocks base method
func (m *MockChimeSDKMediaPipelinesAPI) DeleteMediaCapturePipelineRequest(arg0 *chimesdkmediapipelines.DeleteMediaCapturePipelineInput) (*request.Request, *chimesdkmediapipelines.DeleteMediaCapturePipelineOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMediaCapturePipelineRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*chimesdkmediapipelines.DeleteMediaCapturePipelineOutput)
	return ret0, ret1
}

// DeleteMediaCapturePipelineRequest indicates an expected call of DeleteMediaCapturePipelineRequest
func (mr *MockChimeSDKMediaPipelinesAPIMockRecorder) DeleteMediaCapturePipelineRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMediaCapturePipelineRequest", reflect.TypeOf((*MockChimeSDKMediaPipelinesAPI)(nil).DeleteMediaCapturePipelineRequest), arg0)
}

// DeleteMediaCapturePipelineWithContext mocks base method
func (m *MockChimeSDKMediaPipelinesAPI) DeleteMediaCapturePipelineWithContext(arg0 context.Context, arg1 *chimesdkmediapipelines.DeleteMediaCapturePipelineInput, arg2 ...request.Option) (*chimesdkmediapipelines.DeleteMediaCapturePipelineOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMediaCapturePipelineWithContext", varargs...)
	ret0, _ := ret[0].(*chimesdkmediapipelines.DeleteMediaCapturePipelineOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMediaCapturePipelineWithContext indicates an expected call of DeleteMediaCapturePipelineWithContext
func (mr *MockChimeSDKMediaPipelinesAPIMockRecorder) DeleteMediaCapturePipelineWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMediaCapturePipelineWithContext", reflect.TypeOf((*MockChimeSDKMediaPipelinesAPI)(nil).DeleteMediaCapturePipelineWithContext), varargs...)
}

// GetMediaCapturePipeline mocks base method
func (m *MockChimeSDKMediaPipelinesAPI) GetMediaCapturePipeline(arg0 *chimesdkmediapipelines.GetMediaCapturePipelineInput) (*chimesdkmediapipelines.GetMediaCapturePipelineOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMediaCapturePipeline", arg0)
	ret0, _ := ret[0].(*chimesdkmediapipelines.GetMediaCapturePipelineOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMediaCapturePipeline indicates an expected call of GetMediaCapturePipeline
func (mr *MockChimeSDKMediaPipelinesAPIMockRecorder) GetMediaCapturePipeline(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMediaCapturePipeline", reflect.TypeOf((*MockChimeSDKMediaPipelinesAPI)(nil).GetMediaCapturePipeline), arg0)
}

// GetMediaCapturePipelineRequest mocks base method
func (m *MockChimeSDKMediaPipelinesAPI) GetMediaCapturePipelineRequest(arg0 *chimesdkmediapipelines.GetMediaCapturePipelineInput) (*request.Request, *chimesdkmediapipelines.GetMediaCapturePipelineOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMediaCapturePipelineRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*chimesdkmediapipelines.GetMediaCapturePipelineOutput)
	return ret0, ret1
}

// GetMediaCapturePipelineRequest indicates an expected call of GetMediaCapturePipelineRequest
func (mr *MockChimeSDKMediaPipelinesAPIMockRecorder) GetMediaCapturePipelineRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMediaCapturePipelineRequest", reflect.TypeOf((*MockChimeSDKMediaPipelinesAPI)(nil).GetMediaCapturePipelineRequest), arg0)
}

// GetMediaCapturePipelineWithContext mocks base method
func (m *MockChimeSDKMediaPipelinesAPI) GetMediaCapturePipelineWithContext(arg0 context.Context, arg1 *chimesdkmediapipelines.GetMediaCapturePipelineInput, arg2 ...request.Option) (*chimesdkmediapipelines.GetMediaCapturePipelineOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMediaCapturePipelineWithContext", varargs...)
	ret0, _ := ret[0].(*chimesdkmediapipelines.GetMediaCapturePipelineOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMediaCapturePipelineWithContext indicates an expected call of GetMediaCapturePipelineWithContext
func (mr *MockChimeSDKMediaPipelinesAPIMockRecorder) GetMediaCapturePipelineWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMediaCapturePipelineWithContext", reflect.TypeOf((*MockChimeSDKMediaPipelinesAPI)(nil).GetMediaCapturePipelineWithContext), varargs...)
}

// ListMediaCapturePipelines mocks base method
func (m *MockChimeSDKMediaPipelinesAPI) ListMediaCapturePipelines(arg0 *chimesdkmediapipelines.ListMediaCapturePipelinesInput) (*chimesdkmediapipelines.ListMediaCapturePipelinesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMediaCapturePipelines", arg0)
	ret0, _ := ret[0].(*chimesdkmediapipelines.ListMediaCapturePipelinesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMediaCapturePipelines indicates an expected call of ListMediaCapturePipelines
func (mr *MockChimeSDKMediaPipelinesAPIMockRecorder) ListMediaCapturePipelines(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMediaCapturePipelines", reflect.TypeOf((*MockChimeSDKMediaPipelinesAPI)(nil).ListMediaCapturePipelines), arg0)
}

// ListMediaCapturePipelinesPages mocks base method
func (m *MockChimeSDKMediaPipelinesAPI) ListMediaCapturePipelinesPages(arg0 *chimesdkmediapipelines.ListMediaCapturePipelinesInput, arg1 func(*chimesdkmediapipelines.ListMediaCapturePipelinesOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMediaCapturePipelinesPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListMediaCapturePipelinesPages indicates an expected call of ListMediaCapturePipelinesPages
func (mr *MockChimeSDKMediaPipelinesAPIMockRecorder) ListMediaCapturePipelinesPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMediaCapturePipelinesPages", reflect.TypeOf((*MockChimeSDKMediaPipelinesAPI)(nil).ListMediaCapturePipelinesPages), arg0, arg1)
}

// ListMediaCapturePipelinesPagesWithContext mocks base method
func (m *MockChimeSDKMediaPipelinesAPI) ListMediaCapturePipelinesPagesWithContext(arg0 context.Context, arg1 *chimesdkmediapipelines.ListMediaCapturePipelinesInput, arg2 func(*chimesdkmediapipelines.ListMediaCapturePipelinesOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMediaCapturePipelinesPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListMediaCapturePipelinesPagesWithContext indicates an expected call of ListMediaCapturePipelinesPagesWithContext
func (mr *MockChimeSDKMediaPipelinesAPIMockRecorder) ListMediaCapturePipelinesPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMediaCapturePipelinesPagesWithContext", reflect.TypeOf((*MockChimeSDKMediaPipelinesAPI)(nil).ListMediaCapturePipelinesPagesWithContext), varargs...)
}

// ListMediaCapturePipelinesRequest mocks base method
func (m *MockChimeSDKMediaPipelinesAPI) ListMediaCapturePipelinesRequest(arg0 *chimesdkmediapipelines.ListMediaCapturePipelinesInput) (*request.Request, *chimesdkmediapipelines.ListMediaCapturePipelinesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMediaCapturePipelinesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*chimesdkmediapipelines.ListMediaCapturePipelinesOutput)
	return ret0, ret1
}

// ListMediaCapturePipelinesRequest indicates an expected call of ListMediaCapturePipelinesRequest
func (mr *MockChimeSDKMediaPipelinesAPIMockRecorder) ListMediaCapturePipelinesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMediaCapturePipelinesRequest", reflect.TypeOf((*MockChimeSDKMediaPipelinesAPI)(nil).ListMediaCapturePipelinesRequest), arg0)
}

// ListMediaCapturePipelinesWithContext mocks base method
func (m *MockChimeSDKMediaPipelinesAPI) ListMediaCapturePipelinesWithContext(arg0 context.Context, arg1 *chimesdkmediapipelines.ListMediaCapturePipelinesInput, arg2 ...request.Option) (*chimesdkmediapipelines.ListMediaCapturePipelinesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMediaCapturePipelinesWithContext", varargs...)
	ret0, _ := ret[0].(*chimesdkmediapipelines.ListMediaCapturePipelinesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMediaCapturePipelinesWithContext indicates an expected call of ListMediaCapturePipelinesWithContext
func (mr *MockChimeSDKMediaPipelinesAPIMockRecorder) ListMediaCapturePipelinesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMediaCapturePipelinesWithContext", reflect.TypeOf((*MockChimeSDKMediaPipelinesAPI)(nil).ListMediaCapturePipelinesWithContext), varargs...)
}

// ListTagsForResource mocks base method
func (m *MockChimeSDKMediaPipelinesAPI) ListTagsForResource(arg0 *chimesdkmediapipelines.ListTagsForResourceInput) (*chimesdkmediapipelines.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForResource", arg0)
	ret0, _ := ret[0].(*chimesdkmediapipelines.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource
func (mr *MockChimeSDKMediaPipelinesAPIMockRecorder) ListTagsForResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockChimeSDKMediaPipelinesAPI)(nil).ListTagsForResource), arg0)
}

// ListTagsForResourceRequest mocks base method
func (m *MockChimeSDKMediaPipelinesAPI) ListTagsForResourceRequest(arg0 *chimesdkmediapipelines.ListTagsForResourceInput) (*request.Request, *chimesdkmediapipelines.ListTagsForResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*chimesdkmediapipelines.ListTagsForResourceOutput)
	return ret0, ret1
}

// ListTagsForResourceRequest indicates an expected call of ListTagsForResourceRequest
func (mr *MockChimeSDKMediaPipelinesAPIMockRecorder) ListTagsForResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResourceRequest", reflect.TypeOf((*MockChimeSDKMediaPipelinesAPI)(nil).ListTagsForResourceRequest), arg0)
}

// ListTagsForResourceWithContext mocks base method
func (m *MockChimeSDKMediaPipelinesAPI) ListTagsForResourceWithContext(arg0 context.Context, arg1 *chimesdkmediapipelines.ListTagsForResourceInput, arg2 ...request.Option) (*chimesdkmediapipelines.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResourceWithContext", varargs...)
	ret0, _ := ret[0].(*chimesdkmediapipelines.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResourceWithContext indicates an expected call of ListTagsForResourceWithContext
func (mr *MockChimeSDKMediaPipelinesAPIMockRecorder) ListTagsForResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResourceWithContext", reflect.TypeOf((*MockChimeSDKMediaPipelinesAPI)(nil).ListTagsForResourceWithContext), varargs...)
}

// TagResource mocks base method
func (m *MockChimeSDKMediaPipelinesAPI) TagResource(arg0 *chimesdkmediapipelines.TagResourceInput) (*chimesdkmediapipelines.TagResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagResource", arg0)
	ret0, _ := ret[0].(*chimesdkmediapipelines.TagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagResource indicates an expected call of TagResource
func (mr *MockChimeSDKMediaPipelinesAPIMockRecorder) TagResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResource", reflect.TypeOf((*MockChimeSDKMediaPipelinesAPI)(nil).TagResource), arg0)
}

// TagResourceRequest mocks base method
func (m *MockChimeSDKMediaPipelinesAPI) TagResourceRequest(arg0 *chimesdkmediapipelines.TagResourceInput) (*request.Request, *chimesdkmediapipelines.TagResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*chimesdkmediapipelines.TagResourceOutput)
	return ret0, ret1
}

// TagResourceRequest indicates an expected call of TagResourceRequest
func (mr *MockChimeSDKMediaPipelinesAPIMockRecorder) TagResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResourceRequest", reflect.TypeOf((*MockChimeSDKMediaPipelinesAPI)(nil).TagResourceRequest), arg0)
}

// TagResourceWithContext mocks base method
func (m *MockChimeSDKMediaPipelinesAPI) TagResourceWithContext(arg0 context.Context, arg1 *chimesdkmediapipelines.TagResourceInput, arg2 ...request.Option) (*chimesdkmediapipelines.TagResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TagResourceWithContext", varargs...)
	ret0, _ := ret[0].(*chimesdkmediapipelines.TagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagResourceWithContext indicates an expected call of TagResourceWithContext
func (mr *MockChimeSDKMediaPipelinesAPIMockRecorder) TagResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResourceWithContext", reflect.TypeOf((*MockChimeSDKMediaPipelinesAPI)(nil).TagResourceWithContext), varargs...)
}

// UntagResource mocks base method
func (m *MockChimeSDKMediaPipelinesAPI) UntagResource(arg0 *chimesdkmediapipelines.UntagResourceInput) (*chimesdkmediapipelines.UntagResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagResource", arg0)
	ret0, _ := ret[0].(*chimesdkmediapipelines.UntagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagResource indicates an expected call of UntagResource
func (mr *MockChimeSDKMediaPipelinesAPIMockRecorder) UntagResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResource", reflect.TypeOf((*MockChimeSDKMediaPipelinesAPI)(nil).UntagResource), arg0)
}

// UntagResourceRequest mocks base method
func (m *MockChimeSDKMediaPipelinesAPI) UntagResourceRequest(arg0 *chimesdkmediapipelines.UntagResourceInput) (*request.Request, *chimesdkmediapipelines.UntagResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*chimesdkmediapipelines.UntagResourceOutput)
	return ret0, ret1
}

// UntagResourceRequest indicates an expected call of UntagResourceRequest
func (mr *MockChimeSDKMediaPipelinesAPIMockRecorder) UntagResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResourceRequest", reflect.TypeOf((*MockChimeSDKMediaPipelinesAPI)(nil).UntagResourceRequest), arg0)
}

// UntagResourceWithContext mocks base method
func (m *MockChimeSDKMediaPipelinesAPI) UntagResourceWithContext(arg0 context.Context, arg1 *chimesdkmediapipelines.UntagResourceInput, arg2 ...request.Option) (*chimesdkmediapipelines.UntagResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UntagResourceWithContext", varargs...)
	ret0, _ := ret[0].(*chimesdkmediapipelines.UntagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagResourceWithContext indicates an expected call of UntagResourceWithContext
func (mr *MockChimeSDKMediaPipelinesAPIMockRecorder) UntagResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResourceWithContext", reflect.TypeOf((*MockChimeSDKMediaPipelinesAPI)(nil).UntagResourceWithContext), varargs...)
}