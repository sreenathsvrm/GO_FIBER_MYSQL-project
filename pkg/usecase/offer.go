package usecase

import (
	"context"

	domain "github.com/sreenathsvrm/voix-me/pkg/domain"
	interfaces "github.com/sreenathsvrm/voix-me/pkg/repository/interface"
	"github.com/sreenathsvrm/voix-me/pkg/service"
	services "github.com/sreenathsvrm/voix-me/pkg/usecase/interface"
)

type offerUseCase struct {
	offerRepo  interfaces.OfferRepository
	inMemoryDB service.InMemoryStorage
}

func NewOfferCompanyUseCase(repo interfaces.OfferRepository, inMemoryDB service.InMemoryStorage) services.OfferUseCase {
	return &offerUseCase{
		offerRepo:  repo,
		inMemoryDB: inMemoryDB,
	}

}

func (c *offerUseCase) FindAll(ctx context.Context, countryName string) ([]domain.OfferCompany, error) {

	offerCompanies, ok := c.inMemoryDB.GetOfferCompanyByCountryName(ctx, countryName)
	if ok { // if data is in in memory then return that data
		return offerCompanies, nil
	}

	// if not available then get from database
	offerCompanies, err := c.offerRepo.FindAll(ctx, countryName)

	if err != nil {
		return nil, err
	}
	if offerCompanies == nil {
		return nil, nil
	}

	// use go routines to store it on in memory
	go c.storeOfferCompanyOnInMemory(ctx, countryName, offerCompanies)

	return offerCompanies, err
}

func (c *offerUseCase) storeOfferCompanyOnInMemory(ctx context.Context, countryName string, offerCompanies []domain.OfferCompany) {
	c.inMemoryDB.SetOfferCompany(ctx, countryName, offerCompanies)
}
