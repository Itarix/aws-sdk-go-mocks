// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/efs/efsiface (interfaces: EFSAPI)

// Package efsmock is a generated GoMock package.
package efsmock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	efs "github.com/aws/aws-sdk-go/service/efs"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockEFSAPI is a mock of EFSAPI interface
type MockEFSAPI struct {
	ctrl     *gomock.Controller
	recorder *MockEFSAPIMockRecorder
}

// MockEFSAPIMockRecorder is the mock recorder for MockEFSAPI
type MockEFSAPIMockRecorder struct {
	mock *MockEFSAPI
}

// NewMockEFSAPI creates a new mock instance
func NewMockEFSAPI(ctrl *gomock.Controller) *MockEFSAPI {
	mock := &MockEFSAPI{ctrl: ctrl}
	mock.recorder = &MockEFSAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEFSAPI) EXPECT() *MockEFSAPIMockRecorder {
	return m.recorder
}

// CreateFileSystem mocks base method
func (m *MockEFSAPI) CreateFileSystem(arg0 *efs.CreateFileSystemInput) (*efs.FileSystemDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFileSystem", arg0)
	ret0, _ := ret[0].(*efs.FileSystemDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFileSystem indicates an expected call of CreateFileSystem
func (mr *MockEFSAPIMockRecorder) CreateFileSystem(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFileSystem", reflect.TypeOf((*MockEFSAPI)(nil).CreateFileSystem), arg0)
}

// CreateFileSystemRequest mocks base method
func (m *MockEFSAPI) CreateFileSystemRequest(arg0 *efs.CreateFileSystemInput) (*request.Request, *efs.FileSystemDescription) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFileSystemRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*efs.FileSystemDescription)
	return ret0, ret1
}

// CreateFileSystemRequest indicates an expected call of CreateFileSystemRequest
func (mr *MockEFSAPIMockRecorder) CreateFileSystemRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFileSystemRequest", reflect.TypeOf((*MockEFSAPI)(nil).CreateFileSystemRequest), arg0)
}

// CreateFileSystemWithContext mocks base method
func (m *MockEFSAPI) CreateFileSystemWithContext(arg0 context.Context, arg1 *efs.CreateFileSystemInput, arg2 ...request.Option) (*efs.FileSystemDescription, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateFileSystemWithContext", varargs...)
	ret0, _ := ret[0].(*efs.FileSystemDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFileSystemWithContext indicates an expected call of CreateFileSystemWithContext
func (mr *MockEFSAPIMockRecorder) CreateFileSystemWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFileSystemWithContext", reflect.TypeOf((*MockEFSAPI)(nil).CreateFileSystemWithContext), varargs...)
}

// CreateMountTarget mocks base method
func (m *MockEFSAPI) CreateMountTarget(arg0 *efs.CreateMountTargetInput) (*efs.MountTargetDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMountTarget", arg0)
	ret0, _ := ret[0].(*efs.MountTargetDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMountTarget indicates an expected call of CreateMountTarget
func (mr *MockEFSAPIMockRecorder) CreateMountTarget(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMountTarget", reflect.TypeOf((*MockEFSAPI)(nil).CreateMountTarget), arg0)
}

// CreateMountTargetRequest mocks base method
func (m *MockEFSAPI) CreateMountTargetRequest(arg0 *efs.CreateMountTargetInput) (*request.Request, *efs.MountTargetDescription) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMountTargetRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*efs.MountTargetDescription)
	return ret0, ret1
}

// CreateMountTargetRequest indicates an expected call of CreateMountTargetRequest
func (mr *MockEFSAPIMockRecorder) CreateMountTargetRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMountTargetRequest", reflect.TypeOf((*MockEFSAPI)(nil).CreateMountTargetRequest), arg0)
}

// CreateMountTargetWithContext mocks base method
func (m *MockEFSAPI) CreateMountTargetWithContext(arg0 context.Context, arg1 *efs.CreateMountTargetInput, arg2 ...request.Option) (*efs.MountTargetDescription, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMountTargetWithContext", varargs...)
	ret0, _ := ret[0].(*efs.MountTargetDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMountTargetWithContext indicates an expected call of CreateMountTargetWithContext
func (mr *MockEFSAPIMockRecorder) CreateMountTargetWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMountTargetWithContext", reflect.TypeOf((*MockEFSAPI)(nil).CreateMountTargetWithContext), varargs...)
}

// CreateTags mocks base method
func (m *MockEFSAPI) CreateTags(arg0 *efs.CreateTagsInput) (*efs.CreateTagsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTags", arg0)
	ret0, _ := ret[0].(*efs.CreateTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTags indicates an expected call of CreateTags
func (mr *MockEFSAPIMockRecorder) CreateTags(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTags", reflect.TypeOf((*MockEFSAPI)(nil).CreateTags), arg0)
}

// CreateTagsRequest mocks base method
func (m *MockEFSAPI) CreateTagsRequest(arg0 *efs.CreateTagsInput) (*request.Request, *efs.CreateTagsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTagsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*efs.CreateTagsOutput)
	return ret0, ret1
}

// CreateTagsRequest indicates an expected call of CreateTagsRequest
func (mr *MockEFSAPIMockRecorder) CreateTagsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTagsRequest", reflect.TypeOf((*MockEFSAPI)(nil).CreateTagsRequest), arg0)
}

// CreateTagsWithContext mocks base method
func (m *MockEFSAPI) CreateTagsWithContext(arg0 context.Context, arg1 *efs.CreateTagsInput, arg2 ...request.Option) (*efs.CreateTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTagsWithContext", varargs...)
	ret0, _ := ret[0].(*efs.CreateTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTagsWithContext indicates an expected call of CreateTagsWithContext
func (mr *MockEFSAPIMockRecorder) CreateTagsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTagsWithContext", reflect.TypeOf((*MockEFSAPI)(nil).CreateTagsWithContext), varargs...)
}

// DeleteFileSystem mocks base method
func (m *MockEFSAPI) DeleteFileSystem(arg0 *efs.DeleteFileSystemInput) (*efs.DeleteFileSystemOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFileSystem", arg0)
	ret0, _ := ret[0].(*efs.DeleteFileSystemOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFileSystem indicates an expected call of DeleteFileSystem
func (mr *MockEFSAPIMockRecorder) DeleteFileSystem(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFileSystem", reflect.TypeOf((*MockEFSAPI)(nil).DeleteFileSystem), arg0)
}

// DeleteFileSystemRequest mocks base method
func (m *MockEFSAPI) DeleteFileSystemRequest(arg0 *efs.DeleteFileSystemInput) (*request.Request, *efs.DeleteFileSystemOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFileSystemRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*efs.DeleteFileSystemOutput)
	return ret0, ret1
}

// DeleteFileSystemRequest indicates an expected call of DeleteFileSystemRequest
func (mr *MockEFSAPIMockRecorder) DeleteFileSystemRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFileSystemRequest", reflect.TypeOf((*MockEFSAPI)(nil).DeleteFileSystemRequest), arg0)
}

// DeleteFileSystemWithContext mocks base method
func (m *MockEFSAPI) DeleteFileSystemWithContext(arg0 context.Context, arg1 *efs.DeleteFileSystemInput, arg2 ...request.Option) (*efs.DeleteFileSystemOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteFileSystemWithContext", varargs...)
	ret0, _ := ret[0].(*efs.DeleteFileSystemOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFileSystemWithContext indicates an expected call of DeleteFileSystemWithContext
func (mr *MockEFSAPIMockRecorder) DeleteFileSystemWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFileSystemWithContext", reflect.TypeOf((*MockEFSAPI)(nil).DeleteFileSystemWithContext), varargs...)
}

// DeleteMountTarget mocks base method
func (m *MockEFSAPI) DeleteMountTarget(arg0 *efs.DeleteMountTargetInput) (*efs.DeleteMountTargetOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMountTarget", arg0)
	ret0, _ := ret[0].(*efs.DeleteMountTargetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMountTarget indicates an expected call of DeleteMountTarget
func (mr *MockEFSAPIMockRecorder) DeleteMountTarget(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMountTarget", reflect.TypeOf((*MockEFSAPI)(nil).DeleteMountTarget), arg0)
}

// DeleteMountTargetRequest mocks base method
func (m *MockEFSAPI) DeleteMountTargetRequest(arg0 *efs.DeleteMountTargetInput) (*request.Request, *efs.DeleteMountTargetOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMountTargetRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*efs.DeleteMountTargetOutput)
	return ret0, ret1
}

// DeleteMountTargetRequest indicates an expected call of DeleteMountTargetRequest
func (mr *MockEFSAPIMockRecorder) DeleteMountTargetRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMountTargetRequest", reflect.TypeOf((*MockEFSAPI)(nil).DeleteMountTargetRequest), arg0)
}

// DeleteMountTargetWithContext mocks base method
func (m *MockEFSAPI) DeleteMountTargetWithContext(arg0 context.Context, arg1 *efs.DeleteMountTargetInput, arg2 ...request.Option) (*efs.DeleteMountTargetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMountTargetWithContext", varargs...)
	ret0, _ := ret[0].(*efs.DeleteMountTargetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMountTargetWithContext indicates an expected call of DeleteMountTargetWithContext
func (mr *MockEFSAPIMockRecorder) DeleteMountTargetWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMountTargetWithContext", reflect.TypeOf((*MockEFSAPI)(nil).DeleteMountTargetWithContext), varargs...)
}

// DeleteTags mocks base method
func (m *MockEFSAPI) DeleteTags(arg0 *efs.DeleteTagsInput) (*efs.DeleteTagsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTags", arg0)
	ret0, _ := ret[0].(*efs.DeleteTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTags indicates an expected call of DeleteTags
func (mr *MockEFSAPIMockRecorder) DeleteTags(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTags", reflect.TypeOf((*MockEFSAPI)(nil).DeleteTags), arg0)
}

// DeleteTagsRequest mocks base method
func (m *MockEFSAPI) DeleteTagsRequest(arg0 *efs.DeleteTagsInput) (*request.Request, *efs.DeleteTagsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTagsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*efs.DeleteTagsOutput)
	return ret0, ret1
}

// DeleteTagsRequest indicates an expected call of DeleteTagsRequest
func (mr *MockEFSAPIMockRecorder) DeleteTagsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTagsRequest", reflect.TypeOf((*MockEFSAPI)(nil).DeleteTagsRequest), arg0)
}

// DeleteTagsWithContext mocks base method
func (m *MockEFSAPI) DeleteTagsWithContext(arg0 context.Context, arg1 *efs.DeleteTagsInput, arg2 ...request.Option) (*efs.DeleteTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteTagsWithContext", varargs...)
	ret0, _ := ret[0].(*efs.DeleteTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTagsWithContext indicates an expected call of DeleteTagsWithContext
func (mr *MockEFSAPIMockRecorder) DeleteTagsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTagsWithContext", reflect.TypeOf((*MockEFSAPI)(nil).DeleteTagsWithContext), varargs...)
}

// DescribeFileSystems mocks base method
func (m *MockEFSAPI) DescribeFileSystems(arg0 *efs.DescribeFileSystemsInput) (*efs.DescribeFileSystemsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeFileSystems", arg0)
	ret0, _ := ret[0].(*efs.DescribeFileSystemsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeFileSystems indicates an expected call of DescribeFileSystems
func (mr *MockEFSAPIMockRecorder) DescribeFileSystems(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFileSystems", reflect.TypeOf((*MockEFSAPI)(nil).DescribeFileSystems), arg0)
}

// DescribeFileSystemsRequest mocks base method
func (m *MockEFSAPI) DescribeFileSystemsRequest(arg0 *efs.DescribeFileSystemsInput) (*request.Request, *efs.DescribeFileSystemsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeFileSystemsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*efs.DescribeFileSystemsOutput)
	return ret0, ret1
}

// DescribeFileSystemsRequest indicates an expected call of DescribeFileSystemsRequest
func (mr *MockEFSAPIMockRecorder) DescribeFileSystemsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFileSystemsRequest", reflect.TypeOf((*MockEFSAPI)(nil).DescribeFileSystemsRequest), arg0)
}

// DescribeFileSystemsWithContext mocks base method
func (m *MockEFSAPI) DescribeFileSystemsWithContext(arg0 context.Context, arg1 *efs.DescribeFileSystemsInput, arg2 ...request.Option) (*efs.DescribeFileSystemsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeFileSystemsWithContext", varargs...)
	ret0, _ := ret[0].(*efs.DescribeFileSystemsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeFileSystemsWithContext indicates an expected call of DescribeFileSystemsWithContext
func (mr *MockEFSAPIMockRecorder) DescribeFileSystemsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFileSystemsWithContext", reflect.TypeOf((*MockEFSAPI)(nil).DescribeFileSystemsWithContext), varargs...)
}

// DescribeMountTargetSecurityGroups mocks base method
func (m *MockEFSAPI) DescribeMountTargetSecurityGroups(arg0 *efs.DescribeMountTargetSecurityGroupsInput) (*efs.DescribeMountTargetSecurityGroupsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeMountTargetSecurityGroups", arg0)
	ret0, _ := ret[0].(*efs.DescribeMountTargetSecurityGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeMountTargetSecurityGroups indicates an expected call of DescribeMountTargetSecurityGroups
func (mr *MockEFSAPIMockRecorder) DescribeMountTargetSecurityGroups(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeMountTargetSecurityGroups", reflect.TypeOf((*MockEFSAPI)(nil).DescribeMountTargetSecurityGroups), arg0)
}

// DescribeMountTargetSecurityGroupsRequest mocks base method
func (m *MockEFSAPI) DescribeMountTargetSecurityGroupsRequest(arg0 *efs.DescribeMountTargetSecurityGroupsInput) (*request.Request, *efs.DescribeMountTargetSecurityGroupsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeMountTargetSecurityGroupsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*efs.DescribeMountTargetSecurityGroupsOutput)
	return ret0, ret1
}

// DescribeMountTargetSecurityGroupsRequest indicates an expected call of DescribeMountTargetSecurityGroupsRequest
func (mr *MockEFSAPIMockRecorder) DescribeMountTargetSecurityGroupsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeMountTargetSecurityGroupsRequest", reflect.TypeOf((*MockEFSAPI)(nil).DescribeMountTargetSecurityGroupsRequest), arg0)
}

// DescribeMountTargetSecurityGroupsWithContext mocks base method
func (m *MockEFSAPI) DescribeMountTargetSecurityGroupsWithContext(arg0 context.Context, arg1 *efs.DescribeMountTargetSecurityGroupsInput, arg2 ...request.Option) (*efs.DescribeMountTargetSecurityGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeMountTargetSecurityGroupsWithContext", varargs...)
	ret0, _ := ret[0].(*efs.DescribeMountTargetSecurityGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeMountTargetSecurityGroupsWithContext indicates an expected call of DescribeMountTargetSecurityGroupsWithContext
func (mr *MockEFSAPIMockRecorder) DescribeMountTargetSecurityGroupsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeMountTargetSecurityGroupsWithContext", reflect.TypeOf((*MockEFSAPI)(nil).DescribeMountTargetSecurityGroupsWithContext), varargs...)
}

// DescribeMountTargets mocks base method
func (m *MockEFSAPI) DescribeMountTargets(arg0 *efs.DescribeMountTargetsInput) (*efs.DescribeMountTargetsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeMountTargets", arg0)
	ret0, _ := ret[0].(*efs.DescribeMountTargetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeMountTargets indicates an expected call of DescribeMountTargets
func (mr *MockEFSAPIMockRecorder) DescribeMountTargets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeMountTargets", reflect.TypeOf((*MockEFSAPI)(nil).DescribeMountTargets), arg0)
}

// DescribeMountTargetsRequest mocks base method
func (m *MockEFSAPI) DescribeMountTargetsRequest(arg0 *efs.DescribeMountTargetsInput) (*request.Request, *efs.DescribeMountTargetsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeMountTargetsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*efs.DescribeMountTargetsOutput)
	return ret0, ret1
}

// DescribeMountTargetsRequest indicates an expected call of DescribeMountTargetsRequest
func (mr *MockEFSAPIMockRecorder) DescribeMountTargetsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeMountTargetsRequest", reflect.TypeOf((*MockEFSAPI)(nil).DescribeMountTargetsRequest), arg0)
}

// DescribeMountTargetsWithContext mocks base method
func (m *MockEFSAPI) DescribeMountTargetsWithContext(arg0 context.Context, arg1 *efs.DescribeMountTargetsInput, arg2 ...request.Option) (*efs.DescribeMountTargetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeMountTargetsWithContext", varargs...)
	ret0, _ := ret[0].(*efs.DescribeMountTargetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeMountTargetsWithContext indicates an expected call of DescribeMountTargetsWithContext
func (mr *MockEFSAPIMockRecorder) DescribeMountTargetsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeMountTargetsWithContext", reflect.TypeOf((*MockEFSAPI)(nil).DescribeMountTargetsWithContext), varargs...)
}

// DescribeTags mocks base method
func (m *MockEFSAPI) DescribeTags(arg0 *efs.DescribeTagsInput) (*efs.DescribeTagsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTags", arg0)
	ret0, _ := ret[0].(*efs.DescribeTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTags indicates an expected call of DescribeTags
func (mr *MockEFSAPIMockRecorder) DescribeTags(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTags", reflect.TypeOf((*MockEFSAPI)(nil).DescribeTags), arg0)
}

// DescribeTagsRequest mocks base method
func (m *MockEFSAPI) DescribeTagsRequest(arg0 *efs.DescribeTagsInput) (*request.Request, *efs.DescribeTagsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTagsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*efs.DescribeTagsOutput)
	return ret0, ret1
}

// DescribeTagsRequest indicates an expected call of DescribeTagsRequest
func (mr *MockEFSAPIMockRecorder) DescribeTagsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTagsRequest", reflect.TypeOf((*MockEFSAPI)(nil).DescribeTagsRequest), arg0)
}

// DescribeTagsWithContext mocks base method
func (m *MockEFSAPI) DescribeTagsWithContext(arg0 context.Context, arg1 *efs.DescribeTagsInput, arg2 ...request.Option) (*efs.DescribeTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTagsWithContext", varargs...)
	ret0, _ := ret[0].(*efs.DescribeTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTagsWithContext indicates an expected call of DescribeTagsWithContext
func (mr *MockEFSAPIMockRecorder) DescribeTagsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTagsWithContext", reflect.TypeOf((*MockEFSAPI)(nil).DescribeTagsWithContext), varargs...)
}

// ModifyMountTargetSecurityGroups mocks base method
func (m *MockEFSAPI) ModifyMountTargetSecurityGroups(arg0 *efs.ModifyMountTargetSecurityGroupsInput) (*efs.ModifyMountTargetSecurityGroupsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyMountTargetSecurityGroups", arg0)
	ret0, _ := ret[0].(*efs.ModifyMountTargetSecurityGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyMountTargetSecurityGroups indicates an expected call of ModifyMountTargetSecurityGroups
func (mr *MockEFSAPIMockRecorder) ModifyMountTargetSecurityGroups(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyMountTargetSecurityGroups", reflect.TypeOf((*MockEFSAPI)(nil).ModifyMountTargetSecurityGroups), arg0)
}

// ModifyMountTargetSecurityGroupsRequest mocks base method
func (m *MockEFSAPI) ModifyMountTargetSecurityGroupsRequest(arg0 *efs.ModifyMountTargetSecurityGroupsInput) (*request.Request, *efs.ModifyMountTargetSecurityGroupsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyMountTargetSecurityGroupsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*efs.ModifyMountTargetSecurityGroupsOutput)
	return ret0, ret1
}

// ModifyMountTargetSecurityGroupsRequest indicates an expected call of ModifyMountTargetSecurityGroupsRequest
func (mr *MockEFSAPIMockRecorder) ModifyMountTargetSecurityGroupsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyMountTargetSecurityGroupsRequest", reflect.TypeOf((*MockEFSAPI)(nil).ModifyMountTargetSecurityGroupsRequest), arg0)
}

// ModifyMountTargetSecurityGroupsWithContext mocks base method
func (m *MockEFSAPI) ModifyMountTargetSecurityGroupsWithContext(arg0 context.Context, arg1 *efs.ModifyMountTargetSecurityGroupsInput, arg2 ...request.Option) (*efs.ModifyMountTargetSecurityGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ModifyMountTargetSecurityGroupsWithContext", varargs...)
	ret0, _ := ret[0].(*efs.ModifyMountTargetSecurityGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyMountTargetSecurityGroupsWithContext indicates an expected call of ModifyMountTargetSecurityGroupsWithContext
func (mr *MockEFSAPIMockRecorder) ModifyMountTargetSecurityGroupsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyMountTargetSecurityGroupsWithContext", reflect.TypeOf((*MockEFSAPI)(nil).ModifyMountTargetSecurityGroupsWithContext), varargs...)
}

// UpdateFileSystem mocks base method
func (m *MockEFSAPI) UpdateFileSystem(arg0 *efs.UpdateFileSystemInput) (*efs.UpdateFileSystemOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFileSystem", arg0)
	ret0, _ := ret[0].(*efs.UpdateFileSystemOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFileSystem indicates an expected call of UpdateFileSystem
func (mr *MockEFSAPIMockRecorder) UpdateFileSystem(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFileSystem", reflect.TypeOf((*MockEFSAPI)(nil).UpdateFileSystem), arg0)
}

// UpdateFileSystemRequest mocks base method
func (m *MockEFSAPI) UpdateFileSystemRequest(arg0 *efs.UpdateFileSystemInput) (*request.Request, *efs.UpdateFileSystemOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFileSystemRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*efs.UpdateFileSystemOutput)
	return ret0, ret1
}

// UpdateFileSystemRequest indicates an expected call of UpdateFileSystemRequest
func (mr *MockEFSAPIMockRecorder) UpdateFileSystemRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFileSystemRequest", reflect.TypeOf((*MockEFSAPI)(nil).UpdateFileSystemRequest), arg0)
}

// UpdateFileSystemWithContext mocks base method
func (m *MockEFSAPI) UpdateFileSystemWithContext(arg0 context.Context, arg1 *efs.UpdateFileSystemInput, arg2 ...request.Option) (*efs.UpdateFileSystemOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateFileSystemWithContext", varargs...)
	ret0, _ := ret[0].(*efs.UpdateFileSystemOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFileSystemWithContext indicates an expected call of UpdateFileSystemWithContext
func (mr *MockEFSAPIMockRecorder) UpdateFileSystemWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFileSystemWithContext", reflect.TypeOf((*MockEFSAPI)(nil).UpdateFileSystemWithContext), varargs...)
}
