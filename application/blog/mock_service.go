// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package blog is a generated GoMock package.
package blog

import (
	gomock "github.com/golang/mock/gomock"
	domain "github.com/keitam913/airlog/domain"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetByHTMLFilename mocks base method
func (m *MockService) GetByHTMLFilename(filename string) (*domain.Post, error) {
	ret := m.ctrl.Call(m, "GetByHTMLFilename", filename)
	ret0, _ := ret[0].(*domain.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByHTMLFilename indicates an expected call of GetByHTMLFilename
func (mr *MockServiceMockRecorder) GetByHTMLFilename(filename interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByHTMLFilename", reflect.TypeOf((*MockService)(nil).GetByHTMLFilename), filename)
}

// Recent mocks base method
func (m *MockService) Recent() ([]*domain.Post, error) {
	ret := m.ctrl.Call(m, "Recent")
	ret0, _ := ret[0].([]*domain.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recent indicates an expected call of Recent
func (mr *MockServiceMockRecorder) Recent() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recent", reflect.TypeOf((*MockService)(nil).Recent))
}

// RegisterPost mocks base method
func (m *MockService) RegisterPost(filename, content string) error {
	ret := m.ctrl.Call(m, "RegisterPost", filename, content)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterPost indicates an expected call of RegisterPost
func (mr *MockServiceMockRecorder) RegisterPost(filename, content interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterPost", reflect.TypeOf((*MockService)(nil).RegisterPost), filename, content)
}