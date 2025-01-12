// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/interfaces.go
//
// Generated by this command:
//
//	mockgen -source=internal/repository/interfaces.go -destination=internal/repository/interfaces.mock.gen.go -package=repository
//

// Package repository is a generated GoMock package.
package repository

import (
	model "coding-test-11-01-2025-dating-app/internal/model"
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface.
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
	isgomock struct{}
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface.
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance.
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// CheckMatch mocks base method.
func (m *MockRepositoryInterface) CheckMatch(ctx context.Context, userID, profileID uint) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckMatch", ctx, userID, profileID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckMatch indicates an expected call of CheckMatch.
func (mr *MockRepositoryInterfaceMockRecorder) CheckMatch(ctx, userID, profileID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckMatch", reflect.TypeOf((*MockRepositoryInterface)(nil).CheckMatch), ctx, userID, profileID)
}

// CreateProfile mocks base method.
func (m *MockRepositoryInterface) CreateProfile(ctx context.Context, profile *model.Profile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProfile", ctx, profile)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProfile indicates an expected call of CreateProfile.
func (mr *MockRepositoryInterfaceMockRecorder) CreateProfile(ctx, profile any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProfile", reflect.TypeOf((*MockRepositoryInterface)(nil).CreateProfile), ctx, profile)
}

// CreateSubscription mocks base method.
func (m *MockRepositoryInterface) CreateSubscription(ctx context.Context, sub *model.Subscription) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubscription", ctx, sub)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSubscription indicates an expected call of CreateSubscription.
func (mr *MockRepositoryInterfaceMockRecorder) CreateSubscription(ctx, sub any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubscription", reflect.TypeOf((*MockRepositoryInterface)(nil).CreateSubscription), ctx, sub)
}

// CreateSwipe mocks base method.
func (m *MockRepositoryInterface) CreateSwipe(ctx context.Context, swipe *model.Swipe) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSwipe", ctx, swipe)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSwipe indicates an expected call of CreateSwipe.
func (mr *MockRepositoryInterfaceMockRecorder) CreateSwipe(ctx, swipe any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSwipe", reflect.TypeOf((*MockRepositoryInterface)(nil).CreateSwipe), ctx, swipe)
}

// CreateUser mocks base method.
func (m *MockRepositoryInterface) CreateUser(ctx context.Context, user *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockRepositoryInterfaceMockRecorder) CreateUser(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockRepositoryInterface)(nil).CreateUser), ctx, user)
}

// Delete mocks base method.
func (m *MockRepositoryInterface) Delete(ctx context.Context, user *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryInterfaceMockRecorder) Delete(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepositoryInterface)(nil).Delete), ctx, user)
}

// GetActiveSubscription mocks base method.
func (m *MockRepositoryInterface) GetActiveSubscription(ctx context.Context, userID uint) (*model.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveSubscription", ctx, userID)
	ret0, _ := ret[0].(*model.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActiveSubscription indicates an expected call of GetActiveSubscription.
func (mr *MockRepositoryInterfaceMockRecorder) GetActiveSubscription(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveSubscription", reflect.TypeOf((*MockRepositoryInterface)(nil).GetActiveSubscription), ctx, userID)
}

// GetDailySwipeCount mocks base method.
func (m *MockRepositoryInterface) GetDailySwipeCount(ctx context.Context, userID uint) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDailySwipeCount", ctx, userID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDailySwipeCount indicates an expected call of GetDailySwipeCount.
func (mr *MockRepositoryInterfaceMockRecorder) GetDailySwipeCount(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDailySwipeCount", reflect.TypeOf((*MockRepositoryInterface)(nil).GetDailySwipeCount), ctx, userID)
}

// GetUnswiped mocks base method.
func (m *MockRepositoryInterface) GetUnswiped(ctx context.Context, userID uint, limit int) ([]model.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnswiped", ctx, userID, limit)
	ret0, _ := ret[0].([]model.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnswiped indicates an expected call of GetUnswiped.
func (mr *MockRepositoryInterfaceMockRecorder) GetUnswiped(ctx, userID, limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnswiped", reflect.TypeOf((*MockRepositoryInterface)(nil).GetUnswiped), ctx, userID, limit)
}

// GetUserByID mocks base method.
func (m *MockRepositoryInterface) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockRepositoryInterfaceMockRecorder) GetUserByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockRepositoryInterface)(nil).GetUserByID), ctx, id)
}

// GetUserByUsername mocks base method.
func (m *MockRepositoryInterface) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", ctx, username)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockRepositoryInterfaceMockRecorder) GetUserByUsername(ctx, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockRepositoryInterface)(nil).GetUserByUsername), ctx, username)
}

// IsUserPremium mocks base method.
func (m *MockRepositoryInterface) IsUserPremium(ctx context.Context, userID uint) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUserPremium", ctx, userID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsUserPremium indicates an expected call of IsUserPremium.
func (mr *MockRepositoryInterfaceMockRecorder) IsUserPremium(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUserPremium", reflect.TypeOf((*MockRepositoryInterface)(nil).IsUserPremium), ctx, userID)
}

// VerifyProfile mocks base method.
func (m *MockRepositoryInterface) VerifyProfile(ctx context.Context, userID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyProfile", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyProfile indicates an expected call of VerifyProfile.
func (mr *MockRepositoryInterfaceMockRecorder) VerifyProfile(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyProfile", reflect.TypeOf((*MockRepositoryInterface)(nil).VerifyProfile), ctx, userID)
}