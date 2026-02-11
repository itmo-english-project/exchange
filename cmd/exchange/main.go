package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/itmo-english-project/common/pkg/http/httphandler"
	"github.com/itmo-english-project/common/pkg/http/httpserver"
	"github.com/itmo-english-project/common/pkg/log"
	"github.com/itmo-english-project/common/pkg/storage/postgres"
	"github.com/itmo-english-project/exchange/cmd/config"
	"github.com/itmo-english-project/exchange/internal/adapters/in/http/exchanges/public"
	"github.com/itmo-english-project/exchange/internal/adapters/out/postgres/exchanges"
	"github.com/itmo-english-project/exchange/internal/services/exchange"
)

const (
	configPathEnv = "CONFIG_PATH"
)

func main() {
	cfg, err := config.InitConfig(configPathEnv)
	if err != nil {
		panic(err)
	}
	logger := log.NewLogger(cfg.General.Development)
	defer func() { _ = logger.Sync() }()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = ctx

	exchangeDB, err := postgres.NewDatabase(cfg.Database.Exchanges)
	if err != nil {
		logger.Panic("failed to connect to exchange database", zap.Error(err))
	}
	defer exchangeDB.Close()

	storage, err := exchanges.NewStorage(exchangeDB)
	if err != nil {
		logger.Panic("failed to create exchange storage", zap.Error(err))
	}

	exchangeService := exchange.NewService(storage)

	publicHandler := public.NewHandler(exchangeService)

	publicServer := httpserver.New(
		httphandler.DefaultMiddleware(logger),
		[]httpserver.Provider{publicHandler},
		httphandler.ErrorHandler,
	)
	go func() {
		if err = publicServer.Listen(fmt.Sprintf(":%d", cfg.General.Port)); err != nil {
			logger.Panic("failed to start public server", zap.Error(err))
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	logger.Info("app started")

	<-stop
	logger.Info("app stopped, stopped gracefully")
	signal.Stop(stop)
	close(stop)

	if err = publicServer.Shutdown(); err != nil {
		logger.Panic("failed to stop public server", zap.Error(err))
	}
}
