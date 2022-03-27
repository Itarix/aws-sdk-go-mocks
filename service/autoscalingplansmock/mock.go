// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/autoscalingplans/autoscalingplansiface (interfaces: AutoScalingPlansAPI)

// Package autoscalingplansmock is a generated GoMock package.
package autoscalingplansmock

import (
	context "context"
	reflect "reflect"

	request "github.com/aws/aws-sdk-go/aws/request"
	autoscalingplans "github.com/aws/aws-sdk-go/service/autoscalingplans"
	gomock "github.com/golang/mock/gomock"
)

// MockAutoScalingPlansAPI is a mock of AutoScalingPlansAPI interface.
type MockAutoScalingPlansAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAutoScalingPlansAPIMockRecorder
}

// MockAutoScalingPlansAPIMockRecorder is the mock recorder for MockAutoScalingPlansAPI.
type MockAutoScalingPlansAPIMockRecorder struct {
	mock *MockAutoScalingPlansAPI
}

// NewMockAutoScalingPlansAPI creates a new mock instance.
func NewMockAutoScalingPlansAPI(ctrl *gomock.Controller) *MockAutoScalingPlansAPI {
	mock := &MockAutoScalingPlansAPI{ctrl: ctrl}
	mock.recorder = &MockAutoScalingPlansAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAutoScalingPlansAPI) EXPECT() *MockAutoScalingPlansAPIMockRecorder {
	return m.recorder
}

// CreateScalingPlan mocks base method.
func (m *MockAutoScalingPlansAPI) CreateScalingPlan(arg0 *autoscalingplans.CreateScalingPlanInput) (*autoscalingplans.CreateScalingPlanOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateScalingPlan", arg0)
	ret0, _ := ret[0].(*autoscalingplans.CreateScalingPlanOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateScalingPlan indicates an expected call of CreateScalingPlan.
func (mr *MockAutoScalingPlansAPIMockRecorder) CreateScalingPlan(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateScalingPlan", reflect.TypeOf((*MockAutoScalingPlansAPI)(nil).CreateScalingPlan), arg0)
}

// CreateScalingPlanRequest mocks base method.
func (m *MockAutoScalingPlansAPI) CreateScalingPlanRequest(arg0 *autoscalingplans.CreateScalingPlanInput) (*request.Request, *autoscalingplans.CreateScalingPlanOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateScalingPlanRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*autoscalingplans.CreateScalingPlanOutput)
	return ret0, ret1
}

// CreateScalingPlanRequest indicates an expected call of CreateScalingPlanRequest.
func (mr *MockAutoScalingPlansAPIMockRecorder) CreateScalingPlanRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateScalingPlanRequest", reflect.TypeOf((*MockAutoScalingPlansAPI)(nil).CreateScalingPlanRequest), arg0)
}

// CreateScalingPlanWithContext mocks base method.
func (m *MockAutoScalingPlansAPI) CreateScalingPlanWithContext(arg0 context.Context, arg1 *autoscalingplans.CreateScalingPlanInput, arg2 ...request.Option) (*autoscalingplans.CreateScalingPlanOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateScalingPlanWithContext", varargs...)
	ret0, _ := ret[0].(*autoscalingplans.CreateScalingPlanOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateScalingPlanWithContext indicates an expected call of CreateScalingPlanWithContext.
func (mr *MockAutoScalingPlansAPIMockRecorder) CreateScalingPlanWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateScalingPlanWithContext", reflect.TypeOf((*MockAutoScalingPlansAPI)(nil).CreateScalingPlanWithContext), varargs...)
}

// DeleteScalingPlan mocks base method.
func (m *MockAutoScalingPlansAPI) DeleteScalingPlan(arg0 *autoscalingplans.DeleteScalingPlanInput) (*autoscalingplans.DeleteScalingPlanOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteScalingPlan", arg0)
	ret0, _ := ret[0].(*autoscalingplans.DeleteScalingPlanOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteScalingPlan indicates an expected call of DeleteScalingPlan.
func (mr *MockAutoScalingPlansAPIMockRecorder) DeleteScalingPlan(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteScalingPlan", reflect.TypeOf((*MockAutoScalingPlansAPI)(nil).DeleteScalingPlan), arg0)
}

// DeleteScalingPlanRequest mocks base method.
func (m *MockAutoScalingPlansAPI) DeleteScalingPlanRequest(arg0 *autoscalingplans.DeleteScalingPlanInput) (*request.Request, *autoscalingplans.DeleteScalingPlanOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteScalingPlanRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*autoscalingplans.DeleteScalingPlanOutput)
	return ret0, ret1
}

// DeleteScalingPlanRequest indicates an expected call of DeleteScalingPlanRequest.
func (mr *MockAutoScalingPlansAPIMockRecorder) DeleteScalingPlanRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteScalingPlanRequest", reflect.TypeOf((*MockAutoScalingPlansAPI)(nil).DeleteScalingPlanRequest), arg0)
}

// DeleteScalingPlanWithContext mocks base method.
func (m *MockAutoScalingPlansAPI) DeleteScalingPlanWithContext(arg0 context.Context, arg1 *autoscalingplans.DeleteScalingPlanInput, arg2 ...request.Option) (*autoscalingplans.DeleteScalingPlanOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteScalingPlanWithContext", varargs...)
	ret0, _ := ret[0].(*autoscalingplans.DeleteScalingPlanOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteScalingPlanWithContext indicates an expected call of DeleteScalingPlanWithContext.
func (mr *MockAutoScalingPlansAPIMockRecorder) DeleteScalingPlanWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteScalingPlanWithContext", reflect.TypeOf((*MockAutoScalingPlansAPI)(nil).DeleteScalingPlanWithContext), varargs...)
}

// DescribeScalingPlanResources mocks base method.
func (m *MockAutoScalingPlansAPI) DescribeScalingPlanResources(arg0 *autoscalingplans.DescribeScalingPlanResourcesInput) (*autoscalingplans.DescribeScalingPlanResourcesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeScalingPlanResources", arg0)
	ret0, _ := ret[0].(*autoscalingplans.DescribeScalingPlanResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeScalingPlanResources indicates an expected call of DescribeScalingPlanResources.
func (mr *MockAutoScalingPlansAPIMockRecorder) DescribeScalingPlanResources(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalingPlanResources", reflect.TypeOf((*MockAutoScalingPlansAPI)(nil).DescribeScalingPlanResources), arg0)
}

// DescribeScalingPlanResourcesRequest mocks base method.
func (m *MockAutoScalingPlansAPI) DescribeScalingPlanResourcesRequest(arg0 *autoscalingplans.DescribeScalingPlanResourcesInput) (*request.Request, *autoscalingplans.DescribeScalingPlanResourcesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeScalingPlanResourcesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*autoscalingplans.DescribeScalingPlanResourcesOutput)
	return ret0, ret1
}

// DescribeScalingPlanResourcesRequest indicates an expected call of DescribeScalingPlanResourcesRequest.
func (mr *MockAutoScalingPlansAPIMockRecorder) DescribeScalingPlanResourcesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalingPlanResourcesRequest", reflect.TypeOf((*MockAutoScalingPlansAPI)(nil).DescribeScalingPlanResourcesRequest), arg0)
}

// DescribeScalingPlanResourcesWithContext mocks base method.
func (m *MockAutoScalingPlansAPI) DescribeScalingPlanResourcesWithContext(arg0 context.Context, arg1 *autoscalingplans.DescribeScalingPlanResourcesInput, arg2 ...request.Option) (*autoscalingplans.DescribeScalingPlanResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeScalingPlanResourcesWithContext", varargs...)
	ret0, _ := ret[0].(*autoscalingplans.DescribeScalingPlanResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeScalingPlanResourcesWithContext indicates an expected call of DescribeScalingPlanResourcesWithContext.
func (mr *MockAutoScalingPlansAPIMockRecorder) DescribeScalingPlanResourcesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalingPlanResourcesWithContext", reflect.TypeOf((*MockAutoScalingPlansAPI)(nil).DescribeScalingPlanResourcesWithContext), varargs...)
}

// DescribeScalingPlans mocks base method.
func (m *MockAutoScalingPlansAPI) DescribeScalingPlans(arg0 *autoscalingplans.DescribeScalingPlansInput) (*autoscalingplans.DescribeScalingPlansOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeScalingPlans", arg0)
	ret0, _ := ret[0].(*autoscalingplans.DescribeScalingPlansOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeScalingPlans indicates an expected call of DescribeScalingPlans.
func (mr *MockAutoScalingPlansAPIMockRecorder) DescribeScalingPlans(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalingPlans", reflect.TypeOf((*MockAutoScalingPlansAPI)(nil).DescribeScalingPlans), arg0)
}

// DescribeScalingPlansRequest mocks base method.
func (m *MockAutoScalingPlansAPI) DescribeScalingPlansRequest(arg0 *autoscalingplans.DescribeScalingPlansInput) (*request.Request, *autoscalingplans.DescribeScalingPlansOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeScalingPlansRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*autoscalingplans.DescribeScalingPlansOutput)
	return ret0, ret1
}

// DescribeScalingPlansRequest indicates an expected call of DescribeScalingPlansRequest.
func (mr *MockAutoScalingPlansAPIMockRecorder) DescribeScalingPlansRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalingPlansRequest", reflect.TypeOf((*MockAutoScalingPlansAPI)(nil).DescribeScalingPlansRequest), arg0)
}

// DescribeScalingPlansWithContext mocks base method.
func (m *MockAutoScalingPlansAPI) DescribeScalingPlansWithContext(arg0 context.Context, arg1 *autoscalingplans.DescribeScalingPlansInput, arg2 ...request.Option) (*autoscalingplans.DescribeScalingPlansOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeScalingPlansWithContext", varargs...)
	ret0, _ := ret[0].(*autoscalingplans.DescribeScalingPlansOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeScalingPlansWithContext indicates an expected call of DescribeScalingPlansWithContext.
func (mr *MockAutoScalingPlansAPIMockRecorder) DescribeScalingPlansWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalingPlansWithContext", reflect.TypeOf((*MockAutoScalingPlansAPI)(nil).DescribeScalingPlansWithContext), varargs...)
}

// GetScalingPlanResourceForecastData mocks base method.
func (m *MockAutoScalingPlansAPI) GetScalingPlanResourceForecastData(arg0 *autoscalingplans.GetScalingPlanResourceForecastDataInput) (*autoscalingplans.GetScalingPlanResourceForecastDataOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScalingPlanResourceForecastData", arg0)
	ret0, _ := ret[0].(*autoscalingplans.GetScalingPlanResourceForecastDataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScalingPlanResourceForecastData indicates an expected call of GetScalingPlanResourceForecastData.
func (mr *MockAutoScalingPlansAPIMockRecorder) GetScalingPlanResourceForecastData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScalingPlanResourceForecastData", reflect.TypeOf((*MockAutoScalingPlansAPI)(nil).GetScalingPlanResourceForecastData), arg0)
}

// GetScalingPlanResourceForecastDataRequest mocks base method.
func (m *MockAutoScalingPlansAPI) GetScalingPlanResourceForecastDataRequest(arg0 *autoscalingplans.GetScalingPlanResourceForecastDataInput) (*request.Request, *autoscalingplans.GetScalingPlanResourceForecastDataOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScalingPlanResourceForecastDataRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*autoscalingplans.GetScalingPlanResourceForecastDataOutput)
	return ret0, ret1
}

// GetScalingPlanResourceForecastDataRequest indicates an expected call of GetScalingPlanResourceForecastDataRequest.
func (mr *MockAutoScalingPlansAPIMockRecorder) GetScalingPlanResourceForecastDataRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScalingPlanResourceForecastDataRequest", reflect.TypeOf((*MockAutoScalingPlansAPI)(nil).GetScalingPlanResourceForecastDataRequest), arg0)
}

// GetScalingPlanResourceForecastDataWithContext mocks base method.
func (m *MockAutoScalingPlansAPI) GetScalingPlanResourceForecastDataWithContext(arg0 context.Context, arg1 *autoscalingplans.GetScalingPlanResourceForecastDataInput, arg2 ...request.Option) (*autoscalingplans.GetScalingPlanResourceForecastDataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetScalingPlanResourceForecastDataWithContext", varargs...)
	ret0, _ := ret[0].(*autoscalingplans.GetScalingPlanResourceForecastDataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScalingPlanResourceForecastDataWithContext indicates an expected call of GetScalingPlanResourceForecastDataWithContext.
func (mr *MockAutoScalingPlansAPIMockRecorder) GetScalingPlanResourceForecastDataWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScalingPlanResourceForecastDataWithContext", reflect.TypeOf((*MockAutoScalingPlansAPI)(nil).GetScalingPlanResourceForecastDataWithContext), varargs...)
}

// UpdateScalingPlan mocks base method.
func (m *MockAutoScalingPlansAPI) UpdateScalingPlan(arg0 *autoscalingplans.UpdateScalingPlanInput) (*autoscalingplans.UpdateScalingPlanOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateScalingPlan", arg0)
	ret0, _ := ret[0].(*autoscalingplans.UpdateScalingPlanOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateScalingPlan indicates an expected call of UpdateScalingPlan.
func (mr *MockAutoScalingPlansAPIMockRecorder) UpdateScalingPlan(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateScalingPlan", reflect.TypeOf((*MockAutoScalingPlansAPI)(nil).UpdateScalingPlan), arg0)
}

// UpdateScalingPlanRequest mocks base method.
func (m *MockAutoScalingPlansAPI) UpdateScalingPlanRequest(arg0 *autoscalingplans.UpdateScalingPlanInput) (*request.Request, *autoscalingplans.UpdateScalingPlanOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateScalingPlanRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*autoscalingplans.UpdateScalingPlanOutput)
	return ret0, ret1
}

// UpdateScalingPlanRequest indicates an expected call of UpdateScalingPlanRequest.
func (mr *MockAutoScalingPlansAPIMockRecorder) UpdateScalingPlanRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateScalingPlanRequest", reflect.TypeOf((*MockAutoScalingPlansAPI)(nil).UpdateScalingPlanRequest), arg0)
}

// UpdateScalingPlanWithContext mocks base method.
func (m *MockAutoScalingPlansAPI) UpdateScalingPlanWithContext(arg0 context.Context, arg1 *autoscalingplans.UpdateScalingPlanInput, arg2 ...request.Option) (*autoscalingplans.UpdateScalingPlanOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateScalingPlanWithContext", varargs...)
	ret0, _ := ret[0].(*autoscalingplans.UpdateScalingPlanOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateScalingPlanWithContext indicates an expected call of UpdateScalingPlanWithContext.
func (mr *MockAutoScalingPlansAPIMockRecorder) UpdateScalingPlanWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateScalingPlanWithContext", reflect.TypeOf((*MockAutoScalingPlansAPI)(nil).UpdateScalingPlanWithContext), varargs...)
}
