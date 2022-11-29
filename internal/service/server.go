package service

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jmoiron/sqlx"
	"log"
	"probabilisticTimeSeriesModeling/config"
	"probabilisticTimeSeriesModeling/internal/credit/controller"
	"probabilisticTimeSeriesModeling/pkg/logger"
)

type server struct {
	app    *fiber.App
	ctx    *context.Context
	config *config.Config
	logger logger.Logger
	pgDB   *sqlx.DB
}

type Server interface {
	RunServer() (err error)
}

func NewServer(
	app *fiber.App,
	ctx *context.Context,
	config *config.Config,
	logger logger.Logger,
	database *sqlx.DB,
) (obj Server, err error) {
	return &server{
		app:    app,
		ctx:    ctx,
		config: config,
		logger: logger,
		pgDB:   database,
	}, err
}

func (s *server) RunServer() (err error) {
	newService, err := NewService()
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	//utils
	mw, err := newService.NewMDWManager(s.config, s.logger)
	_, err = newService.NewLogger(s.config, s.logger)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	//controllers
	controllers, err := newService.NewService(s.config, s.logger, s.pgDB)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	s.app.Use(fiberLogger.New())
	s.app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	//routing
	api := s.app.Group("/api")
	creditRoute := api.Group("/credit")

	controller.CreditRoutesGroup(mw.MDWManager, creditRoute, controllers.Credit)

	return
}
