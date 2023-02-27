// Code generated by MockGen. DO NOT EDIT.
// Source: ./service.go

// Package services is a generated GoMock package.
package services

import (
	reflect "reflect"

	entity "github.com/DerryRenaldy/learnFiber/entity"
	forms "github.com/DerryRenaldy/learnFiber/forms"
	gomock "github.com/golang/mock/gomock"
	fasthttp "github.com/valyala/fasthttp"
)

// MockIService is a mock of IService interface.
type MockIService struct {
	ctrl     *gomock.Controller
	recorder *MockIServiceMockRecorder
}

// MockIServiceMockRecorder is the mock recorder for MockIService.
type MockIServiceMockRecorder struct {
	mock *MockIService
}

// NewMockIService creates a new mock instance.
func NewMockIService(ctrl *gomock.Controller) *MockIService {
	mock := &MockIService{ctrl: ctrl}
	mock.recorder = &MockIServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIService) EXPECT() *MockIServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIService) Create(ctx *fasthttp.RequestCtx, req forms.CreateRequest) (*entity.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, req)
	ret0, _ := ret[0].(*entity.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIServiceMockRecorder) Create(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIService)(nil).Create), ctx, req)
}