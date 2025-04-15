// Code generated by MockGen. DO NOT EDIT.
// Source: bg_service_dao.go

// Package mock_bg_service is a generated GoMock package.
package mock_bg_service

import (
	context "context"
	model "github.com/netcracker/qubership-maas/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBgServiceDao is a mock of BgServiceDao interface.
type MockBgServiceDao struct {
	ctrl     *gomock.Controller
	recorder *MockBgServiceDaoMockRecorder
}

// MockBgServiceDaoMockRecorder is the mock recorder for MockBgServiceDao.
type MockBgServiceDaoMockRecorder struct {
	mock *MockBgServiceDao
}

// NewMockBgServiceDao creates a new mock instance.
func NewMockBgServiceDao(ctrl *gomock.Controller) *MockBgServiceDao {
	mock := &MockBgServiceDao{ctrl: ctrl}
	mock.recorder = &MockBgServiceDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBgServiceDao) EXPECT() *MockBgServiceDaoMockRecorder {
	return m.recorder
}

// DeleteCpMessagesByNamespace mocks base method.
func (m *MockBgServiceDao) DeleteCpMessagesByNamespace(ctx context.Context, namespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCpMessagesByNamespace", ctx, namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCpMessagesByNamespace indicates an expected call of DeleteCpMessagesByNamespace.
func (mr *MockBgServiceDaoMockRecorder) DeleteCpMessagesByNamespace(ctx, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCpMessagesByNamespace", reflect.TypeOf((*MockBgServiceDao)(nil).DeleteCpMessagesByNamespace), ctx, namespace)
}

// GetActiveVersionByNamespace mocks base method.
func (m *MockBgServiceDao) GetActiveVersionByNamespace(ctx context.Context, namespace string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveVersionByNamespace", ctx, namespace)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActiveVersionByNamespace indicates an expected call of GetActiveVersionByNamespace.
func (mr *MockBgServiceDaoMockRecorder) GetActiveVersionByNamespace(ctx, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveVersionByNamespace", reflect.TypeOf((*MockBgServiceDao)(nil).GetActiveVersionByNamespace), ctx, namespace)
}

// GetBgStatusByNamespace mocks base method.
func (m *MockBgServiceDao) GetBgStatusByNamespace(arg0 context.Context, arg1 string) (*model.BgStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBgStatusByNamespace", arg0, arg1)
	ret0, _ := ret[0].(*model.BgStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBgStatusByNamespace indicates an expected call of GetBgStatusByNamespace.
func (mr *MockBgServiceDaoMockRecorder) GetBgStatusByNamespace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBgStatusByNamespace", reflect.TypeOf((*MockBgServiceDao)(nil).GetBgStatusByNamespace), arg0, arg1)
}

// InsertBgStatus mocks base method.
func (m *MockBgServiceDao) InsertBgStatus(ctx context.Context, cpMessage *model.BgStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertBgStatus", ctx, cpMessage)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertBgStatus indicates an expected call of InsertBgStatus.
func (mr *MockBgServiceDaoMockRecorder) InsertBgStatus(ctx, cpMessage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertBgStatus", reflect.TypeOf((*MockBgServiceDao)(nil).InsertBgStatus), ctx, cpMessage)
}
