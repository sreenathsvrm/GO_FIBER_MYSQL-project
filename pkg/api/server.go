package http

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"

	_ "github.com/sreenathsvrm/voix-me/cmd/api/docs"
	handler "github.com/sreenathsvrm/voix-me/pkg/api/handler"
	"github.com/sreenathsvrm/voix-me/pkg/api/middleware"
	"github.com/sreenathsvrm/voix-me/pkg/config"
)

type ServerHTTP struct {
	engine *fiber.App
	port   string
	server *http.Server
}

func NewServerHTTP(cfg config.Config, offerCompanyHandler *handler.OfferCompanyHandler) *ServerHTTP {
	engine := fiber.New()

	engine.Use(middleware.LimitMiddlewarehHandler())

	engine.Get("api/offer-company/:country_name", offerCompanyHandler.FindAll)

	// server := &http.Server{
	// 	Addr:    ":" + cfg.Port,
	// 	Handler: engine,
	// }

	return &ServerHTTP{
		engine: engine,
		port:   cfg.Port,
		// server: server,
	}
}

func (sh *ServerHTTP) Start() {
	go func() {
		// if err := sh.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		// 	panic(err)
		// }
		sh.engine.Listen(":8000")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sh.server.Shutdown(ctx); err != nil {
		panic(err)
	}
}
