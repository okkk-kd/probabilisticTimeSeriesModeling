package registry

import (
	"github.com/jmoiron/sqlx"
	"probabilisticTimeSeriesModeling/config"
	"probabilisticTimeSeriesModeling/internal/session/controller"
	"probabilisticTimeSeriesModeling/internal/session/repository"
	"probabilisticTimeSeriesModeling/internal/session/usecase"
)

type sessionReg struct {
}

type SessionReg interface {
	NewSessionCtrl(cfg *config.Config, pgDB *sqlx.DB) (ctrl controller.SessionCtrl, err error)
}

func NewSessionReg() (obj SessionReg, err error) {
	return &sessionReg{}, err
}

func (c *sessionReg) NewSessionCtrl(cfg *config.Config, pgDB *sqlx.DB) (ctrl controller.SessionCtrl, err error) {
	repo, err := c.NewSessionRepo(cfg, pgDB)
	if err != nil {
		return
	}
	uc, err := c.NewSessionUC(repo)
	if err != nil {
		return
	}
	ctrl, err = controller.NewSessionCtrl(uc)
	if err != nil {
		return
	}
	return
}

func (c *sessionReg) NewSessionUC(repo repository.SessionRepo) (uc usecase.SessionUC, err error) {
	uc, err = usecase.NewSessionUC(repo)
	if err != nil {
		return
	}
	return
}

func (c *sessionReg) NewSessionRepo(cfg *config.Config, pgDB *sqlx.DB) (repo repository.SessionRepo, err error) {
	repo, err = repository.NewSessionRepo(cfg, pgDB)
	if err != nil {
		return
	}
	return
}
