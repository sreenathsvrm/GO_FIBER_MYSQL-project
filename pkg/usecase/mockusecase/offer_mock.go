// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/usecase/interface/offer.go

// Package mockusecase is a generated GoMock package.
package mockusecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/sreenathsvrm/voix-me/pkg/domain"
)

// MockOfferUseCase is a mock of OfferUseCase interface.
type MockOfferUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockOfferUseCaseMockRecorder
}

// MockOfferUseCaseMockRecorder is the mock recorder for MockOfferUseCase.
type MockOfferUseCaseMockRecorder struct {
	mock *MockOfferUseCase
}

// NewMockOfferUseCase creates a new mock instance.
func NewMockOfferUseCase(ctrl *gomock.Controller) *MockOfferUseCase {
	mock := &MockOfferUseCase{ctrl: ctrl}
	mock.recorder = &MockOfferUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOfferUseCase) EXPECT() *MockOfferUseCaseMockRecorder {
	return m.recorder
}

// FindAll mocks base method.
func (m *MockOfferUseCase) FindAll(ctx context.Context, countryName string) ([]domain.OfferCompany, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx, countryName)
	ret0, _ := ret[0].([]domain.OfferCompany)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockOfferUseCaseMockRecorder) FindAll(ctx, countryName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockOfferUseCase)(nil).FindAll), ctx, countryName)
}
