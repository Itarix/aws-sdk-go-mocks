// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/licensemanagerlinuxsubscriptions/licensemanagerlinuxsubscriptionsiface (interfaces: LicenseManagerLinuxSubscriptionsAPI)

// Package licensemanagerlinuxsubscriptionsmock is a generated GoMock package.
package licensemanagerlinuxsubscriptionsmock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	licensemanagerlinuxsubscriptions "github.com/aws/aws-sdk-go/service/licensemanagerlinuxsubscriptions"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockLicenseManagerLinuxSubscriptionsAPI is a mock of LicenseManagerLinuxSubscriptionsAPI interface
type MockLicenseManagerLinuxSubscriptionsAPI struct {
	ctrl     *gomock.Controller
	recorder *MockLicenseManagerLinuxSubscriptionsAPIMockRecorder
}

// MockLicenseManagerLinuxSubscriptionsAPIMockRecorder is the mock recorder for MockLicenseManagerLinuxSubscriptionsAPI
type MockLicenseManagerLinuxSubscriptionsAPIMockRecorder struct {
	mock *MockLicenseManagerLinuxSubscriptionsAPI
}

// NewMockLicenseManagerLinuxSubscriptionsAPI creates a new mock instance
func NewMockLicenseManagerLinuxSubscriptionsAPI(ctrl *gomock.Controller) *MockLicenseManagerLinuxSubscriptionsAPI {
	mock := &MockLicenseManagerLinuxSubscriptionsAPI{ctrl: ctrl}
	mock.recorder = &MockLicenseManagerLinuxSubscriptionsAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLicenseManagerLinuxSubscriptionsAPI) EXPECT() *MockLicenseManagerLinuxSubscriptionsAPIMockRecorder {
	return m.recorder
}

// GetServiceSettings mocks base method
func (m *MockLicenseManagerLinuxSubscriptionsAPI) GetServiceSettings(arg0 *licensemanagerlinuxsubscriptions.GetServiceSettingsInput) (*licensemanagerlinuxsubscriptions.GetServiceSettingsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceSettings", arg0)
	ret0, _ := ret[0].(*licensemanagerlinuxsubscriptions.GetServiceSettingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceSettings indicates an expected call of GetServiceSettings
func (mr *MockLicenseManagerLinuxSubscriptionsAPIMockRecorder) GetServiceSettings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceSettings", reflect.TypeOf((*MockLicenseManagerLinuxSubscriptionsAPI)(nil).GetServiceSettings), arg0)
}

// GetServiceSettingsRequest mocks base method
func (m *MockLicenseManagerLinuxSubscriptionsAPI) GetServiceSettingsRequest(arg0 *licensemanagerlinuxsubscriptions.GetServiceSettingsInput) (*request.Request, *licensemanagerlinuxsubscriptions.GetServiceSettingsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceSettingsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*licensemanagerlinuxsubscriptions.GetServiceSettingsOutput)
	return ret0, ret1
}

// GetServiceSettingsRequest indicates an expected call of GetServiceSettingsRequest
func (mr *MockLicenseManagerLinuxSubscriptionsAPIMockRecorder) GetServiceSettingsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceSettingsRequest", reflect.TypeOf((*MockLicenseManagerLinuxSubscriptionsAPI)(nil).GetServiceSettingsRequest), arg0)
}

// GetServiceSettingsWithContext mocks base method
func (m *MockLicenseManagerLinuxSubscriptionsAPI) GetServiceSettingsWithContext(arg0 context.Context, arg1 *licensemanagerlinuxsubscriptions.GetServiceSettingsInput, arg2 ...request.Option) (*licensemanagerlinuxsubscriptions.GetServiceSettingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetServiceSettingsWithContext", varargs...)
	ret0, _ := ret[0].(*licensemanagerlinuxsubscriptions.GetServiceSettingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceSettingsWithContext indicates an expected call of GetServiceSettingsWithContext
func (mr *MockLicenseManagerLinuxSubscriptionsAPIMockRecorder) GetServiceSettingsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceSettingsWithContext", reflect.TypeOf((*MockLicenseManagerLinuxSubscriptionsAPI)(nil).GetServiceSettingsWithContext), varargs...)
}

// ListLinuxSubscriptionInstances mocks base method
func (m *MockLicenseManagerLinuxSubscriptionsAPI) ListLinuxSubscriptionInstances(arg0 *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesInput) (*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLinuxSubscriptionInstances", arg0)
	ret0, _ := ret[0].(*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLinuxSubscriptionInstances indicates an expected call of ListLinuxSubscriptionInstances
func (mr *MockLicenseManagerLinuxSubscriptionsAPIMockRecorder) ListLinuxSubscriptionInstances(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLinuxSubscriptionInstances", reflect.TypeOf((*MockLicenseManagerLinuxSubscriptionsAPI)(nil).ListLinuxSubscriptionInstances), arg0)
}

// ListLinuxSubscriptionInstancesPages mocks base method
func (m *MockLicenseManagerLinuxSubscriptionsAPI) ListLinuxSubscriptionInstancesPages(arg0 *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesInput, arg1 func(*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLinuxSubscriptionInstancesPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListLinuxSubscriptionInstancesPages indicates an expected call of ListLinuxSubscriptionInstancesPages
func (mr *MockLicenseManagerLinuxSubscriptionsAPIMockRecorder) ListLinuxSubscriptionInstancesPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLinuxSubscriptionInstancesPages", reflect.TypeOf((*MockLicenseManagerLinuxSubscriptionsAPI)(nil).ListLinuxSubscriptionInstancesPages), arg0, arg1)
}

// ListLinuxSubscriptionInstancesPagesWithContext mocks base method
func (m *MockLicenseManagerLinuxSubscriptionsAPI) ListLinuxSubscriptionInstancesPagesWithContext(arg0 context.Context, arg1 *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesInput, arg2 func(*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListLinuxSubscriptionInstancesPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListLinuxSubscriptionInstancesPagesWithContext indicates an expected call of ListLinuxSubscriptionInstancesPagesWithContext
func (mr *MockLicenseManagerLinuxSubscriptionsAPIMockRecorder) ListLinuxSubscriptionInstancesPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLinuxSubscriptionInstancesPagesWithContext", reflect.TypeOf((*MockLicenseManagerLinuxSubscriptionsAPI)(nil).ListLinuxSubscriptionInstancesPagesWithContext), varargs...)
}

// ListLinuxSubscriptionInstancesRequest mocks base method
func (m *MockLicenseManagerLinuxSubscriptionsAPI) ListLinuxSubscriptionInstancesRequest(arg0 *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesInput) (*request.Request, *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLinuxSubscriptionInstancesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesOutput)
	return ret0, ret1
}

// ListLinuxSubscriptionInstancesRequest indicates an expected call of ListLinuxSubscriptionInstancesRequest
func (mr *MockLicenseManagerLinuxSubscriptionsAPIMockRecorder) ListLinuxSubscriptionInstancesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLinuxSubscriptionInstancesRequest", reflect.TypeOf((*MockLicenseManagerLinuxSubscriptionsAPI)(nil).ListLinuxSubscriptionInstancesRequest), arg0)
}

// ListLinuxSubscriptionInstancesWithContext mocks base method
func (m *MockLicenseManagerLinuxSubscriptionsAPI) ListLinuxSubscriptionInstancesWithContext(arg0 context.Context, arg1 *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesInput, arg2 ...request.Option) (*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListLinuxSubscriptionInstancesWithContext", varargs...)
	ret0, _ := ret[0].(*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLinuxSubscriptionInstancesWithContext indicates an expected call of ListLinuxSubscriptionInstancesWithContext
func (mr *MockLicenseManagerLinuxSubscriptionsAPIMockRecorder) ListLinuxSubscriptionInstancesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLinuxSubscriptionInstancesWithContext", reflect.TypeOf((*MockLicenseManagerLinuxSubscriptionsAPI)(nil).ListLinuxSubscriptionInstancesWithContext), varargs...)
}

// ListLinuxSubscriptions mocks base method
func (m *MockLicenseManagerLinuxSubscriptionsAPI) ListLinuxSubscriptions(arg0 *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsInput) (*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLinuxSubscriptions", arg0)
	ret0, _ := ret[0].(*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLinuxSubscriptions indicates an expected call of ListLinuxSubscriptions
func (mr *MockLicenseManagerLinuxSubscriptionsAPIMockRecorder) ListLinuxSubscriptions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLinuxSubscriptions", reflect.TypeOf((*MockLicenseManagerLinuxSubscriptionsAPI)(nil).ListLinuxSubscriptions), arg0)
}

// ListLinuxSubscriptionsPages mocks base method
func (m *MockLicenseManagerLinuxSubscriptionsAPI) ListLinuxSubscriptionsPages(arg0 *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsInput, arg1 func(*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLinuxSubscriptionsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListLinuxSubscriptionsPages indicates an expected call of ListLinuxSubscriptionsPages
func (mr *MockLicenseManagerLinuxSubscriptionsAPIMockRecorder) ListLinuxSubscriptionsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLinuxSubscriptionsPages", reflect.TypeOf((*MockLicenseManagerLinuxSubscriptionsAPI)(nil).ListLinuxSubscriptionsPages), arg0, arg1)
}

// ListLinuxSubscriptionsPagesWithContext mocks base method
func (m *MockLicenseManagerLinuxSubscriptionsAPI) ListLinuxSubscriptionsPagesWithContext(arg0 context.Context, arg1 *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsInput, arg2 func(*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListLinuxSubscriptionsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListLinuxSubscriptionsPagesWithContext indicates an expected call of ListLinuxSubscriptionsPagesWithContext
func (mr *MockLicenseManagerLinuxSubscriptionsAPIMockRecorder) ListLinuxSubscriptionsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLinuxSubscriptionsPagesWithContext", reflect.TypeOf((*MockLicenseManagerLinuxSubscriptionsAPI)(nil).ListLinuxSubscriptionsPagesWithContext), varargs...)
}

// ListLinuxSubscriptionsRequest mocks base method
func (m *MockLicenseManagerLinuxSubscriptionsAPI) ListLinuxSubscriptionsRequest(arg0 *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsInput) (*request.Request, *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLinuxSubscriptionsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsOutput)
	return ret0, ret1
}

// ListLinuxSubscriptionsRequest indicates an expected call of ListLinuxSubscriptionsRequest
func (mr *MockLicenseManagerLinuxSubscriptionsAPIMockRecorder) ListLinuxSubscriptionsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLinuxSubscriptionsRequest", reflect.TypeOf((*MockLicenseManagerLinuxSubscriptionsAPI)(nil).ListLinuxSubscriptionsRequest), arg0)
}

// ListLinuxSubscriptionsWithContext mocks base method
func (m *MockLicenseManagerLinuxSubscriptionsAPI) ListLinuxSubscriptionsWithContext(arg0 context.Context, arg1 *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsInput, arg2 ...request.Option) (*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListLinuxSubscriptionsWithContext", varargs...)
	ret0, _ := ret[0].(*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLinuxSubscriptionsWithContext indicates an expected call of ListLinuxSubscriptionsWithContext
func (mr *MockLicenseManagerLinuxSubscriptionsAPIMockRecorder) ListLinuxSubscriptionsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLinuxSubscriptionsWithContext", reflect.TypeOf((*MockLicenseManagerLinuxSubscriptionsAPI)(nil).ListLinuxSubscriptionsWithContext), varargs...)
}

// UpdateServiceSettings mocks base method
func (m *MockLicenseManagerLinuxSubscriptionsAPI) UpdateServiceSettings(arg0 *licensemanagerlinuxsubscriptions.UpdateServiceSettingsInput) (*licensemanagerlinuxsubscriptions.UpdateServiceSettingsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateServiceSettings", arg0)
	ret0, _ := ret[0].(*licensemanagerlinuxsubscriptions.UpdateServiceSettingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateServiceSettings indicates an expected call of UpdateServiceSettings
func (mr *MockLicenseManagerLinuxSubscriptionsAPIMockRecorder) UpdateServiceSettings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateServiceSettings", reflect.TypeOf((*MockLicenseManagerLinuxSubscriptionsAPI)(nil).UpdateServiceSettings), arg0)
}

// UpdateServiceSettingsRequest mocks base method
func (m *MockLicenseManagerLinuxSubscriptionsAPI) UpdateServiceSettingsRequest(arg0 *licensemanagerlinuxsubscriptions.UpdateServiceSettingsInput) (*request.Request, *licensemanagerlinuxsubscriptions.UpdateServiceSettingsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateServiceSettingsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*licensemanagerlinuxsubscriptions.UpdateServiceSettingsOutput)
	return ret0, ret1
}

// UpdateServiceSettingsRequest indicates an expected call of UpdateServiceSettingsRequest
func (mr *MockLicenseManagerLinuxSubscriptionsAPIMockRecorder) UpdateServiceSettingsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateServiceSettingsRequest", reflect.TypeOf((*MockLicenseManagerLinuxSubscriptionsAPI)(nil).UpdateServiceSettingsRequest), arg0)
}

// UpdateServiceSettingsWithContext mocks base method
func (m *MockLicenseManagerLinuxSubscriptionsAPI) UpdateServiceSettingsWithContext(arg0 context.Context, arg1 *licensemanagerlinuxsubscriptions.UpdateServiceSettingsInput, arg2 ...request.Option) (*licensemanagerlinuxsubscriptions.UpdateServiceSettingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateServiceSettingsWithContext", varargs...)
	ret0, _ := ret[0].(*licensemanagerlinuxsubscriptions.UpdateServiceSettingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateServiceSettingsWithContext indicates an expected call of UpdateServiceSettingsWithContext
func (mr *MockLicenseManagerLinuxSubscriptionsAPIMockRecorder) UpdateServiceSettingsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateServiceSettingsWithContext", reflect.TypeOf((*MockLicenseManagerLinuxSubscriptionsAPI)(nil).UpdateServiceSettingsWithContext), varargs...)
}
