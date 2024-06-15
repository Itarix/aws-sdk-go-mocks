// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/personalizeruntime/personalizeruntimeiface (interfaces: PersonalizeRuntimeAPI)

// Package personalizeruntimemock is a generated GoMock package.
package personalizeruntimemock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	personalizeruntime "github.com/aws/aws-sdk-go/service/personalizeruntime"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPersonalizeRuntimeAPI is a mock of PersonalizeRuntimeAPI interface
type MockPersonalizeRuntimeAPI struct {
	ctrl     *gomock.Controller
	recorder *MockPersonalizeRuntimeAPIMockRecorder
}

// MockPersonalizeRuntimeAPIMockRecorder is the mock recorder for MockPersonalizeRuntimeAPI
type MockPersonalizeRuntimeAPIMockRecorder struct {
	mock *MockPersonalizeRuntimeAPI
}

// NewMockPersonalizeRuntimeAPI creates a new mock instance
func NewMockPersonalizeRuntimeAPI(ctrl *gomock.Controller) *MockPersonalizeRuntimeAPI {
	mock := &MockPersonalizeRuntimeAPI{ctrl: ctrl}
	mock.recorder = &MockPersonalizeRuntimeAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPersonalizeRuntimeAPI) EXPECT() *MockPersonalizeRuntimeAPIMockRecorder {
	return m.recorder
}

// GetActionRecommendations mocks base method
func (m *MockPersonalizeRuntimeAPI) GetActionRecommendations(arg0 *personalizeruntime.GetActionRecommendationsInput) (*personalizeruntime.GetActionRecommendationsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActionRecommendations", arg0)
	ret0, _ := ret[0].(*personalizeruntime.GetActionRecommendationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActionRecommendations indicates an expected call of GetActionRecommendations
func (mr *MockPersonalizeRuntimeAPIMockRecorder) GetActionRecommendations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActionRecommendations", reflect.TypeOf((*MockPersonalizeRuntimeAPI)(nil).GetActionRecommendations), arg0)
}

// GetActionRecommendationsRequest mocks base method
func (m *MockPersonalizeRuntimeAPI) GetActionRecommendationsRequest(arg0 *personalizeruntime.GetActionRecommendationsInput) (*request.Request, *personalizeruntime.GetActionRecommendationsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActionRecommendationsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*personalizeruntime.GetActionRecommendationsOutput)
	return ret0, ret1
}

// GetActionRecommendationsRequest indicates an expected call of GetActionRecommendationsRequest
func (mr *MockPersonalizeRuntimeAPIMockRecorder) GetActionRecommendationsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActionRecommendationsRequest", reflect.TypeOf((*MockPersonalizeRuntimeAPI)(nil).GetActionRecommendationsRequest), arg0)
}

// GetActionRecommendationsWithContext mocks base method
func (m *MockPersonalizeRuntimeAPI) GetActionRecommendationsWithContext(arg0 context.Context, arg1 *personalizeruntime.GetActionRecommendationsInput, arg2 ...request.Option) (*personalizeruntime.GetActionRecommendationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetActionRecommendationsWithContext", varargs...)
	ret0, _ := ret[0].(*personalizeruntime.GetActionRecommendationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActionRecommendationsWithContext indicates an expected call of GetActionRecommendationsWithContext
func (mr *MockPersonalizeRuntimeAPIMockRecorder) GetActionRecommendationsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActionRecommendationsWithContext", reflect.TypeOf((*MockPersonalizeRuntimeAPI)(nil).GetActionRecommendationsWithContext), varargs...)
}

// GetPersonalizedRanking mocks base method
func (m *MockPersonalizeRuntimeAPI) GetPersonalizedRanking(arg0 *personalizeruntime.GetPersonalizedRankingInput) (*personalizeruntime.GetPersonalizedRankingOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPersonalizedRanking", arg0)
	ret0, _ := ret[0].(*personalizeruntime.GetPersonalizedRankingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPersonalizedRanking indicates an expected call of GetPersonalizedRanking
func (mr *MockPersonalizeRuntimeAPIMockRecorder) GetPersonalizedRanking(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPersonalizedRanking", reflect.TypeOf((*MockPersonalizeRuntimeAPI)(nil).GetPersonalizedRanking), arg0)
}

// GetPersonalizedRankingRequest mocks base method
func (m *MockPersonalizeRuntimeAPI) GetPersonalizedRankingRequest(arg0 *personalizeruntime.GetPersonalizedRankingInput) (*request.Request, *personalizeruntime.GetPersonalizedRankingOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPersonalizedRankingRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*personalizeruntime.GetPersonalizedRankingOutput)
	return ret0, ret1
}

// GetPersonalizedRankingRequest indicates an expected call of GetPersonalizedRankingRequest
func (mr *MockPersonalizeRuntimeAPIMockRecorder) GetPersonalizedRankingRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPersonalizedRankingRequest", reflect.TypeOf((*MockPersonalizeRuntimeAPI)(nil).GetPersonalizedRankingRequest), arg0)
}

// GetPersonalizedRankingWithContext mocks base method
func (m *MockPersonalizeRuntimeAPI) GetPersonalizedRankingWithContext(arg0 context.Context, arg1 *personalizeruntime.GetPersonalizedRankingInput, arg2 ...request.Option) (*personalizeruntime.GetPersonalizedRankingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPersonalizedRankingWithContext", varargs...)
	ret0, _ := ret[0].(*personalizeruntime.GetPersonalizedRankingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPersonalizedRankingWithContext indicates an expected call of GetPersonalizedRankingWithContext
func (mr *MockPersonalizeRuntimeAPIMockRecorder) GetPersonalizedRankingWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPersonalizedRankingWithContext", reflect.TypeOf((*MockPersonalizeRuntimeAPI)(nil).GetPersonalizedRankingWithContext), varargs...)
}

// GetRecommendations mocks base method
func (m *MockPersonalizeRuntimeAPI) GetRecommendations(arg0 *personalizeruntime.GetRecommendationsInput) (*personalizeruntime.GetRecommendationsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecommendations", arg0)
	ret0, _ := ret[0].(*personalizeruntime.GetRecommendationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecommendations indicates an expected call of GetRecommendations
func (mr *MockPersonalizeRuntimeAPIMockRecorder) GetRecommendations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecommendations", reflect.TypeOf((*MockPersonalizeRuntimeAPI)(nil).GetRecommendations), arg0)
}

// GetRecommendationsRequest mocks base method
func (m *MockPersonalizeRuntimeAPI) GetRecommendationsRequest(arg0 *personalizeruntime.GetRecommendationsInput) (*request.Request, *personalizeruntime.GetRecommendationsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecommendationsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*personalizeruntime.GetRecommendationsOutput)
	return ret0, ret1
}

// GetRecommendationsRequest indicates an expected call of GetRecommendationsRequest
func (mr *MockPersonalizeRuntimeAPIMockRecorder) GetRecommendationsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecommendationsRequest", reflect.TypeOf((*MockPersonalizeRuntimeAPI)(nil).GetRecommendationsRequest), arg0)
}

// GetRecommendationsWithContext mocks base method
func (m *MockPersonalizeRuntimeAPI) GetRecommendationsWithContext(arg0 context.Context, arg1 *personalizeruntime.GetRecommendationsInput, arg2 ...request.Option) (*personalizeruntime.GetRecommendationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRecommendationsWithContext", varargs...)
	ret0, _ := ret[0].(*personalizeruntime.GetRecommendationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecommendationsWithContext indicates an expected call of GetRecommendationsWithContext
func (mr *MockPersonalizeRuntimeAPIMockRecorder) GetRecommendationsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecommendationsWithContext", reflect.TypeOf((*MockPersonalizeRuntimeAPI)(nil).GetRecommendationsWithContext), varargs...)
}
