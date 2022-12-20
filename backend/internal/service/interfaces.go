package service

import (
	creditCtrl "probabilisticTimeSeriesModeling/internal/credit/controller"
	"probabilisticTimeSeriesModeling/internal/middleware"
	sessionCtrl "probabilisticTimeSeriesModeling/internal/session/controller"
	userCtrl "probabilisticTimeSeriesModeling/internal/users/controller"
	"probabilisticTimeSeriesModeling/pkg/logger"
)

type Controllers struct {
	Credit interface {
		creditCtrl.CreditCtrl
	}
	User interface {
		userCtrl.UserCtrl
	}
	Session interface {
		sessionCtrl.SessionCtrl
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
