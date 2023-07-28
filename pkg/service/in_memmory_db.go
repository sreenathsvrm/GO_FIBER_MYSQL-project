package service

import (
	"context"
	"sync"

	"github.com/sreenathsvrm/voix-me/pkg/domain"
)

type InMemoryStorage interface {
	GetOfferCompanyByCountryName(ctx context.Context, countryName string) ([]domain.OfferCompany, bool)
	SetOfferCompany(ctx context.Context, countryName string, companyOffer []domain.OfferCompany)
}

type inMemoryStorage struct {
	db map[string][]domain.OfferCompany
	mu sync.Mutex
}

func NewInMemoryStorage() InMemoryStorage {
	return &inMemoryStorage{
		db: make(map[string][]domain.OfferCompany),
	}
}

func (c *inMemoryStorage) GetOfferCompanyByCountryName(ctx context.Context,
	countryName string) ([]domain.OfferCompany, bool) {

	offerCompany, ok := c.db[countryName]
	return offerCompany, ok
}
func (c *inMemoryStorage) SetOfferCompany(ctx context.Context, countryName string, companyOffer []domain.OfferCompany) {

	c.mu.Lock()
	c.db[countryName] = companyOffer
	c.mu.Unlock()
}
