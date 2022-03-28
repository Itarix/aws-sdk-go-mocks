// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/apigatewaymanagementapi/apigatewaymanagementapiiface (interfaces: ApiGatewayManagementApiAPI)

// Package apigatewaymanagementapimock is a generated GoMock package.
package apigatewaymanagementapimock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	apigatewaymanagementapi "github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockApiGatewayManagementApiAPI is a mock of ApiGatewayManagementApiAPI interface
type MockApiGatewayManagementApiAPI struct {
	ctrl     *gomock.Controller
	recorder *MockApiGatewayManagementApiAPIMockRecorder
}

// MockApiGatewayManagementApiAPIMockRecorder is the mock recorder for MockApiGatewayManagementApiAPI
type MockApiGatewayManagementApiAPIMockRecorder struct {
	mock *MockApiGatewayManagementApiAPI
}

// NewMockApiGatewayManagementApiAPI creates a new mock instance
func NewMockApiGatewayManagementApiAPI(ctrl *gomock.Controller) *MockApiGatewayManagementApiAPI {
	mock := &MockApiGatewayManagementApiAPI{ctrl: ctrl}
	mock.recorder = &MockApiGatewayManagementApiAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApiGatewayManagementApiAPI) EXPECT() *MockApiGatewayManagementApiAPIMockRecorder {
	return m.recorder
}

// PostToConnection mocks base method
func (m *MockApiGatewayManagementApiAPI) PostToConnection(arg0 *apigatewaymanagementapi.PostToConnectionInput) (*apigatewaymanagementapi.PostToConnectionOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostToConnection", arg0)
	ret0, _ := ret[0].(*apigatewaymanagementapi.PostToConnectionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostToConnection indicates an expected call of PostToConnection
func (mr *MockApiGatewayManagementApiAPIMockRecorder) PostToConnection(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostToConnection", reflect.TypeOf((*MockApiGatewayManagementApiAPI)(nil).PostToConnection), arg0)
}

// PostToConnectionRequest mocks base method
func (m *MockApiGatewayManagementApiAPI) PostToConnectionRequest(arg0 *apigatewaymanagementapi.PostToConnectionInput) (*request.Request, *apigatewaymanagementapi.PostToConnectionOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostToConnectionRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*apigatewaymanagementapi.PostToConnectionOutput)
	return ret0, ret1
}

// PostToConnectionRequest indicates an expected call of PostToConnectionRequest
func (mr *MockApiGatewayManagementApiAPIMockRecorder) PostToConnectionRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostToConnectionRequest", reflect.TypeOf((*MockApiGatewayManagementApiAPI)(nil).PostToConnectionRequest), arg0)
}

// PostToConnectionWithContext mocks base method
func (m *MockApiGatewayManagementApiAPI) PostToConnectionWithContext(arg0 context.Context, arg1 *apigatewaymanagementapi.PostToConnectionInput, arg2 ...request.Option) (*apigatewaymanagementapi.PostToConnectionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostToConnectionWithContext", varargs...)
	ret0, _ := ret[0].(*apigatewaymanagementapi.PostToConnectionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostToConnectionWithContext indicates an expected call of PostToConnectionWithContext
func (mr *MockApiGatewayManagementApiAPIMockRecorder) PostToConnectionWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostToConnectionWithContext", reflect.TypeOf((*MockApiGatewayManagementApiAPI)(nil).PostToConnectionWithContext), varargs...)
}
