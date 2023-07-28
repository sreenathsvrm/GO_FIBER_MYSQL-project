package interfaces

import (
	"context"

	"github.com/sreenathsvrm/voix-me/pkg/domain"
)

type OfferRepository interface {
	FindAll(ctx context.Context, countryName string) ([]domain.OfferCompany, error)
}
