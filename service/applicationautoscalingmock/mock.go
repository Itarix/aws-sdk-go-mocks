// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/applicationautoscaling/applicationautoscalingiface (interfaces: ApplicationAutoScalingAPI)

// Package applicationautoscalingmock is a generated GoMock package.
package applicationautoscalingmock

import (
	context "context"
	reflect "reflect"

	request "github.com/aws/aws-sdk-go/aws/request"
	applicationautoscaling "github.com/aws/aws-sdk-go/service/applicationautoscaling"
	gomock "github.com/golang/mock/gomock"
)

// MockApplicationAutoScalingAPI is a mock of ApplicationAutoScalingAPI interface.
type MockApplicationAutoScalingAPI struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationAutoScalingAPIMockRecorder
}

// MockApplicationAutoScalingAPIMockRecorder is the mock recorder for MockApplicationAutoScalingAPI.
type MockApplicationAutoScalingAPIMockRecorder struct {
	mock *MockApplicationAutoScalingAPI
}

// NewMockApplicationAutoScalingAPI creates a new mock instance.
func NewMockApplicationAutoScalingAPI(ctrl *gomock.Controller) *MockApplicationAutoScalingAPI {
	mock := &MockApplicationAutoScalingAPI{ctrl: ctrl}
	mock.recorder = &MockApplicationAutoScalingAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationAutoScalingAPI) EXPECT() *MockApplicationAutoScalingAPIMockRecorder {
	return m.recorder
}

// DeleteScalingPolicy mocks base method.
func (m *MockApplicationAutoScalingAPI) DeleteScalingPolicy(arg0 *applicationautoscaling.DeleteScalingPolicyInput) (*applicationautoscaling.DeleteScalingPolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteScalingPolicy", arg0)
	ret0, _ := ret[0].(*applicationautoscaling.DeleteScalingPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteScalingPolicy indicates an expected call of DeleteScalingPolicy.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DeleteScalingPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteScalingPolicy", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DeleteScalingPolicy), arg0)
}

// DeleteScalingPolicyRequest mocks base method.
func (m *MockApplicationAutoScalingAPI) DeleteScalingPolicyRequest(arg0 *applicationautoscaling.DeleteScalingPolicyInput) (*request.Request, *applicationautoscaling.DeleteScalingPolicyOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteScalingPolicyRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*applicationautoscaling.DeleteScalingPolicyOutput)
	return ret0, ret1
}

// DeleteScalingPolicyRequest indicates an expected call of DeleteScalingPolicyRequest.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DeleteScalingPolicyRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteScalingPolicyRequest", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DeleteScalingPolicyRequest), arg0)
}

// DeleteScalingPolicyWithContext mocks base method.
func (m *MockApplicationAutoScalingAPI) DeleteScalingPolicyWithContext(arg0 context.Context, arg1 *applicationautoscaling.DeleteScalingPolicyInput, arg2 ...request.Option) (*applicationautoscaling.DeleteScalingPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteScalingPolicyWithContext", varargs...)
	ret0, _ := ret[0].(*applicationautoscaling.DeleteScalingPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteScalingPolicyWithContext indicates an expected call of DeleteScalingPolicyWithContext.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DeleteScalingPolicyWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteScalingPolicyWithContext", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DeleteScalingPolicyWithContext), varargs...)
}

// DeleteScheduledAction mocks base method.
func (m *MockApplicationAutoScalingAPI) DeleteScheduledAction(arg0 *applicationautoscaling.DeleteScheduledActionInput) (*applicationautoscaling.DeleteScheduledActionOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteScheduledAction", arg0)
	ret0, _ := ret[0].(*applicationautoscaling.DeleteScheduledActionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteScheduledAction indicates an expected call of DeleteScheduledAction.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DeleteScheduledAction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteScheduledAction", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DeleteScheduledAction), arg0)
}

// DeleteScheduledActionRequest mocks base method.
func (m *MockApplicationAutoScalingAPI) DeleteScheduledActionRequest(arg0 *applicationautoscaling.DeleteScheduledActionInput) (*request.Request, *applicationautoscaling.DeleteScheduledActionOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteScheduledActionRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*applicationautoscaling.DeleteScheduledActionOutput)
	return ret0, ret1
}

// DeleteScheduledActionRequest indicates an expected call of DeleteScheduledActionRequest.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DeleteScheduledActionRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteScheduledActionRequest", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DeleteScheduledActionRequest), arg0)
}

// DeleteScheduledActionWithContext mocks base method.
func (m *MockApplicationAutoScalingAPI) DeleteScheduledActionWithContext(arg0 context.Context, arg1 *applicationautoscaling.DeleteScheduledActionInput, arg2 ...request.Option) (*applicationautoscaling.DeleteScheduledActionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteScheduledActionWithContext", varargs...)
	ret0, _ := ret[0].(*applicationautoscaling.DeleteScheduledActionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteScheduledActionWithContext indicates an expected call of DeleteScheduledActionWithContext.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DeleteScheduledActionWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteScheduledActionWithContext", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DeleteScheduledActionWithContext), varargs...)
}

// DeregisterScalableTarget mocks base method.
func (m *MockApplicationAutoScalingAPI) DeregisterScalableTarget(arg0 *applicationautoscaling.DeregisterScalableTargetInput) (*applicationautoscaling.DeregisterScalableTargetOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeregisterScalableTarget", arg0)
	ret0, _ := ret[0].(*applicationautoscaling.DeregisterScalableTargetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeregisterScalableTarget indicates an expected call of DeregisterScalableTarget.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DeregisterScalableTarget(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeregisterScalableTarget", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DeregisterScalableTarget), arg0)
}

// DeregisterScalableTargetRequest mocks base method.
func (m *MockApplicationAutoScalingAPI) DeregisterScalableTargetRequest(arg0 *applicationautoscaling.DeregisterScalableTargetInput) (*request.Request, *applicationautoscaling.DeregisterScalableTargetOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeregisterScalableTargetRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*applicationautoscaling.DeregisterScalableTargetOutput)
	return ret0, ret1
}

// DeregisterScalableTargetRequest indicates an expected call of DeregisterScalableTargetRequest.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DeregisterScalableTargetRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeregisterScalableTargetRequest", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DeregisterScalableTargetRequest), arg0)
}

// DeregisterScalableTargetWithContext mocks base method.
func (m *MockApplicationAutoScalingAPI) DeregisterScalableTargetWithContext(arg0 context.Context, arg1 *applicationautoscaling.DeregisterScalableTargetInput, arg2 ...request.Option) (*applicationautoscaling.DeregisterScalableTargetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeregisterScalableTargetWithContext", varargs...)
	ret0, _ := ret[0].(*applicationautoscaling.DeregisterScalableTargetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeregisterScalableTargetWithContext indicates an expected call of DeregisterScalableTargetWithContext.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DeregisterScalableTargetWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeregisterScalableTargetWithContext", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DeregisterScalableTargetWithContext), varargs...)
}

// DescribeScalableTargets mocks base method.
func (m *MockApplicationAutoScalingAPI) DescribeScalableTargets(arg0 *applicationautoscaling.DescribeScalableTargetsInput) (*applicationautoscaling.DescribeScalableTargetsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeScalableTargets", arg0)
	ret0, _ := ret[0].(*applicationautoscaling.DescribeScalableTargetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeScalableTargets indicates an expected call of DescribeScalableTargets.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DescribeScalableTargets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalableTargets", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DescribeScalableTargets), arg0)
}

// DescribeScalableTargetsPages mocks base method.
func (m *MockApplicationAutoScalingAPI) DescribeScalableTargetsPages(arg0 *applicationautoscaling.DescribeScalableTargetsInput, arg1 func(*applicationautoscaling.DescribeScalableTargetsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeScalableTargetsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DescribeScalableTargetsPages indicates an expected call of DescribeScalableTargetsPages.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DescribeScalableTargetsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalableTargetsPages", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DescribeScalableTargetsPages), arg0, arg1)
}

// DescribeScalableTargetsPagesWithContext mocks base method.
func (m *MockApplicationAutoScalingAPI) DescribeScalableTargetsPagesWithContext(arg0 context.Context, arg1 *applicationautoscaling.DescribeScalableTargetsInput, arg2 func(*applicationautoscaling.DescribeScalableTargetsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeScalableTargetsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DescribeScalableTargetsPagesWithContext indicates an expected call of DescribeScalableTargetsPagesWithContext.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DescribeScalableTargetsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalableTargetsPagesWithContext", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DescribeScalableTargetsPagesWithContext), varargs...)
}

// DescribeScalableTargetsRequest mocks base method.
func (m *MockApplicationAutoScalingAPI) DescribeScalableTargetsRequest(arg0 *applicationautoscaling.DescribeScalableTargetsInput) (*request.Request, *applicationautoscaling.DescribeScalableTargetsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeScalableTargetsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*applicationautoscaling.DescribeScalableTargetsOutput)
	return ret0, ret1
}

// DescribeScalableTargetsRequest indicates an expected call of DescribeScalableTargetsRequest.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DescribeScalableTargetsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalableTargetsRequest", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DescribeScalableTargetsRequest), arg0)
}

// DescribeScalableTargetsWithContext mocks base method.
func (m *MockApplicationAutoScalingAPI) DescribeScalableTargetsWithContext(arg0 context.Context, arg1 *applicationautoscaling.DescribeScalableTargetsInput, arg2 ...request.Option) (*applicationautoscaling.DescribeScalableTargetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeScalableTargetsWithContext", varargs...)
	ret0, _ := ret[0].(*applicationautoscaling.DescribeScalableTargetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeScalableTargetsWithContext indicates an expected call of DescribeScalableTargetsWithContext.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DescribeScalableTargetsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalableTargetsWithContext", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DescribeScalableTargetsWithContext), varargs...)
}

// DescribeScalingActivities mocks base method.
func (m *MockApplicationAutoScalingAPI) DescribeScalingActivities(arg0 *applicationautoscaling.DescribeScalingActivitiesInput) (*applicationautoscaling.DescribeScalingActivitiesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeScalingActivities", arg0)
	ret0, _ := ret[0].(*applicationautoscaling.DescribeScalingActivitiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeScalingActivities indicates an expected call of DescribeScalingActivities.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DescribeScalingActivities(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalingActivities", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DescribeScalingActivities), arg0)
}

// DescribeScalingActivitiesPages mocks base method.
func (m *MockApplicationAutoScalingAPI) DescribeScalingActivitiesPages(arg0 *applicationautoscaling.DescribeScalingActivitiesInput, arg1 func(*applicationautoscaling.DescribeScalingActivitiesOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeScalingActivitiesPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DescribeScalingActivitiesPages indicates an expected call of DescribeScalingActivitiesPages.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DescribeScalingActivitiesPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalingActivitiesPages", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DescribeScalingActivitiesPages), arg0, arg1)
}

// DescribeScalingActivitiesPagesWithContext mocks base method.
func (m *MockApplicationAutoScalingAPI) DescribeScalingActivitiesPagesWithContext(arg0 context.Context, arg1 *applicationautoscaling.DescribeScalingActivitiesInput, arg2 func(*applicationautoscaling.DescribeScalingActivitiesOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeScalingActivitiesPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DescribeScalingActivitiesPagesWithContext indicates an expected call of DescribeScalingActivitiesPagesWithContext.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DescribeScalingActivitiesPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalingActivitiesPagesWithContext", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DescribeScalingActivitiesPagesWithContext), varargs...)
}

// DescribeScalingActivitiesRequest mocks base method.
func (m *MockApplicationAutoScalingAPI) DescribeScalingActivitiesRequest(arg0 *applicationautoscaling.DescribeScalingActivitiesInput) (*request.Request, *applicationautoscaling.DescribeScalingActivitiesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeScalingActivitiesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*applicationautoscaling.DescribeScalingActivitiesOutput)
	return ret0, ret1
}

// DescribeScalingActivitiesRequest indicates an expected call of DescribeScalingActivitiesRequest.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DescribeScalingActivitiesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalingActivitiesRequest", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DescribeScalingActivitiesRequest), arg0)
}

// DescribeScalingActivitiesWithContext mocks base method.
func (m *MockApplicationAutoScalingAPI) DescribeScalingActivitiesWithContext(arg0 context.Context, arg1 *applicationautoscaling.DescribeScalingActivitiesInput, arg2 ...request.Option) (*applicationautoscaling.DescribeScalingActivitiesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeScalingActivitiesWithContext", varargs...)
	ret0, _ := ret[0].(*applicationautoscaling.DescribeScalingActivitiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeScalingActivitiesWithContext indicates an expected call of DescribeScalingActivitiesWithContext.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DescribeScalingActivitiesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalingActivitiesWithContext", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DescribeScalingActivitiesWithContext), varargs...)
}

// DescribeScalingPolicies mocks base method.
func (m *MockApplicationAutoScalingAPI) DescribeScalingPolicies(arg0 *applicationautoscaling.DescribeScalingPoliciesInput) (*applicationautoscaling.DescribeScalingPoliciesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeScalingPolicies", arg0)
	ret0, _ := ret[0].(*applicationautoscaling.DescribeScalingPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeScalingPolicies indicates an expected call of DescribeScalingPolicies.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DescribeScalingPolicies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalingPolicies", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DescribeScalingPolicies), arg0)
}

// DescribeScalingPoliciesPages mocks base method.
func (m *MockApplicationAutoScalingAPI) DescribeScalingPoliciesPages(arg0 *applicationautoscaling.DescribeScalingPoliciesInput, arg1 func(*applicationautoscaling.DescribeScalingPoliciesOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeScalingPoliciesPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DescribeScalingPoliciesPages indicates an expected call of DescribeScalingPoliciesPages.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DescribeScalingPoliciesPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalingPoliciesPages", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DescribeScalingPoliciesPages), arg0, arg1)
}

// DescribeScalingPoliciesPagesWithContext mocks base method.
func (m *MockApplicationAutoScalingAPI) DescribeScalingPoliciesPagesWithContext(arg0 context.Context, arg1 *applicationautoscaling.DescribeScalingPoliciesInput, arg2 func(*applicationautoscaling.DescribeScalingPoliciesOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeScalingPoliciesPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DescribeScalingPoliciesPagesWithContext indicates an expected call of DescribeScalingPoliciesPagesWithContext.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DescribeScalingPoliciesPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalingPoliciesPagesWithContext", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DescribeScalingPoliciesPagesWithContext), varargs...)
}

// DescribeScalingPoliciesRequest mocks base method.
func (m *MockApplicationAutoScalingAPI) DescribeScalingPoliciesRequest(arg0 *applicationautoscaling.DescribeScalingPoliciesInput) (*request.Request, *applicationautoscaling.DescribeScalingPoliciesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeScalingPoliciesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*applicationautoscaling.DescribeScalingPoliciesOutput)
	return ret0, ret1
}

// DescribeScalingPoliciesRequest indicates an expected call of DescribeScalingPoliciesRequest.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DescribeScalingPoliciesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalingPoliciesRequest", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DescribeScalingPoliciesRequest), arg0)
}

// DescribeScalingPoliciesWithContext mocks base method.
func (m *MockApplicationAutoScalingAPI) DescribeScalingPoliciesWithContext(arg0 context.Context, arg1 *applicationautoscaling.DescribeScalingPoliciesInput, arg2 ...request.Option) (*applicationautoscaling.DescribeScalingPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeScalingPoliciesWithContext", varargs...)
	ret0, _ := ret[0].(*applicationautoscaling.DescribeScalingPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeScalingPoliciesWithContext indicates an expected call of DescribeScalingPoliciesWithContext.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DescribeScalingPoliciesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalingPoliciesWithContext", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DescribeScalingPoliciesWithContext), varargs...)
}

// DescribeScheduledActions mocks base method.
func (m *MockApplicationAutoScalingAPI) DescribeScheduledActions(arg0 *applicationautoscaling.DescribeScheduledActionsInput) (*applicationautoscaling.DescribeScheduledActionsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeScheduledActions", arg0)
	ret0, _ := ret[0].(*applicationautoscaling.DescribeScheduledActionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeScheduledActions indicates an expected call of DescribeScheduledActions.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DescribeScheduledActions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScheduledActions", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DescribeScheduledActions), arg0)
}

// DescribeScheduledActionsPages mocks base method.
func (m *MockApplicationAutoScalingAPI) DescribeScheduledActionsPages(arg0 *applicationautoscaling.DescribeScheduledActionsInput, arg1 func(*applicationautoscaling.DescribeScheduledActionsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeScheduledActionsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DescribeScheduledActionsPages indicates an expected call of DescribeScheduledActionsPages.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DescribeScheduledActionsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScheduledActionsPages", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DescribeScheduledActionsPages), arg0, arg1)
}

// DescribeScheduledActionsPagesWithContext mocks base method.
func (m *MockApplicationAutoScalingAPI) DescribeScheduledActionsPagesWithContext(arg0 context.Context, arg1 *applicationautoscaling.DescribeScheduledActionsInput, arg2 func(*applicationautoscaling.DescribeScheduledActionsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeScheduledActionsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DescribeScheduledActionsPagesWithContext indicates an expected call of DescribeScheduledActionsPagesWithContext.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DescribeScheduledActionsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScheduledActionsPagesWithContext", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DescribeScheduledActionsPagesWithContext), varargs...)
}

// DescribeScheduledActionsRequest mocks base method.
func (m *MockApplicationAutoScalingAPI) DescribeScheduledActionsRequest(arg0 *applicationautoscaling.DescribeScheduledActionsInput) (*request.Request, *applicationautoscaling.DescribeScheduledActionsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeScheduledActionsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*applicationautoscaling.DescribeScheduledActionsOutput)
	return ret0, ret1
}

// DescribeScheduledActionsRequest indicates an expected call of DescribeScheduledActionsRequest.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DescribeScheduledActionsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScheduledActionsRequest", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DescribeScheduledActionsRequest), arg0)
}

// DescribeScheduledActionsWithContext mocks base method.
func (m *MockApplicationAutoScalingAPI) DescribeScheduledActionsWithContext(arg0 context.Context, arg1 *applicationautoscaling.DescribeScheduledActionsInput, arg2 ...request.Option) (*applicationautoscaling.DescribeScheduledActionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeScheduledActionsWithContext", varargs...)
	ret0, _ := ret[0].(*applicationautoscaling.DescribeScheduledActionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeScheduledActionsWithContext indicates an expected call of DescribeScheduledActionsWithContext.
func (mr *MockApplicationAutoScalingAPIMockRecorder) DescribeScheduledActionsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScheduledActionsWithContext", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).DescribeScheduledActionsWithContext), varargs...)
}

// PutScalingPolicy mocks base method.
func (m *MockApplicationAutoScalingAPI) PutScalingPolicy(arg0 *applicationautoscaling.PutScalingPolicyInput) (*applicationautoscaling.PutScalingPolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutScalingPolicy", arg0)
	ret0, _ := ret[0].(*applicationautoscaling.PutScalingPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutScalingPolicy indicates an expected call of PutScalingPolicy.
func (mr *MockApplicationAutoScalingAPIMockRecorder) PutScalingPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutScalingPolicy", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).PutScalingPolicy), arg0)
}

// PutScalingPolicyRequest mocks base method.
func (m *MockApplicationAutoScalingAPI) PutScalingPolicyRequest(arg0 *applicationautoscaling.PutScalingPolicyInput) (*request.Request, *applicationautoscaling.PutScalingPolicyOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutScalingPolicyRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*applicationautoscaling.PutScalingPolicyOutput)
	return ret0, ret1
}

// PutScalingPolicyRequest indicates an expected call of PutScalingPolicyRequest.
func (mr *MockApplicationAutoScalingAPIMockRecorder) PutScalingPolicyRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutScalingPolicyRequest", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).PutScalingPolicyRequest), arg0)
}

// PutScalingPolicyWithContext mocks base method.
func (m *MockApplicationAutoScalingAPI) PutScalingPolicyWithContext(arg0 context.Context, arg1 *applicationautoscaling.PutScalingPolicyInput, arg2 ...request.Option) (*applicationautoscaling.PutScalingPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutScalingPolicyWithContext", varargs...)
	ret0, _ := ret[0].(*applicationautoscaling.PutScalingPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutScalingPolicyWithContext indicates an expected call of PutScalingPolicyWithContext.
func (mr *MockApplicationAutoScalingAPIMockRecorder) PutScalingPolicyWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutScalingPolicyWithContext", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).PutScalingPolicyWithContext), varargs...)
}

// PutScheduledAction mocks base method.
func (m *MockApplicationAutoScalingAPI) PutScheduledAction(arg0 *applicationautoscaling.PutScheduledActionInput) (*applicationautoscaling.PutScheduledActionOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutScheduledAction", arg0)
	ret0, _ := ret[0].(*applicationautoscaling.PutScheduledActionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutScheduledAction indicates an expected call of PutScheduledAction.
func (mr *MockApplicationAutoScalingAPIMockRecorder) PutScheduledAction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutScheduledAction", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).PutScheduledAction), arg0)
}

// PutScheduledActionRequest mocks base method.
func (m *MockApplicationAutoScalingAPI) PutScheduledActionRequest(arg0 *applicationautoscaling.PutScheduledActionInput) (*request.Request, *applicationautoscaling.PutScheduledActionOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutScheduledActionRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*applicationautoscaling.PutScheduledActionOutput)
	return ret0, ret1
}

// PutScheduledActionRequest indicates an expected call of PutScheduledActionRequest.
func (mr *MockApplicationAutoScalingAPIMockRecorder) PutScheduledActionRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutScheduledActionRequest", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).PutScheduledActionRequest), arg0)
}

// PutScheduledActionWithContext mocks base method.
func (m *MockApplicationAutoScalingAPI) PutScheduledActionWithContext(arg0 context.Context, arg1 *applicationautoscaling.PutScheduledActionInput, arg2 ...request.Option) (*applicationautoscaling.PutScheduledActionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutScheduledActionWithContext", varargs...)
	ret0, _ := ret[0].(*applicationautoscaling.PutScheduledActionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutScheduledActionWithContext indicates an expected call of PutScheduledActionWithContext.
func (mr *MockApplicationAutoScalingAPIMockRecorder) PutScheduledActionWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutScheduledActionWithContext", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).PutScheduledActionWithContext), varargs...)
}

// RegisterScalableTarget mocks base method.
func (m *MockApplicationAutoScalingAPI) RegisterScalableTarget(arg0 *applicationautoscaling.RegisterScalableTargetInput) (*applicationautoscaling.RegisterScalableTargetOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterScalableTarget", arg0)
	ret0, _ := ret[0].(*applicationautoscaling.RegisterScalableTargetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterScalableTarget indicates an expected call of RegisterScalableTarget.
func (mr *MockApplicationAutoScalingAPIMockRecorder) RegisterScalableTarget(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterScalableTarget", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).RegisterScalableTarget), arg0)
}

// RegisterScalableTargetRequest mocks base method.
func (m *MockApplicationAutoScalingAPI) RegisterScalableTargetRequest(arg0 *applicationautoscaling.RegisterScalableTargetInput) (*request.Request, *applicationautoscaling.RegisterScalableTargetOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterScalableTargetRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*applicationautoscaling.RegisterScalableTargetOutput)
	return ret0, ret1
}

// RegisterScalableTargetRequest indicates an expected call of RegisterScalableTargetRequest.
func (mr *MockApplicationAutoScalingAPIMockRecorder) RegisterScalableTargetRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterScalableTargetRequest", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).RegisterScalableTargetRequest), arg0)
}

// RegisterScalableTargetWithContext mocks base method.
func (m *MockApplicationAutoScalingAPI) RegisterScalableTargetWithContext(arg0 context.Context, arg1 *applicationautoscaling.RegisterScalableTargetInput, arg2 ...request.Option) (*applicationautoscaling.RegisterScalableTargetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RegisterScalableTargetWithContext", varargs...)
	ret0, _ := ret[0].(*applicationautoscaling.RegisterScalableTargetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterScalableTargetWithContext indicates an expected call of RegisterScalableTargetWithContext.
func (mr *MockApplicationAutoScalingAPIMockRecorder) RegisterScalableTargetWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterScalableTargetWithContext", reflect.TypeOf((*MockApplicationAutoScalingAPI)(nil).RegisterScalableTargetWithContext), varargs...)
}
