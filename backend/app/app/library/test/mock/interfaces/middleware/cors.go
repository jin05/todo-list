// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces/middleware/cors.go

// Package mock_middleware is a generated GoMock package.
package mock_middleware

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCORSMiddleware is a mock of CORSMiddleware interface.
type MockCORSMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MockCORSMiddlewareMockRecorder
}

// MockCORSMiddlewareMockRecorder is the mock recorder for MockCORSMiddleware.
type MockCORSMiddlewareMockRecorder struct {
	mock *MockCORSMiddleware
}

// NewMockCORSMiddleware creates a new mock instance.
func NewMockCORSMiddleware(ctrl *gomock.Controller) *MockCORSMiddleware {
	mock := &MockCORSMiddleware{ctrl: ctrl}
	mock.recorder = &MockCORSMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCORSMiddleware) EXPECT() *MockCORSMiddlewareMockRecorder {
	return m.recorder
}

// Handler mocks base method.
func (m *MockCORSMiddleware) Handler(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handler", w, r)
	ret0, _ := ret[0].(http.ResponseWriter)
	ret1, _ := ret[1].(*http.Request)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Handler indicates an expected call of Handler.
func (mr *MockCORSMiddlewareMockRecorder) Handler(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handler", reflect.TypeOf((*MockCORSMiddleware)(nil).Handler), w, r)
}
