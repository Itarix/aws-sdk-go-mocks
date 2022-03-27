// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/lexruntimev2/lexruntimev2iface (interfaces: LexRuntimeV2API)

// Package lexruntimev2mock is a generated GoMock package.
package lexruntimev2mock

import (
	context "context"
	reflect "reflect"

	request "github.com/aws/aws-sdk-go/aws/request"
	lexruntimev2 "github.com/aws/aws-sdk-go/service/lexruntimev2"
	gomock "github.com/golang/mock/gomock"
)

// MockLexRuntimeV2API is a mock of LexRuntimeV2API interface.
type MockLexRuntimeV2API struct {
	ctrl     *gomock.Controller
	recorder *MockLexRuntimeV2APIMockRecorder
}

// MockLexRuntimeV2APIMockRecorder is the mock recorder for MockLexRuntimeV2API.
type MockLexRuntimeV2APIMockRecorder struct {
	mock *MockLexRuntimeV2API
}

// NewMockLexRuntimeV2API creates a new mock instance.
func NewMockLexRuntimeV2API(ctrl *gomock.Controller) *MockLexRuntimeV2API {
	mock := &MockLexRuntimeV2API{ctrl: ctrl}
	mock.recorder = &MockLexRuntimeV2APIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLexRuntimeV2API) EXPECT() *MockLexRuntimeV2APIMockRecorder {
	return m.recorder
}

// DeleteSession mocks base method.
func (m *MockLexRuntimeV2API) DeleteSession(arg0 *lexruntimev2.DeleteSessionInput) (*lexruntimev2.DeleteSessionOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", arg0)
	ret0, _ := ret[0].(*lexruntimev2.DeleteSessionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockLexRuntimeV2APIMockRecorder) DeleteSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockLexRuntimeV2API)(nil).DeleteSession), arg0)
}

// DeleteSessionRequest mocks base method.
func (m *MockLexRuntimeV2API) DeleteSessionRequest(arg0 *lexruntimev2.DeleteSessionInput) (*request.Request, *lexruntimev2.DeleteSessionOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSessionRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lexruntimev2.DeleteSessionOutput)
	return ret0, ret1
}

// DeleteSessionRequest indicates an expected call of DeleteSessionRequest.
func (mr *MockLexRuntimeV2APIMockRecorder) DeleteSessionRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSessionRequest", reflect.TypeOf((*MockLexRuntimeV2API)(nil).DeleteSessionRequest), arg0)
}

// DeleteSessionWithContext mocks base method.
func (m *MockLexRuntimeV2API) DeleteSessionWithContext(arg0 context.Context, arg1 *lexruntimev2.DeleteSessionInput, arg2 ...request.Option) (*lexruntimev2.DeleteSessionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteSessionWithContext", varargs...)
	ret0, _ := ret[0].(*lexruntimev2.DeleteSessionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSessionWithContext indicates an expected call of DeleteSessionWithContext.
func (mr *MockLexRuntimeV2APIMockRecorder) DeleteSessionWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSessionWithContext", reflect.TypeOf((*MockLexRuntimeV2API)(nil).DeleteSessionWithContext), varargs...)
}

// GetSession mocks base method.
func (m *MockLexRuntimeV2API) GetSession(arg0 *lexruntimev2.GetSessionInput) (*lexruntimev2.GetSessionOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0)
	ret0, _ := ret[0].(*lexruntimev2.GetSessionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockLexRuntimeV2APIMockRecorder) GetSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockLexRuntimeV2API)(nil).GetSession), arg0)
}

// GetSessionRequest mocks base method.
func (m *MockLexRuntimeV2API) GetSessionRequest(arg0 *lexruntimev2.GetSessionInput) (*request.Request, *lexruntimev2.GetSessionOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lexruntimev2.GetSessionOutput)
	return ret0, ret1
}

// GetSessionRequest indicates an expected call of GetSessionRequest.
func (mr *MockLexRuntimeV2APIMockRecorder) GetSessionRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionRequest", reflect.TypeOf((*MockLexRuntimeV2API)(nil).GetSessionRequest), arg0)
}

// GetSessionWithContext mocks base method.
func (m *MockLexRuntimeV2API) GetSessionWithContext(arg0 context.Context, arg1 *lexruntimev2.GetSessionInput, arg2 ...request.Option) (*lexruntimev2.GetSessionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSessionWithContext", varargs...)
	ret0, _ := ret[0].(*lexruntimev2.GetSessionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionWithContext indicates an expected call of GetSessionWithContext.
func (mr *MockLexRuntimeV2APIMockRecorder) GetSessionWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionWithContext", reflect.TypeOf((*MockLexRuntimeV2API)(nil).GetSessionWithContext), varargs...)
}

// PutSession mocks base method.
func (m *MockLexRuntimeV2API) PutSession(arg0 *lexruntimev2.PutSessionInput) (*lexruntimev2.PutSessionOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutSession", arg0)
	ret0, _ := ret[0].(*lexruntimev2.PutSessionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutSession indicates an expected call of PutSession.
func (mr *MockLexRuntimeV2APIMockRecorder) PutSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutSession", reflect.TypeOf((*MockLexRuntimeV2API)(nil).PutSession), arg0)
}

// PutSessionRequest mocks base method.
func (m *MockLexRuntimeV2API) PutSessionRequest(arg0 *lexruntimev2.PutSessionInput) (*request.Request, *lexruntimev2.PutSessionOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutSessionRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lexruntimev2.PutSessionOutput)
	return ret0, ret1
}

// PutSessionRequest indicates an expected call of PutSessionRequest.
func (mr *MockLexRuntimeV2APIMockRecorder) PutSessionRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutSessionRequest", reflect.TypeOf((*MockLexRuntimeV2API)(nil).PutSessionRequest), arg0)
}

// PutSessionWithContext mocks base method.
func (m *MockLexRuntimeV2API) PutSessionWithContext(arg0 context.Context, arg1 *lexruntimev2.PutSessionInput, arg2 ...request.Option) (*lexruntimev2.PutSessionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutSessionWithContext", varargs...)
	ret0, _ := ret[0].(*lexruntimev2.PutSessionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutSessionWithContext indicates an expected call of PutSessionWithContext.
func (mr *MockLexRuntimeV2APIMockRecorder) PutSessionWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutSessionWithContext", reflect.TypeOf((*MockLexRuntimeV2API)(nil).PutSessionWithContext), varargs...)
}

// RecognizeText mocks base method.
func (m *MockLexRuntimeV2API) RecognizeText(arg0 *lexruntimev2.RecognizeTextInput) (*lexruntimev2.RecognizeTextOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecognizeText", arg0)
	ret0, _ := ret[0].(*lexruntimev2.RecognizeTextOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecognizeText indicates an expected call of RecognizeText.
func (mr *MockLexRuntimeV2APIMockRecorder) RecognizeText(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecognizeText", reflect.TypeOf((*MockLexRuntimeV2API)(nil).RecognizeText), arg0)
}

// RecognizeTextRequest mocks base method.
func (m *MockLexRuntimeV2API) RecognizeTextRequest(arg0 *lexruntimev2.RecognizeTextInput) (*request.Request, *lexruntimev2.RecognizeTextOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecognizeTextRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lexruntimev2.RecognizeTextOutput)
	return ret0, ret1
}

// RecognizeTextRequest indicates an expected call of RecognizeTextRequest.
func (mr *MockLexRuntimeV2APIMockRecorder) RecognizeTextRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecognizeTextRequest", reflect.TypeOf((*MockLexRuntimeV2API)(nil).RecognizeTextRequest), arg0)
}

// RecognizeTextWithContext mocks base method.
func (m *MockLexRuntimeV2API) RecognizeTextWithContext(arg0 context.Context, arg1 *lexruntimev2.RecognizeTextInput, arg2 ...request.Option) (*lexruntimev2.RecognizeTextOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RecognizeTextWithContext", varargs...)
	ret0, _ := ret[0].(*lexruntimev2.RecognizeTextOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecognizeTextWithContext indicates an expected call of RecognizeTextWithContext.
func (mr *MockLexRuntimeV2APIMockRecorder) RecognizeTextWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecognizeTextWithContext", reflect.TypeOf((*MockLexRuntimeV2API)(nil).RecognizeTextWithContext), varargs...)
}

// RecognizeUtterance mocks base method.
func (m *MockLexRuntimeV2API) RecognizeUtterance(arg0 *lexruntimev2.RecognizeUtteranceInput) (*lexruntimev2.RecognizeUtteranceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecognizeUtterance", arg0)
	ret0, _ := ret[0].(*lexruntimev2.RecognizeUtteranceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecognizeUtterance indicates an expected call of RecognizeUtterance.
func (mr *MockLexRuntimeV2APIMockRecorder) RecognizeUtterance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecognizeUtterance", reflect.TypeOf((*MockLexRuntimeV2API)(nil).RecognizeUtterance), arg0)
}

// RecognizeUtteranceRequest mocks base method.
func (m *MockLexRuntimeV2API) RecognizeUtteranceRequest(arg0 *lexruntimev2.RecognizeUtteranceInput) (*request.Request, *lexruntimev2.RecognizeUtteranceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecognizeUtteranceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lexruntimev2.RecognizeUtteranceOutput)
	return ret0, ret1
}

// RecognizeUtteranceRequest indicates an expected call of RecognizeUtteranceRequest.
func (mr *MockLexRuntimeV2APIMockRecorder) RecognizeUtteranceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecognizeUtteranceRequest", reflect.TypeOf((*MockLexRuntimeV2API)(nil).RecognizeUtteranceRequest), arg0)
}

// RecognizeUtteranceWithContext mocks base method.
func (m *MockLexRuntimeV2API) RecognizeUtteranceWithContext(arg0 context.Context, arg1 *lexruntimev2.RecognizeUtteranceInput, arg2 ...request.Option) (*lexruntimev2.RecognizeUtteranceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RecognizeUtteranceWithContext", varargs...)
	ret0, _ := ret[0].(*lexruntimev2.RecognizeUtteranceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecognizeUtteranceWithContext indicates an expected call of RecognizeUtteranceWithContext.
func (mr *MockLexRuntimeV2APIMockRecorder) RecognizeUtteranceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecognizeUtteranceWithContext", reflect.TypeOf((*MockLexRuntimeV2API)(nil).RecognizeUtteranceWithContext), varargs...)
}

// StartConversation mocks base method.
func (m *MockLexRuntimeV2API) StartConversation(arg0 *lexruntimev2.StartConversationInput) (*lexruntimev2.StartConversationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartConversation", arg0)
	ret0, _ := ret[0].(*lexruntimev2.StartConversationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartConversation indicates an expected call of StartConversation.
func (mr *MockLexRuntimeV2APIMockRecorder) StartConversation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartConversation", reflect.TypeOf((*MockLexRuntimeV2API)(nil).StartConversation), arg0)
}

// StartConversationRequest mocks base method.
func (m *MockLexRuntimeV2API) StartConversationRequest(arg0 *lexruntimev2.StartConversationInput) (*request.Request, *lexruntimev2.StartConversationOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartConversationRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lexruntimev2.StartConversationOutput)
	return ret0, ret1
}

// StartConversationRequest indicates an expected call of StartConversationRequest.
func (mr *MockLexRuntimeV2APIMockRecorder) StartConversationRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartConversationRequest", reflect.TypeOf((*MockLexRuntimeV2API)(nil).StartConversationRequest), arg0)
}

// StartConversationWithContext mocks base method.
func (m *MockLexRuntimeV2API) StartConversationWithContext(arg0 context.Context, arg1 *lexruntimev2.StartConversationInput, arg2 ...request.Option) (*lexruntimev2.StartConversationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StartConversationWithContext", varargs...)
	ret0, _ := ret[0].(*lexruntimev2.StartConversationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartConversationWithContext indicates an expected call of StartConversationWithContext.
func (mr *MockLexRuntimeV2APIMockRecorder) StartConversationWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartConversationWithContext", reflect.TypeOf((*MockLexRuntimeV2API)(nil).StartConversationWithContext), varargs...)
}
