// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/store (interfaces: APIKeyLister)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	reflect "reflect"
)

// MockAPIKeyLister is a mock of APIKeyLister interface
type MockAPIKeyLister struct {
	ctrl     *gomock.Controller
	recorder *MockAPIKeyListerMockRecorder
}

// MockAPIKeyListerMockRecorder is the mock recorder for MockAPIKeyLister
type MockAPIKeyListerMockRecorder struct {
	mock *MockAPIKeyLister
}

// NewMockAPIKeyLister creates a new mock instance
func NewMockAPIKeyLister(ctrl *gomock.Controller) *MockAPIKeyLister {
	mock := &MockAPIKeyLister{ctrl: ctrl}
	mock.recorder = &MockAPIKeyListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAPIKeyLister) EXPECT() *MockAPIKeyListerMockRecorder {
	return m.recorder
}

// APIKeys mocks base method
func (m *MockAPIKeyLister) APIKeys(arg0 string, arg1 *mongodbatlas.ListOptions) ([]mongodbatlas.APIKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIKeys", arg0, arg1)
	ret0, _ := ret[0].([]mongodbatlas.APIKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// APIKeys indicates an expected call of APIKeys
func (mr *MockAPIKeyListerMockRecorder) APIKeys(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIKeys", reflect.TypeOf((*MockAPIKeyLister)(nil).APIKeys), arg0, arg1)
}