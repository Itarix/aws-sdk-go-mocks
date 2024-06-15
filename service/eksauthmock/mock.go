// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/eksauth/eksauthiface (interfaces: EKSAuthAPI)

// Package eksauthmock is a generated GoMock package.
package eksauthmock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	eksauth "github.com/aws/aws-sdk-go/service/eksauth"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockEKSAuthAPI is a mock of EKSAuthAPI interface
type MockEKSAuthAPI struct {
	ctrl     *gomock.Controller
	recorder *MockEKSAuthAPIMockRecorder
}

// MockEKSAuthAPIMockRecorder is the mock recorder for MockEKSAuthAPI
type MockEKSAuthAPIMockRecorder struct {
	mock *MockEKSAuthAPI
}

// NewMockEKSAuthAPI creates a new mock instance
func NewMockEKSAuthAPI(ctrl *gomock.Controller) *MockEKSAuthAPI {
	mock := &MockEKSAuthAPI{ctrl: ctrl}
	mock.recorder = &MockEKSAuthAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEKSAuthAPI) EXPECT() *MockEKSAuthAPIMockRecorder {
	return m.recorder
}

// AssumeRoleForPodIdentity mocks base method
func (m *MockEKSAuthAPI) AssumeRoleForPodIdentity(arg0 *eksauth.AssumeRoleForPodIdentityInput) (*eksauth.AssumeRoleForPodIdentityOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssumeRoleForPodIdentity", arg0)
	ret0, _ := ret[0].(*eksauth.AssumeRoleForPodIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssumeRoleForPodIdentity indicates an expected call of AssumeRoleForPodIdentity
func (mr *MockEKSAuthAPIMockRecorder) AssumeRoleForPodIdentity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssumeRoleForPodIdentity", reflect.TypeOf((*MockEKSAuthAPI)(nil).AssumeRoleForPodIdentity), arg0)
}

// AssumeRoleForPodIdentityRequest mocks base method
func (m *MockEKSAuthAPI) AssumeRoleForPodIdentityRequest(arg0 *eksauth.AssumeRoleForPodIdentityInput) (*request.Request, *eksauth.AssumeRoleForPodIdentityOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssumeRoleForPodIdentityRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*eksauth.AssumeRoleForPodIdentityOutput)
	return ret0, ret1
}

// AssumeRoleForPodIdentityRequest indicates an expected call of AssumeRoleForPodIdentityRequest
func (mr *MockEKSAuthAPIMockRecorder) AssumeRoleForPodIdentityRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssumeRoleForPodIdentityRequest", reflect.TypeOf((*MockEKSAuthAPI)(nil).AssumeRoleForPodIdentityRequest), arg0)
}

// AssumeRoleForPodIdentityWithContext mocks base method
func (m *MockEKSAuthAPI) AssumeRoleForPodIdentityWithContext(arg0 context.Context, arg1 *eksauth.AssumeRoleForPodIdentityInput, arg2 ...request.Option) (*eksauth.AssumeRoleForPodIdentityOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AssumeRoleForPodIdentityWithContext", varargs...)
	ret0, _ := ret[0].(*eksauth.AssumeRoleForPodIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssumeRoleForPodIdentityWithContext indicates an expected call of AssumeRoleForPodIdentityWithContext
func (mr *MockEKSAuthAPIMockRecorder) AssumeRoleForPodIdentityWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssumeRoleForPodIdentityWithContext", reflect.TypeOf((*MockEKSAuthAPI)(nil).AssumeRoleForPodIdentityWithContext), varargs...)
}
