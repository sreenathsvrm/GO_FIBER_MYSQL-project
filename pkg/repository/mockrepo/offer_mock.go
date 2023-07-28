// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/repository/interface/offer.go

// Package mockrepo is a generated GoMock package.
package mockrepo

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/sreenathsvrm/voix-me/pkg/domain"
)

// MockOfferRepository is a mock of OfferRepository interface.
type MockOfferRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOfferRepositoryMockRecorder
}

// MockOfferRepositoryMockRecorder is the mock recorder for MockOfferRepository.
type MockOfferRepositoryMockRecorder struct {
	mock *MockOfferRepository
}

// NewMockOfferRepository creates a new mock instance.
func NewMockOfferRepository(ctrl *gomock.Controller) *MockOfferRepository {
	mock := &MockOfferRepository{ctrl: ctrl}
	mock.recorder = &MockOfferRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOfferRepository) EXPECT() *MockOfferRepositoryMockRecorder {
	return m.recorder
}

// FindAll mocks base method.
func (m *MockOfferRepository) FindAll(ctx context.Context, countryName string) ([]domain.OfferCompany, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx, countryName)
	ret0, _ := ret[0].([]domain.OfferCompany)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockOfferRepositoryMockRecorder) FindAll(ctx, countryName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockOfferRepository)(nil).FindAll), ctx, countryName)
}