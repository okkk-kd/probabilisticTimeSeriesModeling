package registry

import (
	"probabilisticTimeSeriesModeling/config"
	"probabilisticTimeSeriesModeling/pkg/logger"
)

type logging struct {
}

type Logging interface {
	NewLogging(cfg *config.Config, log logger.Logger) (obj logger.LoggerUC)
}

func NewLoggingReg() Logging {
	return &logging{}
}

func (l *logging) NewLogging(cfg *config.Config, log logger.Logger) (obj logger.LoggerUC) {
	obj = logger.NewLoggerUC(cfg, log)
	return
}
