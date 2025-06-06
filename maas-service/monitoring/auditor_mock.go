// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package monitoring is a generated GoMock package.
package monitoring

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuditor is a mock of Auditor interface.
type MockAuditor struct {
	ctrl     *gomock.Controller
	recorder *MockAuditorMockRecorder
}

// MockAuditorMockRecorder is the mock recorder for MockAuditor.
type MockAuditorMockRecorder struct {
	mock *MockAuditor
}

// NewMockAuditor creates a new mock instance.
func NewMockAuditor(ctrl *gomock.Controller) *MockAuditor {
	mock := &MockAuditor{ctrl: ctrl}
	mock.recorder = &MockAuditorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuditor) EXPECT() *MockAuditorMockRecorder {
	return m.recorder
}

// AddEntityRequestStat mocks base method.
func (m *MockAuditor) AddEntityRequestStat(ctx context.Context, entityType EntityType, name, instance string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddEntityRequestStat", ctx, entityType, name, instance)
}

// AddEntityRequestStat indicates an expected call of AddEntityRequestStat.
func (mr *MockAuditorMockRecorder) AddEntityRequestStat(ctx, entityType, name, instance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEntityRequestStat", reflect.TypeOf((*MockAuditor)(nil).AddEntityRequestStat), ctx, entityType, name, instance)
}

// CleanupData mocks base method.
func (m *MockAuditor) CleanupData(ctx context.Context, namespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanupData", ctx, namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanupData indicates an expected call of CleanupData.
func (mr *MockAuditorMockRecorder) CleanupData(ctx, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanupData", reflect.TypeOf((*MockAuditor)(nil).CleanupData), ctx, namespace)
}

// GetAllEntityRequestsStat mocks base method.
func (m *MockAuditor) GetAllEntityRequestsStat(ctx context.Context) (*[]EntityRequestStatDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllEntityRequestsStat", ctx)
	ret0, _ := ret[0].(*[]EntityRequestStatDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllEntityRequestsStat indicates an expected call of GetAllEntityRequestsStat.
func (mr *MockAuditorMockRecorder) GetAllEntityRequestsStat(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllEntityRequestsStat", reflect.TypeOf((*MockAuditor)(nil).GetAllEntityRequestsStat), ctx)
}

// GetKafkaMonitoringEntities mocks base method.
func (m *MockAuditor) GetKafkaMonitoringEntities(ctx context.Context) (*[]KafkaStatLine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKafkaMonitoringEntities", ctx)
	ret0, _ := ret[0].(*[]KafkaStatLine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKafkaMonitoringEntities indicates an expected call of GetKafkaMonitoringEntities.
func (mr *MockAuditorMockRecorder) GetKafkaMonitoringEntities(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKafkaMonitoringEntities", reflect.TypeOf((*MockAuditor)(nil).GetKafkaMonitoringEntities), ctx)
}

// GetRabbitMonitoringEntities mocks base method.
func (m *MockAuditor) GetRabbitMonitoringEntities(ctx context.Context) (*[]VhostStatLine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRabbitMonitoringEntities", ctx)
	ret0, _ := ret[0].(*[]VhostStatLine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRabbitMonitoringEntities indicates an expected call of GetRabbitMonitoringEntities.
func (mr *MockAuditorMockRecorder) GetRabbitMonitoringEntities(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRabbitMonitoringEntities", reflect.TypeOf((*MockAuditor)(nil).GetRabbitMonitoringEntities), ctx)
}
