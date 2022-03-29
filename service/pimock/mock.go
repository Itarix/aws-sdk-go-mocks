// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/pi/piiface (interfaces: PIAPI)

// Package pimock is a generated GoMock package.
package pimock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	pi "github.com/aws/aws-sdk-go/service/pi"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPIAPI is a mock of PIAPI interface
type MockPIAPI struct {
	ctrl     *gomock.Controller
	recorder *MockPIAPIMockRecorder
}

// MockPIAPIMockRecorder is the mock recorder for MockPIAPI
type MockPIAPIMockRecorder struct {
	mock *MockPIAPI
}

// NewMockPIAPI creates a new mock instance
func NewMockPIAPI(ctrl *gomock.Controller) *MockPIAPI {
	mock := &MockPIAPI{ctrl: ctrl}
	mock.recorder = &MockPIAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPIAPI) EXPECT() *MockPIAPIMockRecorder {
	return m.recorder
}

// DescribeDimensionKeys mocks base method
func (m *MockPIAPI) DescribeDimensionKeys(arg0 *pi.DescribeDimensionKeysInput) (*pi.DescribeDimensionKeysOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeDimensionKeys", arg0)
	ret0, _ := ret[0].(*pi.DescribeDimensionKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDimensionKeys indicates an expected call of DescribeDimensionKeys
func (mr *MockPIAPIMockRecorder) DescribeDimensionKeys(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDimensionKeys", reflect.TypeOf((*MockPIAPI)(nil).DescribeDimensionKeys), arg0)
}

// DescribeDimensionKeysRequest mocks base method
func (m *MockPIAPI) DescribeDimensionKeysRequest(arg0 *pi.DescribeDimensionKeysInput) (*request.Request, *pi.DescribeDimensionKeysOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeDimensionKeysRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*pi.DescribeDimensionKeysOutput)
	return ret0, ret1
}

// DescribeDimensionKeysRequest indicates an expected call of DescribeDimensionKeysRequest
func (mr *MockPIAPIMockRecorder) DescribeDimensionKeysRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDimensionKeysRequest", reflect.TypeOf((*MockPIAPI)(nil).DescribeDimensionKeysRequest), arg0)
}

// DescribeDimensionKeysWithContext mocks base method
func (m *MockPIAPI) DescribeDimensionKeysWithContext(arg0 context.Context, arg1 *pi.DescribeDimensionKeysInput, arg2 ...request.Option) (*pi.DescribeDimensionKeysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDimensionKeysWithContext", varargs...)
	ret0, _ := ret[0].(*pi.DescribeDimensionKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDimensionKeysWithContext indicates an expected call of DescribeDimensionKeysWithContext
func (mr *MockPIAPIMockRecorder) DescribeDimensionKeysWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDimensionKeysWithContext", reflect.TypeOf((*MockPIAPI)(nil).DescribeDimensionKeysWithContext), varargs...)
}

// GetResourceMetrics mocks base method
func (m *MockPIAPI) GetResourceMetrics(arg0 *pi.GetResourceMetricsInput) (*pi.GetResourceMetricsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceMetrics", arg0)
	ret0, _ := ret[0].(*pi.GetResourceMetricsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourceMetrics indicates an expected call of GetResourceMetrics
func (mr *MockPIAPIMockRecorder) GetResourceMetrics(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceMetrics", reflect.TypeOf((*MockPIAPI)(nil).GetResourceMetrics), arg0)
}

// GetResourceMetricsRequest mocks base method
func (m *MockPIAPI) GetResourceMetricsRequest(arg0 *pi.GetResourceMetricsInput) (*request.Request, *pi.GetResourceMetricsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceMetricsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*pi.GetResourceMetricsOutput)
	return ret0, ret1
}

// GetResourceMetricsRequest indicates an expected call of GetResourceMetricsRequest
func (mr *MockPIAPIMockRecorder) GetResourceMetricsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceMetricsRequest", reflect.TypeOf((*MockPIAPI)(nil).GetResourceMetricsRequest), arg0)
}

// GetResourceMetricsWithContext mocks base method
func (m *MockPIAPI) GetResourceMetricsWithContext(arg0 context.Context, arg1 *pi.GetResourceMetricsInput, arg2 ...request.Option) (*pi.GetResourceMetricsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetResourceMetricsWithContext", varargs...)
	ret0, _ := ret[0].(*pi.GetResourceMetricsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourceMetricsWithContext indicates an expected call of GetResourceMetricsWithContext
func (mr *MockPIAPIMockRecorder) GetResourceMetricsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceMetricsWithContext", reflect.TypeOf((*MockPIAPI)(nil).GetResourceMetricsWithContext), varargs...)
}