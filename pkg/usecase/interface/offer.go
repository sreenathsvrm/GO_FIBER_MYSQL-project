package interfaces

import (
	"context"

	domain "github.com/sreenathsvrm/voix-me/pkg/domain"
)

type OfferUseCase interface {
	FindAll(ctx context.Context, countryName string) ([]domain.OfferCompany, error)
}
