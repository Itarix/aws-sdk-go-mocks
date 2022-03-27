// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/identitystore/identitystoreiface (interfaces: IdentityStoreAPI)

// Package identitystoremock is a generated GoMock package.
package identitystoremock

import (
	context "context"
	reflect "reflect"

	request "github.com/aws/aws-sdk-go/aws/request"
	identitystore "github.com/aws/aws-sdk-go/service/identitystore"
	gomock "github.com/golang/mock/gomock"
)

// MockIdentityStoreAPI is a mock of IdentityStoreAPI interface.
type MockIdentityStoreAPI struct {
	ctrl     *gomock.Controller
	recorder *MockIdentityStoreAPIMockRecorder
}

// MockIdentityStoreAPIMockRecorder is the mock recorder for MockIdentityStoreAPI.
type MockIdentityStoreAPIMockRecorder struct {
	mock *MockIdentityStoreAPI
}

// NewMockIdentityStoreAPI creates a new mock instance.
func NewMockIdentityStoreAPI(ctrl *gomock.Controller) *MockIdentityStoreAPI {
	mock := &MockIdentityStoreAPI{ctrl: ctrl}
	mock.recorder = &MockIdentityStoreAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIdentityStoreAPI) EXPECT() *MockIdentityStoreAPIMockRecorder {
	return m.recorder
}

// DescribeGroup mocks base method.
func (m *MockIdentityStoreAPI) DescribeGroup(arg0 *identitystore.DescribeGroupInput) (*identitystore.DescribeGroupOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeGroup", arg0)
	ret0, _ := ret[0].(*identitystore.DescribeGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeGroup indicates an expected call of DescribeGroup.
func (mr *MockIdentityStoreAPIMockRecorder) DescribeGroup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeGroup", reflect.TypeOf((*MockIdentityStoreAPI)(nil).DescribeGroup), arg0)
}

// DescribeGroupRequest mocks base method.
func (m *MockIdentityStoreAPI) DescribeGroupRequest(arg0 *identitystore.DescribeGroupInput) (*request.Request, *identitystore.DescribeGroupOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeGroupRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*identitystore.DescribeGroupOutput)
	return ret0, ret1
}

// DescribeGroupRequest indicates an expected call of DescribeGroupRequest.
func (mr *MockIdentityStoreAPIMockRecorder) DescribeGroupRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeGroupRequest", reflect.TypeOf((*MockIdentityStoreAPI)(nil).DescribeGroupRequest), arg0)
}

// DescribeGroupWithContext mocks base method.
func (m *MockIdentityStoreAPI) DescribeGroupWithContext(arg0 context.Context, arg1 *identitystore.DescribeGroupInput, arg2 ...request.Option) (*identitystore.DescribeGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeGroupWithContext", varargs...)
	ret0, _ := ret[0].(*identitystore.DescribeGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeGroupWithContext indicates an expected call of DescribeGroupWithContext.
func (mr *MockIdentityStoreAPIMockRecorder) DescribeGroupWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeGroupWithContext", reflect.TypeOf((*MockIdentityStoreAPI)(nil).DescribeGroupWithContext), varargs...)
}

// DescribeUser mocks base method.
func (m *MockIdentityStoreAPI) DescribeUser(arg0 *identitystore.DescribeUserInput) (*identitystore.DescribeUserOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeUser", arg0)
	ret0, _ := ret[0].(*identitystore.DescribeUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeUser indicates an expected call of DescribeUser.
func (mr *MockIdentityStoreAPIMockRecorder) DescribeUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeUser", reflect.TypeOf((*MockIdentityStoreAPI)(nil).DescribeUser), arg0)
}

// DescribeUserRequest mocks base method.
func (m *MockIdentityStoreAPI) DescribeUserRequest(arg0 *identitystore.DescribeUserInput) (*request.Request, *identitystore.DescribeUserOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeUserRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*identitystore.DescribeUserOutput)
	return ret0, ret1
}

// DescribeUserRequest indicates an expected call of DescribeUserRequest.
func (mr *MockIdentityStoreAPIMockRecorder) DescribeUserRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeUserRequest", reflect.TypeOf((*MockIdentityStoreAPI)(nil).DescribeUserRequest), arg0)
}

// DescribeUserWithContext mocks base method.
func (m *MockIdentityStoreAPI) DescribeUserWithContext(arg0 context.Context, arg1 *identitystore.DescribeUserInput, arg2 ...request.Option) (*identitystore.DescribeUserOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeUserWithContext", varargs...)
	ret0, _ := ret[0].(*identitystore.DescribeUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeUserWithContext indicates an expected call of DescribeUserWithContext.
func (mr *MockIdentityStoreAPIMockRecorder) DescribeUserWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeUserWithContext", reflect.TypeOf((*MockIdentityStoreAPI)(nil).DescribeUserWithContext), varargs...)
}

// ListGroups mocks base method.
func (m *MockIdentityStoreAPI) ListGroups(arg0 *identitystore.ListGroupsInput) (*identitystore.ListGroupsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGroups", arg0)
	ret0, _ := ret[0].(*identitystore.ListGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroups indicates an expected call of ListGroups.
func (mr *MockIdentityStoreAPIMockRecorder) ListGroups(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroups", reflect.TypeOf((*MockIdentityStoreAPI)(nil).ListGroups), arg0)
}

// ListGroupsPages mocks base method.
func (m *MockIdentityStoreAPI) ListGroupsPages(arg0 *identitystore.ListGroupsInput, arg1 func(*identitystore.ListGroupsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGroupsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListGroupsPages indicates an expected call of ListGroupsPages.
func (mr *MockIdentityStoreAPIMockRecorder) ListGroupsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupsPages", reflect.TypeOf((*MockIdentityStoreAPI)(nil).ListGroupsPages), arg0, arg1)
}

// ListGroupsPagesWithContext mocks base method.
func (m *MockIdentityStoreAPI) ListGroupsPagesWithContext(arg0 context.Context, arg1 *identitystore.ListGroupsInput, arg2 func(*identitystore.ListGroupsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGroupsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListGroupsPagesWithContext indicates an expected call of ListGroupsPagesWithContext.
func (mr *MockIdentityStoreAPIMockRecorder) ListGroupsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupsPagesWithContext", reflect.TypeOf((*MockIdentityStoreAPI)(nil).ListGroupsPagesWithContext), varargs...)
}

// ListGroupsRequest mocks base method.
func (m *MockIdentityStoreAPI) ListGroupsRequest(arg0 *identitystore.ListGroupsInput) (*request.Request, *identitystore.ListGroupsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGroupsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*identitystore.ListGroupsOutput)
	return ret0, ret1
}

// ListGroupsRequest indicates an expected call of ListGroupsRequest.
func (mr *MockIdentityStoreAPIMockRecorder) ListGroupsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupsRequest", reflect.TypeOf((*MockIdentityStoreAPI)(nil).ListGroupsRequest), arg0)
}

// ListGroupsWithContext mocks base method.
func (m *MockIdentityStoreAPI) ListGroupsWithContext(arg0 context.Context, arg1 *identitystore.ListGroupsInput, arg2 ...request.Option) (*identitystore.ListGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGroupsWithContext", varargs...)
	ret0, _ := ret[0].(*identitystore.ListGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroupsWithContext indicates an expected call of ListGroupsWithContext.
func (mr *MockIdentityStoreAPIMockRecorder) ListGroupsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupsWithContext", reflect.TypeOf((*MockIdentityStoreAPI)(nil).ListGroupsWithContext), varargs...)
}

// ListUsers mocks base method.
func (m *MockIdentityStoreAPI) ListUsers(arg0 *identitystore.ListUsersInput) (*identitystore.ListUsersOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", arg0)
	ret0, _ := ret[0].(*identitystore.ListUsersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockIdentityStoreAPIMockRecorder) ListUsers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockIdentityStoreAPI)(nil).ListUsers), arg0)
}

// ListUsersPages mocks base method.
func (m *MockIdentityStoreAPI) ListUsersPages(arg0 *identitystore.ListUsersInput, arg1 func(*identitystore.ListUsersOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsersPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListUsersPages indicates an expected call of ListUsersPages.
func (mr *MockIdentityStoreAPIMockRecorder) ListUsersPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsersPages", reflect.TypeOf((*MockIdentityStoreAPI)(nil).ListUsersPages), arg0, arg1)
}

// ListUsersPagesWithContext mocks base method.
func (m *MockIdentityStoreAPI) ListUsersPagesWithContext(arg0 context.Context, arg1 *identitystore.ListUsersInput, arg2 func(*identitystore.ListUsersOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListUsersPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListUsersPagesWithContext indicates an expected call of ListUsersPagesWithContext.
func (mr *MockIdentityStoreAPIMockRecorder) ListUsersPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsersPagesWithContext", reflect.TypeOf((*MockIdentityStoreAPI)(nil).ListUsersPagesWithContext), varargs...)
}

// ListUsersRequest mocks base method.
func (m *MockIdentityStoreAPI) ListUsersRequest(arg0 *identitystore.ListUsersInput) (*request.Request, *identitystore.ListUsersOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsersRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*identitystore.ListUsersOutput)
	return ret0, ret1
}

// ListUsersRequest indicates an expected call of ListUsersRequest.
func (mr *MockIdentityStoreAPIMockRecorder) ListUsersRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsersRequest", reflect.TypeOf((*MockIdentityStoreAPI)(nil).ListUsersRequest), arg0)
}

// ListUsersWithContext mocks base method.
func (m *MockIdentityStoreAPI) ListUsersWithContext(arg0 context.Context, arg1 *identitystore.ListUsersInput, arg2 ...request.Option) (*identitystore.ListUsersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListUsersWithContext", varargs...)
	ret0, _ := ret[0].(*identitystore.ListUsersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsersWithContext indicates an expected call of ListUsersWithContext.
func (mr *MockIdentityStoreAPIMockRecorder) ListUsersWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsersWithContext", reflect.TypeOf((*MockIdentityStoreAPI)(nil).ListUsersWithContext), varargs...)
}
