package service

import (
	"probabilisticTimeSeriesModeling/config"
	"probabilisticTimeSeriesModeling/internal/credit/controller"
	"probabilisticTimeSeriesModeling/internal/middleware"
	registry "probabilisticTimeSeriesModeling/internal/service/registry"
	"probabilisticTimeSeriesModeling/pkg/fhttp"
	"probabilisticTimeSeriesModeling/pkg/logger"
)

type serviceRegistry struct {
	creditReg registry.CreditReg
	mw        registry.MDWManager
	logging   registry.Logging
}

type ServiceRegistry interface {
	NewCreditReg(*config.Config, logger.Logger) (
		controller.CreditCtrl,
		error,
	)
	NewMDWManager(*config.Config, logger.Logger) (middleware.MDWManager, error)
	NewLogging(cfg *config.Config, logger logger.Logger) (obj logger.LoggerUC, err error)
}

func NewRegistry() (obj ServiceRegistry, err error) {
	creditReg, err := registry.NewCreditReg()
	if err != nil {
		return
	}
	mw := registry.NewMDWManagerReg()
	logger := registry.NewLoggingReg()
	return &serviceRegistry{
		creditReg: creditReg,
		mw:        mw,
		logging:   logger,
	}, err
}

func (r *serviceRegistry) NewCreditReg(cfg *config.Config, logger logger.Logger) (
	ctrl controller.CreditCtrl,
	err error,
) {
	fhttpClient := fhttp.NewClient(cfg, logger)
	log := r.logging.NewLogging(cfg, logger)
	ctrl, err = r.creditReg.NewCreditCtrl(cfg, fhttpClient, log)
	if err != nil {
		return
	}
	return
}

func (r *serviceRegistry) NewMDWManager(cfg *config.Config, logger logger.Logger) (obj middleware.MDWManager, err error) {
	obj = r.mw.NewMDWManager(cfg, logger)
	return
}

func (r *serviceRegistry) NewLogging(cfg *config.Config, logger logger.Logger) (obj logger.LoggerUC, err error) {
	obj = r.logging.NewLogging(cfg, logger)
	return
}
