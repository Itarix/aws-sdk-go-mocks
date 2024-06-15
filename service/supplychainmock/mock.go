// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/supplychain/supplychainiface (interfaces: SupplyChainAPI)

// Package supplychainmock is a generated GoMock package.
package supplychainmock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	supplychain "github.com/aws/aws-sdk-go/service/supplychain"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSupplyChainAPI is a mock of SupplyChainAPI interface
type MockSupplyChainAPI struct {
	ctrl     *gomock.Controller
	recorder *MockSupplyChainAPIMockRecorder
}

// MockSupplyChainAPIMockRecorder is the mock recorder for MockSupplyChainAPI
type MockSupplyChainAPIMockRecorder struct {
	mock *MockSupplyChainAPI
}

// NewMockSupplyChainAPI creates a new mock instance
func NewMockSupplyChainAPI(ctrl *gomock.Controller) *MockSupplyChainAPI {
	mock := &MockSupplyChainAPI{ctrl: ctrl}
	mock.recorder = &MockSupplyChainAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSupplyChainAPI) EXPECT() *MockSupplyChainAPIMockRecorder {
	return m.recorder
}

// CreateBillOfMaterialsImportJob mocks base method
func (m *MockSupplyChainAPI) CreateBillOfMaterialsImportJob(arg0 *supplychain.CreateBillOfMaterialsImportJobInput) (*supplychain.CreateBillOfMaterialsImportJobOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBillOfMaterialsImportJob", arg0)
	ret0, _ := ret[0].(*supplychain.CreateBillOfMaterialsImportJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBillOfMaterialsImportJob indicates an expected call of CreateBillOfMaterialsImportJob
func (mr *MockSupplyChainAPIMockRecorder) CreateBillOfMaterialsImportJob(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBillOfMaterialsImportJob", reflect.TypeOf((*MockSupplyChainAPI)(nil).CreateBillOfMaterialsImportJob), arg0)
}

// CreateBillOfMaterialsImportJobRequest mocks base method
func (m *MockSupplyChainAPI) CreateBillOfMaterialsImportJobRequest(arg0 *supplychain.CreateBillOfMaterialsImportJobInput) (*request.Request, *supplychain.CreateBillOfMaterialsImportJobOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBillOfMaterialsImportJobRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*supplychain.CreateBillOfMaterialsImportJobOutput)
	return ret0, ret1
}

// CreateBillOfMaterialsImportJobRequest indicates an expected call of CreateBillOfMaterialsImportJobRequest
func (mr *MockSupplyChainAPIMockRecorder) CreateBillOfMaterialsImportJobRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBillOfMaterialsImportJobRequest", reflect.TypeOf((*MockSupplyChainAPI)(nil).CreateBillOfMaterialsImportJobRequest), arg0)
}

// CreateBillOfMaterialsImportJobWithContext mocks base method
func (m *MockSupplyChainAPI) CreateBillOfMaterialsImportJobWithContext(arg0 context.Context, arg1 *supplychain.CreateBillOfMaterialsImportJobInput, arg2 ...request.Option) (*supplychain.CreateBillOfMaterialsImportJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateBillOfMaterialsImportJobWithContext", varargs...)
	ret0, _ := ret[0].(*supplychain.CreateBillOfMaterialsImportJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBillOfMaterialsImportJobWithContext indicates an expected call of CreateBillOfMaterialsImportJobWithContext
func (mr *MockSupplyChainAPIMockRecorder) CreateBillOfMaterialsImportJobWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBillOfMaterialsImportJobWithContext", reflect.TypeOf((*MockSupplyChainAPI)(nil).CreateBillOfMaterialsImportJobWithContext), varargs...)
}

// GetBillOfMaterialsImportJob mocks base method
func (m *MockSupplyChainAPI) GetBillOfMaterialsImportJob(arg0 *supplychain.GetBillOfMaterialsImportJobInput) (*supplychain.GetBillOfMaterialsImportJobOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBillOfMaterialsImportJob", arg0)
	ret0, _ := ret[0].(*supplychain.GetBillOfMaterialsImportJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBillOfMaterialsImportJob indicates an expected call of GetBillOfMaterialsImportJob
func (mr *MockSupplyChainAPIMockRecorder) GetBillOfMaterialsImportJob(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBillOfMaterialsImportJob", reflect.TypeOf((*MockSupplyChainAPI)(nil).GetBillOfMaterialsImportJob), arg0)
}

// GetBillOfMaterialsImportJobRequest mocks base method
func (m *MockSupplyChainAPI) GetBillOfMaterialsImportJobRequest(arg0 *supplychain.GetBillOfMaterialsImportJobInput) (*request.Request, *supplychain.GetBillOfMaterialsImportJobOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBillOfMaterialsImportJobRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*supplychain.GetBillOfMaterialsImportJobOutput)
	return ret0, ret1
}

// GetBillOfMaterialsImportJobRequest indicates an expected call of GetBillOfMaterialsImportJobRequest
func (mr *MockSupplyChainAPIMockRecorder) GetBillOfMaterialsImportJobRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBillOfMaterialsImportJobRequest", reflect.TypeOf((*MockSupplyChainAPI)(nil).GetBillOfMaterialsImportJobRequest), arg0)
}

// GetBillOfMaterialsImportJobWithContext mocks base method
func (m *MockSupplyChainAPI) GetBillOfMaterialsImportJobWithContext(arg0 context.Context, arg1 *supplychain.GetBillOfMaterialsImportJobInput, arg2 ...request.Option) (*supplychain.GetBillOfMaterialsImportJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBillOfMaterialsImportJobWithContext", varargs...)
	ret0, _ := ret[0].(*supplychain.GetBillOfMaterialsImportJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBillOfMaterialsImportJobWithContext indicates an expected call of GetBillOfMaterialsImportJobWithContext
func (mr *MockSupplyChainAPIMockRecorder) GetBillOfMaterialsImportJobWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBillOfMaterialsImportJobWithContext", reflect.TypeOf((*MockSupplyChainAPI)(nil).GetBillOfMaterialsImportJobWithContext), varargs...)
}

// SendDataIntegrationEvent mocks base method
func (m *MockSupplyChainAPI) SendDataIntegrationEvent(arg0 *supplychain.SendDataIntegrationEventInput) (*supplychain.SendDataIntegrationEventOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendDataIntegrationEvent", arg0)
	ret0, _ := ret[0].(*supplychain.SendDataIntegrationEventOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendDataIntegrationEvent indicates an expected call of SendDataIntegrationEvent
func (mr *MockSupplyChainAPIMockRecorder) SendDataIntegrationEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendDataIntegrationEvent", reflect.TypeOf((*MockSupplyChainAPI)(nil).SendDataIntegrationEvent), arg0)
}

// SendDataIntegrationEventRequest mocks base method
func (m *MockSupplyChainAPI) SendDataIntegrationEventRequest(arg0 *supplychain.SendDataIntegrationEventInput) (*request.Request, *supplychain.SendDataIntegrationEventOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendDataIntegrationEventRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*supplychain.SendDataIntegrationEventOutput)
	return ret0, ret1
}

// SendDataIntegrationEventRequest indicates an expected call of SendDataIntegrationEventRequest
func (mr *MockSupplyChainAPIMockRecorder) SendDataIntegrationEventRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendDataIntegrationEventRequest", reflect.TypeOf((*MockSupplyChainAPI)(nil).SendDataIntegrationEventRequest), arg0)
}

// SendDataIntegrationEventWithContext mocks base method
func (m *MockSupplyChainAPI) SendDataIntegrationEventWithContext(arg0 context.Context, arg1 *supplychain.SendDataIntegrationEventInput, arg2 ...request.Option) (*supplychain.SendDataIntegrationEventOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendDataIntegrationEventWithContext", varargs...)
	ret0, _ := ret[0].(*supplychain.SendDataIntegrationEventOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendDataIntegrationEventWithContext indicates an expected call of SendDataIntegrationEventWithContext
func (mr *MockSupplyChainAPIMockRecorder) SendDataIntegrationEventWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendDataIntegrationEventWithContext", reflect.TypeOf((*MockSupplyChainAPI)(nil).SendDataIntegrationEventWithContext), varargs...)
}
