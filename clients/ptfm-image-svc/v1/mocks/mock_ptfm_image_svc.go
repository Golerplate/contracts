// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -source interface.go -destination mocks/mock_ptfm_image_svc.go -package ptfm_image_svc_v1_mocks
//

// Package ptfm_image_svc_v1_mocks is a generated GoMock package.
package ptfm_image_svc_v1_mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPtfmImageSvc is a mock of PtfmImageSvc interface.
type MockPtfmImageSvc struct {
	ctrl     *gomock.Controller
	recorder *MockPtfmImageSvcMockRecorder
}

// MockPtfmImageSvcMockRecorder is the mock recorder for MockPtfmImageSvc.
type MockPtfmImageSvcMockRecorder struct {
	mock *MockPtfmImageSvc
}

// NewMockPtfmImageSvc creates a new mock instance.
func NewMockPtfmImageSvc(ctrl *gomock.Controller) *MockPtfmImageSvc {
	mock := &MockPtfmImageSvc{ctrl: ctrl}
	mock.recorder = &MockPtfmImageSvcMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPtfmImageSvc) EXPECT() *MockPtfmImageSvcMockRecorder {
	return m.recorder
}

// CreateSignedProfilePictureUrl mocks base method.
func (m *MockPtfmImageSvc) CreateSignedProfilePictureUrl(ctx context.Context, userID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSignedProfilePictureUrl", ctx, userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSignedProfilePictureUrl indicates an expected call of CreateSignedProfilePictureUrl.
func (mr *MockPtfmImageSvcMockRecorder) CreateSignedProfilePictureUrl(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSignedProfilePictureUrl", reflect.TypeOf((*MockPtfmImageSvc)(nil).CreateSignedProfilePictureUrl), ctx, userID)
}
