// Code generated by MockGen. DO NOT EDIT.
// Source: helper.go

// Package mock_helper is a generated GoMock package.
package mock_helper

import (
	context "context"
	model "github.com/netcracker/qubership-maas/model"
	reflect "reflect"

	sarama "github.com/IBM/sarama"
	gomock "github.com/golang/mock/gomock"
)

// MockHelper is a mock of Helper interface.
type MockHelper struct {
	ctrl     *gomock.Controller
	recorder *MockHelperMockRecorder
}

// MockHelperMockRecorder is the mock recorder for MockHelper.
type MockHelperMockRecorder struct {
	mock *MockHelper
}

// NewMockHelper creates a new mock instance.
func NewMockHelper(ctrl *gomock.Controller) *MockHelper {
	mock := &MockHelper{ctrl: ctrl}
	mock.recorder = &MockHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHelper) EXPECT() *MockHelperMockRecorder {
	return m.recorder
}

// BulkGetTopicSettings mocks base method.
func (m *MockHelper) BulkGetTopicSettings(ctx context.Context, topics []*model.TopicRegistrationRespDto) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkGetTopicSettings", ctx, topics)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkGetTopicSettings indicates an expected call of BulkGetTopicSettings.
func (mr *MockHelperMockRecorder) BulkGetTopicSettings(ctx, topics interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkGetTopicSettings", reflect.TypeOf((*MockHelper)(nil).BulkGetTopicSettings), ctx, topics)
}

// CheckHealth mocks base method.
func (m *MockHelper) CheckHealth(ctx context.Context, instance *model.KafkaInstance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckHealth", ctx, instance)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckHealth indicates an expected call of CheckHealth.
func (mr *MockHelperMockRecorder) CheckHealth(ctx, instance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckHealth", reflect.TypeOf((*MockHelper)(nil).CheckHealth), ctx, instance)
}

// CreateTopic mocks base method.
func (m *MockHelper) CreateTopic(ctx context.Context, topic *model.TopicRegistration) (*model.TopicRegistrationRespDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTopic", ctx, topic)
	ret0, _ := ret[0].(*model.TopicRegistrationRespDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTopic indicates an expected call of CreateTopic.
func (mr *MockHelperMockRecorder) CreateTopic(ctx, topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTopic", reflect.TypeOf((*MockHelper)(nil).CreateTopic), ctx, topic)
}

// DeleteTopic mocks base method.
func (m *MockHelper) DeleteTopic(ctx context.Context, topic *model.TopicRegistration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTopic", ctx, topic)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTopic indicates an expected call of DeleteTopic.
func (mr *MockHelperMockRecorder) DeleteTopic(ctx, topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTopic", reflect.TypeOf((*MockHelper)(nil).DeleteTopic), ctx, topic)
}

// DoesTopicExistOnKafka mocks base method.
func (m *MockHelper) DoesTopicExistOnKafka(ctx context.Context, instance *model.KafkaInstance, topicName string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoesTopicExistOnKafka", ctx, instance, topicName)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoesTopicExistOnKafka indicates an expected call of DoesTopicExistOnKafka.
func (mr *MockHelperMockRecorder) DoesTopicExistOnKafka(ctx, instance, topicName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoesTopicExistOnKafka", reflect.TypeOf((*MockHelper)(nil).DoesTopicExistOnKafka), ctx, instance, topicName)
}

// GetListTopics mocks base method.
func (m *MockHelper) GetListTopics(ctx context.Context, instance *model.KafkaInstance) (map[string]sarama.TopicDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListTopics", ctx, instance)
	ret0, _ := ret[0].(map[string]sarama.TopicDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListTopics indicates an expected call of GetListTopics.
func (mr *MockHelperMockRecorder) GetListTopics(ctx, instance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListTopics", reflect.TypeOf((*MockHelper)(nil).GetListTopics), ctx, instance)
}

// GetTopicSettings mocks base method.
func (m *MockHelper) GetTopicSettings(ctx context.Context, topic *model.TopicRegistrationRespDto) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTopicSettings", ctx, topic)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetTopicSettings indicates an expected call of GetTopicSettings.
func (mr *MockHelperMockRecorder) GetTopicSettings(ctx, topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopicSettings", reflect.TypeOf((*MockHelper)(nil).GetTopicSettings), ctx, topic)
}

// UpdateTopicSettings mocks base method.
func (m *MockHelper) UpdateTopicSettings(ctx context.Context, topic *model.TopicRegistrationRespDto) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTopicSettings", ctx, topic)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTopicSettings indicates an expected call of UpdateTopicSettings.
func (mr *MockHelperMockRecorder) UpdateTopicSettings(ctx, topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTopicSettings", reflect.TypeOf((*MockHelper)(nil).UpdateTopicSettings), ctx, topic)
}
