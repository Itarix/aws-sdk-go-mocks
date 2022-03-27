// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/opsworkscm/opsworkscmiface (interfaces: OpsWorksCMAPI)

// Package opsworkscmmock is a generated GoMock package.
package opsworkscmmock

import (
	aws "github.com/aws/aws-sdk-go/aws"
	request "github.com/aws/aws-sdk-go/aws/request"
	opsworkscm "github.com/aws/aws-sdk-go/service/opsworkscm"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockOpsWorksCMAPI is a mock of OpsWorksCMAPI interface
type MockOpsWorksCMAPI struct {
	ctrl     *gomock.Controller
	recorder *MockOpsWorksCMAPIMockRecorder
}

// MockOpsWorksCMAPIMockRecorder is the mock recorder for MockOpsWorksCMAPI
type MockOpsWorksCMAPIMockRecorder struct {
	mock *MockOpsWorksCMAPI
}

// NewMockOpsWorksCMAPI creates a new mock instance
func NewMockOpsWorksCMAPI(ctrl *gomock.Controller) *MockOpsWorksCMAPI {
	mock := &MockOpsWorksCMAPI{ctrl: ctrl}
	mock.recorder = &MockOpsWorksCMAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOpsWorksCMAPI) EXPECT() *MockOpsWorksCMAPIMockRecorder {
	return m.recorder
}

// AssociateNode mocks base method
func (m *MockOpsWorksCMAPI) AssociateNode(arg0 *opsworkscm.AssociateNodeInput) (*opsworkscm.AssociateNodeOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssociateNode", arg0)
	ret0, _ := ret[0].(*opsworkscm.AssociateNodeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssociateNode indicates an expected call of AssociateNode
func (mr *MockOpsWorksCMAPIMockRecorder) AssociateNode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssociateNode", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).AssociateNode), arg0)
}

// AssociateNodeRequest mocks base method
func (m *MockOpsWorksCMAPI) AssociateNodeRequest(arg0 *opsworkscm.AssociateNodeInput) (*request.Request, *opsworkscm.AssociateNodeOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssociateNodeRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*opsworkscm.AssociateNodeOutput)
	return ret0, ret1
}

// AssociateNodeRequest indicates an expected call of AssociateNodeRequest
func (mr *MockOpsWorksCMAPIMockRecorder) AssociateNodeRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssociateNodeRequest", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).AssociateNodeRequest), arg0)
}

// AssociateNodeWithContext mocks base method
func (m *MockOpsWorksCMAPI) AssociateNodeWithContext(arg0 aws.Context, arg1 *opsworkscm.AssociateNodeInput, arg2 ...request.Option) (*opsworkscm.AssociateNodeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AssociateNodeWithContext", varargs...)
	ret0, _ := ret[0].(*opsworkscm.AssociateNodeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssociateNodeWithContext indicates an expected call of AssociateNodeWithContext
func (mr *MockOpsWorksCMAPIMockRecorder) AssociateNodeWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssociateNodeWithContext", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).AssociateNodeWithContext), varargs...)
}

// CreateBackup mocks base method
func (m *MockOpsWorksCMAPI) CreateBackup(arg0 *opsworkscm.CreateBackupInput) (*opsworkscm.CreateBackupOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBackup", arg0)
	ret0, _ := ret[0].(*opsworkscm.CreateBackupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBackup indicates an expected call of CreateBackup
func (mr *MockOpsWorksCMAPIMockRecorder) CreateBackup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBackup", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).CreateBackup), arg0)
}

// CreateBackupRequest mocks base method
func (m *MockOpsWorksCMAPI) CreateBackupRequest(arg0 *opsworkscm.CreateBackupInput) (*request.Request, *opsworkscm.CreateBackupOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBackupRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*opsworkscm.CreateBackupOutput)
	return ret0, ret1
}

// CreateBackupRequest indicates an expected call of CreateBackupRequest
func (mr *MockOpsWorksCMAPIMockRecorder) CreateBackupRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBackupRequest", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).CreateBackupRequest), arg0)
}

// CreateBackupWithContext mocks base method
func (m *MockOpsWorksCMAPI) CreateBackupWithContext(arg0 aws.Context, arg1 *opsworkscm.CreateBackupInput, arg2 ...request.Option) (*opsworkscm.CreateBackupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateBackupWithContext", varargs...)
	ret0, _ := ret[0].(*opsworkscm.CreateBackupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBackupWithContext indicates an expected call of CreateBackupWithContext
func (mr *MockOpsWorksCMAPIMockRecorder) CreateBackupWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBackupWithContext", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).CreateBackupWithContext), varargs...)
}

// CreateServer mocks base method
func (m *MockOpsWorksCMAPI) CreateServer(arg0 *opsworkscm.CreateServerInput) (*opsworkscm.CreateServerOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateServer", arg0)
	ret0, _ := ret[0].(*opsworkscm.CreateServerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateServer indicates an expected call of CreateServer
func (mr *MockOpsWorksCMAPIMockRecorder) CreateServer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateServer", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).CreateServer), arg0)
}

// CreateServerRequest mocks base method
func (m *MockOpsWorksCMAPI) CreateServerRequest(arg0 *opsworkscm.CreateServerInput) (*request.Request, *opsworkscm.CreateServerOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateServerRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*opsworkscm.CreateServerOutput)
	return ret0, ret1
}

// CreateServerRequest indicates an expected call of CreateServerRequest
func (mr *MockOpsWorksCMAPIMockRecorder) CreateServerRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateServerRequest", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).CreateServerRequest), arg0)
}

// CreateServerWithContext mocks base method
func (m *MockOpsWorksCMAPI) CreateServerWithContext(arg0 aws.Context, arg1 *opsworkscm.CreateServerInput, arg2 ...request.Option) (*opsworkscm.CreateServerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateServerWithContext", varargs...)
	ret0, _ := ret[0].(*opsworkscm.CreateServerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateServerWithContext indicates an expected call of CreateServerWithContext
func (mr *MockOpsWorksCMAPIMockRecorder) CreateServerWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateServerWithContext", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).CreateServerWithContext), varargs...)
}

// DeleteBackup mocks base method
func (m *MockOpsWorksCMAPI) DeleteBackup(arg0 *opsworkscm.DeleteBackupInput) (*opsworkscm.DeleteBackupOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBackup", arg0)
	ret0, _ := ret[0].(*opsworkscm.DeleteBackupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBackup indicates an expected call of DeleteBackup
func (mr *MockOpsWorksCMAPIMockRecorder) DeleteBackup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBackup", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).DeleteBackup), arg0)
}

// DeleteBackupRequest mocks base method
func (m *MockOpsWorksCMAPI) DeleteBackupRequest(arg0 *opsworkscm.DeleteBackupInput) (*request.Request, *opsworkscm.DeleteBackupOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBackupRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*opsworkscm.DeleteBackupOutput)
	return ret0, ret1
}

// DeleteBackupRequest indicates an expected call of DeleteBackupRequest
func (mr *MockOpsWorksCMAPIMockRecorder) DeleteBackupRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBackupRequest", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).DeleteBackupRequest), arg0)
}

// DeleteBackupWithContext mocks base method
func (m *MockOpsWorksCMAPI) DeleteBackupWithContext(arg0 aws.Context, arg1 *opsworkscm.DeleteBackupInput, arg2 ...request.Option) (*opsworkscm.DeleteBackupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteBackupWithContext", varargs...)
	ret0, _ := ret[0].(*opsworkscm.DeleteBackupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBackupWithContext indicates an expected call of DeleteBackupWithContext
func (mr *MockOpsWorksCMAPIMockRecorder) DeleteBackupWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBackupWithContext", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).DeleteBackupWithContext), varargs...)
}

// DeleteServer mocks base method
func (m *MockOpsWorksCMAPI) DeleteServer(arg0 *opsworkscm.DeleteServerInput) (*opsworkscm.DeleteServerOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteServer", arg0)
	ret0, _ := ret[0].(*opsworkscm.DeleteServerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteServer indicates an expected call of DeleteServer
func (mr *MockOpsWorksCMAPIMockRecorder) DeleteServer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServer", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).DeleteServer), arg0)
}

// DeleteServerRequest mocks base method
func (m *MockOpsWorksCMAPI) DeleteServerRequest(arg0 *opsworkscm.DeleteServerInput) (*request.Request, *opsworkscm.DeleteServerOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteServerRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*opsworkscm.DeleteServerOutput)
	return ret0, ret1
}

// DeleteServerRequest indicates an expected call of DeleteServerRequest
func (mr *MockOpsWorksCMAPIMockRecorder) DeleteServerRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServerRequest", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).DeleteServerRequest), arg0)
}

// DeleteServerWithContext mocks base method
func (m *MockOpsWorksCMAPI) DeleteServerWithContext(arg0 aws.Context, arg1 *opsworkscm.DeleteServerInput, arg2 ...request.Option) (*opsworkscm.DeleteServerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteServerWithContext", varargs...)
	ret0, _ := ret[0].(*opsworkscm.DeleteServerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteServerWithContext indicates an expected call of DeleteServerWithContext
func (mr *MockOpsWorksCMAPIMockRecorder) DeleteServerWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServerWithContext", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).DeleteServerWithContext), varargs...)
}

// DescribeAccountAttributes mocks base method
func (m *MockOpsWorksCMAPI) DescribeAccountAttributes(arg0 *opsworkscm.DescribeAccountAttributesInput) (*opsworkscm.DescribeAccountAttributesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeAccountAttributes", arg0)
	ret0, _ := ret[0].(*opsworkscm.DescribeAccountAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAccountAttributes indicates an expected call of DescribeAccountAttributes
func (mr *MockOpsWorksCMAPIMockRecorder) DescribeAccountAttributes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAccountAttributes", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).DescribeAccountAttributes), arg0)
}

// DescribeAccountAttributesRequest mocks base method
func (m *MockOpsWorksCMAPI) DescribeAccountAttributesRequest(arg0 *opsworkscm.DescribeAccountAttributesInput) (*request.Request, *opsworkscm.DescribeAccountAttributesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeAccountAttributesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*opsworkscm.DescribeAccountAttributesOutput)
	return ret0, ret1
}

// DescribeAccountAttributesRequest indicates an expected call of DescribeAccountAttributesRequest
func (mr *MockOpsWorksCMAPIMockRecorder) DescribeAccountAttributesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAccountAttributesRequest", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).DescribeAccountAttributesRequest), arg0)
}

// DescribeAccountAttributesWithContext mocks base method
func (m *MockOpsWorksCMAPI) DescribeAccountAttributesWithContext(arg0 aws.Context, arg1 *opsworkscm.DescribeAccountAttributesInput, arg2 ...request.Option) (*opsworkscm.DescribeAccountAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAccountAttributesWithContext", varargs...)
	ret0, _ := ret[0].(*opsworkscm.DescribeAccountAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAccountAttributesWithContext indicates an expected call of DescribeAccountAttributesWithContext
func (mr *MockOpsWorksCMAPIMockRecorder) DescribeAccountAttributesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAccountAttributesWithContext", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).DescribeAccountAttributesWithContext), varargs...)
}

// DescribeBackups mocks base method
func (m *MockOpsWorksCMAPI) DescribeBackups(arg0 *opsworkscm.DescribeBackupsInput) (*opsworkscm.DescribeBackupsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeBackups", arg0)
	ret0, _ := ret[0].(*opsworkscm.DescribeBackupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeBackups indicates an expected call of DescribeBackups
func (mr *MockOpsWorksCMAPIMockRecorder) DescribeBackups(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeBackups", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).DescribeBackups), arg0)
}

// DescribeBackupsRequest mocks base method
func (m *MockOpsWorksCMAPI) DescribeBackupsRequest(arg0 *opsworkscm.DescribeBackupsInput) (*request.Request, *opsworkscm.DescribeBackupsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeBackupsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*opsworkscm.DescribeBackupsOutput)
	return ret0, ret1
}

// DescribeBackupsRequest indicates an expected call of DescribeBackupsRequest
func (mr *MockOpsWorksCMAPIMockRecorder) DescribeBackupsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeBackupsRequest", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).DescribeBackupsRequest), arg0)
}

// DescribeBackupsWithContext mocks base method
func (m *MockOpsWorksCMAPI) DescribeBackupsWithContext(arg0 aws.Context, arg1 *opsworkscm.DescribeBackupsInput, arg2 ...request.Option) (*opsworkscm.DescribeBackupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeBackupsWithContext", varargs...)
	ret0, _ := ret[0].(*opsworkscm.DescribeBackupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeBackupsWithContext indicates an expected call of DescribeBackupsWithContext
func (mr *MockOpsWorksCMAPIMockRecorder) DescribeBackupsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeBackupsWithContext", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).DescribeBackupsWithContext), varargs...)
}

// DescribeEvents mocks base method
func (m *MockOpsWorksCMAPI) DescribeEvents(arg0 *opsworkscm.DescribeEventsInput) (*opsworkscm.DescribeEventsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeEvents", arg0)
	ret0, _ := ret[0].(*opsworkscm.DescribeEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEvents indicates an expected call of DescribeEvents
func (mr *MockOpsWorksCMAPIMockRecorder) DescribeEvents(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEvents", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).DescribeEvents), arg0)
}

// DescribeEventsRequest mocks base method
func (m *MockOpsWorksCMAPI) DescribeEventsRequest(arg0 *opsworkscm.DescribeEventsInput) (*request.Request, *opsworkscm.DescribeEventsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeEventsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*opsworkscm.DescribeEventsOutput)
	return ret0, ret1
}

// DescribeEventsRequest indicates an expected call of DescribeEventsRequest
func (mr *MockOpsWorksCMAPIMockRecorder) DescribeEventsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEventsRequest", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).DescribeEventsRequest), arg0)
}

// DescribeEventsWithContext mocks base method
func (m *MockOpsWorksCMAPI) DescribeEventsWithContext(arg0 aws.Context, arg1 *opsworkscm.DescribeEventsInput, arg2 ...request.Option) (*opsworkscm.DescribeEventsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeEventsWithContext", varargs...)
	ret0, _ := ret[0].(*opsworkscm.DescribeEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEventsWithContext indicates an expected call of DescribeEventsWithContext
func (mr *MockOpsWorksCMAPIMockRecorder) DescribeEventsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEventsWithContext", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).DescribeEventsWithContext), varargs...)
}

// DescribeNodeAssociationStatus mocks base method
func (m *MockOpsWorksCMAPI) DescribeNodeAssociationStatus(arg0 *opsworkscm.DescribeNodeAssociationStatusInput) (*opsworkscm.DescribeNodeAssociationStatusOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeNodeAssociationStatus", arg0)
	ret0, _ := ret[0].(*opsworkscm.DescribeNodeAssociationStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeNodeAssociationStatus indicates an expected call of DescribeNodeAssociationStatus
func (mr *MockOpsWorksCMAPIMockRecorder) DescribeNodeAssociationStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeNodeAssociationStatus", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).DescribeNodeAssociationStatus), arg0)
}

// DescribeNodeAssociationStatusRequest mocks base method
func (m *MockOpsWorksCMAPI) DescribeNodeAssociationStatusRequest(arg0 *opsworkscm.DescribeNodeAssociationStatusInput) (*request.Request, *opsworkscm.DescribeNodeAssociationStatusOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeNodeAssociationStatusRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*opsworkscm.DescribeNodeAssociationStatusOutput)
	return ret0, ret1
}

// DescribeNodeAssociationStatusRequest indicates an expected call of DescribeNodeAssociationStatusRequest
func (mr *MockOpsWorksCMAPIMockRecorder) DescribeNodeAssociationStatusRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeNodeAssociationStatusRequest", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).DescribeNodeAssociationStatusRequest), arg0)
}

// DescribeNodeAssociationStatusWithContext mocks base method
func (m *MockOpsWorksCMAPI) DescribeNodeAssociationStatusWithContext(arg0 aws.Context, arg1 *opsworkscm.DescribeNodeAssociationStatusInput, arg2 ...request.Option) (*opsworkscm.DescribeNodeAssociationStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeNodeAssociationStatusWithContext", varargs...)
	ret0, _ := ret[0].(*opsworkscm.DescribeNodeAssociationStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeNodeAssociationStatusWithContext indicates an expected call of DescribeNodeAssociationStatusWithContext
func (mr *MockOpsWorksCMAPIMockRecorder) DescribeNodeAssociationStatusWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeNodeAssociationStatusWithContext", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).DescribeNodeAssociationStatusWithContext), varargs...)
}

// DescribeServers mocks base method
func (m *MockOpsWorksCMAPI) DescribeServers(arg0 *opsworkscm.DescribeServersInput) (*opsworkscm.DescribeServersOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeServers", arg0)
	ret0, _ := ret[0].(*opsworkscm.DescribeServersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeServers indicates an expected call of DescribeServers
func (mr *MockOpsWorksCMAPIMockRecorder) DescribeServers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeServers", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).DescribeServers), arg0)
}

// DescribeServersRequest mocks base method
func (m *MockOpsWorksCMAPI) DescribeServersRequest(arg0 *opsworkscm.DescribeServersInput) (*request.Request, *opsworkscm.DescribeServersOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeServersRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*opsworkscm.DescribeServersOutput)
	return ret0, ret1
}

// DescribeServersRequest indicates an expected call of DescribeServersRequest
func (mr *MockOpsWorksCMAPIMockRecorder) DescribeServersRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeServersRequest", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).DescribeServersRequest), arg0)
}

// DescribeServersWithContext mocks base method
func (m *MockOpsWorksCMAPI) DescribeServersWithContext(arg0 aws.Context, arg1 *opsworkscm.DescribeServersInput, arg2 ...request.Option) (*opsworkscm.DescribeServersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeServersWithContext", varargs...)
	ret0, _ := ret[0].(*opsworkscm.DescribeServersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeServersWithContext indicates an expected call of DescribeServersWithContext
func (mr *MockOpsWorksCMAPIMockRecorder) DescribeServersWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeServersWithContext", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).DescribeServersWithContext), varargs...)
}

// DisassociateNode mocks base method
func (m *MockOpsWorksCMAPI) DisassociateNode(arg0 *opsworkscm.DisassociateNodeInput) (*opsworkscm.DisassociateNodeOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisassociateNode", arg0)
	ret0, _ := ret[0].(*opsworkscm.DisassociateNodeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisassociateNode indicates an expected call of DisassociateNode
func (mr *MockOpsWorksCMAPIMockRecorder) DisassociateNode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisassociateNode", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).DisassociateNode), arg0)
}

// DisassociateNodeRequest mocks base method
func (m *MockOpsWorksCMAPI) DisassociateNodeRequest(arg0 *opsworkscm.DisassociateNodeInput) (*request.Request, *opsworkscm.DisassociateNodeOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisassociateNodeRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*opsworkscm.DisassociateNodeOutput)
	return ret0, ret1
}

// DisassociateNodeRequest indicates an expected call of DisassociateNodeRequest
func (mr *MockOpsWorksCMAPIMockRecorder) DisassociateNodeRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisassociateNodeRequest", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).DisassociateNodeRequest), arg0)
}

// DisassociateNodeWithContext mocks base method
func (m *MockOpsWorksCMAPI) DisassociateNodeWithContext(arg0 aws.Context, arg1 *opsworkscm.DisassociateNodeInput, arg2 ...request.Option) (*opsworkscm.DisassociateNodeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DisassociateNodeWithContext", varargs...)
	ret0, _ := ret[0].(*opsworkscm.DisassociateNodeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisassociateNodeWithContext indicates an expected call of DisassociateNodeWithContext
func (mr *MockOpsWorksCMAPIMockRecorder) DisassociateNodeWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisassociateNodeWithContext", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).DisassociateNodeWithContext), varargs...)
}

// RestoreServer mocks base method
func (m *MockOpsWorksCMAPI) RestoreServer(arg0 *opsworkscm.RestoreServerInput) (*opsworkscm.RestoreServerOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreServer", arg0)
	ret0, _ := ret[0].(*opsworkscm.RestoreServerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RestoreServer indicates an expected call of RestoreServer
func (mr *MockOpsWorksCMAPIMockRecorder) RestoreServer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreServer", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).RestoreServer), arg0)
}

// RestoreServerRequest mocks base method
func (m *MockOpsWorksCMAPI) RestoreServerRequest(arg0 *opsworkscm.RestoreServerInput) (*request.Request, *opsworkscm.RestoreServerOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreServerRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*opsworkscm.RestoreServerOutput)
	return ret0, ret1
}

// RestoreServerRequest indicates an expected call of RestoreServerRequest
func (mr *MockOpsWorksCMAPIMockRecorder) RestoreServerRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreServerRequest", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).RestoreServerRequest), arg0)
}

// RestoreServerWithContext mocks base method
func (m *MockOpsWorksCMAPI) RestoreServerWithContext(arg0 aws.Context, arg1 *opsworkscm.RestoreServerInput, arg2 ...request.Option) (*opsworkscm.RestoreServerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RestoreServerWithContext", varargs...)
	ret0, _ := ret[0].(*opsworkscm.RestoreServerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RestoreServerWithContext indicates an expected call of RestoreServerWithContext
func (mr *MockOpsWorksCMAPIMockRecorder) RestoreServerWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreServerWithContext", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).RestoreServerWithContext), varargs...)
}

// StartMaintenance mocks base method
func (m *MockOpsWorksCMAPI) StartMaintenance(arg0 *opsworkscm.StartMaintenanceInput) (*opsworkscm.StartMaintenanceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartMaintenance", arg0)
	ret0, _ := ret[0].(*opsworkscm.StartMaintenanceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartMaintenance indicates an expected call of StartMaintenance
func (mr *MockOpsWorksCMAPIMockRecorder) StartMaintenance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartMaintenance", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).StartMaintenance), arg0)
}

// StartMaintenanceRequest mocks base method
func (m *MockOpsWorksCMAPI) StartMaintenanceRequest(arg0 *opsworkscm.StartMaintenanceInput) (*request.Request, *opsworkscm.StartMaintenanceOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartMaintenanceRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*opsworkscm.StartMaintenanceOutput)
	return ret0, ret1
}

// StartMaintenanceRequest indicates an expected call of StartMaintenanceRequest
func (mr *MockOpsWorksCMAPIMockRecorder) StartMaintenanceRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartMaintenanceRequest", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).StartMaintenanceRequest), arg0)
}

// StartMaintenanceWithContext mocks base method
func (m *MockOpsWorksCMAPI) StartMaintenanceWithContext(arg0 aws.Context, arg1 *opsworkscm.StartMaintenanceInput, arg2 ...request.Option) (*opsworkscm.StartMaintenanceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StartMaintenanceWithContext", varargs...)
	ret0, _ := ret[0].(*opsworkscm.StartMaintenanceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartMaintenanceWithContext indicates an expected call of StartMaintenanceWithContext
func (mr *MockOpsWorksCMAPIMockRecorder) StartMaintenanceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartMaintenanceWithContext", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).StartMaintenanceWithContext), varargs...)
}

// UpdateServer mocks base method
func (m *MockOpsWorksCMAPI) UpdateServer(arg0 *opsworkscm.UpdateServerInput) (*opsworkscm.UpdateServerOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateServer", arg0)
	ret0, _ := ret[0].(*opsworkscm.UpdateServerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateServer indicates an expected call of UpdateServer
func (mr *MockOpsWorksCMAPIMockRecorder) UpdateServer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateServer", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).UpdateServer), arg0)
}

// UpdateServerEngineAttributes mocks base method
func (m *MockOpsWorksCMAPI) UpdateServerEngineAttributes(arg0 *opsworkscm.UpdateServerEngineAttributesInput) (*opsworkscm.UpdateServerEngineAttributesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateServerEngineAttributes", arg0)
	ret0, _ := ret[0].(*opsworkscm.UpdateServerEngineAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateServerEngineAttributes indicates an expected call of UpdateServerEngineAttributes
func (mr *MockOpsWorksCMAPIMockRecorder) UpdateServerEngineAttributes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateServerEngineAttributes", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).UpdateServerEngineAttributes), arg0)
}

// UpdateServerEngineAttributesRequest mocks base method
func (m *MockOpsWorksCMAPI) UpdateServerEngineAttributesRequest(arg0 *opsworkscm.UpdateServerEngineAttributesInput) (*request.Request, *opsworkscm.UpdateServerEngineAttributesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateServerEngineAttributesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*opsworkscm.UpdateServerEngineAttributesOutput)
	return ret0, ret1
}

// UpdateServerEngineAttributesRequest indicates an expected call of UpdateServerEngineAttributesRequest
func (mr *MockOpsWorksCMAPIMockRecorder) UpdateServerEngineAttributesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateServerEngineAttributesRequest", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).UpdateServerEngineAttributesRequest), arg0)
}

// UpdateServerEngineAttributesWithContext mocks base method
func (m *MockOpsWorksCMAPI) UpdateServerEngineAttributesWithContext(arg0 aws.Context, arg1 *opsworkscm.UpdateServerEngineAttributesInput, arg2 ...request.Option) (*opsworkscm.UpdateServerEngineAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateServerEngineAttributesWithContext", varargs...)
	ret0, _ := ret[0].(*opsworkscm.UpdateServerEngineAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateServerEngineAttributesWithContext indicates an expected call of UpdateServerEngineAttributesWithContext
func (mr *MockOpsWorksCMAPIMockRecorder) UpdateServerEngineAttributesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateServerEngineAttributesWithContext", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).UpdateServerEngineAttributesWithContext), varargs...)
}

// UpdateServerRequest mocks base method
func (m *MockOpsWorksCMAPI) UpdateServerRequest(arg0 *opsworkscm.UpdateServerInput) (*request.Request, *opsworkscm.UpdateServerOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateServerRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*opsworkscm.UpdateServerOutput)
	return ret0, ret1
}

// UpdateServerRequest indicates an expected call of UpdateServerRequest
func (mr *MockOpsWorksCMAPIMockRecorder) UpdateServerRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateServerRequest", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).UpdateServerRequest), arg0)
}

// UpdateServerWithContext mocks base method
func (m *MockOpsWorksCMAPI) UpdateServerWithContext(arg0 aws.Context, arg1 *opsworkscm.UpdateServerInput, arg2 ...request.Option) (*opsworkscm.UpdateServerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateServerWithContext", varargs...)
	ret0, _ := ret[0].(*opsworkscm.UpdateServerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateServerWithContext indicates an expected call of UpdateServerWithContext
func (mr *MockOpsWorksCMAPIMockRecorder) UpdateServerWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateServerWithContext", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).UpdateServerWithContext), varargs...)
}

// WaitUntilNodeAssociated mocks base method
func (m *MockOpsWorksCMAPI) WaitUntilNodeAssociated(arg0 *opsworkscm.DescribeNodeAssociationStatusInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitUntilNodeAssociated", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitUntilNodeAssociated indicates an expected call of WaitUntilNodeAssociated
func (mr *MockOpsWorksCMAPIMockRecorder) WaitUntilNodeAssociated(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitUntilNodeAssociated", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).WaitUntilNodeAssociated), arg0)
}

// WaitUntilNodeAssociatedWithContext mocks base method
func (m *MockOpsWorksCMAPI) WaitUntilNodeAssociatedWithContext(arg0 aws.Context, arg1 *opsworkscm.DescribeNodeAssociationStatusInput, arg2 ...request.WaiterOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WaitUntilNodeAssociatedWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitUntilNodeAssociatedWithContext indicates an expected call of WaitUntilNodeAssociatedWithContext
func (mr *MockOpsWorksCMAPIMockRecorder) WaitUntilNodeAssociatedWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitUntilNodeAssociatedWithContext", reflect.TypeOf((*MockOpsWorksCMAPI)(nil).WaitUntilNodeAssociatedWithContext), varargs...)
}
