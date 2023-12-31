// Code generated by MockGen. DO NOT EDIT.
// Source: reposiotry.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"
	time "time"

	domain "github.com/begenov/register-service/internal/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockUsers is a mock of Users interface.
type MockUsers struct {
	ctrl     *gomock.Controller
	recorder *MockUsersMockRecorder
}

// MockUsersMockRecorder is the mock recorder for MockUsers.
type MockUsersMockRecorder struct {
	mock *MockUsers
}

// NewMockUsers creates a new mock instance.
func NewMockUsers(ctrl *gomock.Controller) *MockUsers {
	mock := &MockUsers{ctrl: ctrl}
	mock.recorder = &MockUsersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsers) EXPECT() *MockUsersMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUsers) Create(ctx context.Context, user domain.Register) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUsersMockRecorder) Create(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUsers)(nil).Create), ctx, user)
}

// GetByEmail mocks base method.
func (m *MockUsers) GetByEmail(ctx context.Context, email string) (domain.Register, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", ctx, email)
	ret0, _ := ret[0].(domain.Register)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail.
func (mr *MockUsersMockRecorder) GetByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockUsers)(nil).GetByEmail), ctx, email)
}

// GetUserById mocks base method.
func (m *MockUsers) GetUserById(ctx context.Context, id int) (domain.Register, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", ctx, id)
	ret0, _ := ret[0].(domain.Register)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUsersMockRecorder) GetUserById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUsers)(nil).GetUserById), ctx, id)
}

// GetUserByRefreshToken mocks base method.
func (m *MockUsers) GetUserByRefreshToken(ctx context.Context, token string) (domain.Register, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByRefreshToken", ctx, token)
	ret0, _ := ret[0].(domain.Register)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByRefreshToken indicates an expected call of GetUserByRefreshToken.
func (mr *MockUsersMockRecorder) GetUserByRefreshToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByRefreshToken", reflect.TypeOf((*MockUsers)(nil).GetUserByRefreshToken), ctx, token)
}

// SetSession mocks base method.
func (m *MockUsers) SetSession(ctx context.Context, userID int, token string, expiresAt time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSession", ctx, userID, token, expiresAt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSession indicates an expected call of SetSession.
func (mr *MockUsersMockRecorder) SetSession(ctx, userID, token, expiresAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSession", reflect.TypeOf((*MockUsers)(nil).SetSession), ctx, userID, token, expiresAt)
}

// MockCourier is a mock of Courier interface.
type MockCourier struct {
	ctrl     *gomock.Controller
	recorder *MockCourierMockRecorder
}

// MockCourierMockRecorder is the mock recorder for MockCourier.
type MockCourierMockRecorder struct {
	mock *MockCourier
}

// NewMockCourier creates a new mock instance.
func NewMockCourier(ctrl *gomock.Controller) *MockCourier {
	mock := &MockCourier{ctrl: ctrl}
	mock.recorder = &MockCourierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCourier) EXPECT() *MockCourierMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCourier) Create(ctx context.Context, courier domain.Register) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, courier)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockCourierMockRecorder) Create(ctx, courier interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCourier)(nil).Create), ctx, courier)
}

// GetCourierByEmail mocks base method.
func (m *MockCourier) GetCourierByEmail(ctx context.Context, email string) (domain.Register, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourierByEmail", ctx, email)
	ret0, _ := ret[0].(domain.Register)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourierByEmail indicates an expected call of GetCourierByEmail.
func (mr *MockCourierMockRecorder) GetCourierByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourierByEmail", reflect.TypeOf((*MockCourier)(nil).GetCourierByEmail), ctx, email)
}

// GetCourierByID mocks base method.
func (m *MockCourier) GetCourierByID(ctx context.Context, id int) (domain.Register, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourierByID", ctx, id)
	ret0, _ := ret[0].(domain.Register)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourierByID indicates an expected call of GetCourierByID.
func (mr *MockCourierMockRecorder) GetCourierByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourierByID", reflect.TypeOf((*MockCourier)(nil).GetCourierByID), ctx, id)
}

// GetCourierByRefreshToken mocks base method.
func (m *MockCourier) GetCourierByRefreshToken(ctx context.Context, token string) (domain.Register, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourierByRefreshToken", ctx, token)
	ret0, _ := ret[0].(domain.Register)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourierByRefreshToken indicates an expected call of GetCourierByRefreshToken.
func (mr *MockCourierMockRecorder) GetCourierByRefreshToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourierByRefreshToken", reflect.TypeOf((*MockCourier)(nil).GetCourierByRefreshToken), ctx, token)
}

// SetSession mocks base method.
func (m *MockCourier) SetSession(ctx context.Context, token string, expiredAt time.Time, courierId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSession", ctx, token, expiredAt, courierId)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSession indicates an expected call of SetSession.
func (mr *MockCourierMockRecorder) SetSession(ctx, token, expiredAt, courierId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSession", reflect.TypeOf((*MockCourier)(nil).SetSession), ctx, token, expiredAt, courierId)
}

// MockRestaurant is a mock of Restaurant interface.
type MockRestaurant struct {
	ctrl     *gomock.Controller
	recorder *MockRestaurantMockRecorder
}

// MockRestaurantMockRecorder is the mock recorder for MockRestaurant.
type MockRestaurantMockRecorder struct {
	mock *MockRestaurant
}

// NewMockRestaurant creates a new mock instance.
func NewMockRestaurant(ctrl *gomock.Controller) *MockRestaurant {
	mock := &MockRestaurant{ctrl: ctrl}
	mock.recorder = &MockRestaurantMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRestaurant) EXPECT() *MockRestaurantMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRestaurant) Create(ctx context.Context, user domain.Register) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRestaurantMockRecorder) Create(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRestaurant)(nil).Create), ctx, user)
}

// GetRestaurantByEmail mocks base method.
func (m *MockRestaurant) GetRestaurantByEmail(ctx context.Context, email string) (domain.Register, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestaurantByEmail", ctx, email)
	ret0, _ := ret[0].(domain.Register)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestaurantByEmail indicates an expected call of GetRestaurantByEmail.
func (mr *MockRestaurantMockRecorder) GetRestaurantByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestaurantByEmail", reflect.TypeOf((*MockRestaurant)(nil).GetRestaurantByEmail), ctx, email)
}

// GetRestaurantByID mocks base method.
func (m *MockRestaurant) GetRestaurantByID(ctx context.Context, id int) (domain.Register, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestaurantByID", ctx, id)
	ret0, _ := ret[0].(domain.Register)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestaurantByID indicates an expected call of GetRestaurantByID.
func (mr *MockRestaurantMockRecorder) GetRestaurantByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestaurantByID", reflect.TypeOf((*MockRestaurant)(nil).GetRestaurantByID), ctx, id)
}

// GetRestaurantByRefreshToken mocks base method.
func (m *MockRestaurant) GetRestaurantByRefreshToken(ctx context.Context, token string) (domain.Register, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestaurantByRefreshToken", ctx, token)
	ret0, _ := ret[0].(domain.Register)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestaurantByRefreshToken indicates an expected call of GetRestaurantByRefreshToken.
func (mr *MockRestaurantMockRecorder) GetRestaurantByRefreshToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestaurantByRefreshToken", reflect.TypeOf((*MockRestaurant)(nil).GetRestaurantByRefreshToken), ctx, token)
}

// SetSession mocks base method.
func (m *MockRestaurant) SetSession(ctx context.Context, token string, expiredAt time.Time, restaurantId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSession", ctx, token, expiredAt, restaurantId)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSession indicates an expected call of SetSession.
func (mr *MockRestaurantMockRecorder) SetSession(ctx, token, expiredAt, restaurantId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSession", reflect.TypeOf((*MockRestaurant)(nil).SetSession), ctx, token, expiredAt, restaurantId)
}
