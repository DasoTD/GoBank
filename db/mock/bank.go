// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dasotd/gobank/db/sqlc (interfaces: Bank)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	sqlc "github.com/dasotd/gobank/db/sqlc"
	gomock "github.com/golang/mock/gomock"
)

// MockBank is a mock of Bank interface.
type MockBank struct {
	ctrl     *gomock.Controller
	recorder *MockBankMockRecorder
}

// MockBankMockRecorder is the mock recorder for MockBank.
type MockBankMockRecorder struct {
	mock *MockBank
}

// NewMockBank creates a new mock instance.
func NewMockBank(ctrl *gomock.Controller) *MockBank {
	mock := &MockBank{ctrl: ctrl}
	mock.recorder = &MockBankMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBank) EXPECT() *MockBankMockRecorder {
	return m.recorder
}

// AddAccountBalance mocks base method.
func (m *MockBank) AddAccountBalance(arg0 context.Context, arg1 sqlc.AddAccountBalanceParams) (sqlc.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAccountBalance", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAccountBalance indicates an expected call of AddAccountBalance.
func (mr *MockBankMockRecorder) AddAccountBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAccountBalance", reflect.TypeOf((*MockBank)(nil).AddAccountBalance), arg0, arg1)
}

// CreateAccount mocks base method.
func (m *MockBank) CreateAccount(arg0 context.Context, arg1 sqlc.CreateAccountParams) (sqlc.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockBankMockRecorder) CreateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockBank)(nil).CreateAccount), arg0, arg1)
}

// CreateEntry mocks base method.
func (m *MockBank) CreateEntry(arg0 context.Context, arg1 sqlc.CreateEntryParams) (sqlc.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEntry", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEntry indicates an expected call of CreateEntry.
func (mr *MockBankMockRecorder) CreateEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEntry", reflect.TypeOf((*MockBank)(nil).CreateEntry), arg0, arg1)
}

// CreateTransfer mocks base method.
func (m *MockBank) CreateTransfer(arg0 context.Context, arg1 sqlc.CreateTransferParams) (sqlc.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransfer", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransfer indicates an expected call of CreateTransfer.
func (mr *MockBankMockRecorder) CreateTransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransfer", reflect.TypeOf((*MockBank)(nil).CreateTransfer), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockBank) CreateUser(arg0 context.Context, arg1 sqlc.CreateUserParams) (sqlc.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(sqlc.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockBankMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockBank)(nil).CreateUser), arg0, arg1)
}

// DeleteAccount mocks base method.
func (m *MockBank) DeleteAccount(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockBankMockRecorder) DeleteAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockBank)(nil).DeleteAccount), arg0, arg1)
}

// DeleteEntry mocks base method.
func (m *MockBank) DeleteEntry(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEntry", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEntry indicates an expected call of DeleteEntry.
func (mr *MockBankMockRecorder) DeleteEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEntry", reflect.TypeOf((*MockBank)(nil).DeleteEntry), arg0, arg1)
}

// GetAccount mocks base method.
func (m *MockBank) GetAccount(arg0 context.Context, arg1 int64) (sqlc.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockBankMockRecorder) GetAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockBank)(nil).GetAccount), arg0, arg1)
}

// GetAccountForUpdate mocks base method.
func (m *MockBank) GetAccountForUpdate(arg0 context.Context, arg1 int64) (sqlc.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountForUpdate", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountForUpdate indicates an expected call of GetAccountForUpdate.
func (mr *MockBankMockRecorder) GetAccountForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountForUpdate", reflect.TypeOf((*MockBank)(nil).GetAccountForUpdate), arg0, arg1)
}

// GetEntry mocks base method.
func (m *MockBank) GetEntry(arg0 context.Context, arg1 int64) (sqlc.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntry", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntry indicates an expected call of GetEntry.
func (mr *MockBankMockRecorder) GetEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntry", reflect.TypeOf((*MockBank)(nil).GetEntry), arg0, arg1)
}

// GetTransfer mocks base method.
func (m *MockBank) GetTransfer(arg0 context.Context, arg1 int64) (sqlc.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransfer", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransfer indicates an expected call of GetTransfer.
func (mr *MockBankMockRecorder) GetTransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransfer", reflect.TypeOf((*MockBank)(nil).GetTransfer), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockBank) GetUser(arg0 context.Context, arg1 string) (sqlc.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(sqlc.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockBankMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockBank)(nil).GetUser), arg0, arg1)
}

// ListAccount mocks base method.
func (m *MockBank) ListAccount(arg0 context.Context, arg1 sqlc.ListAccountParams) ([]sqlc.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccount", arg0, arg1)
	ret0, _ := ret[0].([]sqlc.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccount indicates an expected call of ListAccount.
func (mr *MockBankMockRecorder) ListAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccount", reflect.TypeOf((*MockBank)(nil).ListAccount), arg0, arg1)
}

// ListEntries mocks base method.
func (m *MockBank) ListEntries(arg0 context.Context, arg1 sqlc.ListEntriesParams) ([]sqlc.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEntries", arg0, arg1)
	ret0, _ := ret[0].([]sqlc.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEntries indicates an expected call of ListEntries.
func (mr *MockBankMockRecorder) ListEntries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntries", reflect.TypeOf((*MockBank)(nil).ListEntries), arg0, arg1)
}

// ListTransfers mocks base method.
func (m *MockBank) ListTransfers(arg0 context.Context, arg1 sqlc.ListTransfersParams) ([]sqlc.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTransfers", arg0, arg1)
	ret0, _ := ret[0].([]sqlc.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTransfers indicates an expected call of ListTransfers.
func (mr *MockBankMockRecorder) ListTransfers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTransfers", reflect.TypeOf((*MockBank)(nil).ListTransfers), arg0, arg1)
}

// TransferTx mocks base method.
func (m *MockBank) TransferTx(arg0 context.Context, arg1 sqlc.TransferTxParams) (sqlc.TransferTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferTx", arg0, arg1)
	ret0, _ := ret[0].(sqlc.TransferTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransferTx indicates an expected call of TransferTx.
func (mr *MockBankMockRecorder) TransferTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferTx", reflect.TypeOf((*MockBank)(nil).TransferTx), arg0, arg1)
}

// UpdateAccount mocks base method.
func (m *MockBank) UpdateAccount(arg0 context.Context, arg1 sqlc.UpdateAccountParams) (sqlc.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockBankMockRecorder) UpdateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockBank)(nil).UpdateAccount), arg0, arg1)
}

// UpdateEntry mocks base method.
func (m *MockBank) UpdateEntry(arg0 context.Context, arg1 sqlc.UpdateEntryParams) (sqlc.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEntry", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEntry indicates an expected call of UpdateEntry.
func (mr *MockBankMockRecorder) UpdateEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEntry", reflect.TypeOf((*MockBank)(nil).UpdateEntry), arg0, arg1)
}
