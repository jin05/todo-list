// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces/api/todo.go

// Package mock_api is a generated GoMock package.
package mock_api

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTodoAPI is a mock of TodoAPI interface.
type MockTodoAPI struct {
	ctrl     *gomock.Controller
	recorder *MockTodoAPIMockRecorder
}

// MockTodoAPIMockRecorder is the mock recorder for MockTodoAPI.
type MockTodoAPIMockRecorder struct {
	mock *MockTodoAPI
}

// NewMockTodoAPI creates a new mock instance.
func NewMockTodoAPI(ctrl *gomock.Controller) *MockTodoAPI {
	mock := &MockTodoAPI{ctrl: ctrl}
	mock.recorder = &MockTodoAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoAPI) EXPECT() *MockTodoAPIMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTodoAPI) Create(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Create", w, r)
}

// Create indicates an expected call of Create.
func (mr *MockTodoAPIMockRecorder) Create(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTodoAPI)(nil).Create), w, r)
}

// Delete mocks base method.
func (m *MockTodoAPI) Delete(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", w, r)
}

// Delete indicates an expected call of Delete.
func (mr *MockTodoAPIMockRecorder) Delete(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTodoAPI)(nil).Delete), w, r)
}

// Get mocks base method.
func (m *MockTodoAPI) Get(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Get", w, r)
}

// Get indicates an expected call of Get.
func (mr *MockTodoAPIMockRecorder) Get(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTodoAPI)(nil).Get), w, r)
}

// Handler mocks base method.
func (m *MockTodoAPI) Handler(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Handler", w, r)
}

// Handler indicates an expected call of Handler.
func (mr *MockTodoAPIMockRecorder) Handler(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handler", reflect.TypeOf((*MockTodoAPI)(nil).Handler), w, r)
}

// List mocks base method.
func (m *MockTodoAPI) List(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "List", w, r)
}

// List indicates an expected call of List.
func (mr *MockTodoAPIMockRecorder) List(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTodoAPI)(nil).List), w, r)
}

// Update mocks base method.
func (m *MockTodoAPI) Update(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Update", w, r)
}

// Update indicates an expected call of Update.
func (mr *MockTodoAPIMockRecorder) Update(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTodoAPI)(nil).Update), w, r)
}