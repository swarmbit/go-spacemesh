// Code generated by MockGen. DO NOT EDIT.
// Source: ./interface.go
//
// Generated by this command:
//
//	mockgen -typed -package=sql -destination=./mocks.go -source=./interface.go
//
// Package sql is a generated GoMock package.
package sql

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockMigration is a mock of Migration interface.
type MockMigration struct {
	ctrl     *gomock.Controller
	recorder *MockMigrationMockRecorder
}

// MockMigrationMockRecorder is the mock recorder for MockMigration.
type MockMigrationMockRecorder struct {
	mock *MockMigration
}

// NewMockMigration creates a new mock instance.
func NewMockMigration(ctrl *gomock.Controller) *MockMigration {
	mock := &MockMigration{ctrl: ctrl}
	mock.recorder = &MockMigrationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMigration) EXPECT() *MockMigrationMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockMigration) Apply(db Executor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", db)
	ret0, _ := ret[0].(error)
	return ret0
}

// Apply indicates an expected call of Apply.
func (mr *MockMigrationMockRecorder) Apply(db any) *MigrationApplyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockMigration)(nil).Apply), db)
	return &MigrationApplyCall{Call: call}
}

// MigrationApplyCall wrap *gomock.Call
type MigrationApplyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MigrationApplyCall) Return(arg0 error) *MigrationApplyCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MigrationApplyCall) Do(f func(Executor) error) *MigrationApplyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MigrationApplyCall) DoAndReturn(f func(Executor) error) *MigrationApplyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Name mocks base method.
func (m *MockMigration) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockMigrationMockRecorder) Name() *MigrationNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockMigration)(nil).Name))
	return &MigrationNameCall{Call: call}
}

// MigrationNameCall wrap *gomock.Call
type MigrationNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MigrationNameCall) Return(arg0 string) *MigrationNameCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MigrationNameCall) Do(f func() string) *MigrationNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MigrationNameCall) DoAndReturn(f func() string) *MigrationNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Order mocks base method.
func (m *MockMigration) Order() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Order")
	ret0, _ := ret[0].(int)
	return ret0
}

// Order indicates an expected call of Order.
func (mr *MockMigrationMockRecorder) Order() *MigrationOrderCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Order", reflect.TypeOf((*MockMigration)(nil).Order))
	return &MigrationOrderCall{Call: call}
}

// MigrationOrderCall wrap *gomock.Call
type MigrationOrderCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MigrationOrderCall) Return(arg0 int) *MigrationOrderCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MigrationOrderCall) Do(f func() int) *MigrationOrderCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MigrationOrderCall) DoAndReturn(f func() int) *MigrationOrderCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Rollback mocks base method.
func (m *MockMigration) Rollback() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback")
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockMigrationMockRecorder) Rollback() *MigrationRollbackCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockMigration)(nil).Rollback))
	return &MigrationRollbackCall{Call: call}
}

// MigrationRollbackCall wrap *gomock.Call
type MigrationRollbackCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MigrationRollbackCall) Return(arg0 error) *MigrationRollbackCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MigrationRollbackCall) Do(f func() error) *MigrationRollbackCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MigrationRollbackCall) DoAndReturn(f func() error) *MigrationRollbackCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}