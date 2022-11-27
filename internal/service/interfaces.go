package service

import (
	"probabilisticTimeSeriesModeling/internal/credit/controller"
	"probabilisticTimeSeriesModeling/internal/middleware"
	"probabilisticTimeSeriesModeling/pkg/logger"
)

type Controllers struct {
	Credit interface {
		controller.CreditCtrl
	}
}

type MDWManager struct {
	MDWManager interface {
		middleware.MDWManager
	}
}

type Logger struct {
	Logger interface {
		logger.LoggerUC
	}
}
