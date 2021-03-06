// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/iotsecuretunneling/iotsecuretunnelingiface (interfaces: IoTSecureTunnelingAPI)

// Package iotsecuretunnelingmock is a generated GoMock package.
package iotsecuretunnelingmock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	iotsecuretunneling "github.com/aws/aws-sdk-go/service/iotsecuretunneling"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIoTSecureTunnelingAPI is a mock of IoTSecureTunnelingAPI interface
type MockIoTSecureTunnelingAPI struct {
	ctrl     *gomock.Controller
	recorder *MockIoTSecureTunnelingAPIMockRecorder
}

// MockIoTSecureTunnelingAPIMockRecorder is the mock recorder for MockIoTSecureTunnelingAPI
type MockIoTSecureTunnelingAPIMockRecorder struct {
	mock *MockIoTSecureTunnelingAPI
}

// NewMockIoTSecureTunnelingAPI creates a new mock instance
func NewMockIoTSecureTunnelingAPI(ctrl *gomock.Controller) *MockIoTSecureTunnelingAPI {
	mock := &MockIoTSecureTunnelingAPI{ctrl: ctrl}
	mock.recorder = &MockIoTSecureTunnelingAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIoTSecureTunnelingAPI) EXPECT() *MockIoTSecureTunnelingAPIMockRecorder {
	return m.recorder
}

// CloseTunnel mocks base method
func (m *MockIoTSecureTunnelingAPI) CloseTunnel(arg0 *iotsecuretunneling.CloseTunnelInput) (*iotsecuretunneling.CloseTunnelOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseTunnel", arg0)
	ret0, _ := ret[0].(*iotsecuretunneling.CloseTunnelOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloseTunnel indicates an expected call of CloseTunnel
func (mr *MockIoTSecureTunnelingAPIMockRecorder) CloseTunnel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseTunnel", reflect.TypeOf((*MockIoTSecureTunnelingAPI)(nil).CloseTunnel), arg0)
}

// CloseTunnelRequest mocks base method
func (m *MockIoTSecureTunnelingAPI) CloseTunnelRequest(arg0 *iotsecuretunneling.CloseTunnelInput) (*request.Request, *iotsecuretunneling.CloseTunnelOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseTunnelRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iotsecuretunneling.CloseTunnelOutput)
	return ret0, ret1
}

// CloseTunnelRequest indicates an expected call of CloseTunnelRequest
func (mr *MockIoTSecureTunnelingAPIMockRecorder) CloseTunnelRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseTunnelRequest", reflect.TypeOf((*MockIoTSecureTunnelingAPI)(nil).CloseTunnelRequest), arg0)
}

// CloseTunnelWithContext mocks base method
func (m *MockIoTSecureTunnelingAPI) CloseTunnelWithContext(arg0 context.Context, arg1 *iotsecuretunneling.CloseTunnelInput, arg2 ...request.Option) (*iotsecuretunneling.CloseTunnelOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CloseTunnelWithContext", varargs...)
	ret0, _ := ret[0].(*iotsecuretunneling.CloseTunnelOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloseTunnelWithContext indicates an expected call of CloseTunnelWithContext
func (mr *MockIoTSecureTunnelingAPIMockRecorder) CloseTunnelWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseTunnelWithContext", reflect.TypeOf((*MockIoTSecureTunnelingAPI)(nil).CloseTunnelWithContext), varargs...)
}

// DescribeTunnel mocks base method
func (m *MockIoTSecureTunnelingAPI) DescribeTunnel(arg0 *iotsecuretunneling.DescribeTunnelInput) (*iotsecuretunneling.DescribeTunnelOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTunnel", arg0)
	ret0, _ := ret[0].(*iotsecuretunneling.DescribeTunnelOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTunnel indicates an expected call of DescribeTunnel
func (mr *MockIoTSecureTunnelingAPIMockRecorder) DescribeTunnel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTunnel", reflect.TypeOf((*MockIoTSecureTunnelingAPI)(nil).DescribeTunnel), arg0)
}

// DescribeTunnelRequest mocks base method
func (m *MockIoTSecureTunnelingAPI) DescribeTunnelRequest(arg0 *iotsecuretunneling.DescribeTunnelInput) (*request.Request, *iotsecuretunneling.DescribeTunnelOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTunnelRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iotsecuretunneling.DescribeTunnelOutput)
	return ret0, ret1
}

// DescribeTunnelRequest indicates an expected call of DescribeTunnelRequest
func (mr *MockIoTSecureTunnelingAPIMockRecorder) DescribeTunnelRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTunnelRequest", reflect.TypeOf((*MockIoTSecureTunnelingAPI)(nil).DescribeTunnelRequest), arg0)
}

// DescribeTunnelWithContext mocks base method
func (m *MockIoTSecureTunnelingAPI) DescribeTunnelWithContext(arg0 context.Context, arg1 *iotsecuretunneling.DescribeTunnelInput, arg2 ...request.Option) (*iotsecuretunneling.DescribeTunnelOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTunnelWithContext", varargs...)
	ret0, _ := ret[0].(*iotsecuretunneling.DescribeTunnelOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTunnelWithContext indicates an expected call of DescribeTunnelWithContext
func (mr *MockIoTSecureTunnelingAPIMockRecorder) DescribeTunnelWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTunnelWithContext", reflect.TypeOf((*MockIoTSecureTunnelingAPI)(nil).DescribeTunnelWithContext), varargs...)
}

// ListTagsForResource mocks base method
func (m *MockIoTSecureTunnelingAPI) ListTagsForResource(arg0 *iotsecuretunneling.ListTagsForResourceInput) (*iotsecuretunneling.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForResource", arg0)
	ret0, _ := ret[0].(*iotsecuretunneling.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource
func (mr *MockIoTSecureTunnelingAPIMockRecorder) ListTagsForResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockIoTSecureTunnelingAPI)(nil).ListTagsForResource), arg0)
}

// ListTagsForResourceRequest mocks base method
func (m *MockIoTSecureTunnelingAPI) ListTagsForResourceRequest(arg0 *iotsecuretunneling.ListTagsForResourceInput) (*request.Request, *iotsecuretunneling.ListTagsForResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iotsecuretunneling.ListTagsForResourceOutput)
	return ret0, ret1
}

// ListTagsForResourceRequest indicates an expected call of ListTagsForResourceRequest
func (mr *MockIoTSecureTunnelingAPIMockRecorder) ListTagsForResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResourceRequest", reflect.TypeOf((*MockIoTSecureTunnelingAPI)(nil).ListTagsForResourceRequest), arg0)
}

// ListTagsForResourceWithContext mocks base method
func (m *MockIoTSecureTunnelingAPI) ListTagsForResourceWithContext(arg0 context.Context, arg1 *iotsecuretunneling.ListTagsForResourceInput, arg2 ...request.Option) (*iotsecuretunneling.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResourceWithContext", varargs...)
	ret0, _ := ret[0].(*iotsecuretunneling.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResourceWithContext indicates an expected call of ListTagsForResourceWithContext
func (mr *MockIoTSecureTunnelingAPIMockRecorder) ListTagsForResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResourceWithContext", reflect.TypeOf((*MockIoTSecureTunnelingAPI)(nil).ListTagsForResourceWithContext), varargs...)
}

// ListTunnels mocks base method
func (m *MockIoTSecureTunnelingAPI) ListTunnels(arg0 *iotsecuretunneling.ListTunnelsInput) (*iotsecuretunneling.ListTunnelsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTunnels", arg0)
	ret0, _ := ret[0].(*iotsecuretunneling.ListTunnelsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTunnels indicates an expected call of ListTunnels
func (mr *MockIoTSecureTunnelingAPIMockRecorder) ListTunnels(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTunnels", reflect.TypeOf((*MockIoTSecureTunnelingAPI)(nil).ListTunnels), arg0)
}

// ListTunnelsPages mocks base method
func (m *MockIoTSecureTunnelingAPI) ListTunnelsPages(arg0 *iotsecuretunneling.ListTunnelsInput, arg1 func(*iotsecuretunneling.ListTunnelsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTunnelsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListTunnelsPages indicates an expected call of ListTunnelsPages
func (mr *MockIoTSecureTunnelingAPIMockRecorder) ListTunnelsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTunnelsPages", reflect.TypeOf((*MockIoTSecureTunnelingAPI)(nil).ListTunnelsPages), arg0, arg1)
}

// ListTunnelsPagesWithContext mocks base method
func (m *MockIoTSecureTunnelingAPI) ListTunnelsPagesWithContext(arg0 context.Context, arg1 *iotsecuretunneling.ListTunnelsInput, arg2 func(*iotsecuretunneling.ListTunnelsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTunnelsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListTunnelsPagesWithContext indicates an expected call of ListTunnelsPagesWithContext
func (mr *MockIoTSecureTunnelingAPIMockRecorder) ListTunnelsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTunnelsPagesWithContext", reflect.TypeOf((*MockIoTSecureTunnelingAPI)(nil).ListTunnelsPagesWithContext), varargs...)
}

// ListTunnelsRequest mocks base method
func (m *MockIoTSecureTunnelingAPI) ListTunnelsRequest(arg0 *iotsecuretunneling.ListTunnelsInput) (*request.Request, *iotsecuretunneling.ListTunnelsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTunnelsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iotsecuretunneling.ListTunnelsOutput)
	return ret0, ret1
}

// ListTunnelsRequest indicates an expected call of ListTunnelsRequest
func (mr *MockIoTSecureTunnelingAPIMockRecorder) ListTunnelsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTunnelsRequest", reflect.TypeOf((*MockIoTSecureTunnelingAPI)(nil).ListTunnelsRequest), arg0)
}

// ListTunnelsWithContext mocks base method
func (m *MockIoTSecureTunnelingAPI) ListTunnelsWithContext(arg0 context.Context, arg1 *iotsecuretunneling.ListTunnelsInput, arg2 ...request.Option) (*iotsecuretunneling.ListTunnelsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTunnelsWithContext", varargs...)
	ret0, _ := ret[0].(*iotsecuretunneling.ListTunnelsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTunnelsWithContext indicates an expected call of ListTunnelsWithContext
func (mr *MockIoTSecureTunnelingAPIMockRecorder) ListTunnelsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTunnelsWithContext", reflect.TypeOf((*MockIoTSecureTunnelingAPI)(nil).ListTunnelsWithContext), varargs...)
}

// OpenTunnel mocks base method
func (m *MockIoTSecureTunnelingAPI) OpenTunnel(arg0 *iotsecuretunneling.OpenTunnelInput) (*iotsecuretunneling.OpenTunnelOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenTunnel", arg0)
	ret0, _ := ret[0].(*iotsecuretunneling.OpenTunnelOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenTunnel indicates an expected call of OpenTunnel
func (mr *MockIoTSecureTunnelingAPIMockRecorder) OpenTunnel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenTunnel", reflect.TypeOf((*MockIoTSecureTunnelingAPI)(nil).OpenTunnel), arg0)
}

// OpenTunnelRequest mocks base method
func (m *MockIoTSecureTunnelingAPI) OpenTunnelRequest(arg0 *iotsecuretunneling.OpenTunnelInput) (*request.Request, *iotsecuretunneling.OpenTunnelOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenTunnelRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iotsecuretunneling.OpenTunnelOutput)
	return ret0, ret1
}

// OpenTunnelRequest indicates an expected call of OpenTunnelRequest
func (mr *MockIoTSecureTunnelingAPIMockRecorder) OpenTunnelRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenTunnelRequest", reflect.TypeOf((*MockIoTSecureTunnelingAPI)(nil).OpenTunnelRequest), arg0)
}

// OpenTunnelWithContext mocks base method
func (m *MockIoTSecureTunnelingAPI) OpenTunnelWithContext(arg0 context.Context, arg1 *iotsecuretunneling.OpenTunnelInput, arg2 ...request.Option) (*iotsecuretunneling.OpenTunnelOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "OpenTunnelWithContext", varargs...)
	ret0, _ := ret[0].(*iotsecuretunneling.OpenTunnelOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenTunnelWithContext indicates an expected call of OpenTunnelWithContext
func (mr *MockIoTSecureTunnelingAPIMockRecorder) OpenTunnelWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenTunnelWithContext", reflect.TypeOf((*MockIoTSecureTunnelingAPI)(nil).OpenTunnelWithContext), varargs...)
}

// RotateTunnelAccessToken mocks base method
func (m *MockIoTSecureTunnelingAPI) RotateTunnelAccessToken(arg0 *iotsecuretunneling.RotateTunnelAccessTokenInput) (*iotsecuretunneling.RotateTunnelAccessTokenOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RotateTunnelAccessToken", arg0)
	ret0, _ := ret[0].(*iotsecuretunneling.RotateTunnelAccessTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RotateTunnelAccessToken indicates an expected call of RotateTunnelAccessToken
func (mr *MockIoTSecureTunnelingAPIMockRecorder) RotateTunnelAccessToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RotateTunnelAccessToken", reflect.TypeOf((*MockIoTSecureTunnelingAPI)(nil).RotateTunnelAccessToken), arg0)
}

// RotateTunnelAccessTokenRequest mocks base method
func (m *MockIoTSecureTunnelingAPI) RotateTunnelAccessTokenRequest(arg0 *iotsecuretunneling.RotateTunnelAccessTokenInput) (*request.Request, *iotsecuretunneling.RotateTunnelAccessTokenOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RotateTunnelAccessTokenRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iotsecuretunneling.RotateTunnelAccessTokenOutput)
	return ret0, ret1
}

// RotateTunnelAccessTokenRequest indicates an expected call of RotateTunnelAccessTokenRequest
func (mr *MockIoTSecureTunnelingAPIMockRecorder) RotateTunnelAccessTokenRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RotateTunnelAccessTokenRequest", reflect.TypeOf((*MockIoTSecureTunnelingAPI)(nil).RotateTunnelAccessTokenRequest), arg0)
}

// RotateTunnelAccessTokenWithContext mocks base method
func (m *MockIoTSecureTunnelingAPI) RotateTunnelAccessTokenWithContext(arg0 context.Context, arg1 *iotsecuretunneling.RotateTunnelAccessTokenInput, arg2 ...request.Option) (*iotsecuretunneling.RotateTunnelAccessTokenOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RotateTunnelAccessTokenWithContext", varargs...)
	ret0, _ := ret[0].(*iotsecuretunneling.RotateTunnelAccessTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RotateTunnelAccessTokenWithContext indicates an expected call of RotateTunnelAccessTokenWithContext
func (mr *MockIoTSecureTunnelingAPIMockRecorder) RotateTunnelAccessTokenWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RotateTunnelAccessTokenWithContext", reflect.TypeOf((*MockIoTSecureTunnelingAPI)(nil).RotateTunnelAccessTokenWithContext), varargs...)
}

// TagResource mocks base method
func (m *MockIoTSecureTunnelingAPI) TagResource(arg0 *iotsecuretunneling.TagResourceInput) (*iotsecuretunneling.TagResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagResource", arg0)
	ret0, _ := ret[0].(*iotsecuretunneling.TagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagResource indicates an expected call of TagResource
func (mr *MockIoTSecureTunnelingAPIMockRecorder) TagResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResource", reflect.TypeOf((*MockIoTSecureTunnelingAPI)(nil).TagResource), arg0)
}

// TagResourceRequest mocks base method
func (m *MockIoTSecureTunnelingAPI) TagResourceRequest(arg0 *iotsecuretunneling.TagResourceInput) (*request.Request, *iotsecuretunneling.TagResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iotsecuretunneling.TagResourceOutput)
	return ret0, ret1
}

// TagResourceRequest indicates an expected call of TagResourceRequest
func (mr *MockIoTSecureTunnelingAPIMockRecorder) TagResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResourceRequest", reflect.TypeOf((*MockIoTSecureTunnelingAPI)(nil).TagResourceRequest), arg0)
}

// TagResourceWithContext mocks base method
func (m *MockIoTSecureTunnelingAPI) TagResourceWithContext(arg0 context.Context, arg1 *iotsecuretunneling.TagResourceInput, arg2 ...request.Option) (*iotsecuretunneling.TagResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TagResourceWithContext", varargs...)
	ret0, _ := ret[0].(*iotsecuretunneling.TagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagResourceWithContext indicates an expected call of TagResourceWithContext
func (mr *MockIoTSecureTunnelingAPIMockRecorder) TagResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResourceWithContext", reflect.TypeOf((*MockIoTSecureTunnelingAPI)(nil).TagResourceWithContext), varargs...)
}

// UntagResource mocks base method
func (m *MockIoTSecureTunnelingAPI) UntagResource(arg0 *iotsecuretunneling.UntagResourceInput) (*iotsecuretunneling.UntagResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagResource", arg0)
	ret0, _ := ret[0].(*iotsecuretunneling.UntagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagResource indicates an expected call of UntagResource
func (mr *MockIoTSecureTunnelingAPIMockRecorder) UntagResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResource", reflect.TypeOf((*MockIoTSecureTunnelingAPI)(nil).UntagResource), arg0)
}

// UntagResourceRequest mocks base method
func (m *MockIoTSecureTunnelingAPI) UntagResourceRequest(arg0 *iotsecuretunneling.UntagResourceInput) (*request.Request, *iotsecuretunneling.UntagResourceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagResourceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*iotsecuretunneling.UntagResourceOutput)
	return ret0, ret1
}

// UntagResourceRequest indicates an expected call of UntagResourceRequest
func (mr *MockIoTSecureTunnelingAPIMockRecorder) UntagResourceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResourceRequest", reflect.TypeOf((*MockIoTSecureTunnelingAPI)(nil).UntagResourceRequest), arg0)
}

// UntagResourceWithContext mocks base method
func (m *MockIoTSecureTunnelingAPI) UntagResourceWithContext(arg0 context.Context, arg1 *iotsecuretunneling.UntagResourceInput, arg2 ...request.Option) (*iotsecuretunneling.UntagResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UntagResourceWithContext", varargs...)
	ret0, _ := ret[0].(*iotsecuretunneling.UntagResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagResourceWithContext indicates an expected call of UntagResourceWithContext
func (mr *MockIoTSecureTunnelingAPIMockRecorder) UntagResourceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResourceWithContext", reflect.TypeOf((*MockIoTSecureTunnelingAPI)(nil).UntagResourceWithContext), varargs...)
}
