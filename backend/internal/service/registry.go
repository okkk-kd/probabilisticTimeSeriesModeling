package service

import (
	"github.com/jmoiron/sqlx"
	"probabilisticTimeSeriesModeling/config"
	"probabilisticTimeSeriesModeling/internal/credit/controller"
	"probabilisticTimeSeriesModeling/internal/middleware"
	registry "probabilisticTimeSeriesModeling/internal/service/registry"
	sessionCtrl "probabilisticTimeSeriesModeling/internal/session/controller"
	userCtrl "probabilisticTimeSeriesModeling/internal/users/controller"
	"probabilisticTimeSeriesModeling/pkg/fhttp"
	"probabilisticTimeSeriesModeling/pkg/logger"
)

type serviceRegistry struct {
	creditReg registry.CreditReg
	mw        registry.MDWManager
	logging   registry.Logging
	session   registry.SessionReg
	user      registry.UserReg
}

type ServiceRegistry interface {
	NewCreditReg(*config.Config, logger.Logger, *sqlx.DB) (
		controller.CreditCtrl,
		error,
	)
	NewMDWManager(*config.Config, logger.Logger, *sqlx.DB) (middleware.MDWManager, error)
	NewLogging(cfg *config.Config, logger logger.Logger) (obj logger.LoggerUC, err error)
	NewUserReg(cfg *config.Config, pgDB *sqlx.DB, logger logger.Logger) (obj userCtrl.UserCtrl, err error)
	NewSessionReg(cfg *config.Config, pgDB *sqlx.DB) (obj sessionCtrl.SessionCtrl, err error)
}

func NewRegistry() (obj ServiceRegistry, err error) {
	creditReg, err := registry.NewCreditReg()
	if err != nil {
		return
	}
	mw := registry.NewMDWManagerReg()
	logger := registry.NewLoggingReg()
	user, err := registry.NewUserReg()
	if err != nil {
		return
	}
	session, err := registry.NewSessionReg()
	if err != nil {
		return
	}
	return &serviceRegistry{
		creditReg: creditReg,
		mw:        mw,
		logging:   logger,
		user:      user,
		session:   session,
	}, err
}

func (r *serviceRegistry) NewCreditReg(cfg *config.Config, logger logger.Logger, pgDB *sqlx.DB) (
	ctrl controller.CreditCtrl,
	err error,
) {
	fhttpClient := fhttp.NewClient(cfg, logger)
	log := r.logging.NewLogging(cfg, logger)
	ctrl, err = r.creditReg.NewCreditCtrl(cfg, fhttpClient, log, pgDB)
	if err != nil {
		return
	}
	return
}

func (r *serviceRegistry) NewMDWManager(cfg *config.Config, logger logger.Logger, pgDB *sqlx.DB) (obj middleware.MDWManager, err error) {
	obj = r.mw.NewMDWManager(cfg, logger, pgDB)
	return
}

func (r *serviceRegistry) NewLogging(cfg *config.Config, logger logger.Logger) (obj logger.LoggerUC, err error) {
	obj = r.logging.NewLogging(cfg, logger)
	return
}

func (r *serviceRegistry) NewUserReg(cfg *config.Config, pgDB *sqlx.DB, logger logger.Logger) (obj userCtrl.UserCtrl, err error) {
	log := r.logging.NewLogging(cfg, logger)
	obj, err = r.user.NewUserCtrl(cfg, pgDB, log)
	if err != nil {
		return
	}
	return
}

func (r *serviceRegistry) NewSessionReg(cfg *config.Config, pgDB *sqlx.DB) (obj sessionCtrl.SessionCtrl, err error) {
	obj, err = r.session.NewSessionCtrl(cfg, pgDB)
	if err != nil {
		return
	}
	return
}
