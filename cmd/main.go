package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"os"
	"os/signal"
	"probabilisticTimeSeriesModeling/config"
	"probabilisticTimeSeriesModeling/internal/service"
	"probabilisticTimeSeriesModeling/pkg/httperrors"
	"probabilisticTimeSeriesModeling/pkg/logger"
	"probabilisticTimeSeriesModeling/pkg/storage/postgres"
	"syscall"
)

func main() {
	ctx := context.Background()

	cfg, err := config.LoadConfig()
	if err != nil {
		return
	}

	logger := logger.NewLogger(&cfg)
	err = logger.InitLogger()
	if err != nil {
		return
	}

	app := fiber.New(fiber.Config{ErrorHandler: httperrors.Init(&cfg, logger)})

	psqlDB, err := postgres.InitPsqlDB(ctx, &cfg)
	if err != nil {
		logger.Fatalf("PostgreSQL init error: %s", err)
	} else {
		logger.Infof("PostgreSQL connected, status: %#v", psqlDB.Stats())
	}

	defer func(psqlDB *sqlx.DB) {
		err = psqlDB.Close()
		if err != nil {
			logger.Infof(err.Error())
		} else {
			logger.Info("PostgreSQL closed properly")
		}
	}(psqlDB)

	s, err := service.NewServer(app, &ctx, &cfg, logger, psqlDB)
	if err != nil {
		return
	}

	err = s.RunServer()
	if err != nil {
		return
	}

	go func() {
		logger.Infof("Start server on port: %s:%s", cfg.Server.IP, cfg.Server.Port)
		if err := app.Listen(fmt.Sprintf("%s:%s", cfg.Server.IP, cfg.Server.Port)); err != nil {
			logger.Fatalf("Error starting Server: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	err = app.Shutdown()
	if err != nil {
		logger.Error(err)
	} else {
		logger.Info("Fiber server exited properly")
	}
}
