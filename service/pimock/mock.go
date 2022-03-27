// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/pi/piiface (interfaces: PIAPI)

// Package pimock is a generated GoMock package.
package pimock

import (
	context "context"
	reflect "reflect"

	request "github.com/aws/aws-sdk-go/aws/request"
	pi "github.com/aws/aws-sdk-go/service/pi"
	gomock "github.com/golang/mock/gomock"
)

// MockPIAPI is a mock of PIAPI interface.
type MockPIAPI struct {
	ctrl     *gomock.Controller
	recorder *MockPIAPIMockRecorder
}

// MockPIAPIMockRecorder is the mock recorder for MockPIAPI.
type MockPIAPIMockRecorder struct {
	mock *MockPIAPI
}

// NewMockPIAPI creates a new mock instance.
func NewMockPIAPI(ctrl *gomock.Controller) *MockPIAPI {
	mock := &MockPIAPI{ctrl: ctrl}
	mock.recorder = &MockPIAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPIAPI) EXPECT() *MockPIAPIMockRecorder {
	return m.recorder
}

// DescribeDimensionKeys mocks base method.
func (m *MockPIAPI) DescribeDimensionKeys(arg0 *pi.DescribeDimensionKeysInput) (*pi.DescribeDimensionKeysOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeDimensionKeys", arg0)
	ret0, _ := ret[0].(*pi.DescribeDimensionKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDimensionKeys indicates an expected call of DescribeDimensionKeys.
func (mr *MockPIAPIMockRecorder) DescribeDimensionKeys(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDimensionKeys", reflect.TypeOf((*MockPIAPI)(nil).DescribeDimensionKeys), arg0)
}

// DescribeDimensionKeysPages mocks base method.
func (m *MockPIAPI) DescribeDimensionKeysPages(arg0 *pi.DescribeDimensionKeysInput, arg1 func(*pi.DescribeDimensionKeysOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeDimensionKeysPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DescribeDimensionKeysPages indicates an expected call of DescribeDimensionKeysPages.
func (mr *MockPIAPIMockRecorder) DescribeDimensionKeysPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDimensionKeysPages", reflect.TypeOf((*MockPIAPI)(nil).DescribeDimensionKeysPages), arg0, arg1)
}

// DescribeDimensionKeysPagesWithContext mocks base method.
func (m *MockPIAPI) DescribeDimensionKeysPagesWithContext(arg0 context.Context, arg1 *pi.DescribeDimensionKeysInput, arg2 func(*pi.DescribeDimensionKeysOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDimensionKeysPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DescribeDimensionKeysPagesWithContext indicates an expected call of DescribeDimensionKeysPagesWithContext.
func (mr *MockPIAPIMockRecorder) DescribeDimensionKeysPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDimensionKeysPagesWithContext", reflect.TypeOf((*MockPIAPI)(nil).DescribeDimensionKeysPagesWithContext), varargs...)
}

// DescribeDimensionKeysRequest mocks base method.
func (m *MockPIAPI) DescribeDimensionKeysRequest(arg0 *pi.DescribeDimensionKeysInput) (*request.Request, *pi.DescribeDimensionKeysOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeDimensionKeysRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*pi.DescribeDimensionKeysOutput)
	return ret0, ret1
}

// DescribeDimensionKeysRequest indicates an expected call of DescribeDimensionKeysRequest.
func (mr *MockPIAPIMockRecorder) DescribeDimensionKeysRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDimensionKeysRequest", reflect.TypeOf((*MockPIAPI)(nil).DescribeDimensionKeysRequest), arg0)
}

// DescribeDimensionKeysWithContext mocks base method.
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

// DescribeDimensionKeysWithContext indicates an expected call of DescribeDimensionKeysWithContext.
func (mr *MockPIAPIMockRecorder) DescribeDimensionKeysWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDimensionKeysWithContext", reflect.TypeOf((*MockPIAPI)(nil).DescribeDimensionKeysWithContext), varargs...)
}

// GetDimensionKeyDetails mocks base method.
func (m *MockPIAPI) GetDimensionKeyDetails(arg0 *pi.GetDimensionKeyDetailsInput) (*pi.GetDimensionKeyDetailsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDimensionKeyDetails", arg0)
	ret0, _ := ret[0].(*pi.GetDimensionKeyDetailsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDimensionKeyDetails indicates an expected call of GetDimensionKeyDetails.
func (mr *MockPIAPIMockRecorder) GetDimensionKeyDetails(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDimensionKeyDetails", reflect.TypeOf((*MockPIAPI)(nil).GetDimensionKeyDetails), arg0)
}

// GetDimensionKeyDetailsRequest mocks base method.
func (m *MockPIAPI) GetDimensionKeyDetailsRequest(arg0 *pi.GetDimensionKeyDetailsInput) (*request.Request, *pi.GetDimensionKeyDetailsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDimensionKeyDetailsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*pi.GetDimensionKeyDetailsOutput)
	return ret0, ret1
}

// GetDimensionKeyDetailsRequest indicates an expected call of GetDimensionKeyDetailsRequest.
func (mr *MockPIAPIMockRecorder) GetDimensionKeyDetailsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDimensionKeyDetailsRequest", reflect.TypeOf((*MockPIAPI)(nil).GetDimensionKeyDetailsRequest), arg0)
}

// GetDimensionKeyDetailsWithContext mocks base method.
func (m *MockPIAPI) GetDimensionKeyDetailsWithContext(arg0 context.Context, arg1 *pi.GetDimensionKeyDetailsInput, arg2 ...request.Option) (*pi.GetDimensionKeyDetailsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDimensionKeyDetailsWithContext", varargs...)
	ret0, _ := ret[0].(*pi.GetDimensionKeyDetailsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDimensionKeyDetailsWithContext indicates an expected call of GetDimensionKeyDetailsWithContext.
func (mr *MockPIAPIMockRecorder) GetDimensionKeyDetailsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDimensionKeyDetailsWithContext", reflect.TypeOf((*MockPIAPI)(nil).GetDimensionKeyDetailsWithContext), varargs...)
}

// GetResourceMetadata mocks base method.
func (m *MockPIAPI) GetResourceMetadata(arg0 *pi.GetResourceMetadataInput) (*pi.GetResourceMetadataOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceMetadata", arg0)
	ret0, _ := ret[0].(*pi.GetResourceMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourceMetadata indicates an expected call of GetResourceMetadata.
func (mr *MockPIAPIMockRecorder) GetResourceMetadata(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceMetadata", reflect.TypeOf((*MockPIAPI)(nil).GetResourceMetadata), arg0)
}

// GetResourceMetadataRequest mocks base method.
func (m *MockPIAPI) GetResourceMetadataRequest(arg0 *pi.GetResourceMetadataInput) (*request.Request, *pi.GetResourceMetadataOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceMetadataRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*pi.GetResourceMetadataOutput)
	return ret0, ret1
}

// GetResourceMetadataRequest indicates an expected call of GetResourceMetadataRequest.
func (mr *MockPIAPIMockRecorder) GetResourceMetadataRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceMetadataRequest", reflect.TypeOf((*MockPIAPI)(nil).GetResourceMetadataRequest), arg0)
}

// GetResourceMetadataWithContext mocks base method.
func (m *MockPIAPI) GetResourceMetadataWithContext(arg0 context.Context, arg1 *pi.GetResourceMetadataInput, arg2 ...request.Option) (*pi.GetResourceMetadataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetResourceMetadataWithContext", varargs...)
	ret0, _ := ret[0].(*pi.GetResourceMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourceMetadataWithContext indicates an expected call of GetResourceMetadataWithContext.
func (mr *MockPIAPIMockRecorder) GetResourceMetadataWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceMetadataWithContext", reflect.TypeOf((*MockPIAPI)(nil).GetResourceMetadataWithContext), varargs...)
}

// GetResourceMetrics mocks base method.
func (m *MockPIAPI) GetResourceMetrics(arg0 *pi.GetResourceMetricsInput) (*pi.GetResourceMetricsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceMetrics", arg0)
	ret0, _ := ret[0].(*pi.GetResourceMetricsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourceMetrics indicates an expected call of GetResourceMetrics.
func (mr *MockPIAPIMockRecorder) GetResourceMetrics(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceMetrics", reflect.TypeOf((*MockPIAPI)(nil).GetResourceMetrics), arg0)
}

// GetResourceMetricsPages mocks base method.
func (m *MockPIAPI) GetResourceMetricsPages(arg0 *pi.GetResourceMetricsInput, arg1 func(*pi.GetResourceMetricsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceMetricsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetResourceMetricsPages indicates an expected call of GetResourceMetricsPages.
func (mr *MockPIAPIMockRecorder) GetResourceMetricsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceMetricsPages", reflect.TypeOf((*MockPIAPI)(nil).GetResourceMetricsPages), arg0, arg1)
}

// GetResourceMetricsPagesWithContext mocks base method.
func (m *MockPIAPI) GetResourceMetricsPagesWithContext(arg0 context.Context, arg1 *pi.GetResourceMetricsInput, arg2 func(*pi.GetResourceMetricsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetResourceMetricsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetResourceMetricsPagesWithContext indicates an expected call of GetResourceMetricsPagesWithContext.
func (mr *MockPIAPIMockRecorder) GetResourceMetricsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceMetricsPagesWithContext", reflect.TypeOf((*MockPIAPI)(nil).GetResourceMetricsPagesWithContext), varargs...)
}

// GetResourceMetricsRequest mocks base method.
func (m *MockPIAPI) GetResourceMetricsRequest(arg0 *pi.GetResourceMetricsInput) (*request.Request, *pi.GetResourceMetricsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceMetricsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*pi.GetResourceMetricsOutput)
	return ret0, ret1
}

// GetResourceMetricsRequest indicates an expected call of GetResourceMetricsRequest.
func (mr *MockPIAPIMockRecorder) GetResourceMetricsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceMetricsRequest", reflect.TypeOf((*MockPIAPI)(nil).GetResourceMetricsRequest), arg0)
}

// GetResourceMetricsWithContext mocks base method.
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

// GetResourceMetricsWithContext indicates an expected call of GetResourceMetricsWithContext.
func (mr *MockPIAPIMockRecorder) GetResourceMetricsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceMetricsWithContext", reflect.TypeOf((*MockPIAPI)(nil).GetResourceMetricsWithContext), varargs...)
}

// ListAvailableResourceDimensions mocks base method.
func (m *MockPIAPI) ListAvailableResourceDimensions(arg0 *pi.ListAvailableResourceDimensionsInput) (*pi.ListAvailableResourceDimensionsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAvailableResourceDimensions", arg0)
	ret0, _ := ret[0].(*pi.ListAvailableResourceDimensionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAvailableResourceDimensions indicates an expected call of ListAvailableResourceDimensions.
func (mr *MockPIAPIMockRecorder) ListAvailableResourceDimensions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAvailableResourceDimensions", reflect.TypeOf((*MockPIAPI)(nil).ListAvailableResourceDimensions), arg0)
}

// ListAvailableResourceDimensionsPages mocks base method.
func (m *MockPIAPI) ListAvailableResourceDimensionsPages(arg0 *pi.ListAvailableResourceDimensionsInput, arg1 func(*pi.ListAvailableResourceDimensionsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAvailableResourceDimensionsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListAvailableResourceDimensionsPages indicates an expected call of ListAvailableResourceDimensionsPages.
func (mr *MockPIAPIMockRecorder) ListAvailableResourceDimensionsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAvailableResourceDimensionsPages", reflect.TypeOf((*MockPIAPI)(nil).ListAvailableResourceDimensionsPages), arg0, arg1)
}

// ListAvailableResourceDimensionsPagesWithContext mocks base method.
func (m *MockPIAPI) ListAvailableResourceDimensionsPagesWithContext(arg0 context.Context, arg1 *pi.ListAvailableResourceDimensionsInput, arg2 func(*pi.ListAvailableResourceDimensionsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAvailableResourceDimensionsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListAvailableResourceDimensionsPagesWithContext indicates an expected call of ListAvailableResourceDimensionsPagesWithContext.
func (mr *MockPIAPIMockRecorder) ListAvailableResourceDimensionsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAvailableResourceDimensionsPagesWithContext", reflect.TypeOf((*MockPIAPI)(nil).ListAvailableResourceDimensionsPagesWithContext), varargs...)
}

// ListAvailableResourceDimensionsRequest mocks base method.
func (m *MockPIAPI) ListAvailableResourceDimensionsRequest(arg0 *pi.ListAvailableResourceDimensionsInput) (*request.Request, *pi.ListAvailableResourceDimensionsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAvailableResourceDimensionsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*pi.ListAvailableResourceDimensionsOutput)
	return ret0, ret1
}

// ListAvailableResourceDimensionsRequest indicates an expected call of ListAvailableResourceDimensionsRequest.
func (mr *MockPIAPIMockRecorder) ListAvailableResourceDimensionsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAvailableResourceDimensionsRequest", reflect.TypeOf((*MockPIAPI)(nil).ListAvailableResourceDimensionsRequest), arg0)
}

// ListAvailableResourceDimensionsWithContext mocks base method.
func (m *MockPIAPI) ListAvailableResourceDimensionsWithContext(arg0 context.Context, arg1 *pi.ListAvailableResourceDimensionsInput, arg2 ...request.Option) (*pi.ListAvailableResourceDimensionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAvailableResourceDimensionsWithContext", varargs...)
	ret0, _ := ret[0].(*pi.ListAvailableResourceDimensionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAvailableResourceDimensionsWithContext indicates an expected call of ListAvailableResourceDimensionsWithContext.
func (mr *MockPIAPIMockRecorder) ListAvailableResourceDimensionsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAvailableResourceDimensionsWithContext", reflect.TypeOf((*MockPIAPI)(nil).ListAvailableResourceDimensionsWithContext), varargs...)
}

// ListAvailableResourceMetrics mocks base method.
func (m *MockPIAPI) ListAvailableResourceMetrics(arg0 *pi.ListAvailableResourceMetricsInput) (*pi.ListAvailableResourceMetricsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAvailableResourceMetrics", arg0)
	ret0, _ := ret[0].(*pi.ListAvailableResourceMetricsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAvailableResourceMetrics indicates an expected call of ListAvailableResourceMetrics.
func (mr *MockPIAPIMockRecorder) ListAvailableResourceMetrics(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAvailableResourceMetrics", reflect.TypeOf((*MockPIAPI)(nil).ListAvailableResourceMetrics), arg0)
}

// ListAvailableResourceMetricsPages mocks base method.
func (m *MockPIAPI) ListAvailableResourceMetricsPages(arg0 *pi.ListAvailableResourceMetricsInput, arg1 func(*pi.ListAvailableResourceMetricsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAvailableResourceMetricsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListAvailableResourceMetricsPages indicates an expected call of ListAvailableResourceMetricsPages.
func (mr *MockPIAPIMockRecorder) ListAvailableResourceMetricsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAvailableResourceMetricsPages", reflect.TypeOf((*MockPIAPI)(nil).ListAvailableResourceMetricsPages), arg0, arg1)
}

// ListAvailableResourceMetricsPagesWithContext mocks base method.
func (m *MockPIAPI) ListAvailableResourceMetricsPagesWithContext(arg0 context.Context, arg1 *pi.ListAvailableResourceMetricsInput, arg2 func(*pi.ListAvailableResourceMetricsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAvailableResourceMetricsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListAvailableResourceMetricsPagesWithContext indicates an expected call of ListAvailableResourceMetricsPagesWithContext.
func (mr *MockPIAPIMockRecorder) ListAvailableResourceMetricsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAvailableResourceMetricsPagesWithContext", reflect.TypeOf((*MockPIAPI)(nil).ListAvailableResourceMetricsPagesWithContext), varargs...)
}

// ListAvailableResourceMetricsRequest mocks base method.
func (m *MockPIAPI) ListAvailableResourceMetricsRequest(arg0 *pi.ListAvailableResourceMetricsInput) (*request.Request, *pi.ListAvailableResourceMetricsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAvailableResourceMetricsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*pi.ListAvailableResourceMetricsOutput)
	return ret0, ret1
}

// ListAvailableResourceMetricsRequest indicates an expected call of ListAvailableResourceMetricsRequest.
func (mr *MockPIAPIMockRecorder) ListAvailableResourceMetricsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAvailableResourceMetricsRequest", reflect.TypeOf((*MockPIAPI)(nil).ListAvailableResourceMetricsRequest), arg0)
}

// ListAvailableResourceMetricsWithContext mocks base method.
func (m *MockPIAPI) ListAvailableResourceMetricsWithContext(arg0 context.Context, arg1 *pi.ListAvailableResourceMetricsInput, arg2 ...request.Option) (*pi.ListAvailableResourceMetricsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAvailableResourceMetricsWithContext", varargs...)
	ret0, _ := ret[0].(*pi.ListAvailableResourceMetricsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAvailableResourceMetricsWithContext indicates an expected call of ListAvailableResourceMetricsWithContext.
func (mr *MockPIAPIMockRecorder) ListAvailableResourceMetricsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAvailableResourceMetricsWithContext", reflect.TypeOf((*MockPIAPI)(nil).ListAvailableResourceMetricsWithContext), varargs...)
}
