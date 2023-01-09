// Code generated by MockGen. DO NOT EDIT.
// Source: crypto_stream_manager.go

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/hktalent/quic-go/internal/protocol"
)

// MockCryptoDataHandler is a mock of CryptoDataHandler interface.
type MockCryptoDataHandler struct {
	ctrl     *gomock.Controller
	recorder *MockCryptoDataHandlerMockRecorder
}

// MockCryptoDataHandlerMockRecorder is the mock recorder for MockCryptoDataHandler.
type MockCryptoDataHandlerMockRecorder struct {
	mock *MockCryptoDataHandler
}

// NewMockCryptoDataHandler creates a new mock instance.
func NewMockCryptoDataHandler(ctrl *gomock.Controller) *MockCryptoDataHandler {
	mock := &MockCryptoDataHandler{ctrl: ctrl}
	mock.recorder = &MockCryptoDataHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCryptoDataHandler) EXPECT() *MockCryptoDataHandlerMockRecorder {
	return m.recorder
}

// HandleMessage mocks base method.
func (m *MockCryptoDataHandler) HandleMessage(arg0 []byte, arg1 protocol.EncryptionLevel) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleMessage", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HandleMessage indicates an expected call of HandleMessage.
func (mr *MockCryptoDataHandlerMockRecorder) HandleMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleMessage", reflect.TypeOf((*MockCryptoDataHandler)(nil).HandleMessage), arg0, arg1)
}
