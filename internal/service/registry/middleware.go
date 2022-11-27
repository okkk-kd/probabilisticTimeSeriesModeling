package registry

import (
	"probabilisticTimeSeriesModeling/config"
	"probabilisticTimeSeriesModeling/internal/middleware"
	log "probabilisticTimeSeriesModeling/pkg/logger"
)

type mdwManager struct {
}

type MDWManager interface {
	NewMDWManager(
		*config.Config,
		log.Logger,
	) middleware.MDWManager
}

func NewMDWManagerReg() MDWManager {
	return &mdwManager{}
}

func (mw *mdwManager) NewMDWManager(
	cfg *config.Config,
	logger log.Logger,
) (obj middleware.MDWManager) {
	obj = middleware.NewMDWManager(cfg, logger)
	return
}
