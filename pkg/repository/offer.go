package repository

import (
	"context"
	"fmt"

	domain "github.com/sreenathsvrm/voix-me/pkg/domain"
	interfaces "github.com/sreenathsvrm/voix-me/pkg/repository/interface"
	"gorm.io/gorm"
)

type offerDatabase struct {
	DB *gorm.DB
}

func NewOfferCompanyRepository(DB *gorm.DB) interfaces.OfferRepository {
	return &offerDatabase{DB}
}

func (c *offerDatabase) FindAll(ctx context.Context, countryName string) ([]domain.OfferCompany, error) {
	fmt.Println("called db")
	var offerCompanies []domain.OfferCompany

	query := `SELECT * FROM offer_company WHERE country = ?`
	err := c.DB.Raw(query, countryName).Scan(&offerCompanies).Error

	return offerCompanies, err
}
