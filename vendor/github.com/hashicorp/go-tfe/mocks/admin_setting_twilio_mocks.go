// Code generated by MockGen. DO NOT EDIT.
// Source: admin_setting_twilio.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	tfe "github.com/hashicorp/go-tfe"
)

// MockTwilioSettings is a mock of TwilioSettings interface.
type MockTwilioSettings struct {
	ctrl     *gomock.Controller
	recorder *MockTwilioSettingsMockRecorder
}

// MockTwilioSettingsMockRecorder is the mock recorder for MockTwilioSettings.
type MockTwilioSettingsMockRecorder struct {
	mock *MockTwilioSettings
}

// NewMockTwilioSettings creates a new mock instance.
func NewMockTwilioSettings(ctrl *gomock.Controller) *MockTwilioSettings {
	mock := &MockTwilioSettings{ctrl: ctrl}
	mock.recorder = &MockTwilioSettingsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTwilioSettings) EXPECT() *MockTwilioSettingsMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockTwilioSettings) Read(ctx context.Context) (*tfe.AdminTwilioSetting, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx)
	ret0, _ := ret[0].(*tfe.AdminTwilioSetting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockTwilioSettingsMockRecorder) Read(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockTwilioSettings)(nil).Read), ctx)
}

// Update mocks base method.
func (m *MockTwilioSettings) Update(ctx context.Context, options tfe.AdminTwilioSettingsUpdateOptions) (*tfe.AdminTwilioSetting, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, options)
	ret0, _ := ret[0].(*tfe.AdminTwilioSetting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockTwilioSettingsMockRecorder) Update(ctx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTwilioSettings)(nil).Update), ctx, options)
}

// Verify mocks base method.
func (m *MockTwilioSettings) Verify(ctx context.Context, options tfe.AdminTwilioSettingsVerifyOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", ctx, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Verify indicates an expected call of Verify.
func (mr *MockTwilioSettingsMockRecorder) Verify(ctx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockTwilioSettings)(nil).Verify), ctx, options)
}
