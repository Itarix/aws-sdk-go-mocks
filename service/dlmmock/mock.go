// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/dlm/dlmiface (interfaces: DLMAPI)

// Package dlmmock is a generated GoMock package.
package dlmmock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	dlm "github.com/aws/aws-sdk-go/service/dlm"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDLMAPI is a mock of DLMAPI interface
type MockDLMAPI struct {
	ctrl     *gomock.Controller
	recorder *MockDLMAPIMockRecorder
}

// MockDLMAPIMockRecorder is the mock recorder for MockDLMAPI
type MockDLMAPIMockRecorder struct {
	mock *MockDLMAPI
}

// NewMockDLMAPI creates a new mock instance
func NewMockDLMAPI(ctrl *gomock.Controller) *MockDLMAPI {
	mock := &MockDLMAPI{ctrl: ctrl}
	mock.recorder = &MockDLMAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDLMAPI) EXPECT() *MockDLMAPIMockRecorder {
	return m.recorder
}

// CreateLifecyclePolicy mocks base method
func (m *MockDLMAPI) CreateLifecyclePolicy(arg0 *dlm.CreateLifecyclePolicyInput) (*dlm.CreateLifecyclePolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLifecyclePolicy", arg0)
	ret0, _ := ret[0].(*dlm.CreateLifecyclePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLifecyclePolicy indicates an expected call of CreateLifecyclePolicy
func (mr *MockDLMAPIMockRecorder) CreateLifecyclePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLifecyclePolicy", reflect.TypeOf((*MockDLMAPI)(nil).CreateLifecyclePolicy), arg0)
}

// CreateLifecyclePolicyRequest mocks base method
func (m *MockDLMAPI) CreateLifecyclePolicyRequest(arg0 *dlm.CreateLifecyclePolicyInput) (*request.Request, *dlm.CreateLifecyclePolicyOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLifecyclePolicyRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*dlm.CreateLifecyclePolicyOutput)
	return ret0, ret1
}

// CreateLifecyclePolicyRequest indicates an expected call of CreateLifecyclePolicyRequest
func (mr *MockDLMAPIMockRecorder) CreateLifecyclePolicyRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLifecyclePolicyRequest", reflect.TypeOf((*MockDLMAPI)(nil).CreateLifecyclePolicyRequest), arg0)
}

// CreateLifecyclePolicyWithContext mocks base method
func (m *MockDLMAPI) CreateLifecyclePolicyWithContext(arg0 context.Context, arg1 *dlm.CreateLifecyclePolicyInput, arg2 ...request.Option) (*dlm.CreateLifecyclePolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateLifecyclePolicyWithContext", varargs...)
	ret0, _ := ret[0].(*dlm.CreateLifecyclePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLifecyclePolicyWithContext indicates an expected call of CreateLifecyclePolicyWithContext
func (mr *MockDLMAPIMockRecorder) CreateLifecyclePolicyWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLifecyclePolicyWithContext", reflect.TypeOf((*MockDLMAPI)(nil).CreateLifecyclePolicyWithContext), varargs...)
}

// DeleteLifecyclePolicy mocks base method
func (m *MockDLMAPI) DeleteLifecyclePolicy(arg0 *dlm.DeleteLifecyclePolicyInput) (*dlm.DeleteLifecyclePolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLifecyclePolicy", arg0)
	ret0, _ := ret[0].(*dlm.DeleteLifecyclePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteLifecyclePolicy indicates an expected call of DeleteLifecyclePolicy
func (mr *MockDLMAPIMockRecorder) DeleteLifecyclePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLifecyclePolicy", reflect.TypeOf((*MockDLMAPI)(nil).DeleteLifecyclePolicy), arg0)
}

// DeleteLifecyclePolicyRequest mocks base method
func (m *MockDLMAPI) DeleteLifecyclePolicyRequest(arg0 *dlm.DeleteLifecyclePolicyInput) (*request.Request, *dlm.DeleteLifecyclePolicyOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLifecyclePolicyRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*dlm.DeleteLifecyclePolicyOutput)
	return ret0, ret1
}

// DeleteLifecyclePolicyRequest indicates an expected call of DeleteLifecyclePolicyRequest
func (mr *MockDLMAPIMockRecorder) DeleteLifecyclePolicyRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLifecyclePolicyRequest", reflect.TypeOf((*MockDLMAPI)(nil).DeleteLifecyclePolicyRequest), arg0)
}

// DeleteLifecyclePolicyWithContext mocks base method
func (m *MockDLMAPI) DeleteLifecyclePolicyWithContext(arg0 context.Context, arg1 *dlm.DeleteLifecyclePolicyInput, arg2 ...request.Option) (*dlm.DeleteLifecyclePolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteLifecyclePolicyWithContext", varargs...)
	ret0, _ := ret[0].(*dlm.DeleteLifecyclePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteLifecyclePolicyWithContext indicates an expected call of DeleteLifecyclePolicyWithContext
func (mr *MockDLMAPIMockRecorder) DeleteLifecyclePolicyWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLifecyclePolicyWithContext", reflect.TypeOf((*MockDLMAPI)(nil).DeleteLifecyclePolicyWithContext), varargs...)
}

// GetLifecyclePolicies mocks base method
func (m *MockDLMAPI) GetLifecyclePolicies(arg0 *dlm.GetLifecyclePoliciesInput) (*dlm.GetLifecyclePoliciesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLifecyclePolicies", arg0)
	ret0, _ := ret[0].(*dlm.GetLifecyclePoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLifecyclePolicies indicates an expected call of GetLifecyclePolicies
func (mr *MockDLMAPIMockRecorder) GetLifecyclePolicies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLifecyclePolicies", reflect.TypeOf((*MockDLMAPI)(nil).GetLifecyclePolicies), arg0)
}

// GetLifecyclePoliciesRequest mocks base method
func (m *MockDLMAPI) GetLifecyclePoliciesRequest(arg0 *dlm.GetLifecyclePoliciesInput) (*request.Request, *dlm.GetLifecyclePoliciesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLifecyclePoliciesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*dlm.GetLifecyclePoliciesOutput)
	return ret0, ret1
}

// GetLifecyclePoliciesRequest indicates an expected call of GetLifecyclePoliciesRequest
func (mr *MockDLMAPIMockRecorder) GetLifecyclePoliciesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLifecyclePoliciesRequest", reflect.TypeOf((*MockDLMAPI)(nil).GetLifecyclePoliciesRequest), arg0)
}

// GetLifecyclePoliciesWithContext mocks base method
func (m *MockDLMAPI) GetLifecyclePoliciesWithContext(arg0 context.Context, arg1 *dlm.GetLifecyclePoliciesInput, arg2 ...request.Option) (*dlm.GetLifecyclePoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLifecyclePoliciesWithContext", varargs...)
	ret0, _ := ret[0].(*dlm.GetLifecyclePoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLifecyclePoliciesWithContext indicates an expected call of GetLifecyclePoliciesWithContext
func (mr *MockDLMAPIMockRecorder) GetLifecyclePoliciesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLifecyclePoliciesWithContext", reflect.TypeOf((*MockDLMAPI)(nil).GetLifecyclePoliciesWithContext), varargs...)
}

// GetLifecyclePolicy mocks base method
func (m *MockDLMAPI) GetLifecyclePolicy(arg0 *dlm.GetLifecyclePolicyInput) (*dlm.GetLifecyclePolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLifecyclePolicy", arg0)
	ret0, _ := ret[0].(*dlm.GetLifecyclePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLifecyclePolicy indicates an expected call of GetLifecyclePolicy
func (mr *MockDLMAPIMockRecorder) GetLifecyclePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLifecyclePolicy", reflect.TypeOf((*MockDLMAPI)(nil).GetLifecyclePolicy), arg0)
}

// GetLifecyclePolicyRequest mocks base method
func (m *MockDLMAPI) GetLifecyclePolicyRequest(arg0 *dlm.GetLifecyclePolicyInput) (*request.Request, *dlm.GetLifecyclePolicyOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLifecyclePolicyRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*dlm.GetLifecyclePolicyOutput)
	return ret0, ret1
}

// GetLifecyclePolicyRequest indicates an expected call of GetLifecyclePolicyRequest
func (mr *MockDLMAPIMockRecorder) GetLifecyclePolicyRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLifecyclePolicyRequest", reflect.TypeOf((*MockDLMAPI)(nil).GetLifecyclePolicyRequest), arg0)
}

// GetLifecyclePolicyWithContext mocks base method
func (m *MockDLMAPI) GetLifecyclePolicyWithContext(arg0 context.Context, arg1 *dlm.GetLifecyclePolicyInput, arg2 ...request.Option) (*dlm.GetLifecyclePolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLifecyclePolicyWithContext", varargs...)
	ret0, _ := ret[0].(*dlm.GetLifecyclePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLifecyclePolicyWithContext indicates an expected call of GetLifecyclePolicyWithContext
func (mr *MockDLMAPIMockRecorder) GetLifecyclePolicyWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLifecyclePolicyWithContext", reflect.TypeOf((*MockDLMAPI)(nil).GetLifecyclePolicyWithContext), varargs...)
}

// ListTagsForResource mocks base method
func (m *MockDLMAPI) ListTagsForResource(arg0 *dlm.ListTagsForResourceInput) (*dlm.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForResource", arg0)
	ret0, _ := ret[0].(*dlm.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource
func (mr *MockDLMAPIMockRecorder) ListTagsForResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockDLMAPI)(nil).ListTagsForResource), arg0)
}

// ListTagsForResourceRequest mocks base method
func (m *MockDLMAPI) ListTagsForResourceRequest(arg0 *dlm.ListTagsForResourceInput) (*request.Request, *dlm.ListTagsForResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*dlm.ListTagsForResourceOutput)
	return ret0, ret1
}

// ListTagsForResourceRequest indicates an expected call of ListTagsForResourceRequest
func (mr *MockDLMAPIMockRecorder) ListTagsForResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResourceRequest", reflect.TypeOf((*MockDLMAPI)(nil).ListTagsForResourceRequest), arg0)
}

// ListTagsForResourceWithContext mocks base method
func (m *MockDLMAPI) ListTagsForResourceWithContext(arg0 context.Context, arg1 *dlm.ListTagsForResourceInput, arg2 ...request.Option) (*dlm.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResourceWithContext", varargs...)
	ret0, _ := ret[0].(*dlm.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResourceWithContext indicates an expected call of ListTagsForResourceWithContext
func (mr *MockDLMAPIMockRecorder) ListTagsForResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResourceWithContext", reflect.TypeOf((*MockDLMAPI)(nil).ListTagsForResourceWithContext), varargs...)
}

// TagResource mocks base method
func (m *MockDLMAPI) TagResource(arg0 *dlm.TagResourceInput) (*dlm.TagResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagResource", arg0)
	ret0, _ := ret[0].(*dlm.TagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagResource indicates an expected call of TagResource
func (mr *MockDLMAPIMockRecorder) TagResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResource", reflect.TypeOf((*MockDLMAPI)(nil).TagResource), arg0)
}

// TagResourceRequest mocks base method
func (m *MockDLMAPI) TagResourceRequest(arg0 *dlm.TagResourceInput) (*request.Request, *dlm.TagResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*dlm.TagResourceOutput)
	return ret0, ret1
}

// TagResourceRequest indicates an expected call of TagResourceRequest
func (mr *MockDLMAPIMockRecorder) TagResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResourceRequest", reflect.TypeOf((*MockDLMAPI)(nil).TagResourceRequest), arg0)
}

// TagResourceWithContext mocks base method
func (m *MockDLMAPI) TagResourceWithContext(arg0 context.Context, arg1 *dlm.TagResourceInput, arg2 ...request.Option) (*dlm.TagResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TagResourceWithContext", varargs...)
	ret0, _ := ret[0].(*dlm.TagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagResourceWithContext indicates an expected call of TagResourceWithContext
func (mr *MockDLMAPIMockRecorder) TagResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResourceWithContext", reflect.TypeOf((*MockDLMAPI)(nil).TagResourceWithContext), varargs...)
}

// UntagResource mocks base method
func (m *MockDLMAPI) UntagResource(arg0 *dlm.UntagResourceInput) (*dlm.UntagResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagResource", arg0)
	ret0, _ := ret[0].(*dlm.UntagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagResource indicates an expected call of UntagResource
func (mr *MockDLMAPIMockRecorder) UntagResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResource", reflect.TypeOf((*MockDLMAPI)(nil).UntagResource), arg0)
}

// UntagResourceRequest mocks base method
func (m *MockDLMAPI) UntagResourceRequest(arg0 *dlm.UntagResourceInput) (*request.Request, *dlm.UntagResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*dlm.UntagResourceOutput)
	return ret0, ret1
}

// UntagResourceRequest indicates an expected call of UntagResourceRequest
func (mr *MockDLMAPIMockRecorder) UntagResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResourceRequest", reflect.TypeOf((*MockDLMAPI)(nil).UntagResourceRequest), arg0)
}

// UntagResourceWithContext mocks base method
func (m *MockDLMAPI) UntagResourceWithContext(arg0 context.Context, arg1 *dlm.UntagResourceInput, arg2 ...request.Option) (*dlm.UntagResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UntagResourceWithContext", varargs...)
	ret0, _ := ret[0].(*dlm.UntagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagResourceWithContext indicates an expected call of UntagResourceWithContext
func (mr *MockDLMAPIMockRecorder) UntagResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResourceWithContext", reflect.TypeOf((*MockDLMAPI)(nil).UntagResourceWithContext), varargs...)
}

// UpdateLifecyclePolicy mocks base method
func (m *MockDLMAPI) UpdateLifecyclePolicy(arg0 *dlm.UpdateLifecyclePolicyInput) (*dlm.UpdateLifecyclePolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLifecyclePolicy", arg0)
	ret0, _ := ret[0].(*dlm.UpdateLifecyclePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateLifecyclePolicy indicates an expected call of UpdateLifecyclePolicy
func (mr *MockDLMAPIMockRecorder) UpdateLifecyclePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLifecyclePolicy", reflect.TypeOf((*MockDLMAPI)(nil).UpdateLifecyclePolicy), arg0)
}

// UpdateLifecyclePolicyRequest mocks base method
func (m *MockDLMAPI) UpdateLifecyclePolicyRequest(arg0 *dlm.UpdateLifecyclePolicyInput) (*request.Request, *dlm.UpdateLifecyclePolicyOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLifecyclePolicyRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*dlm.UpdateLifecyclePolicyOutput)
	return ret0, ret1
}

// UpdateLifecyclePolicyRequest indicates an expected call of UpdateLifecyclePolicyRequest
func (mr *MockDLMAPIMockRecorder) UpdateLifecyclePolicyRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLifecyclePolicyRequest", reflect.TypeOf((*MockDLMAPI)(nil).UpdateLifecyclePolicyRequest), arg0)
}

// UpdateLifecyclePolicyWithContext mocks base method
func (m *MockDLMAPI) UpdateLifecyclePolicyWithContext(arg0 context.Context, arg1 *dlm.UpdateLifecyclePolicyInput, arg2 ...request.Option) (*dlm.UpdateLifecyclePolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateLifecyclePolicyWithContext", varargs...)
	ret0, _ := ret[0].(*dlm.UpdateLifecyclePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateLifecyclePolicyWithContext indicates an expected call of UpdateLifecyclePolicyWithContext
func (mr *MockDLMAPIMockRecorder) UpdateLifecyclePolicyWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLifecyclePolicyWithContext", reflect.TypeOf((*MockDLMAPI)(nil).UpdateLifecyclePolicyWithContext), varargs...)
}
