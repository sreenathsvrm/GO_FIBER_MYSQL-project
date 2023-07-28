// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/service/in_memmory_db.go

// Package mockinmemmory_db is a generated GoMock package.
package mockinmemmory_db

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/sreenathsvrm/voix-me/pkg/domain"
)

// MockInMemoryStorage is a mock of InMemoryStorage interface.
type MockInMemoryStorage struct {
	ctrl     *gomock.Controller
	recorder *MockInMemoryStorageMockRecorder
}

// MockInMemoryStorageMockRecorder is the mock recorder for MockInMemoryStorage.
type MockInMemoryStorageMockRecorder struct {
	mock *MockInMemoryStorage
}

// NewMockInMemoryStorage creates a new mock instance.
func NewMockInMemoryStorage(ctrl *gomock.Controller) *MockInMemoryStorage {
	mock := &MockInMemoryStorage{ctrl: ctrl}
	mock.recorder = &MockInMemoryStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInMemoryStorage) EXPECT() *MockInMemoryStorageMockRecorder {
	return m.recorder
}

// GetOfferCompanyByCountryName mocks base method.
func (m *MockInMemoryStorage) GetOfferCompanyByCountryName(ctx context.Context, countryName string) ([]domain.OfferCompany, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOfferCompanyByCountryName", ctx, countryName)
	ret0, _ := ret[0].([]domain.OfferCompany)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetOfferCompanyByCountryName indicates an expected call of GetOfferCompanyByCountryName.
func (mr *MockInMemoryStorageMockRecorder) GetOfferCompanyByCountryName(ctx, countryName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOfferCompanyByCountryName", reflect.TypeOf((*MockInMemoryStorage)(nil).GetOfferCompanyByCountryName), ctx, countryName)
}

// SetOfferCompany mocks base method.
func (m *MockInMemoryStorage) SetOfferCompany(ctx context.Context, countryName string, companyOffer []domain.OfferCompany) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetOfferCompany", ctx, countryName, companyOffer)
}

// SetOfferCompany indicates an expected call of SetOfferCompany.
func (mr *MockInMemoryStorageMockRecorder) SetOfferCompany(ctx, countryName, companyOffer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOfferCompany", reflect.TypeOf((*MockInMemoryStorage)(nil).SetOfferCompany), ctx, countryName, companyOffer)
}
