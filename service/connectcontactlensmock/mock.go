// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/connectcontactlens/connectcontactlensiface (interfaces: ConnectContactLensAPI)

// Package connectcontactlensmock is a generated GoMock package.
package connectcontactlensmock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	connectcontactlens "github.com/aws/aws-sdk-go/service/connectcontactlens"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockConnectContactLensAPI is a mock of ConnectContactLensAPI interface
type MockConnectContactLensAPI struct {
	ctrl     *gomock.Controller
	recorder *MockConnectContactLensAPIMockRecorder
}

// MockConnectContactLensAPIMockRecorder is the mock recorder for MockConnectContactLensAPI
type MockConnectContactLensAPIMockRecorder struct {
	mock *MockConnectContactLensAPI
}

// NewMockConnectContactLensAPI creates a new mock instance
func NewMockConnectContactLensAPI(ctrl *gomock.Controller) *MockConnectContactLensAPI {
	mock := &MockConnectContactLensAPI{ctrl: ctrl}
	mock.recorder = &MockConnectContactLensAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConnectContactLensAPI) EXPECT() *MockConnectContactLensAPIMockRecorder {
	return m.recorder
}

// ListRealtimeContactAnalysisSegments mocks base method
func (m *MockConnectContactLensAPI) ListRealtimeContactAnalysisSegments(arg0 *connectcontactlens.ListRealtimeContactAnalysisSegmentsInput) (*connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRealtimeContactAnalysisSegments", arg0)
	ret0, _ := ret[0].(*connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRealtimeContactAnalysisSegments indicates an expected call of ListRealtimeContactAnalysisSegments
func (mr *MockConnectContactLensAPIMockRecorder) ListRealtimeContactAnalysisSegments(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRealtimeContactAnalysisSegments", reflect.TypeOf((*MockConnectContactLensAPI)(nil).ListRealtimeContactAnalysisSegments), arg0)
}

// ListRealtimeContactAnalysisSegmentsPages mocks base method
func (m *MockConnectContactLensAPI) ListRealtimeContactAnalysisSegmentsPages(arg0 *connectcontactlens.ListRealtimeContactAnalysisSegmentsInput, arg1 func(*connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRealtimeContactAnalysisSegmentsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListRealtimeContactAnalysisSegmentsPages indicates an expected call of ListRealtimeContactAnalysisSegmentsPages
func (mr *MockConnectContactLensAPIMockRecorder) ListRealtimeContactAnalysisSegmentsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRealtimeContactAnalysisSegmentsPages", reflect.TypeOf((*MockConnectContactLensAPI)(nil).ListRealtimeContactAnalysisSegmentsPages), arg0, arg1)
}

// ListRealtimeContactAnalysisSegmentsPagesWithContext mocks base method
func (m *MockConnectContactLensAPI) ListRealtimeContactAnalysisSegmentsPagesWithContext(arg0 context.Context, arg1 *connectcontactlens.ListRealtimeContactAnalysisSegmentsInput, arg2 func(*connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRealtimeContactAnalysisSegmentsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListRealtimeContactAnalysisSegmentsPagesWithContext indicates an expected call of ListRealtimeContactAnalysisSegmentsPagesWithContext
func (mr *MockConnectContactLensAPIMockRecorder) ListRealtimeContactAnalysisSegmentsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRealtimeContactAnalysisSegmentsPagesWithContext", reflect.TypeOf((*MockConnectContactLensAPI)(nil).ListRealtimeContactAnalysisSegmentsPagesWithContext), varargs...)
}

// ListRealtimeContactAnalysisSegmentsRequest mocks base method
func (m *MockConnectContactLensAPI) ListRealtimeContactAnalysisSegmentsRequest(arg0 *connectcontactlens.ListRealtimeContactAnalysisSegmentsInput) (*request.Request, *connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRealtimeContactAnalysisSegmentsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput)
	return ret0, ret1
}

// ListRealtimeContactAnalysisSegmentsRequest indicates an expected call of ListRealtimeContactAnalysisSegmentsRequest
func (mr *MockConnectContactLensAPIMockRecorder) ListRealtimeContactAnalysisSegmentsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRealtimeContactAnalysisSegmentsRequest", reflect.TypeOf((*MockConnectContactLensAPI)(nil).ListRealtimeContactAnalysisSegmentsRequest), arg0)
}

// ListRealtimeContactAnalysisSegmentsWithContext mocks base method
func (m *MockConnectContactLensAPI) ListRealtimeContactAnalysisSegmentsWithContext(arg0 context.Context, arg1 *connectcontactlens.ListRealtimeContactAnalysisSegmentsInput, arg2 ...request.Option) (*connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRealtimeContactAnalysisSegmentsWithContext", varargs...)
	ret0, _ := ret[0].(*connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRealtimeContactAnalysisSegmentsWithContext indicates an expected call of ListRealtimeContactAnalysisSegmentsWithContext
func (mr *MockConnectContactLensAPIMockRecorder) ListRealtimeContactAnalysisSegmentsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRealtimeContactAnalysisSegmentsWithContext", reflect.TypeOf((*MockConnectContactLensAPI)(nil).ListRealtimeContactAnalysisSegmentsWithContext), varargs...)
}
