// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/mediastoredata/mediastoredataiface (interfaces: MediaStoreDataAPI)

// Package mediastoredatamock is a generated GoMock package.
package mediastoredatamock

import (
	context "context"
	reflect "reflect"

	request "github.com/aws/aws-sdk-go/aws/request"
	mediastoredata "github.com/aws/aws-sdk-go/service/mediastoredata"
	gomock "github.com/golang/mock/gomock"
)

// MockMediaStoreDataAPI is a mock of MediaStoreDataAPI interface.
type MockMediaStoreDataAPI struct {
	ctrl     *gomock.Controller
	recorder *MockMediaStoreDataAPIMockRecorder
}

// MockMediaStoreDataAPIMockRecorder is the mock recorder for MockMediaStoreDataAPI.
type MockMediaStoreDataAPIMockRecorder struct {
	mock *MockMediaStoreDataAPI
}

// NewMockMediaStoreDataAPI creates a new mock instance.
func NewMockMediaStoreDataAPI(ctrl *gomock.Controller) *MockMediaStoreDataAPI {
	mock := &MockMediaStoreDataAPI{ctrl: ctrl}
	mock.recorder = &MockMediaStoreDataAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMediaStoreDataAPI) EXPECT() *MockMediaStoreDataAPIMockRecorder {
	return m.recorder
}

// DeleteObject mocks base method.
func (m *MockMediaStoreDataAPI) DeleteObject(arg0 *mediastoredata.DeleteObjectInput) (*mediastoredata.DeleteObjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObject", arg0)
	ret0, _ := ret[0].(*mediastoredata.DeleteObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteObject indicates an expected call of DeleteObject.
func (mr *MockMediaStoreDataAPIMockRecorder) DeleteObject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObject", reflect.TypeOf((*MockMediaStoreDataAPI)(nil).DeleteObject), arg0)
}

// DeleteObjectRequest mocks base method.
func (m *MockMediaStoreDataAPI) DeleteObjectRequest(arg0 *mediastoredata.DeleteObjectInput) (*request.Request, *mediastoredata.DeleteObjectOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObjectRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mediastoredata.DeleteObjectOutput)
	return ret0, ret1
}

// DeleteObjectRequest indicates an expected call of DeleteObjectRequest.
func (mr *MockMediaStoreDataAPIMockRecorder) DeleteObjectRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObjectRequest", reflect.TypeOf((*MockMediaStoreDataAPI)(nil).DeleteObjectRequest), arg0)
}

// DeleteObjectWithContext mocks base method.
func (m *MockMediaStoreDataAPI) DeleteObjectWithContext(arg0 context.Context, arg1 *mediastoredata.DeleteObjectInput, arg2 ...request.Option) (*mediastoredata.DeleteObjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteObjectWithContext", varargs...)
	ret0, _ := ret[0].(*mediastoredata.DeleteObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteObjectWithContext indicates an expected call of DeleteObjectWithContext.
func (mr *MockMediaStoreDataAPIMockRecorder) DeleteObjectWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObjectWithContext", reflect.TypeOf((*MockMediaStoreDataAPI)(nil).DeleteObjectWithContext), varargs...)
}

// DescribeObject mocks base method.
func (m *MockMediaStoreDataAPI) DescribeObject(arg0 *mediastoredata.DescribeObjectInput) (*mediastoredata.DescribeObjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeObject", arg0)
	ret0, _ := ret[0].(*mediastoredata.DescribeObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeObject indicates an expected call of DescribeObject.
func (mr *MockMediaStoreDataAPIMockRecorder) DescribeObject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeObject", reflect.TypeOf((*MockMediaStoreDataAPI)(nil).DescribeObject), arg0)
}

// DescribeObjectRequest mocks base method.
func (m *MockMediaStoreDataAPI) DescribeObjectRequest(arg0 *mediastoredata.DescribeObjectInput) (*request.Request, *mediastoredata.DescribeObjectOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeObjectRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mediastoredata.DescribeObjectOutput)
	return ret0, ret1
}

// DescribeObjectRequest indicates an expected call of DescribeObjectRequest.
func (mr *MockMediaStoreDataAPIMockRecorder) DescribeObjectRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeObjectRequest", reflect.TypeOf((*MockMediaStoreDataAPI)(nil).DescribeObjectRequest), arg0)
}

// DescribeObjectWithContext mocks base method.
func (m *MockMediaStoreDataAPI) DescribeObjectWithContext(arg0 context.Context, arg1 *mediastoredata.DescribeObjectInput, arg2 ...request.Option) (*mediastoredata.DescribeObjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeObjectWithContext", varargs...)
	ret0, _ := ret[0].(*mediastoredata.DescribeObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeObjectWithContext indicates an expected call of DescribeObjectWithContext.
func (mr *MockMediaStoreDataAPIMockRecorder) DescribeObjectWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeObjectWithContext", reflect.TypeOf((*MockMediaStoreDataAPI)(nil).DescribeObjectWithContext), varargs...)
}

// GetObject mocks base method.
func (m *MockMediaStoreDataAPI) GetObject(arg0 *mediastoredata.GetObjectInput) (*mediastoredata.GetObjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObject", arg0)
	ret0, _ := ret[0].(*mediastoredata.GetObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObject indicates an expected call of GetObject.
func (mr *MockMediaStoreDataAPIMockRecorder) GetObject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*MockMediaStoreDataAPI)(nil).GetObject), arg0)
}

// GetObjectRequest mocks base method.
func (m *MockMediaStoreDataAPI) GetObjectRequest(arg0 *mediastoredata.GetObjectInput) (*request.Request, *mediastoredata.GetObjectOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mediastoredata.GetObjectOutput)
	return ret0, ret1
}

// GetObjectRequest indicates an expected call of GetObjectRequest.
func (mr *MockMediaStoreDataAPIMockRecorder) GetObjectRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectRequest", reflect.TypeOf((*MockMediaStoreDataAPI)(nil).GetObjectRequest), arg0)
}

// GetObjectWithContext mocks base method.
func (m *MockMediaStoreDataAPI) GetObjectWithContext(arg0 context.Context, arg1 *mediastoredata.GetObjectInput, arg2 ...request.Option) (*mediastoredata.GetObjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObjectWithContext", varargs...)
	ret0, _ := ret[0].(*mediastoredata.GetObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectWithContext indicates an expected call of GetObjectWithContext.
func (mr *MockMediaStoreDataAPIMockRecorder) GetObjectWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectWithContext", reflect.TypeOf((*MockMediaStoreDataAPI)(nil).GetObjectWithContext), varargs...)
}

// ListItems mocks base method.
func (m *MockMediaStoreDataAPI) ListItems(arg0 *mediastoredata.ListItemsInput) (*mediastoredata.ListItemsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListItems", arg0)
	ret0, _ := ret[0].(*mediastoredata.ListItemsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListItems indicates an expected call of ListItems.
func (mr *MockMediaStoreDataAPIMockRecorder) ListItems(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListItems", reflect.TypeOf((*MockMediaStoreDataAPI)(nil).ListItems), arg0)
}

// ListItemsPages mocks base method.
func (m *MockMediaStoreDataAPI) ListItemsPages(arg0 *mediastoredata.ListItemsInput, arg1 func(*mediastoredata.ListItemsOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListItemsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListItemsPages indicates an expected call of ListItemsPages.
func (mr *MockMediaStoreDataAPIMockRecorder) ListItemsPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListItemsPages", reflect.TypeOf((*MockMediaStoreDataAPI)(nil).ListItemsPages), arg0, arg1)
}

// ListItemsPagesWithContext mocks base method.
func (m *MockMediaStoreDataAPI) ListItemsPagesWithContext(arg0 context.Context, arg1 *mediastoredata.ListItemsInput, arg2 func(*mediastoredata.ListItemsOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListItemsPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListItemsPagesWithContext indicates an expected call of ListItemsPagesWithContext.
func (mr *MockMediaStoreDataAPIMockRecorder) ListItemsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListItemsPagesWithContext", reflect.TypeOf((*MockMediaStoreDataAPI)(nil).ListItemsPagesWithContext), varargs...)
}

// ListItemsRequest mocks base method.
func (m *MockMediaStoreDataAPI) ListItemsRequest(arg0 *mediastoredata.ListItemsInput) (*request.Request, *mediastoredata.ListItemsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListItemsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mediastoredata.ListItemsOutput)
	return ret0, ret1
}

// ListItemsRequest indicates an expected call of ListItemsRequest.
func (mr *MockMediaStoreDataAPIMockRecorder) ListItemsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListItemsRequest", reflect.TypeOf((*MockMediaStoreDataAPI)(nil).ListItemsRequest), arg0)
}

// ListItemsWithContext mocks base method.
func (m *MockMediaStoreDataAPI) ListItemsWithContext(arg0 context.Context, arg1 *mediastoredata.ListItemsInput, arg2 ...request.Option) (*mediastoredata.ListItemsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListItemsWithContext", varargs...)
	ret0, _ := ret[0].(*mediastoredata.ListItemsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListItemsWithContext indicates an expected call of ListItemsWithContext.
func (mr *MockMediaStoreDataAPIMockRecorder) ListItemsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListItemsWithContext", reflect.TypeOf((*MockMediaStoreDataAPI)(nil).ListItemsWithContext), varargs...)
}

// PutObject mocks base method.
func (m *MockMediaStoreDataAPI) PutObject(arg0 *mediastoredata.PutObjectInput) (*mediastoredata.PutObjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutObject", arg0)
	ret0, _ := ret[0].(*mediastoredata.PutObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutObject indicates an expected call of PutObject.
func (mr *MockMediaStoreDataAPIMockRecorder) PutObject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObject", reflect.TypeOf((*MockMediaStoreDataAPI)(nil).PutObject), arg0)
}

// PutObjectRequest mocks base method.
func (m *MockMediaStoreDataAPI) PutObjectRequest(arg0 *mediastoredata.PutObjectInput) (*request.Request, *mediastoredata.PutObjectOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutObjectRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*mediastoredata.PutObjectOutput)
	return ret0, ret1
}

// PutObjectRequest indicates an expected call of PutObjectRequest.
func (mr *MockMediaStoreDataAPIMockRecorder) PutObjectRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObjectRequest", reflect.TypeOf((*MockMediaStoreDataAPI)(nil).PutObjectRequest), arg0)
}

// PutObjectWithContext mocks base method.
func (m *MockMediaStoreDataAPI) PutObjectWithContext(arg0 context.Context, arg1 *mediastoredata.PutObjectInput, arg2 ...request.Option) (*mediastoredata.PutObjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutObjectWithContext", varargs...)
	ret0, _ := ret[0].(*mediastoredata.PutObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutObjectWithContext indicates an expected call of PutObjectWithContext.
func (mr *MockMediaStoreDataAPIMockRecorder) PutObjectWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObjectWithContext", reflect.TypeOf((*MockMediaStoreDataAPI)(nil).PutObjectWithContext), varargs...)
}
