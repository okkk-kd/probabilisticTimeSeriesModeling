package registry

import (
	"github.com/jmoiron/sqlx"
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
		*sqlx.DB,
	) middleware.MDWManager
}

func NewMDWManagerReg() MDWManager {
	return &mdwManager{}
}

func (mw *mdwManager) NewMDWManager(
	cfg *config.Config,
	logger log.Logger,
	db *sqlx.DB,
) (obj middleware.MDWManager) {
	obj = middleware.NewMDWManager(cfg, logger, db)
	return
}
