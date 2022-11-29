package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"probabilisticTimeSeriesModeling/config"
	"probabilisticTimeSeriesModeling/pkg/logger"
)

type service struct {
	registry ServiceRegistry
	srv      Server
}

type Service interface {
	NewService(*config.Config, logger.Logger, *sqlx.DB) (Controllers, error)
	NewMDWManager(*config.Config, logger.Logger) (MDWManager, error)
	NewLogger(*config.Config, logger.Logger) (Logger, error)
}

func NewService() (obj Service, err error) {
	registry, err := NewRegistry()
	if err != nil {
		return
	}
	return &service{
		registry: registry,
	}, err
}

func (s *service) NewService(cfg *config.Config, logger logger.Logger, pgDB *sqlx.DB) (_ Controllers, err error) {
	creditCtrl, err := s.registry.NewCreditReg(cfg, logger, pgDB)
	if err != nil {
		return
	}
	userCtrl, err := s.registry.NewUserReg(cfg, pgDB)
	if err != nil {
		return
	}
	sessionCtrl, err := s.registry.NewSessionReg(cfg, pgDB)
	if err != nil {
		return
	}
	return Controllers{
		creditCtrl,
		userCtrl,
		sessionCtrl,
	}, err
}

func (s *service) NewMDWManager(cfg *config.Config, logger logger.Logger) (mw MDWManager, err error) {
	obj, err := s.registry.NewMDWManager(cfg, logger)
	if err != nil {
		err = errors.Wrapf(err, "service.NewMDWManager()")
		return
	}
	return MDWManager{
		obj,
	}, err
}

func (s *service) NewLogger(cfg *config.Config, log logger.Logger) (logger Logger, err error) {
	obj, err := s.registry.NewLogging(cfg, log)
	if err != nil {
		err = errors.Wrapf(err, "service.NewLogger()")
		return
	}
	return Logger{
		obj,
	}, err
}
