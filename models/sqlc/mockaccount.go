// Code generated by MockGen. DO NOT EDIT.
// Source: accounts.sql.go

// Package db is a generated GoMock package.
package db

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockStore) CreateAccount(arg0 context.Context, arg1 CreateAccountParams) (Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", arg0, arg1)
	ret0, _ := ret[0].(Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockStoreMockRecorder) CreateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockStore)(nil).CreateAccount), arg0, arg1)
}

// CreateAccountStatementEntry mocks base method.
func (m *MockStore) CreateAccountStatementEntry(arg0 context.Context, arg1 CreateAccountStatementEntryParams) (AccountTransactionsEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccountStatementEntry", arg0, arg1)
	ret0, _ := ret[0].(AccountTransactionsEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccountStatementEntry indicates an expected call of CreateAccountStatementEntry.
func (mr *MockStoreMockRecorder) CreateAccountStatementEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccountStatementEntry", reflect.TypeOf((*MockStore)(nil).CreateAccountStatementEntry), arg0, arg1)
}

// CreateTransferRecord mocks base method.
func (m *MockStore) CreateTransferRecord(arg0 context.Context, arg1 CreateTransferRecordParams) (Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransferRecord", arg0, arg1)
	ret0, _ := ret[0].(Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransferRecord indicates an expected call of CreateTransferRecord.
func (mr *MockStoreMockRecorder) CreateTransferRecord(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransferRecord", reflect.TypeOf((*MockStore)(nil).CreateTransferRecord), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 CreateUserParams) (CreateUserRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(CreateUserRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteAccount mocks base method.
func (m *MockStore) DeleteAccount(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockStoreMockRecorder) DeleteAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockStore)(nil).DeleteAccount), arg0, arg1)
}

// GetAccountById mocks base method.
func (m *MockStore) GetAccountById(arg0 context.Context, arg1 int64) (Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountById", arg0, arg1)
	ret0, _ := ret[0].(Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountById indicates an expected call of GetAccountById.
func (mr *MockStoreMockRecorder) GetAccountById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountById", reflect.TypeOf((*MockStore)(nil).GetAccountById), arg0, arg1)
}

// GetAccountByOwnerEmail mocks base method.
func (m *MockStore) GetAccountByOwnerEmail(arg0 context.Context, arg1 string) (Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByOwnerEmail", arg0, arg1)
	ret0, _ := ret[0].(Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByOwnerEmail indicates an expected call of GetAccountByOwnerEmail.
func (mr *MockStoreMockRecorder) GetAccountByOwnerEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByOwnerEmail", reflect.TypeOf((*MockStore)(nil).GetAccountByOwnerEmail), arg0, arg1)
}

// GetBalanceByAccountId mocks base method.
func (m *MockStore) GetBalanceByAccountId(arg0 context.Context, arg1 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalanceByAccountId", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalanceByAccountId indicates an expected call of GetBalanceByAccountId.
func (mr *MockStoreMockRecorder) GetBalanceByAccountId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalanceByAccountId", reflect.TypeOf((*MockStore)(nil).GetBalanceByAccountId), arg0, arg1)
}

// GetBalanceByOwnerEmail mocks base method.
func (m *MockStore) GetBalanceByOwnerEmail(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalanceByOwnerEmail", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalanceByOwnerEmail indicates an expected call of GetBalanceByOwnerEmail.
func (mr *MockStoreMockRecorder) GetBalanceByOwnerEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalanceByOwnerEmail", reflect.TypeOf((*MockStore)(nil).GetBalanceByOwnerEmail), arg0, arg1)
}

// GetUserDetails mocks base method.
func (m *MockStore) GetUserDetails(arg0 context.Context, arg1 string) (GetUserDetailsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserDetails", arg0, arg1)
	ret0, _ := ret[0].(GetUserDetailsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserDetails indicates an expected call of GetUserDetails.
func (mr *MockStoreMockRecorder) GetUserDetails(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserDetails", reflect.TypeOf((*MockStore)(nil).GetUserDetails), arg0, arg1)
}

// ListAccounts mocks base method.
func (m *MockStore) ListAccounts(arg0 context.Context, arg1 ListAccountsParams) ([]Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccounts", arg0, arg1)
	ret0, _ := ret[0].([]Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccounts indicates an expected call of ListAccounts.
func (mr *MockStoreMockRecorder) ListAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccounts", reflect.TypeOf((*MockStore)(nil).ListAccounts), arg0, arg1)
}

// UpdateAccountBalanceById mocks base method.
func (m *MockStore) UpdateAccountBalanceById(arg0 context.Context, arg1 UpdateAccountBalanceByIdParams) (Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccountBalanceById", arg0, arg1)
	ret0, _ := ret[0].(Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccountBalanceById indicates an expected call of UpdateAccountBalanceById.
func (mr *MockStoreMockRecorder) UpdateAccountBalanceById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccountBalanceById", reflect.TypeOf((*MockStore)(nil).UpdateAccountBalanceById), arg0, arg1)
}
