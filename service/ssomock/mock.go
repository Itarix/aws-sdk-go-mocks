// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/sso/ssoiface (interfaces: SSOAPI)

// Package ssomock is a generated GoMock package.
package ssomock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	sso "github.com/aws/aws-sdk-go/service/sso"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSSOAPI is a mock of SSOAPI interface
type MockSSOAPI struct {
	ctrl     *gomock.Controller
	recorder *MockSSOAPIMockRecorder
}

// MockSSOAPIMockRecorder is the mock recorder for MockSSOAPI
type MockSSOAPIMockRecorder struct {
	mock *MockSSOAPI
}

// NewMockSSOAPI creates a new mock instance
func NewMockSSOAPI(ctrl *gomock.Controller) *MockSSOAPI {
	mock := &MockSSOAPI{ctrl: ctrl}
	mock.recorder = &MockSSOAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSSOAPI) EXPECT() *MockSSOAPIMockRecorder {
	return m.recorder
}

// GetRoleCredentials mocks base method
func (m *MockSSOAPI) GetRoleCredentials(arg0 *sso.GetRoleCredentialsInput) (*sso.GetRoleCredentialsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleCredentials", arg0)
	ret0, _ := ret[0].(*sso.GetRoleCredentialsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoleCredentials indicates an expected call of GetRoleCredentials
func (mr *MockSSOAPIMockRecorder) GetRoleCredentials(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleCredentials", reflect.TypeOf((*MockSSOAPI)(nil).GetRoleCredentials), arg0)
}

// GetRoleCredentialsRequest mocks base method
func (m *MockSSOAPI) GetRoleCredentialsRequest(arg0 *sso.GetRoleCredentialsInput) (*request.Request, *sso.GetRoleCredentialsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleCredentialsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sso.GetRoleCredentialsOutput)
	return ret0, ret1
}

// GetRoleCredentialsRequest indicates an expected call of GetRoleCredentialsRequest
func (mr *MockSSOAPIMockRecorder) GetRoleCredentialsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleCredentialsRequest", reflect.TypeOf((*MockSSOAPI)(nil).GetRoleCredentialsRequest), arg0)
}

// GetRoleCredentialsWithContext mocks base method
func (m *MockSSOAPI) GetRoleCredentialsWithContext(arg0 context.Context, arg1 *sso.GetRoleCredentialsInput, arg2 ...request.Option) (*sso.GetRoleCredentialsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRoleCredentialsWithContext", varargs...)
	ret0, _ := ret[0].(*sso.GetRoleCredentialsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoleCredentialsWithContext indicates an expected call of GetRoleCredentialsWithContext
func (mr *MockSSOAPIMockRecorder) GetRoleCredentialsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleCredentialsWithContext", reflect.TypeOf((*MockSSOAPI)(nil).GetRoleCredentialsWithContext), varargs...)
}

// ListAccountRoles mocks base method
func (m *MockSSOAPI) ListAccountRoles(arg0 *sso.ListAccountRolesInput) (*sso.ListAccountRolesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccountRoles", arg0)
	ret0, _ := ret[0].(*sso.ListAccountRolesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccountRoles indicates an expected call of ListAccountRoles
func (mr *MockSSOAPIMockRecorder) ListAccountRoles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccountRoles", reflect.TypeOf((*MockSSOAPI)(nil).ListAccountRoles), arg0)
}

// ListAccountRolesPages mocks base method
func (m *MockSSOAPI) ListAccountRolesPages(arg0 *sso.ListAccountRolesInput, arg1 func(*sso.ListAccountRolesOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccountRolesPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListAccountRolesPages indicates an expected call of ListAccountRolesPages
func (mr *MockSSOAPIMockRecorder) ListAccountRolesPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccountRolesPages", reflect.TypeOf((*MockSSOAPI)(nil).ListAccountRolesPages), arg0, arg1)
}

// ListAccountRolesPagesWithContext mocks base method
func (m *MockSSOAPI) ListAccountRolesPagesWithContext(arg0 context.Context, arg1 *sso.ListAccountRolesInput, arg2 func(*sso.ListAccountRolesOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAccountRolesPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListAccountRolesPagesWithContext indicates an expected call of ListAccountRolesPagesWithContext
func (mr *MockSSOAPIMockRecorder) ListAccountRolesPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccountRolesPagesWithContext", reflect.TypeOf((*MockSSOAPI)(nil).ListAccountRolesPagesWithContext), varargs...)
}

// ListAccountRolesRequest mocks base method
func (m *MockSSOAPI) ListAccountRolesRequest(arg0 *sso.ListAccountRolesInput) (*request.Request, *sso.ListAccountRolesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccountRolesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sso.ListAccountRolesOutput)
	return ret0, ret1
}

// ListAccountRolesRequest indicates an expected call of ListAccountRolesRequest
func (mr *MockSSOAPIMockRecorder) ListAccountRolesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccountRolesRequest", reflect.TypeOf((*MockSSOAPI)(nil).ListAccountRolesRequest), arg0)
}

// ListAccountRolesWithContext mocks base method
func (m *MockSSOAPI) ListAccountRolesWithContext(arg0 context.Context, arg1 *sso.ListAccountRolesInput, arg2 ...request.Option) (*sso.ListAccountRolesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAccountRolesWithContext", varargs...)
	ret0, _ := ret[0].(*sso.ListAccountRolesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccountRolesWithContext indicates an expected call of ListAccountRolesWithContext
func (mr *MockSSOAPIMockRecorder) ListAccountRolesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccountRolesWithContext", reflect.TypeOf((*MockSSOAPI)(nil).ListAccountRolesWithContext), varargs...)
}

// ListAccounts mocks base method
func (m *MockSSOAPI) ListAccounts(arg0 *sso.ListAccountsInput) (*sso.ListAccountsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccounts", arg0)
	ret0, _ := ret[0].(*sso.ListAccountsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccounts indicates an expected call of ListAccounts
func (mr *MockSSOAPIMockRecorder) ListAccounts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccounts", reflect.TypeOf((*MockSSOAPI)(nil).ListAccounts), arg0)
}

// ListAccountsPages mocks base method
func (m *MockSSOAPI) ListAccountsPages(arg0 *sso.ListAccountsInput, arg1 func(*sso.ListAccountsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccountsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListAccountsPages indicates an expected call of ListAccountsPages
func (mr *MockSSOAPIMockRecorder) ListAccountsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccountsPages", reflect.TypeOf((*MockSSOAPI)(nil).ListAccountsPages), arg0, arg1)
}

// ListAccountsPagesWithContext mocks base method
func (m *MockSSOAPI) ListAccountsPagesWithContext(arg0 context.Context, arg1 *sso.ListAccountsInput, arg2 func(*sso.ListAccountsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAccountsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListAccountsPagesWithContext indicates an expected call of ListAccountsPagesWithContext
func (mr *MockSSOAPIMockRecorder) ListAccountsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccountsPagesWithContext", reflect.TypeOf((*MockSSOAPI)(nil).ListAccountsPagesWithContext), varargs...)
}

// ListAccountsRequest mocks base method
func (m *MockSSOAPI) ListAccountsRequest(arg0 *sso.ListAccountsInput) (*request.Request, *sso.ListAccountsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccountsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sso.ListAccountsOutput)
	return ret0, ret1
}

// ListAccountsRequest indicates an expected call of ListAccountsRequest
func (mr *MockSSOAPIMockRecorder) ListAccountsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccountsRequest", reflect.TypeOf((*MockSSOAPI)(nil).ListAccountsRequest), arg0)
}

// ListAccountsWithContext mocks base method
func (m *MockSSOAPI) ListAccountsWithContext(arg0 context.Context, arg1 *sso.ListAccountsInput, arg2 ...request.Option) (*sso.ListAccountsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAccountsWithContext", varargs...)
	ret0, _ := ret[0].(*sso.ListAccountsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccountsWithContext indicates an expected call of ListAccountsWithContext
func (mr *MockSSOAPIMockRecorder) ListAccountsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccountsWithContext", reflect.TypeOf((*MockSSOAPI)(nil).ListAccountsWithContext), varargs...)
}

// Logout mocks base method
func (m *MockSSOAPI) Logout(arg0 *sso.LogoutInput) (*sso.LogoutOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", arg0)
	ret0, _ := ret[0].(*sso.LogoutOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Logout indicates an expected call of Logout
func (mr *MockSSOAPIMockRecorder) Logout(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockSSOAPI)(nil).Logout), arg0)
}

// LogoutRequest mocks base method
func (m *MockSSOAPI) LogoutRequest(arg0 *sso.LogoutInput) (*request.Request, *sso.LogoutOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogoutRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sso.LogoutOutput)
	return ret0, ret1
}

// LogoutRequest indicates an expected call of LogoutRequest
func (mr *MockSSOAPIMockRecorder) LogoutRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogoutRequest", reflect.TypeOf((*MockSSOAPI)(nil).LogoutRequest), arg0)
}

// LogoutWithContext mocks base method
func (m *MockSSOAPI) LogoutWithContext(arg0 context.Context, arg1 *sso.LogoutInput, arg2 ...request.Option) (*sso.LogoutOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LogoutWithContext", varargs...)
	ret0, _ := ret[0].(*sso.LogoutOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LogoutWithContext indicates an expected call of LogoutWithContext
func (mr *MockSSOAPIMockRecorder) LogoutWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogoutWithContext", reflect.TypeOf((*MockSSOAPI)(nil).LogoutWithContext), varargs...)
}
