// Code generated by MockGen. DO NOT EDIT.
// Source: kafka_instances_service.go

// Package mock_instance is a generated GoMock package.
package mock_instance

import (
	context "context"
	model "github.com/netcracker/qubership-maas/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockKafkaInstanceService is a mock of KafkaInstanceService interface.
type MockKafkaInstanceService struct {
	ctrl     *gomock.Controller
	recorder *MockKafkaInstanceServiceMockRecorder
}

// MockKafkaInstanceServiceMockRecorder is the mock recorder for MockKafkaInstanceService.
type MockKafkaInstanceServiceMockRecorder struct {
	mock *MockKafkaInstanceService
}

// NewMockKafkaInstanceService creates a new mock instance.
func NewMockKafkaInstanceService(ctrl *gomock.Controller) *MockKafkaInstanceService {
	mock := &MockKafkaInstanceService{ctrl: ctrl}
	mock.recorder = &MockKafkaInstanceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKafkaInstanceService) EXPECT() *MockKafkaInstanceServiceMockRecorder {
	return m.recorder
}

// AddInstanceUpdateCallback mocks base method.
func (m *MockKafkaInstanceService) AddInstanceUpdateCallback(arg0 func(context.Context, *model.KafkaInstance)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddInstanceUpdateCallback", arg0)
}

// AddInstanceUpdateCallback indicates an expected call of AddInstanceUpdateCallback.
func (mr *MockKafkaInstanceServiceMockRecorder) AddInstanceUpdateCallback(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddInstanceUpdateCallback", reflect.TypeOf((*MockKafkaInstanceService)(nil).AddInstanceUpdateCallback), arg0)
}

// CheckHealth mocks base method.
func (m *MockKafkaInstanceService) CheckHealth(arg0 context.Context, arg1 *model.KafkaInstance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckHealth", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckHealth indicates an expected call of CheckHealth.
func (mr *MockKafkaInstanceServiceMockRecorder) CheckHealth(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckHealth", reflect.TypeOf((*MockKafkaInstanceService)(nil).CheckHealth), arg0, arg1)
}

// DeleteKafkaInstanceDesignatorByNamespace mocks base method.
func (m *MockKafkaInstanceService) DeleteKafkaInstanceDesignatorByNamespace(ctx context.Context, namespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteKafkaInstanceDesignatorByNamespace", ctx, namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteKafkaInstanceDesignatorByNamespace indicates an expected call of DeleteKafkaInstanceDesignatorByNamespace.
func (mr *MockKafkaInstanceServiceMockRecorder) DeleteKafkaInstanceDesignatorByNamespace(ctx, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteKafkaInstanceDesignatorByNamespace", reflect.TypeOf((*MockKafkaInstanceService)(nil).DeleteKafkaInstanceDesignatorByNamespace), ctx, namespace)
}

// GetById mocks base method.
func (m *MockKafkaInstanceService) GetById(ctx context.Context, id string) (*model.KafkaInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(*model.KafkaInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockKafkaInstanceServiceMockRecorder) GetById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockKafkaInstanceService)(nil).GetById), ctx, id)
}

// GetDefault mocks base method.
func (m *MockKafkaInstanceService) GetDefault(ctx context.Context) (*model.KafkaInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefault", ctx)
	ret0, _ := ret[0].(*model.KafkaInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDefault indicates an expected call of GetDefault.
func (mr *MockKafkaInstanceServiceMockRecorder) GetDefault(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefault", reflect.TypeOf((*MockKafkaInstanceService)(nil).GetDefault), ctx)
}

// GetKafkaInstanceDesignatorByNamespace mocks base method.
func (m *MockKafkaInstanceService) GetKafkaInstanceDesignatorByNamespace(ctx context.Context, namespace string) (*model.InstanceDesignatorKafka, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKafkaInstanceDesignatorByNamespace", ctx, namespace)
	ret0, _ := ret[0].(*model.InstanceDesignatorKafka)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKafkaInstanceDesignatorByNamespace indicates an expected call of GetKafkaInstanceDesignatorByNamespace.
func (mr *MockKafkaInstanceServiceMockRecorder) GetKafkaInstanceDesignatorByNamespace(ctx, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKafkaInstanceDesignatorByNamespace", reflect.TypeOf((*MockKafkaInstanceService)(nil).GetKafkaInstanceDesignatorByNamespace), ctx, namespace)
}

// GetKafkaInstances mocks base method.
func (m *MockKafkaInstanceService) GetKafkaInstances(ctx context.Context) (*[]model.KafkaInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKafkaInstances", ctx)
	ret0, _ := ret[0].(*[]model.KafkaInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKafkaInstances indicates an expected call of GetKafkaInstances.
func (mr *MockKafkaInstanceServiceMockRecorder) GetKafkaInstances(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKafkaInstances", reflect.TypeOf((*MockKafkaInstanceService)(nil).GetKafkaInstances), ctx)
}

// Register mocks base method.
func (m *MockKafkaInstanceService) Register(ctx context.Context, instance *model.KafkaInstance) (*model.KafkaInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, instance)
	ret0, _ := ret[0].(*model.KafkaInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockKafkaInstanceServiceMockRecorder) Register(ctx, instance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockKafkaInstanceService)(nil).Register), ctx, instance)
}

// SetDefault mocks base method.
func (m *MockKafkaInstanceService) SetDefault(ctx context.Context, instance *model.KafkaInstance) (*model.KafkaInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDefault", ctx, instance)
	ret0, _ := ret[0].(*model.KafkaInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetDefault indicates an expected call of SetDefault.
func (mr *MockKafkaInstanceServiceMockRecorder) SetDefault(ctx, instance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDefault", reflect.TypeOf((*MockKafkaInstanceService)(nil).SetDefault), ctx, instance)
}

// Unregister mocks base method.
func (m *MockKafkaInstanceService) Unregister(ctx context.Context, instanceId string) (*model.KafkaInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unregister", ctx, instanceId)
	ret0, _ := ret[0].(*model.KafkaInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unregister indicates an expected call of Unregister.
func (mr *MockKafkaInstanceServiceMockRecorder) Unregister(ctx, instanceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unregister", reflect.TypeOf((*MockKafkaInstanceService)(nil).Unregister), ctx, instanceId)
}

// Update mocks base method.
func (m *MockKafkaInstanceService) Update(ctx context.Context, instance *model.KafkaInstance) (*model.KafkaInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, instance)
	ret0, _ := ret[0].(*model.KafkaInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockKafkaInstanceServiceMockRecorder) Update(ctx, instance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockKafkaInstanceService)(nil).Update), ctx, instance)
}

// UpsertKafkaInstanceDesignator mocks base method.
func (m *MockKafkaInstanceService) UpsertKafkaInstanceDesignator(ctx context.Context, instanceDesignator model.InstanceDesignatorKafka, namespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertKafkaInstanceDesignator", ctx, instanceDesignator, namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertKafkaInstanceDesignator indicates an expected call of UpsertKafkaInstanceDesignator.
func (mr *MockKafkaInstanceServiceMockRecorder) UpsertKafkaInstanceDesignator(ctx, instanceDesignator, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertKafkaInstanceDesignator", reflect.TypeOf((*MockKafkaInstanceService)(nil).UpsertKafkaInstanceDesignator), ctx, instanceDesignator, namespace)
}
