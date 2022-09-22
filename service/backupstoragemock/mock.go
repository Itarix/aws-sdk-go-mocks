// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/backupstorage/backupstorageiface (interfaces: BackupStorageAPI)

// Package backupstoragemock is a generated GoMock package.
package backupstoragemock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	backupstorage "github.com/aws/aws-sdk-go/service/backupstorage"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBackupStorageAPI is a mock of BackupStorageAPI interface
type MockBackupStorageAPI struct {
	ctrl     *gomock.Controller
	recorder *MockBackupStorageAPIMockRecorder
}

// MockBackupStorageAPIMockRecorder is the mock recorder for MockBackupStorageAPI
type MockBackupStorageAPIMockRecorder struct {
	mock *MockBackupStorageAPI
}

// NewMockBackupStorageAPI creates a new mock instance
func NewMockBackupStorageAPI(ctrl *gomock.Controller) *MockBackupStorageAPI {
	mock := &MockBackupStorageAPI{ctrl: ctrl}
	mock.recorder = &MockBackupStorageAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBackupStorageAPI) EXPECT() *MockBackupStorageAPIMockRecorder {
	return m.recorder
}

// DeleteObject mocks base method
func (m *MockBackupStorageAPI) DeleteObject(arg0 *backupstorage.DeleteObjectInput) (*backupstorage.DeleteObjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObject", arg0)
	ret0, _ := ret[0].(*backupstorage.DeleteObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteObject indicates an expected call of DeleteObject
func (mr *MockBackupStorageAPIMockRecorder) DeleteObject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObject", reflect.TypeOf((*MockBackupStorageAPI)(nil).DeleteObject), arg0)
}

// DeleteObjectRequest mocks base method
func (m *MockBackupStorageAPI) DeleteObjectRequest(arg0 *backupstorage.DeleteObjectInput) (*request.Request, *backupstorage.DeleteObjectOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObjectRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*backupstorage.DeleteObjectOutput)
	return ret0, ret1
}

// DeleteObjectRequest indicates an expected call of DeleteObjectRequest
func (mr *MockBackupStorageAPIMockRecorder) DeleteObjectRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObjectRequest", reflect.TypeOf((*MockBackupStorageAPI)(nil).DeleteObjectRequest), arg0)
}

// DeleteObjectWithContext mocks base method
func (m *MockBackupStorageAPI) DeleteObjectWithContext(arg0 context.Context, arg1 *backupstorage.DeleteObjectInput, arg2 ...request.Option) (*backupstorage.DeleteObjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteObjectWithContext", varargs...)
	ret0, _ := ret[0].(*backupstorage.DeleteObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteObjectWithContext indicates an expected call of DeleteObjectWithContext
func (mr *MockBackupStorageAPIMockRecorder) DeleteObjectWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObjectWithContext", reflect.TypeOf((*MockBackupStorageAPI)(nil).DeleteObjectWithContext), varargs...)
}

// GetChunk mocks base method
func (m *MockBackupStorageAPI) GetChunk(arg0 *backupstorage.GetChunkInput) (*backupstorage.GetChunkOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChunk", arg0)
	ret0, _ := ret[0].(*backupstorage.GetChunkOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChunk indicates an expected call of GetChunk
func (mr *MockBackupStorageAPIMockRecorder) GetChunk(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChunk", reflect.TypeOf((*MockBackupStorageAPI)(nil).GetChunk), arg0)
}

// GetChunkRequest mocks base method
func (m *MockBackupStorageAPI) GetChunkRequest(arg0 *backupstorage.GetChunkInput) (*request.Request, *backupstorage.GetChunkOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChunkRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*backupstorage.GetChunkOutput)
	return ret0, ret1
}

// GetChunkRequest indicates an expected call of GetChunkRequest
func (mr *MockBackupStorageAPIMockRecorder) GetChunkRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChunkRequest", reflect.TypeOf((*MockBackupStorageAPI)(nil).GetChunkRequest), arg0)
}

// GetChunkWithContext mocks base method
func (m *MockBackupStorageAPI) GetChunkWithContext(arg0 context.Context, arg1 *backupstorage.GetChunkInput, arg2 ...request.Option) (*backupstorage.GetChunkOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetChunkWithContext", varargs...)
	ret0, _ := ret[0].(*backupstorage.GetChunkOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChunkWithContext indicates an expected call of GetChunkWithContext
func (mr *MockBackupStorageAPIMockRecorder) GetChunkWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChunkWithContext", reflect.TypeOf((*MockBackupStorageAPI)(nil).GetChunkWithContext), varargs...)
}

// GetObjectMetadata mocks base method
func (m *MockBackupStorageAPI) GetObjectMetadata(arg0 *backupstorage.GetObjectMetadataInput) (*backupstorage.GetObjectMetadataOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectMetadata", arg0)
	ret0, _ := ret[0].(*backupstorage.GetObjectMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectMetadata indicates an expected call of GetObjectMetadata
func (mr *MockBackupStorageAPIMockRecorder) GetObjectMetadata(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectMetadata", reflect.TypeOf((*MockBackupStorageAPI)(nil).GetObjectMetadata), arg0)
}

// GetObjectMetadataRequest mocks base method
func (m *MockBackupStorageAPI) GetObjectMetadataRequest(arg0 *backupstorage.GetObjectMetadataInput) (*request.Request, *backupstorage.GetObjectMetadataOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectMetadataRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*backupstorage.GetObjectMetadataOutput)
	return ret0, ret1
}

// GetObjectMetadataRequest indicates an expected call of GetObjectMetadataRequest
func (mr *MockBackupStorageAPIMockRecorder) GetObjectMetadataRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectMetadataRequest", reflect.TypeOf((*MockBackupStorageAPI)(nil).GetObjectMetadataRequest), arg0)
}

// GetObjectMetadataWithContext mocks base method
func (m *MockBackupStorageAPI) GetObjectMetadataWithContext(arg0 context.Context, arg1 *backupstorage.GetObjectMetadataInput, arg2 ...request.Option) (*backupstorage.GetObjectMetadataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObjectMetadataWithContext", varargs...)
	ret0, _ := ret[0].(*backupstorage.GetObjectMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectMetadataWithContext indicates an expected call of GetObjectMetadataWithContext
func (mr *MockBackupStorageAPIMockRecorder) GetObjectMetadataWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectMetadataWithContext", reflect.TypeOf((*MockBackupStorageAPI)(nil).GetObjectMetadataWithContext), varargs...)
}

// ListChunks mocks base method
func (m *MockBackupStorageAPI) ListChunks(arg0 *backupstorage.ListChunksInput) (*backupstorage.ListChunksOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListChunks", arg0)
	ret0, _ := ret[0].(*backupstorage.ListChunksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListChunks indicates an expected call of ListChunks
func (mr *MockBackupStorageAPIMockRecorder) ListChunks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListChunks", reflect.TypeOf((*MockBackupStorageAPI)(nil).ListChunks), arg0)
}

// ListChunksPages mocks base method
func (m *MockBackupStorageAPI) ListChunksPages(arg0 *backupstorage.ListChunksInput, arg1 func(*backupstorage.ListChunksOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListChunksPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListChunksPages indicates an expected call of ListChunksPages
func (mr *MockBackupStorageAPIMockRecorder) ListChunksPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListChunksPages", reflect.TypeOf((*MockBackupStorageAPI)(nil).ListChunksPages), arg0, arg1)
}

// ListChunksPagesWithContext mocks base method
func (m *MockBackupStorageAPI) ListChunksPagesWithContext(arg0 context.Context, arg1 *backupstorage.ListChunksInput, arg2 func(*backupstorage.ListChunksOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListChunksPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListChunksPagesWithContext indicates an expected call of ListChunksPagesWithContext
func (mr *MockBackupStorageAPIMockRecorder) ListChunksPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListChunksPagesWithContext", reflect.TypeOf((*MockBackupStorageAPI)(nil).ListChunksPagesWithContext), varargs...)
}

// ListChunksRequest mocks base method
func (m *MockBackupStorageAPI) ListChunksRequest(arg0 *backupstorage.ListChunksInput) (*request.Request, *backupstorage.ListChunksOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListChunksRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*backupstorage.ListChunksOutput)
	return ret0, ret1
}

// ListChunksRequest indicates an expected call of ListChunksRequest
func (mr *MockBackupStorageAPIMockRecorder) ListChunksRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListChunksRequest", reflect.TypeOf((*MockBackupStorageAPI)(nil).ListChunksRequest), arg0)
}

// ListChunksWithContext mocks base method
func (m *MockBackupStorageAPI) ListChunksWithContext(arg0 context.Context, arg1 *backupstorage.ListChunksInput, arg2 ...request.Option) (*backupstorage.ListChunksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListChunksWithContext", varargs...)
	ret0, _ := ret[0].(*backupstorage.ListChunksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListChunksWithContext indicates an expected call of ListChunksWithContext
func (mr *MockBackupStorageAPIMockRecorder) ListChunksWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListChunksWithContext", reflect.TypeOf((*MockBackupStorageAPI)(nil).ListChunksWithContext), varargs...)
}

// ListObjects mocks base method
func (m *MockBackupStorageAPI) ListObjects(arg0 *backupstorage.ListObjectsInput) (*backupstorage.ListObjectsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListObjects", arg0)
	ret0, _ := ret[0].(*backupstorage.ListObjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListObjects indicates an expected call of ListObjects
func (mr *MockBackupStorageAPIMockRecorder) ListObjects(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjects", reflect.TypeOf((*MockBackupStorageAPI)(nil).ListObjects), arg0)
}

// ListObjectsPages mocks base method
func (m *MockBackupStorageAPI) ListObjectsPages(arg0 *backupstorage.ListObjectsInput, arg1 func(*backupstorage.ListObjectsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListObjectsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListObjectsPages indicates an expected call of ListObjectsPages
func (mr *MockBackupStorageAPIMockRecorder) ListObjectsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjectsPages", reflect.TypeOf((*MockBackupStorageAPI)(nil).ListObjectsPages), arg0, arg1)
}

// ListObjectsPagesWithContext mocks base method
func (m *MockBackupStorageAPI) ListObjectsPagesWithContext(arg0 context.Context, arg1 *backupstorage.ListObjectsInput, arg2 func(*backupstorage.ListObjectsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListObjectsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListObjectsPagesWithContext indicates an expected call of ListObjectsPagesWithContext
func (mr *MockBackupStorageAPIMockRecorder) ListObjectsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjectsPagesWithContext", reflect.TypeOf((*MockBackupStorageAPI)(nil).ListObjectsPagesWithContext), varargs...)
}

// ListObjectsRequest mocks base method
func (m *MockBackupStorageAPI) ListObjectsRequest(arg0 *backupstorage.ListObjectsInput) (*request.Request, *backupstorage.ListObjectsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListObjectsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*backupstorage.ListObjectsOutput)
	return ret0, ret1
}

// ListObjectsRequest indicates an expected call of ListObjectsRequest
func (mr *MockBackupStorageAPIMockRecorder) ListObjectsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjectsRequest", reflect.TypeOf((*MockBackupStorageAPI)(nil).ListObjectsRequest), arg0)
}

// ListObjectsWithContext mocks base method
func (m *MockBackupStorageAPI) ListObjectsWithContext(arg0 context.Context, arg1 *backupstorage.ListObjectsInput, arg2 ...request.Option) (*backupstorage.ListObjectsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListObjectsWithContext", varargs...)
	ret0, _ := ret[0].(*backupstorage.ListObjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListObjectsWithContext indicates an expected call of ListObjectsWithContext
func (mr *MockBackupStorageAPIMockRecorder) ListObjectsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjectsWithContext", reflect.TypeOf((*MockBackupStorageAPI)(nil).ListObjectsWithContext), varargs...)
}

// NotifyObjectComplete mocks base method
func (m *MockBackupStorageAPI) NotifyObjectComplete(arg0 *backupstorage.NotifyObjectCompleteInput) (*backupstorage.NotifyObjectCompleteOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyObjectComplete", arg0)
	ret0, _ := ret[0].(*backupstorage.NotifyObjectCompleteOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotifyObjectComplete indicates an expected call of NotifyObjectComplete
func (mr *MockBackupStorageAPIMockRecorder) NotifyObjectComplete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyObjectComplete", reflect.TypeOf((*MockBackupStorageAPI)(nil).NotifyObjectComplete), arg0)
}

// NotifyObjectCompleteRequest mocks base method
func (m *MockBackupStorageAPI) NotifyObjectCompleteRequest(arg0 *backupstorage.NotifyObjectCompleteInput) (*request.Request, *backupstorage.NotifyObjectCompleteOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyObjectCompleteRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*backupstorage.NotifyObjectCompleteOutput)
	return ret0, ret1
}

// NotifyObjectCompleteRequest indicates an expected call of NotifyObjectCompleteRequest
func (mr *MockBackupStorageAPIMockRecorder) NotifyObjectCompleteRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyObjectCompleteRequest", reflect.TypeOf((*MockBackupStorageAPI)(nil).NotifyObjectCompleteRequest), arg0)
}

// NotifyObjectCompleteWithContext mocks base method
func (m *MockBackupStorageAPI) NotifyObjectCompleteWithContext(arg0 context.Context, arg1 *backupstorage.NotifyObjectCompleteInput, arg2 ...request.Option) (*backupstorage.NotifyObjectCompleteOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NotifyObjectCompleteWithContext", varargs...)
	ret0, _ := ret[0].(*backupstorage.NotifyObjectCompleteOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotifyObjectCompleteWithContext indicates an expected call of NotifyObjectCompleteWithContext
func (mr *MockBackupStorageAPIMockRecorder) NotifyObjectCompleteWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyObjectCompleteWithContext", reflect.TypeOf((*MockBackupStorageAPI)(nil).NotifyObjectCompleteWithContext), varargs...)
}

// PutChunk mocks base method
func (m *MockBackupStorageAPI) PutChunk(arg0 *backupstorage.PutChunkInput) (*backupstorage.PutChunkOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutChunk", arg0)
	ret0, _ := ret[0].(*backupstorage.PutChunkOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutChunk indicates an expected call of PutChunk
func (mr *MockBackupStorageAPIMockRecorder) PutChunk(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutChunk", reflect.TypeOf((*MockBackupStorageAPI)(nil).PutChunk), arg0)
}

// PutChunkRequest mocks base method
func (m *MockBackupStorageAPI) PutChunkRequest(arg0 *backupstorage.PutChunkInput) (*request.Request, *backupstorage.PutChunkOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutChunkRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*backupstorage.PutChunkOutput)
	return ret0, ret1
}

// PutChunkRequest indicates an expected call of PutChunkRequest
func (mr *MockBackupStorageAPIMockRecorder) PutChunkRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutChunkRequest", reflect.TypeOf((*MockBackupStorageAPI)(nil).PutChunkRequest), arg0)
}

// PutChunkWithContext mocks base method
func (m *MockBackupStorageAPI) PutChunkWithContext(arg0 context.Context, arg1 *backupstorage.PutChunkInput, arg2 ...request.Option) (*backupstorage.PutChunkOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutChunkWithContext", varargs...)
	ret0, _ := ret[0].(*backupstorage.PutChunkOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutChunkWithContext indicates an expected call of PutChunkWithContext
func (mr *MockBackupStorageAPIMockRecorder) PutChunkWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutChunkWithContext", reflect.TypeOf((*MockBackupStorageAPI)(nil).PutChunkWithContext), varargs...)
}

// PutObject mocks base method
func (m *MockBackupStorageAPI) PutObject(arg0 *backupstorage.PutObjectInput) (*backupstorage.PutObjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutObject", arg0)
	ret0, _ := ret[0].(*backupstorage.PutObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutObject indicates an expected call of PutObject
func (mr *MockBackupStorageAPIMockRecorder) PutObject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObject", reflect.TypeOf((*MockBackupStorageAPI)(nil).PutObject), arg0)
}

// PutObjectRequest mocks base method
func (m *MockBackupStorageAPI) PutObjectRequest(arg0 *backupstorage.PutObjectInput) (*request.Request, *backupstorage.PutObjectOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutObjectRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*backupstorage.PutObjectOutput)
	return ret0, ret1
}

// PutObjectRequest indicates an expected call of PutObjectRequest
func (mr *MockBackupStorageAPIMockRecorder) PutObjectRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObjectRequest", reflect.TypeOf((*MockBackupStorageAPI)(nil).PutObjectRequest), arg0)
}

// PutObjectWithContext mocks base method
func (m *MockBackupStorageAPI) PutObjectWithContext(arg0 context.Context, arg1 *backupstorage.PutObjectInput, arg2 ...request.Option) (*backupstorage.PutObjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutObjectWithContext", varargs...)
	ret0, _ := ret[0].(*backupstorage.PutObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutObjectWithContext indicates an expected call of PutObjectWithContext
func (mr *MockBackupStorageAPIMockRecorder) PutObjectWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObjectWithContext", reflect.TypeOf((*MockBackupStorageAPI)(nil).PutObjectWithContext), varargs...)
}

// StartObject mocks base method
func (m *MockBackupStorageAPI) StartObject(arg0 *backupstorage.StartObjectInput) (*backupstorage.StartObjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartObject", arg0)
	ret0, _ := ret[0].(*backupstorage.StartObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartObject indicates an expected call of StartObject
func (mr *MockBackupStorageAPIMockRecorder) StartObject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartObject", reflect.TypeOf((*MockBackupStorageAPI)(nil).StartObject), arg0)
}

// StartObjectRequest mocks base method
func (m *MockBackupStorageAPI) StartObjectRequest(arg0 *backupstorage.StartObjectInput) (*request.Request, *backupstorage.StartObjectOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartObjectRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*backupstorage.StartObjectOutput)
	return ret0, ret1
}

// StartObjectRequest indicates an expected call of StartObjectRequest
func (mr *MockBackupStorageAPIMockRecorder) StartObjectRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartObjectRequest", reflect.TypeOf((*MockBackupStorageAPI)(nil).StartObjectRequest), arg0)
}

// StartObjectWithContext mocks base method
func (m *MockBackupStorageAPI) StartObjectWithContext(arg0 context.Context, arg1 *backupstorage.StartObjectInput, arg2 ...request.Option) (*backupstorage.StartObjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StartObjectWithContext", varargs...)
	ret0, _ := ret[0].(*backupstorage.StartObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartObjectWithContext indicates an expected call of StartObjectWithContext
func (mr *MockBackupStorageAPIMockRecorder) StartObjectWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartObjectWithContext", reflect.TypeOf((*MockBackupStorageAPI)(nil).StartObjectWithContext), varargs...)
}