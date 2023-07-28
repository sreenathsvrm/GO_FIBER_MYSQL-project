//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "github.com/sreenathsvrm/voix-me/pkg/api"
	handler "github.com/sreenathsvrm/voix-me/pkg/api/handler"
	config "github.com/sreenathsvrm/voix-me/pkg/config"
	db "github.com/sreenathsvrm/voix-me/pkg/db"
	repository "github.com/sreenathsvrm/voix-me/pkg/repository"
	"github.com/sreenathsvrm/voix-me/pkg/service"
	usecase "github.com/sreenathsvrm/voix-me/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase, repository.NewOfferCompanyRepository,
		service.NewInMemoryStorage,
		usecase.NewOfferCompanyUseCase,
		handler.NewOfferCompanyHandler,
		http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
