// Code generated by MockGen. DO NOT EDIT.
// Source: server.go

// Package mock_http is a generated GoMock package.
package mock_http

import (
	context "context"
	reflect "reflect"
	domain "siteAccess/internal/domain"

	gomock "github.com/golang/mock/gomock"
)

// Mockservice is a mock of service interface.
type Mockservice struct {
	ctrl     *gomock.Controller
	recorder *MockserviceMockRecorder
}

// MockserviceMockRecorder is the mock recorder for Mockservice.
type MockserviceMockRecorder struct {
	mock *Mockservice
}

// NewMockservice creates a new mock instance.
func NewMockservice(ctrl *gomock.Controller) *Mockservice {
	mock := &Mockservice{ctrl: ctrl}
	mock.recorder = &MockserviceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockservice) EXPECT() *MockserviceMockRecorder {
	return m.recorder
}

// GetMaxTime mocks base method.
func (m *Mockservice) GetMaxTime(ctx context.Context) (*domain.Site, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxTime", ctx)
	ret0, _ := ret[0].(*domain.Site)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMaxTime indicates an expected call of GetMaxTime.
func (mr *MockserviceMockRecorder) GetMaxTime(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxTime", reflect.TypeOf((*Mockservice)(nil).GetMaxTime), ctx)
}

// GetMinTime mocks base method.
func (m *Mockservice) GetMinTime(ctx context.Context) (*domain.Site, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMinTime", ctx)
	ret0, _ := ret[0].(*domain.Site)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMinTime indicates an expected call of GetMinTime.
func (mr *MockserviceMockRecorder) GetMinTime(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMinTime", reflect.TypeOf((*Mockservice)(nil).GetMinTime), ctx)
}

// GetTime mocks base method.
func (m *Mockservice) GetTime(ctx context.Context, site string) (*domain.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTime", ctx, site)
	ret0, _ := ret[0].(*domain.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTime indicates an expected call of GetTime.
func (mr *MockserviceMockRecorder) GetTime(ctx, site interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTime", reflect.TypeOf((*Mockservice)(nil).GetTime), ctx, site)
}