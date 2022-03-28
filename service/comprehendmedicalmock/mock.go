// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/comprehendmedical/comprehendmedicaliface (interfaces: ComprehendMedicalAPI)

// Package comprehendmedicalmock is a generated GoMock package.
package comprehendmedicalmock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	comprehendmedical "github.com/aws/aws-sdk-go/service/comprehendmedical"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockComprehendMedicalAPI is a mock of ComprehendMedicalAPI interface
type MockComprehendMedicalAPI struct {
	ctrl     *gomock.Controller
	recorder *MockComprehendMedicalAPIMockRecorder
}

// MockComprehendMedicalAPIMockRecorder is the mock recorder for MockComprehendMedicalAPI
type MockComprehendMedicalAPIMockRecorder struct {
	mock *MockComprehendMedicalAPI
}

// NewMockComprehendMedicalAPI creates a new mock instance
func NewMockComprehendMedicalAPI(ctrl *gomock.Controller) *MockComprehendMedicalAPI {
	mock := &MockComprehendMedicalAPI{ctrl: ctrl}
	mock.recorder = &MockComprehendMedicalAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockComprehendMedicalAPI) EXPECT() *MockComprehendMedicalAPIMockRecorder {
	return m.recorder
}

// DescribeEntitiesDetectionV2Job mocks base method
func (m *MockComprehendMedicalAPI) DescribeEntitiesDetectionV2Job(arg0 *comprehendmedical.DescribeEntitiesDetectionV2JobInput) (*comprehendmedical.DescribeEntitiesDetectionV2JobOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeEntitiesDetectionV2Job", arg0)
	ret0, _ := ret[0].(*comprehendmedical.DescribeEntitiesDetectionV2JobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEntitiesDetectionV2Job indicates an expected call of DescribeEntitiesDetectionV2Job
func (mr *MockComprehendMedicalAPIMockRecorder) DescribeEntitiesDetectionV2Job(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEntitiesDetectionV2Job", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).DescribeEntitiesDetectionV2Job), arg0)
}

// DescribeEntitiesDetectionV2JobRequest mocks base method
func (m *MockComprehendMedicalAPI) DescribeEntitiesDetectionV2JobRequest(arg0 *comprehendmedical.DescribeEntitiesDetectionV2JobInput) (*request.Request, *comprehendmedical.DescribeEntitiesDetectionV2JobOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeEntitiesDetectionV2JobRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*comprehendmedical.DescribeEntitiesDetectionV2JobOutput)
	return ret0, ret1
}

// DescribeEntitiesDetectionV2JobRequest indicates an expected call of DescribeEntitiesDetectionV2JobRequest
func (mr *MockComprehendMedicalAPIMockRecorder) DescribeEntitiesDetectionV2JobRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEntitiesDetectionV2JobRequest", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).DescribeEntitiesDetectionV2JobRequest), arg0)
}

// DescribeEntitiesDetectionV2JobWithContext mocks base method
func (m *MockComprehendMedicalAPI) DescribeEntitiesDetectionV2JobWithContext(arg0 context.Context, arg1 *comprehendmedical.DescribeEntitiesDetectionV2JobInput, arg2 ...request.Option) (*comprehendmedical.DescribeEntitiesDetectionV2JobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeEntitiesDetectionV2JobWithContext", varargs...)
	ret0, _ := ret[0].(*comprehendmedical.DescribeEntitiesDetectionV2JobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEntitiesDetectionV2JobWithContext indicates an expected call of DescribeEntitiesDetectionV2JobWithContext
func (mr *MockComprehendMedicalAPIMockRecorder) DescribeEntitiesDetectionV2JobWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEntitiesDetectionV2JobWithContext", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).DescribeEntitiesDetectionV2JobWithContext), varargs...)
}

// DescribePHIDetectionJob mocks base method
func (m *MockComprehendMedicalAPI) DescribePHIDetectionJob(arg0 *comprehendmedical.DescribePHIDetectionJobInput) (*comprehendmedical.DescribePHIDetectionJobOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribePHIDetectionJob", arg0)
	ret0, _ := ret[0].(*comprehendmedical.DescribePHIDetectionJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribePHIDetectionJob indicates an expected call of DescribePHIDetectionJob
func (mr *MockComprehendMedicalAPIMockRecorder) DescribePHIDetectionJob(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribePHIDetectionJob", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).DescribePHIDetectionJob), arg0)
}

// DescribePHIDetectionJobRequest mocks base method
func (m *MockComprehendMedicalAPI) DescribePHIDetectionJobRequest(arg0 *comprehendmedical.DescribePHIDetectionJobInput) (*request.Request, *comprehendmedical.DescribePHIDetectionJobOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribePHIDetectionJobRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*comprehendmedical.DescribePHIDetectionJobOutput)
	return ret0, ret1
}

// DescribePHIDetectionJobRequest indicates an expected call of DescribePHIDetectionJobRequest
func (mr *MockComprehendMedicalAPIMockRecorder) DescribePHIDetectionJobRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribePHIDetectionJobRequest", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).DescribePHIDetectionJobRequest), arg0)
}

// DescribePHIDetectionJobWithContext mocks base method
func (m *MockComprehendMedicalAPI) DescribePHIDetectionJobWithContext(arg0 context.Context, arg1 *comprehendmedical.DescribePHIDetectionJobInput, arg2 ...request.Option) (*comprehendmedical.DescribePHIDetectionJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribePHIDetectionJobWithContext", varargs...)
	ret0, _ := ret[0].(*comprehendmedical.DescribePHIDetectionJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribePHIDetectionJobWithContext indicates an expected call of DescribePHIDetectionJobWithContext
func (mr *MockComprehendMedicalAPIMockRecorder) DescribePHIDetectionJobWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribePHIDetectionJobWithContext", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).DescribePHIDetectionJobWithContext), varargs...)
}

// DetectEntities mocks base method
func (m *MockComprehendMedicalAPI) DetectEntities(arg0 *comprehendmedical.DetectEntitiesInput) (*comprehendmedical.DetectEntitiesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetectEntities", arg0)
	ret0, _ := ret[0].(*comprehendmedical.DetectEntitiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetectEntities indicates an expected call of DetectEntities
func (mr *MockComprehendMedicalAPIMockRecorder) DetectEntities(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetectEntities", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).DetectEntities), arg0)
}

// DetectEntitiesRequest mocks base method
func (m *MockComprehendMedicalAPI) DetectEntitiesRequest(arg0 *comprehendmedical.DetectEntitiesInput) (*request.Request, *comprehendmedical.DetectEntitiesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetectEntitiesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*comprehendmedical.DetectEntitiesOutput)
	return ret0, ret1
}

// DetectEntitiesRequest indicates an expected call of DetectEntitiesRequest
func (mr *MockComprehendMedicalAPIMockRecorder) DetectEntitiesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetectEntitiesRequest", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).DetectEntitiesRequest), arg0)
}

// DetectEntitiesV2 mocks base method
func (m *MockComprehendMedicalAPI) DetectEntitiesV2(arg0 *comprehendmedical.DetectEntitiesV2Input) (*comprehendmedical.DetectEntitiesV2Output, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetectEntitiesV2", arg0)
	ret0, _ := ret[0].(*comprehendmedical.DetectEntitiesV2Output)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetectEntitiesV2 indicates an expected call of DetectEntitiesV2
func (mr *MockComprehendMedicalAPIMockRecorder) DetectEntitiesV2(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetectEntitiesV2", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).DetectEntitiesV2), arg0)
}

// DetectEntitiesV2Request mocks base method
func (m *MockComprehendMedicalAPI) DetectEntitiesV2Request(arg0 *comprehendmedical.DetectEntitiesV2Input) (*request.Request, *comprehendmedical.DetectEntitiesV2Output) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetectEntitiesV2Request", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*comprehendmedical.DetectEntitiesV2Output)
	return ret0, ret1
}

// DetectEntitiesV2Request indicates an expected call of DetectEntitiesV2Request
func (mr *MockComprehendMedicalAPIMockRecorder) DetectEntitiesV2Request(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetectEntitiesV2Request", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).DetectEntitiesV2Request), arg0)
}

// DetectEntitiesV2WithContext mocks base method
func (m *MockComprehendMedicalAPI) DetectEntitiesV2WithContext(arg0 context.Context, arg1 *comprehendmedical.DetectEntitiesV2Input, arg2 ...request.Option) (*comprehendmedical.DetectEntitiesV2Output, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DetectEntitiesV2WithContext", varargs...)
	ret0, _ := ret[0].(*comprehendmedical.DetectEntitiesV2Output)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetectEntitiesV2WithContext indicates an expected call of DetectEntitiesV2WithContext
func (mr *MockComprehendMedicalAPIMockRecorder) DetectEntitiesV2WithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetectEntitiesV2WithContext", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).DetectEntitiesV2WithContext), varargs...)
}

// DetectEntitiesWithContext mocks base method
func (m *MockComprehendMedicalAPI) DetectEntitiesWithContext(arg0 context.Context, arg1 *comprehendmedical.DetectEntitiesInput, arg2 ...request.Option) (*comprehendmedical.DetectEntitiesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DetectEntitiesWithContext", varargs...)
	ret0, _ := ret[0].(*comprehendmedical.DetectEntitiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetectEntitiesWithContext indicates an expected call of DetectEntitiesWithContext
func (mr *MockComprehendMedicalAPIMockRecorder) DetectEntitiesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetectEntitiesWithContext", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).DetectEntitiesWithContext), varargs...)
}

// DetectPHI mocks base method
func (m *MockComprehendMedicalAPI) DetectPHI(arg0 *comprehendmedical.DetectPHIInput) (*comprehendmedical.DetectPHIOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetectPHI", arg0)
	ret0, _ := ret[0].(*comprehendmedical.DetectPHIOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetectPHI indicates an expected call of DetectPHI
func (mr *MockComprehendMedicalAPIMockRecorder) DetectPHI(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetectPHI", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).DetectPHI), arg0)
}

// DetectPHIRequest mocks base method
func (m *MockComprehendMedicalAPI) DetectPHIRequest(arg0 *comprehendmedical.DetectPHIInput) (*request.Request, *comprehendmedical.DetectPHIOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetectPHIRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*comprehendmedical.DetectPHIOutput)
	return ret0, ret1
}

// DetectPHIRequest indicates an expected call of DetectPHIRequest
func (mr *MockComprehendMedicalAPIMockRecorder) DetectPHIRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetectPHIRequest", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).DetectPHIRequest), arg0)
}

// DetectPHIWithContext mocks base method
func (m *MockComprehendMedicalAPI) DetectPHIWithContext(arg0 context.Context, arg1 *comprehendmedical.DetectPHIInput, arg2 ...request.Option) (*comprehendmedical.DetectPHIOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DetectPHIWithContext", varargs...)
	ret0, _ := ret[0].(*comprehendmedical.DetectPHIOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetectPHIWithContext indicates an expected call of DetectPHIWithContext
func (mr *MockComprehendMedicalAPIMockRecorder) DetectPHIWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetectPHIWithContext", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).DetectPHIWithContext), varargs...)
}

// InferICD10CM mocks base method
func (m *MockComprehendMedicalAPI) InferICD10CM(arg0 *comprehendmedical.InferICD10CMInput) (*comprehendmedical.InferICD10CMOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InferICD10CM", arg0)
	ret0, _ := ret[0].(*comprehendmedical.InferICD10CMOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InferICD10CM indicates an expected call of InferICD10CM
func (mr *MockComprehendMedicalAPIMockRecorder) InferICD10CM(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InferICD10CM", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).InferICD10CM), arg0)
}

// InferICD10CMRequest mocks base method
func (m *MockComprehendMedicalAPI) InferICD10CMRequest(arg0 *comprehendmedical.InferICD10CMInput) (*request.Request, *comprehendmedical.InferICD10CMOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InferICD10CMRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*comprehendmedical.InferICD10CMOutput)
	return ret0, ret1
}

// InferICD10CMRequest indicates an expected call of InferICD10CMRequest
func (mr *MockComprehendMedicalAPIMockRecorder) InferICD10CMRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InferICD10CMRequest", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).InferICD10CMRequest), arg0)
}

// InferICD10CMWithContext mocks base method
func (m *MockComprehendMedicalAPI) InferICD10CMWithContext(arg0 context.Context, arg1 *comprehendmedical.InferICD10CMInput, arg2 ...request.Option) (*comprehendmedical.InferICD10CMOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InferICD10CMWithContext", varargs...)
	ret0, _ := ret[0].(*comprehendmedical.InferICD10CMOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InferICD10CMWithContext indicates an expected call of InferICD10CMWithContext
func (mr *MockComprehendMedicalAPIMockRecorder) InferICD10CMWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InferICD10CMWithContext", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).InferICD10CMWithContext), varargs...)
}

// InferRxNorm mocks base method
func (m *MockComprehendMedicalAPI) InferRxNorm(arg0 *comprehendmedical.InferRxNormInput) (*comprehendmedical.InferRxNormOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InferRxNorm", arg0)
	ret0, _ := ret[0].(*comprehendmedical.InferRxNormOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InferRxNorm indicates an expected call of InferRxNorm
func (mr *MockComprehendMedicalAPIMockRecorder) InferRxNorm(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InferRxNorm", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).InferRxNorm), arg0)
}

// InferRxNormRequest mocks base method
func (m *MockComprehendMedicalAPI) InferRxNormRequest(arg0 *comprehendmedical.InferRxNormInput) (*request.Request, *comprehendmedical.InferRxNormOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InferRxNormRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*comprehendmedical.InferRxNormOutput)
	return ret0, ret1
}

// InferRxNormRequest indicates an expected call of InferRxNormRequest
func (mr *MockComprehendMedicalAPIMockRecorder) InferRxNormRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InferRxNormRequest", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).InferRxNormRequest), arg0)
}

// InferRxNormWithContext mocks base method
func (m *MockComprehendMedicalAPI) InferRxNormWithContext(arg0 context.Context, arg1 *comprehendmedical.InferRxNormInput, arg2 ...request.Option) (*comprehendmedical.InferRxNormOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InferRxNormWithContext", varargs...)
	ret0, _ := ret[0].(*comprehendmedical.InferRxNormOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InferRxNormWithContext indicates an expected call of InferRxNormWithContext
func (mr *MockComprehendMedicalAPIMockRecorder) InferRxNormWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InferRxNormWithContext", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).InferRxNormWithContext), varargs...)
}

// ListEntitiesDetectionV2Jobs mocks base method
func (m *MockComprehendMedicalAPI) ListEntitiesDetectionV2Jobs(arg0 *comprehendmedical.ListEntitiesDetectionV2JobsInput) (*comprehendmedical.ListEntitiesDetectionV2JobsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEntitiesDetectionV2Jobs", arg0)
	ret0, _ := ret[0].(*comprehendmedical.ListEntitiesDetectionV2JobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEntitiesDetectionV2Jobs indicates an expected call of ListEntitiesDetectionV2Jobs
func (mr *MockComprehendMedicalAPIMockRecorder) ListEntitiesDetectionV2Jobs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntitiesDetectionV2Jobs", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).ListEntitiesDetectionV2Jobs), arg0)
}

// ListEntitiesDetectionV2JobsRequest mocks base method
func (m *MockComprehendMedicalAPI) ListEntitiesDetectionV2JobsRequest(arg0 *comprehendmedical.ListEntitiesDetectionV2JobsInput) (*request.Request, *comprehendmedical.ListEntitiesDetectionV2JobsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEntitiesDetectionV2JobsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*comprehendmedical.ListEntitiesDetectionV2JobsOutput)
	return ret0, ret1
}

// ListEntitiesDetectionV2JobsRequest indicates an expected call of ListEntitiesDetectionV2JobsRequest
func (mr *MockComprehendMedicalAPIMockRecorder) ListEntitiesDetectionV2JobsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntitiesDetectionV2JobsRequest", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).ListEntitiesDetectionV2JobsRequest), arg0)
}

// ListEntitiesDetectionV2JobsWithContext mocks base method
func (m *MockComprehendMedicalAPI) ListEntitiesDetectionV2JobsWithContext(arg0 context.Context, arg1 *comprehendmedical.ListEntitiesDetectionV2JobsInput, arg2 ...request.Option) (*comprehendmedical.ListEntitiesDetectionV2JobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEntitiesDetectionV2JobsWithContext", varargs...)
	ret0, _ := ret[0].(*comprehendmedical.ListEntitiesDetectionV2JobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEntitiesDetectionV2JobsWithContext indicates an expected call of ListEntitiesDetectionV2JobsWithContext
func (mr *MockComprehendMedicalAPIMockRecorder) ListEntitiesDetectionV2JobsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntitiesDetectionV2JobsWithContext", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).ListEntitiesDetectionV2JobsWithContext), varargs...)
}

// ListPHIDetectionJobs mocks base method
func (m *MockComprehendMedicalAPI) ListPHIDetectionJobs(arg0 *comprehendmedical.ListPHIDetectionJobsInput) (*comprehendmedical.ListPHIDetectionJobsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPHIDetectionJobs", arg0)
	ret0, _ := ret[0].(*comprehendmedical.ListPHIDetectionJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPHIDetectionJobs indicates an expected call of ListPHIDetectionJobs
func (mr *MockComprehendMedicalAPIMockRecorder) ListPHIDetectionJobs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPHIDetectionJobs", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).ListPHIDetectionJobs), arg0)
}

// ListPHIDetectionJobsRequest mocks base method
func (m *MockComprehendMedicalAPI) ListPHIDetectionJobsRequest(arg0 *comprehendmedical.ListPHIDetectionJobsInput) (*request.Request, *comprehendmedical.ListPHIDetectionJobsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPHIDetectionJobsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*comprehendmedical.ListPHIDetectionJobsOutput)
	return ret0, ret1
}

// ListPHIDetectionJobsRequest indicates an expected call of ListPHIDetectionJobsRequest
func (mr *MockComprehendMedicalAPIMockRecorder) ListPHIDetectionJobsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPHIDetectionJobsRequest", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).ListPHIDetectionJobsRequest), arg0)
}

// ListPHIDetectionJobsWithContext mocks base method
func (m *MockComprehendMedicalAPI) ListPHIDetectionJobsWithContext(arg0 context.Context, arg1 *comprehendmedical.ListPHIDetectionJobsInput, arg2 ...request.Option) (*comprehendmedical.ListPHIDetectionJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPHIDetectionJobsWithContext", varargs...)
	ret0, _ := ret[0].(*comprehendmedical.ListPHIDetectionJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPHIDetectionJobsWithContext indicates an expected call of ListPHIDetectionJobsWithContext
func (mr *MockComprehendMedicalAPIMockRecorder) ListPHIDetectionJobsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPHIDetectionJobsWithContext", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).ListPHIDetectionJobsWithContext), varargs...)
}

// StartEntitiesDetectionV2Job mocks base method
func (m *MockComprehendMedicalAPI) StartEntitiesDetectionV2Job(arg0 *comprehendmedical.StartEntitiesDetectionV2JobInput) (*comprehendmedical.StartEntitiesDetectionV2JobOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartEntitiesDetectionV2Job", arg0)
	ret0, _ := ret[0].(*comprehendmedical.StartEntitiesDetectionV2JobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartEntitiesDetectionV2Job indicates an expected call of StartEntitiesDetectionV2Job
func (mr *MockComprehendMedicalAPIMockRecorder) StartEntitiesDetectionV2Job(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartEntitiesDetectionV2Job", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).StartEntitiesDetectionV2Job), arg0)
}

// StartEntitiesDetectionV2JobRequest mocks base method
func (m *MockComprehendMedicalAPI) StartEntitiesDetectionV2JobRequest(arg0 *comprehendmedical.StartEntitiesDetectionV2JobInput) (*request.Request, *comprehendmedical.StartEntitiesDetectionV2JobOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartEntitiesDetectionV2JobRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*comprehendmedical.StartEntitiesDetectionV2JobOutput)
	return ret0, ret1
}

// StartEntitiesDetectionV2JobRequest indicates an expected call of StartEntitiesDetectionV2JobRequest
func (mr *MockComprehendMedicalAPIMockRecorder) StartEntitiesDetectionV2JobRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartEntitiesDetectionV2JobRequest", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).StartEntitiesDetectionV2JobRequest), arg0)
}

// StartEntitiesDetectionV2JobWithContext mocks base method
func (m *MockComprehendMedicalAPI) StartEntitiesDetectionV2JobWithContext(arg0 context.Context, arg1 *comprehendmedical.StartEntitiesDetectionV2JobInput, arg2 ...request.Option) (*comprehendmedical.StartEntitiesDetectionV2JobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StartEntitiesDetectionV2JobWithContext", varargs...)
	ret0, _ := ret[0].(*comprehendmedical.StartEntitiesDetectionV2JobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartEntitiesDetectionV2JobWithContext indicates an expected call of StartEntitiesDetectionV2JobWithContext
func (mr *MockComprehendMedicalAPIMockRecorder) StartEntitiesDetectionV2JobWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartEntitiesDetectionV2JobWithContext", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).StartEntitiesDetectionV2JobWithContext), varargs...)
}

// StartPHIDetectionJob mocks base method
func (m *MockComprehendMedicalAPI) StartPHIDetectionJob(arg0 *comprehendmedical.StartPHIDetectionJobInput) (*comprehendmedical.StartPHIDetectionJobOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartPHIDetectionJob", arg0)
	ret0, _ := ret[0].(*comprehendmedical.StartPHIDetectionJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartPHIDetectionJob indicates an expected call of StartPHIDetectionJob
func (mr *MockComprehendMedicalAPIMockRecorder) StartPHIDetectionJob(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartPHIDetectionJob", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).StartPHIDetectionJob), arg0)
}

// StartPHIDetectionJobRequest mocks base method
func (m *MockComprehendMedicalAPI) StartPHIDetectionJobRequest(arg0 *comprehendmedical.StartPHIDetectionJobInput) (*request.Request, *comprehendmedical.StartPHIDetectionJobOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartPHIDetectionJobRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*comprehendmedical.StartPHIDetectionJobOutput)
	return ret0, ret1
}

// StartPHIDetectionJobRequest indicates an expected call of StartPHIDetectionJobRequest
func (mr *MockComprehendMedicalAPIMockRecorder) StartPHIDetectionJobRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartPHIDetectionJobRequest", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).StartPHIDetectionJobRequest), arg0)
}

// StartPHIDetectionJobWithContext mocks base method
func (m *MockComprehendMedicalAPI) StartPHIDetectionJobWithContext(arg0 context.Context, arg1 *comprehendmedical.StartPHIDetectionJobInput, arg2 ...request.Option) (*comprehendmedical.StartPHIDetectionJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StartPHIDetectionJobWithContext", varargs...)
	ret0, _ := ret[0].(*comprehendmedical.StartPHIDetectionJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartPHIDetectionJobWithContext indicates an expected call of StartPHIDetectionJobWithContext
func (mr *MockComprehendMedicalAPIMockRecorder) StartPHIDetectionJobWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartPHIDetectionJobWithContext", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).StartPHIDetectionJobWithContext), varargs...)
}

// StopEntitiesDetectionV2Job mocks base method
func (m *MockComprehendMedicalAPI) StopEntitiesDetectionV2Job(arg0 *comprehendmedical.StopEntitiesDetectionV2JobInput) (*comprehendmedical.StopEntitiesDetectionV2JobOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopEntitiesDetectionV2Job", arg0)
	ret0, _ := ret[0].(*comprehendmedical.StopEntitiesDetectionV2JobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopEntitiesDetectionV2Job indicates an expected call of StopEntitiesDetectionV2Job
func (mr *MockComprehendMedicalAPIMockRecorder) StopEntitiesDetectionV2Job(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopEntitiesDetectionV2Job", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).StopEntitiesDetectionV2Job), arg0)
}

// StopEntitiesDetectionV2JobRequest mocks base method
func (m *MockComprehendMedicalAPI) StopEntitiesDetectionV2JobRequest(arg0 *comprehendmedical.StopEntitiesDetectionV2JobInput) (*request.Request, *comprehendmedical.StopEntitiesDetectionV2JobOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopEntitiesDetectionV2JobRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*comprehendmedical.StopEntitiesDetectionV2JobOutput)
	return ret0, ret1
}

// StopEntitiesDetectionV2JobRequest indicates an expected call of StopEntitiesDetectionV2JobRequest
func (mr *MockComprehendMedicalAPIMockRecorder) StopEntitiesDetectionV2JobRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopEntitiesDetectionV2JobRequest", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).StopEntitiesDetectionV2JobRequest), arg0)
}

// StopEntitiesDetectionV2JobWithContext mocks base method
func (m *MockComprehendMedicalAPI) StopEntitiesDetectionV2JobWithContext(arg0 context.Context, arg1 *comprehendmedical.StopEntitiesDetectionV2JobInput, arg2 ...request.Option) (*comprehendmedical.StopEntitiesDetectionV2JobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StopEntitiesDetectionV2JobWithContext", varargs...)
	ret0, _ := ret[0].(*comprehendmedical.StopEntitiesDetectionV2JobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopEntitiesDetectionV2JobWithContext indicates an expected call of StopEntitiesDetectionV2JobWithContext
func (mr *MockComprehendMedicalAPIMockRecorder) StopEntitiesDetectionV2JobWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopEntitiesDetectionV2JobWithContext", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).StopEntitiesDetectionV2JobWithContext), varargs...)
}

// StopPHIDetectionJob mocks base method
func (m *MockComprehendMedicalAPI) StopPHIDetectionJob(arg0 *comprehendmedical.StopPHIDetectionJobInput) (*comprehendmedical.StopPHIDetectionJobOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopPHIDetectionJob", arg0)
	ret0, _ := ret[0].(*comprehendmedical.StopPHIDetectionJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopPHIDetectionJob indicates an expected call of StopPHIDetectionJob
func (mr *MockComprehendMedicalAPIMockRecorder) StopPHIDetectionJob(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopPHIDetectionJob", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).StopPHIDetectionJob), arg0)
}

// StopPHIDetectionJobRequest mocks base method
func (m *MockComprehendMedicalAPI) StopPHIDetectionJobRequest(arg0 *comprehendmedical.StopPHIDetectionJobInput) (*request.Request, *comprehendmedical.StopPHIDetectionJobOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopPHIDetectionJobRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*comprehendmedical.StopPHIDetectionJobOutput)
	return ret0, ret1
}

// StopPHIDetectionJobRequest indicates an expected call of StopPHIDetectionJobRequest
func (mr *MockComprehendMedicalAPIMockRecorder) StopPHIDetectionJobRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopPHIDetectionJobRequest", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).StopPHIDetectionJobRequest), arg0)
}

// StopPHIDetectionJobWithContext mocks base method
func (m *MockComprehendMedicalAPI) StopPHIDetectionJobWithContext(arg0 context.Context, arg1 *comprehendmedical.StopPHIDetectionJobInput, arg2 ...request.Option) (*comprehendmedical.StopPHIDetectionJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StopPHIDetectionJobWithContext", varargs...)
	ret0, _ := ret[0].(*comprehendmedical.StopPHIDetectionJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopPHIDetectionJobWithContext indicates an expected call of StopPHIDetectionJobWithContext
func (mr *MockComprehendMedicalAPIMockRecorder) StopPHIDetectionJobWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopPHIDetectionJobWithContext", reflect.TypeOf((*MockComprehendMedicalAPI)(nil).StopPHIDetectionJobWithContext), varargs...)
}
