// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/mediastore/mediastoreiface (interfaces: MediaStoreAPI)

// Package mediastoremock is a generated GoMock package.
package mediastoremock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	mediastore "github.com/aws/aws-sdk-go/service/mediastore"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMediaStoreAPI is a mock of MediaStoreAPI interface
type MockMediaStoreAPI struct {
	ctrl     *gomock.Controller
	recorder *MockMediaStoreAPIMockRecorder
}

// MockMediaStoreAPIMockRecorder is the mock recorder for MockMediaStoreAPI
type MockMediaStoreAPIMockRecorder struct {
	mock *MockMediaStoreAPI
}

// NewMockMediaStoreAPI creates a new mock instance
func NewMockMediaStoreAPI(ctrl *gomock.Controller) *MockMediaStoreAPI {
	mock := &MockMediaStoreAPI{ctrl: ctrl}
	mock.recorder = &MockMediaStoreAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMediaStoreAPI) EXPECT() *MockMediaStoreAPIMockRecorder {
	return m.recorder
}

// CreateContainer mocks base method
func (m *MockMediaStoreAPI) CreateContainer(arg0 *mediastore.CreateContainerInput) (*mediastore.CreateContainerOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateContainer", arg0)
	ret0, _ := ret[0].(*mediastore.CreateContainerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateContainer indicates an expected call of CreateContainer
func (mr *MockMediaStoreAPIMockRecorder) CreateContainer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContainer", reflect.TypeOf((*MockMediaStoreAPI)(nil).CreateContainer), arg0)
}

// CreateContainerRequest mocks base method
func (m *MockMediaStoreAPI) CreateContainerRequest(arg0 *mediastore.CreateContainerInput) (*request.Request, *mediastore.CreateContainerOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateContainerRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mediastore.CreateContainerOutput)
	return ret0, ret1
}

// CreateContainerRequest indicates an expected call of CreateContainerRequest
func (mr *MockMediaStoreAPIMockRecorder) CreateContainerRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContainerRequest", reflect.TypeOf((*MockMediaStoreAPI)(nil).CreateContainerRequest), arg0)
}

// CreateContainerWithContext mocks base method
func (m *MockMediaStoreAPI) CreateContainerWithContext(arg0 context.Context, arg1 *mediastore.CreateContainerInput, arg2 ...request.Option) (*mediastore.CreateContainerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateContainerWithContext", varargs...)
	ret0, _ := ret[0].(*mediastore.CreateContainerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateContainerWithContext indicates an expected call of CreateContainerWithContext
func (mr *MockMediaStoreAPIMockRecorder) CreateContainerWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContainerWithContext", reflect.TypeOf((*MockMediaStoreAPI)(nil).CreateContainerWithContext), varargs...)
}

// DeleteContainer mocks base method
func (m *MockMediaStoreAPI) DeleteContainer(arg0 *mediastore.DeleteContainerInput) (*mediastore.DeleteContainerOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteContainer", arg0)
	ret0, _ := ret[0].(*mediastore.DeleteContainerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteContainer indicates an expected call of DeleteContainer
func (mr *MockMediaStoreAPIMockRecorder) DeleteContainer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteContainer", reflect.TypeOf((*MockMediaStoreAPI)(nil).DeleteContainer), arg0)
}

// DeleteContainerPolicy mocks base method
func (m *MockMediaStoreAPI) DeleteContainerPolicy(arg0 *mediastore.DeleteContainerPolicyInput) (*mediastore.DeleteContainerPolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteContainerPolicy", arg0)
	ret0, _ := ret[0].(*mediastore.DeleteContainerPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteContainerPolicy indicates an expected call of DeleteContainerPolicy
func (mr *MockMediaStoreAPIMockRecorder) DeleteContainerPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteContainerPolicy", reflect.TypeOf((*MockMediaStoreAPI)(nil).DeleteContainerPolicy), arg0)
}

// DeleteContainerPolicyRequest mocks base method
func (m *MockMediaStoreAPI) DeleteContainerPolicyRequest(arg0 *mediastore.DeleteContainerPolicyInput) (*request.Request, *mediastore.DeleteContainerPolicyOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteContainerPolicyRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mediastore.DeleteContainerPolicyOutput)
	return ret0, ret1
}

// DeleteContainerPolicyRequest indicates an expected call of DeleteContainerPolicyRequest
func (mr *MockMediaStoreAPIMockRecorder) DeleteContainerPolicyRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteContainerPolicyRequest", reflect.TypeOf((*MockMediaStoreAPI)(nil).DeleteContainerPolicyRequest), arg0)
}

// DeleteContainerPolicyWithContext mocks base method
func (m *MockMediaStoreAPI) DeleteContainerPolicyWithContext(arg0 context.Context, arg1 *mediastore.DeleteContainerPolicyInput, arg2 ...request.Option) (*mediastore.DeleteContainerPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteContainerPolicyWithContext", varargs...)
	ret0, _ := ret[0].(*mediastore.DeleteContainerPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteContainerPolicyWithContext indicates an expected call of DeleteContainerPolicyWithContext
func (mr *MockMediaStoreAPIMockRecorder) DeleteContainerPolicyWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteContainerPolicyWithContext", reflect.TypeOf((*MockMediaStoreAPI)(nil).DeleteContainerPolicyWithContext), varargs...)
}

// DeleteContainerRequest mocks base method
func (m *MockMediaStoreAPI) DeleteContainerRequest(arg0 *mediastore.DeleteContainerInput) (*request.Request, *mediastore.DeleteContainerOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteContainerRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mediastore.DeleteContainerOutput)
	return ret0, ret1
}

// DeleteContainerRequest indicates an expected call of DeleteContainerRequest
func (mr *MockMediaStoreAPIMockRecorder) DeleteContainerRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteContainerRequest", reflect.TypeOf((*MockMediaStoreAPI)(nil).DeleteContainerRequest), arg0)
}

// DeleteContainerWithContext mocks base method
func (m *MockMediaStoreAPI) DeleteContainerWithContext(arg0 context.Context, arg1 *mediastore.DeleteContainerInput, arg2 ...request.Option) (*mediastore.DeleteContainerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteContainerWithContext", varargs...)
	ret0, _ := ret[0].(*mediastore.DeleteContainerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteContainerWithContext indicates an expected call of DeleteContainerWithContext
func (mr *MockMediaStoreAPIMockRecorder) DeleteContainerWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteContainerWithContext", reflect.TypeOf((*MockMediaStoreAPI)(nil).DeleteContainerWithContext), varargs...)
}

// DeleteCorsPolicy mocks base method
func (m *MockMediaStoreAPI) DeleteCorsPolicy(arg0 *mediastore.DeleteCorsPolicyInput) (*mediastore.DeleteCorsPolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCorsPolicy", arg0)
	ret0, _ := ret[0].(*mediastore.DeleteCorsPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCorsPolicy indicates an expected call of DeleteCorsPolicy
func (mr *MockMediaStoreAPIMockRecorder) DeleteCorsPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCorsPolicy", reflect.TypeOf((*MockMediaStoreAPI)(nil).DeleteCorsPolicy), arg0)
}

// DeleteCorsPolicyRequest mocks base method
func (m *MockMediaStoreAPI) DeleteCorsPolicyRequest(arg0 *mediastore.DeleteCorsPolicyInput) (*request.Request, *mediastore.DeleteCorsPolicyOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCorsPolicyRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mediastore.DeleteCorsPolicyOutput)
	return ret0, ret1
}

// DeleteCorsPolicyRequest indicates an expected call of DeleteCorsPolicyRequest
func (mr *MockMediaStoreAPIMockRecorder) DeleteCorsPolicyRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCorsPolicyRequest", reflect.TypeOf((*MockMediaStoreAPI)(nil).DeleteCorsPolicyRequest), arg0)
}

// DeleteCorsPolicyWithContext mocks base method
func (m *MockMediaStoreAPI) DeleteCorsPolicyWithContext(arg0 context.Context, arg1 *mediastore.DeleteCorsPolicyInput, arg2 ...request.Option) (*mediastore.DeleteCorsPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteCorsPolicyWithContext", varargs...)
	ret0, _ := ret[0].(*mediastore.DeleteCorsPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCorsPolicyWithContext indicates an expected call of DeleteCorsPolicyWithContext
func (mr *MockMediaStoreAPIMockRecorder) DeleteCorsPolicyWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCorsPolicyWithContext", reflect.TypeOf((*MockMediaStoreAPI)(nil).DeleteCorsPolicyWithContext), varargs...)
}

// DeleteLifecyclePolicy mocks base method
func (m *MockMediaStoreAPI) DeleteLifecyclePolicy(arg0 *mediastore.DeleteLifecyclePolicyInput) (*mediastore.DeleteLifecyclePolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLifecyclePolicy", arg0)
	ret0, _ := ret[0].(*mediastore.DeleteLifecyclePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteLifecyclePolicy indicates an expected call of DeleteLifecyclePolicy
func (mr *MockMediaStoreAPIMockRecorder) DeleteLifecyclePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLifecyclePolicy", reflect.TypeOf((*MockMediaStoreAPI)(nil).DeleteLifecyclePolicy), arg0)
}

// DeleteLifecyclePolicyRequest mocks base method
func (m *MockMediaStoreAPI) DeleteLifecyclePolicyRequest(arg0 *mediastore.DeleteLifecyclePolicyInput) (*request.Request, *mediastore.DeleteLifecyclePolicyOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLifecyclePolicyRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mediastore.DeleteLifecyclePolicyOutput)
	return ret0, ret1
}

// DeleteLifecyclePolicyRequest indicates an expected call of DeleteLifecyclePolicyRequest
func (mr *MockMediaStoreAPIMockRecorder) DeleteLifecyclePolicyRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLifecyclePolicyRequest", reflect.TypeOf((*MockMediaStoreAPI)(nil).DeleteLifecyclePolicyRequest), arg0)
}

// DeleteLifecyclePolicyWithContext mocks base method
func (m *MockMediaStoreAPI) DeleteLifecyclePolicyWithContext(arg0 context.Context, arg1 *mediastore.DeleteLifecyclePolicyInput, arg2 ...request.Option) (*mediastore.DeleteLifecyclePolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteLifecyclePolicyWithContext", varargs...)
	ret0, _ := ret[0].(*mediastore.DeleteLifecyclePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteLifecyclePolicyWithContext indicates an expected call of DeleteLifecyclePolicyWithContext
func (mr *MockMediaStoreAPIMockRecorder) DeleteLifecyclePolicyWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLifecyclePolicyWithContext", reflect.TypeOf((*MockMediaStoreAPI)(nil).DeleteLifecyclePolicyWithContext), varargs...)
}

// DescribeContainer mocks base method
func (m *MockMediaStoreAPI) DescribeContainer(arg0 *mediastore.DescribeContainerInput) (*mediastore.DescribeContainerOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeContainer", arg0)
	ret0, _ := ret[0].(*mediastore.DescribeContainerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeContainer indicates an expected call of DescribeContainer
func (mr *MockMediaStoreAPIMockRecorder) DescribeContainer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeContainer", reflect.TypeOf((*MockMediaStoreAPI)(nil).DescribeContainer), arg0)
}

// DescribeContainerRequest mocks base method
func (m *MockMediaStoreAPI) DescribeContainerRequest(arg0 *mediastore.DescribeContainerInput) (*request.Request, *mediastore.DescribeContainerOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeContainerRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mediastore.DescribeContainerOutput)
	return ret0, ret1
}

// DescribeContainerRequest indicates an expected call of DescribeContainerRequest
func (mr *MockMediaStoreAPIMockRecorder) DescribeContainerRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeContainerRequest", reflect.TypeOf((*MockMediaStoreAPI)(nil).DescribeContainerRequest), arg0)
}

// DescribeContainerWithContext mocks base method
func (m *MockMediaStoreAPI) DescribeContainerWithContext(arg0 context.Context, arg1 *mediastore.DescribeContainerInput, arg2 ...request.Option) (*mediastore.DescribeContainerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeContainerWithContext", varargs...)
	ret0, _ := ret[0].(*mediastore.DescribeContainerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeContainerWithContext indicates an expected call of DescribeContainerWithContext
func (mr *MockMediaStoreAPIMockRecorder) DescribeContainerWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeContainerWithContext", reflect.TypeOf((*MockMediaStoreAPI)(nil).DescribeContainerWithContext), varargs...)
}

// GetContainerPolicy mocks base method
func (m *MockMediaStoreAPI) GetContainerPolicy(arg0 *mediastore.GetContainerPolicyInput) (*mediastore.GetContainerPolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainerPolicy", arg0)
	ret0, _ := ret[0].(*mediastore.GetContainerPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContainerPolicy indicates an expected call of GetContainerPolicy
func (mr *MockMediaStoreAPIMockRecorder) GetContainerPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainerPolicy", reflect.TypeOf((*MockMediaStoreAPI)(nil).GetContainerPolicy), arg0)
}

// GetContainerPolicyRequest mocks base method
func (m *MockMediaStoreAPI) GetContainerPolicyRequest(arg0 *mediastore.GetContainerPolicyInput) (*request.Request, *mediastore.GetContainerPolicyOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainerPolicyRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mediastore.GetContainerPolicyOutput)
	return ret0, ret1
}

// GetContainerPolicyRequest indicates an expected call of GetContainerPolicyRequest
func (mr *MockMediaStoreAPIMockRecorder) GetContainerPolicyRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainerPolicyRequest", reflect.TypeOf((*MockMediaStoreAPI)(nil).GetContainerPolicyRequest), arg0)
}

// GetContainerPolicyWithContext mocks base method
func (m *MockMediaStoreAPI) GetContainerPolicyWithContext(arg0 context.Context, arg1 *mediastore.GetContainerPolicyInput, arg2 ...request.Option) (*mediastore.GetContainerPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetContainerPolicyWithContext", varargs...)
	ret0, _ := ret[0].(*mediastore.GetContainerPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContainerPolicyWithContext indicates an expected call of GetContainerPolicyWithContext
func (mr *MockMediaStoreAPIMockRecorder) GetContainerPolicyWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainerPolicyWithContext", reflect.TypeOf((*MockMediaStoreAPI)(nil).GetContainerPolicyWithContext), varargs...)
}

// GetCorsPolicy mocks base method
func (m *MockMediaStoreAPI) GetCorsPolicy(arg0 *mediastore.GetCorsPolicyInput) (*mediastore.GetCorsPolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCorsPolicy", arg0)
	ret0, _ := ret[0].(*mediastore.GetCorsPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCorsPolicy indicates an expected call of GetCorsPolicy
func (mr *MockMediaStoreAPIMockRecorder) GetCorsPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCorsPolicy", reflect.TypeOf((*MockMediaStoreAPI)(nil).GetCorsPolicy), arg0)
}

// GetCorsPolicyRequest mocks base method
func (m *MockMediaStoreAPI) GetCorsPolicyRequest(arg0 *mediastore.GetCorsPolicyInput) (*request.Request, *mediastore.GetCorsPolicyOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCorsPolicyRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mediastore.GetCorsPolicyOutput)
	return ret0, ret1
}

// GetCorsPolicyRequest indicates an expected call of GetCorsPolicyRequest
func (mr *MockMediaStoreAPIMockRecorder) GetCorsPolicyRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCorsPolicyRequest", reflect.TypeOf((*MockMediaStoreAPI)(nil).GetCorsPolicyRequest), arg0)
}

// GetCorsPolicyWithContext mocks base method
func (m *MockMediaStoreAPI) GetCorsPolicyWithContext(arg0 context.Context, arg1 *mediastore.GetCorsPolicyInput, arg2 ...request.Option) (*mediastore.GetCorsPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCorsPolicyWithContext", varargs...)
	ret0, _ := ret[0].(*mediastore.GetCorsPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCorsPolicyWithContext indicates an expected call of GetCorsPolicyWithContext
func (mr *MockMediaStoreAPIMockRecorder) GetCorsPolicyWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCorsPolicyWithContext", reflect.TypeOf((*MockMediaStoreAPI)(nil).GetCorsPolicyWithContext), varargs...)
}

// GetLifecyclePolicy mocks base method
func (m *MockMediaStoreAPI) GetLifecyclePolicy(arg0 *mediastore.GetLifecyclePolicyInput) (*mediastore.GetLifecyclePolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLifecyclePolicy", arg0)
	ret0, _ := ret[0].(*mediastore.GetLifecyclePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLifecyclePolicy indicates an expected call of GetLifecyclePolicy
func (mr *MockMediaStoreAPIMockRecorder) GetLifecyclePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLifecyclePolicy", reflect.TypeOf((*MockMediaStoreAPI)(nil).GetLifecyclePolicy), arg0)
}

// GetLifecyclePolicyRequest mocks base method
func (m *MockMediaStoreAPI) GetLifecyclePolicyRequest(arg0 *mediastore.GetLifecyclePolicyInput) (*request.Request, *mediastore.GetLifecyclePolicyOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLifecyclePolicyRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mediastore.GetLifecyclePolicyOutput)
	return ret0, ret1
}

// GetLifecyclePolicyRequest indicates an expected call of GetLifecyclePolicyRequest
func (mr *MockMediaStoreAPIMockRecorder) GetLifecyclePolicyRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLifecyclePolicyRequest", reflect.TypeOf((*MockMediaStoreAPI)(nil).GetLifecyclePolicyRequest), arg0)
}

// GetLifecyclePolicyWithContext mocks base method
func (m *MockMediaStoreAPI) GetLifecyclePolicyWithContext(arg0 context.Context, arg1 *mediastore.GetLifecyclePolicyInput, arg2 ...request.Option) (*mediastore.GetLifecyclePolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLifecyclePolicyWithContext", varargs...)
	ret0, _ := ret[0].(*mediastore.GetLifecyclePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLifecyclePolicyWithContext indicates an expected call of GetLifecyclePolicyWithContext
func (mr *MockMediaStoreAPIMockRecorder) GetLifecyclePolicyWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLifecyclePolicyWithContext", reflect.TypeOf((*MockMediaStoreAPI)(nil).GetLifecyclePolicyWithContext), varargs...)
}

// ListContainers mocks base method
func (m *MockMediaStoreAPI) ListContainers(arg0 *mediastore.ListContainersInput) (*mediastore.ListContainersOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListContainers", arg0)
	ret0, _ := ret[0].(*mediastore.ListContainersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContainers indicates an expected call of ListContainers
func (mr *MockMediaStoreAPIMockRecorder) ListContainers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContainers", reflect.TypeOf((*MockMediaStoreAPI)(nil).ListContainers), arg0)
}

// ListContainersRequest mocks base method
func (m *MockMediaStoreAPI) ListContainersRequest(arg0 *mediastore.ListContainersInput) (*request.Request, *mediastore.ListContainersOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListContainersRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mediastore.ListContainersOutput)
	return ret0, ret1
}

// ListContainersRequest indicates an expected call of ListContainersRequest
func (mr *MockMediaStoreAPIMockRecorder) ListContainersRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContainersRequest", reflect.TypeOf((*MockMediaStoreAPI)(nil).ListContainersRequest), arg0)
}

// ListContainersWithContext mocks base method
func (m *MockMediaStoreAPI) ListContainersWithContext(arg0 context.Context, arg1 *mediastore.ListContainersInput, arg2 ...request.Option) (*mediastore.ListContainersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListContainersWithContext", varargs...)
	ret0, _ := ret[0].(*mediastore.ListContainersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContainersWithContext indicates an expected call of ListContainersWithContext
func (mr *MockMediaStoreAPIMockRecorder) ListContainersWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContainersWithContext", reflect.TypeOf((*MockMediaStoreAPI)(nil).ListContainersWithContext), varargs...)
}

// PutContainerPolicy mocks base method
func (m *MockMediaStoreAPI) PutContainerPolicy(arg0 *mediastore.PutContainerPolicyInput) (*mediastore.PutContainerPolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutContainerPolicy", arg0)
	ret0, _ := ret[0].(*mediastore.PutContainerPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutContainerPolicy indicates an expected call of PutContainerPolicy
func (mr *MockMediaStoreAPIMockRecorder) PutContainerPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutContainerPolicy", reflect.TypeOf((*MockMediaStoreAPI)(nil).PutContainerPolicy), arg0)
}

// PutContainerPolicyRequest mocks base method
func (m *MockMediaStoreAPI) PutContainerPolicyRequest(arg0 *mediastore.PutContainerPolicyInput) (*request.Request, *mediastore.PutContainerPolicyOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutContainerPolicyRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mediastore.PutContainerPolicyOutput)
	return ret0, ret1
}

// PutContainerPolicyRequest indicates an expected call of PutContainerPolicyRequest
func (mr *MockMediaStoreAPIMockRecorder) PutContainerPolicyRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutContainerPolicyRequest", reflect.TypeOf((*MockMediaStoreAPI)(nil).PutContainerPolicyRequest), arg0)
}

// PutContainerPolicyWithContext mocks base method
func (m *MockMediaStoreAPI) PutContainerPolicyWithContext(arg0 context.Context, arg1 *mediastore.PutContainerPolicyInput, arg2 ...request.Option) (*mediastore.PutContainerPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutContainerPolicyWithContext", varargs...)
	ret0, _ := ret[0].(*mediastore.PutContainerPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutContainerPolicyWithContext indicates an expected call of PutContainerPolicyWithContext
func (mr *MockMediaStoreAPIMockRecorder) PutContainerPolicyWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutContainerPolicyWithContext", reflect.TypeOf((*MockMediaStoreAPI)(nil).PutContainerPolicyWithContext), varargs...)
}

// PutCorsPolicy mocks base method
func (m *MockMediaStoreAPI) PutCorsPolicy(arg0 *mediastore.PutCorsPolicyInput) (*mediastore.PutCorsPolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutCorsPolicy", arg0)
	ret0, _ := ret[0].(*mediastore.PutCorsPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutCorsPolicy indicates an expected call of PutCorsPolicy
func (mr *MockMediaStoreAPIMockRecorder) PutCorsPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutCorsPolicy", reflect.TypeOf((*MockMediaStoreAPI)(nil).PutCorsPolicy), arg0)
}

// PutCorsPolicyRequest mocks base method
func (m *MockMediaStoreAPI) PutCorsPolicyRequest(arg0 *mediastore.PutCorsPolicyInput) (*request.Request, *mediastore.PutCorsPolicyOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutCorsPolicyRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mediastore.PutCorsPolicyOutput)
	return ret0, ret1
}

// PutCorsPolicyRequest indicates an expected call of PutCorsPolicyRequest
func (mr *MockMediaStoreAPIMockRecorder) PutCorsPolicyRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutCorsPolicyRequest", reflect.TypeOf((*MockMediaStoreAPI)(nil).PutCorsPolicyRequest), arg0)
}

// PutCorsPolicyWithContext mocks base method
func (m *MockMediaStoreAPI) PutCorsPolicyWithContext(arg0 context.Context, arg1 *mediastore.PutCorsPolicyInput, arg2 ...request.Option) (*mediastore.PutCorsPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutCorsPolicyWithContext", varargs...)
	ret0, _ := ret[0].(*mediastore.PutCorsPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutCorsPolicyWithContext indicates an expected call of PutCorsPolicyWithContext
func (mr *MockMediaStoreAPIMockRecorder) PutCorsPolicyWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutCorsPolicyWithContext", reflect.TypeOf((*MockMediaStoreAPI)(nil).PutCorsPolicyWithContext), varargs...)
}

// PutLifecyclePolicy mocks base method
func (m *MockMediaStoreAPI) PutLifecyclePolicy(arg0 *mediastore.PutLifecyclePolicyInput) (*mediastore.PutLifecyclePolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutLifecyclePolicy", arg0)
	ret0, _ := ret[0].(*mediastore.PutLifecyclePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutLifecyclePolicy indicates an expected call of PutLifecyclePolicy
func (mr *MockMediaStoreAPIMockRecorder) PutLifecyclePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutLifecyclePolicy", reflect.TypeOf((*MockMediaStoreAPI)(nil).PutLifecyclePolicy), arg0)
}

// PutLifecyclePolicyRequest mocks base method
func (m *MockMediaStoreAPI) PutLifecyclePolicyRequest(arg0 *mediastore.PutLifecyclePolicyInput) (*request.Request, *mediastore.PutLifecyclePolicyOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutLifecyclePolicyRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mediastore.PutLifecyclePolicyOutput)
	return ret0, ret1
}

// PutLifecyclePolicyRequest indicates an expected call of PutLifecyclePolicyRequest
func (mr *MockMediaStoreAPIMockRecorder) PutLifecyclePolicyRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutLifecyclePolicyRequest", reflect.TypeOf((*MockMediaStoreAPI)(nil).PutLifecyclePolicyRequest), arg0)
}

// PutLifecyclePolicyWithContext mocks base method
func (m *MockMediaStoreAPI) PutLifecyclePolicyWithContext(arg0 context.Context, arg1 *mediastore.PutLifecyclePolicyInput, arg2 ...request.Option) (*mediastore.PutLifecyclePolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutLifecyclePolicyWithContext", varargs...)
	ret0, _ := ret[0].(*mediastore.PutLifecyclePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutLifecyclePolicyWithContext indicates an expected call of PutLifecyclePolicyWithContext
func (mr *MockMediaStoreAPIMockRecorder) PutLifecyclePolicyWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutLifecyclePolicyWithContext", reflect.TypeOf((*MockMediaStoreAPI)(nil).PutLifecyclePolicyWithContext), varargs...)
}

// StartAccessLogging mocks base method
func (m *MockMediaStoreAPI) StartAccessLogging(arg0 *mediastore.StartAccessLoggingInput) (*mediastore.StartAccessLoggingOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartAccessLogging", arg0)
	ret0, _ := ret[0].(*mediastore.StartAccessLoggingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartAccessLogging indicates an expected call of StartAccessLogging
func (mr *MockMediaStoreAPIMockRecorder) StartAccessLogging(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartAccessLogging", reflect.TypeOf((*MockMediaStoreAPI)(nil).StartAccessLogging), arg0)
}

// StartAccessLoggingRequest mocks base method
func (m *MockMediaStoreAPI) StartAccessLoggingRequest(arg0 *mediastore.StartAccessLoggingInput) (*request.Request, *mediastore.StartAccessLoggingOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartAccessLoggingRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mediastore.StartAccessLoggingOutput)
	return ret0, ret1
}

// StartAccessLoggingRequest indicates an expected call of StartAccessLoggingRequest
func (mr *MockMediaStoreAPIMockRecorder) StartAccessLoggingRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartAccessLoggingRequest", reflect.TypeOf((*MockMediaStoreAPI)(nil).StartAccessLoggingRequest), arg0)
}

// StartAccessLoggingWithContext mocks base method
func (m *MockMediaStoreAPI) StartAccessLoggingWithContext(arg0 context.Context, arg1 *mediastore.StartAccessLoggingInput, arg2 ...request.Option) (*mediastore.StartAccessLoggingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StartAccessLoggingWithContext", varargs...)
	ret0, _ := ret[0].(*mediastore.StartAccessLoggingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartAccessLoggingWithContext indicates an expected call of StartAccessLoggingWithContext
func (mr *MockMediaStoreAPIMockRecorder) StartAccessLoggingWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartAccessLoggingWithContext", reflect.TypeOf((*MockMediaStoreAPI)(nil).StartAccessLoggingWithContext), varargs...)
}

// StopAccessLogging mocks base method
func (m *MockMediaStoreAPI) StopAccessLogging(arg0 *mediastore.StopAccessLoggingInput) (*mediastore.StopAccessLoggingOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopAccessLogging", arg0)
	ret0, _ := ret[0].(*mediastore.StopAccessLoggingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopAccessLogging indicates an expected call of StopAccessLogging
func (mr *MockMediaStoreAPIMockRecorder) StopAccessLogging(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopAccessLogging", reflect.TypeOf((*MockMediaStoreAPI)(nil).StopAccessLogging), arg0)
}

// StopAccessLoggingRequest mocks base method
func (m *MockMediaStoreAPI) StopAccessLoggingRequest(arg0 *mediastore.StopAccessLoggingInput) (*request.Request, *mediastore.StopAccessLoggingOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopAccessLoggingRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mediastore.StopAccessLoggingOutput)
	return ret0, ret1
}

// StopAccessLoggingRequest indicates an expected call of StopAccessLoggingRequest
func (mr *MockMediaStoreAPIMockRecorder) StopAccessLoggingRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopAccessLoggingRequest", reflect.TypeOf((*MockMediaStoreAPI)(nil).StopAccessLoggingRequest), arg0)
}

// StopAccessLoggingWithContext mocks base method
func (m *MockMediaStoreAPI) StopAccessLoggingWithContext(arg0 context.Context, arg1 *mediastore.StopAccessLoggingInput, arg2 ...request.Option) (*mediastore.StopAccessLoggingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StopAccessLoggingWithContext", varargs...)
	ret0, _ := ret[0].(*mediastore.StopAccessLoggingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopAccessLoggingWithContext indicates an expected call of StopAccessLoggingWithContext
func (mr *MockMediaStoreAPIMockRecorder) StopAccessLoggingWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopAccessLoggingWithContext", reflect.TypeOf((*MockMediaStoreAPI)(nil).StopAccessLoggingWithContext), varargs...)
}