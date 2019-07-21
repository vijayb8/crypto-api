package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/vijayb8/crypto-api/pkg/ordering"
	"github.com/vijayb8/crypto-api/pkg/platform/config"
	bhttp "github.com/vijayb8/crypto-api/pkg/platform/http"
	"github.com/vijayb8/crypto-api/pkg/platform/logger"
	"github.com/vijayb8/crypto-api/pkg/pricing"
)

// Version to return for homepage
var Version = "unset"

// App contains service configs and dependencies
type App struct {
	config         *config.Config
	log            *log.Logger
	pricingClient  *pricing.Client
	orderingClient *ordering.Client
}

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatalf("failed to load config: %s", err)
	}

	log, err := logger.Get(cfg.LogLevel)
	if err != nil {
		log.Fatalf("failed to create logger: %s", err)
	}

	pricingClient, err := pricing.NewClient(cfg.CoinMarket.URL, cfg.CoinMarket.ApiKey)
	if err != nil {
		log.Fatalf("can't get connection to CoinMarket API: %s", err)
	}

	orderingClient, err := ordering.NewClient(cfg.CryptoCompare.URL, cfg.CryptoCompare.ApiKey)
	if err != nil {
		log.Fatalf("can't get connection to CoinMarket API: %s", err)
	}

	app := &App{
		config:         cfg,
		pricingClient:  pricingClient,
		orderingClient: orderingClient,
	}

	addr := fmt.Sprintf(":%v", cfg.Port)
	log.Infof("listen and serve on %s", addr)

	srv := &http.Server{
		Addr:         addr,
		Handler:      app.getRouter(),
		ReadTimeout:  cfg.HTTPTimeouts.ServerRead,
		WriteTimeout: cfg.HTTPTimeouts.ServerWrite,
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGTERM)
		<-sigint

		if err := srv.Shutdown(context.Background()); err != nil {
			log.Infof("https-server shutdown error: %s", err)
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Infof("listener error: %v", err)
	}

	<-idleConnsClosed
}

// initRouter returns router with set params
func (a *App) getRouter() *chi.Mux {

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(bhttp.CORSMiddleware())

	r.Route("/v1", func(r chi.Router) {
		r.Get("/pricing", ordering.GetTopList(a.orderingClient, a.log))
	})

	return r
}
