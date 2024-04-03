// Code generated by MockGen. DO NOT EDIT.
// Source: ./syncer.go
//
// Generated by this command:
//
//	mockgen -typed -package=mocks -destination=./mocks/mocks.go -source=./syncer.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	types "github.com/spacemeshos/go-spacemesh/common/types"
	p2p "github.com/spacemeshos/go-spacemesh/p2p"
	gomock "go.uber.org/mock/gomock"
)

// Mockfetcher is a mock of fetcher interface.
type Mockfetcher struct {
	ctrl     *gomock.Controller
	recorder *MockfetcherMockRecorder
}

// MockfetcherMockRecorder is the mock recorder for Mockfetcher.
type MockfetcherMockRecorder struct {
	mock *Mockfetcher
}

// NewMockfetcher creates a new mock instance.
func NewMockfetcher(ctrl *gomock.Controller) *Mockfetcher {
	mock := &Mockfetcher{ctrl: ctrl}
	mock.recorder = &MockfetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockfetcher) EXPECT() *MockfetcherMockRecorder {
	return m.recorder
}

// GetMalfeasanceProofs mocks base method.
func (m *Mockfetcher) GetMalfeasanceProofs(arg0 context.Context, arg1 []types.NodeID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMalfeasanceProofs", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetMalfeasanceProofs indicates an expected call of GetMalfeasanceProofs.
func (mr *MockfetcherMockRecorder) GetMalfeasanceProofs(arg0, arg1 any) *MockfetcherGetMalfeasanceProofsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMalfeasanceProofs", reflect.TypeOf((*Mockfetcher)(nil).GetMalfeasanceProofs), arg0, arg1)
	return &MockfetcherGetMalfeasanceProofsCall{Call: call}
}

// MockfetcherGetMalfeasanceProofsCall wrap *gomock.Call
type MockfetcherGetMalfeasanceProofsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockfetcherGetMalfeasanceProofsCall) Return(arg0 error) *MockfetcherGetMalfeasanceProofsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockfetcherGetMalfeasanceProofsCall) Do(f func(context.Context, []types.NodeID) error) *MockfetcherGetMalfeasanceProofsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockfetcherGetMalfeasanceProofsCall) DoAndReturn(f func(context.Context, []types.NodeID) error) *MockfetcherGetMalfeasanceProofsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetMaliciousIDs mocks base method.
func (m *Mockfetcher) GetMaliciousIDs(arg0 context.Context, arg1 p2p.Peer) ([]types.NodeID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaliciousIDs", arg0, arg1)
	ret0, _ := ret[0].([]types.NodeID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMaliciousIDs indicates an expected call of GetMaliciousIDs.
func (mr *MockfetcherMockRecorder) GetMaliciousIDs(arg0, arg1 any) *MockfetcherGetMaliciousIDsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaliciousIDs", reflect.TypeOf((*Mockfetcher)(nil).GetMaliciousIDs), arg0, arg1)
	return &MockfetcherGetMaliciousIDsCall{Call: call}
}

// MockfetcherGetMaliciousIDsCall wrap *gomock.Call
type MockfetcherGetMaliciousIDsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockfetcherGetMaliciousIDsCall) Return(arg0 []types.NodeID, arg1 error) *MockfetcherGetMaliciousIDsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockfetcherGetMaliciousIDsCall) Do(f func(context.Context, p2p.Peer) ([]types.NodeID, error)) *MockfetcherGetMaliciousIDsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockfetcherGetMaliciousIDsCall) DoAndReturn(f func(context.Context, p2p.Peer) ([]types.NodeID, error)) *MockfetcherGetMaliciousIDsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SelectBestShuffled mocks base method.
func (m *Mockfetcher) SelectBestShuffled(arg0 int) []p2p.Peer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectBestShuffled", arg0)
	ret0, _ := ret[0].([]p2p.Peer)
	return ret0
}

// SelectBestShuffled indicates an expected call of SelectBestShuffled.
func (mr *MockfetcherMockRecorder) SelectBestShuffled(arg0 any) *MockfetcherSelectBestShuffledCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectBestShuffled", reflect.TypeOf((*Mockfetcher)(nil).SelectBestShuffled), arg0)
	return &MockfetcherSelectBestShuffledCall{Call: call}
}

// MockfetcherSelectBestShuffledCall wrap *gomock.Call
type MockfetcherSelectBestShuffledCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockfetcherSelectBestShuffledCall) Return(arg0 []p2p.Peer) *MockfetcherSelectBestShuffledCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockfetcherSelectBestShuffledCall) Do(f func(int) []p2p.Peer) *MockfetcherSelectBestShuffledCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockfetcherSelectBestShuffledCall) DoAndReturn(f func(int) []p2p.Peer) *MockfetcherSelectBestShuffledCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Mockcounter is a mock of counter interface.
type Mockcounter struct {
	ctrl     *gomock.Controller
	recorder *MockcounterMockRecorder
}

// MockcounterMockRecorder is the mock recorder for Mockcounter.
type MockcounterMockRecorder struct {
	mock *Mockcounter
}

// NewMockcounter creates a new mock instance.
func NewMockcounter(ctrl *gomock.Controller) *Mockcounter {
	mock := &Mockcounter{ctrl: ctrl}
	mock.recorder = &MockcounterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockcounter) EXPECT() *MockcounterMockRecorder {
	return m.recorder
}

// Inc mocks base method.
func (m *Mockcounter) Inc() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Inc")
}

// Inc indicates an expected call of Inc.
func (mr *MockcounterMockRecorder) Inc() *MockcounterIncCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inc", reflect.TypeOf((*Mockcounter)(nil).Inc))
	return &MockcounterIncCall{Call: call}
}

// MockcounterIncCall wrap *gomock.Call
type MockcounterIncCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockcounterIncCall) Return() *MockcounterIncCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockcounterIncCall) Do(f func()) *MockcounterIncCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockcounterIncCall) DoAndReturn(f func()) *MockcounterIncCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}