package registry

import (
	"github.com/jmoiron/sqlx"
	"probabilisticTimeSeriesModeling/config"
	"probabilisticTimeSeriesModeling/internal/users/controller"
	"probabilisticTimeSeriesModeling/internal/users/repository"
	"probabilisticTimeSeriesModeling/internal/users/usecase"
	"probabilisticTimeSeriesModeling/pkg/logger"
)

type userReg struct {
}

type UserReg interface {
	NewUserCtrl(cfg *config.Config, pgDB *sqlx.DB, loggingUC logger.LoggerUC) (ctrl controller.UserCtrl, err error)
}

func NewUserReg() (obj UserReg, err error) {
	return &userReg{}, err
}

func (c *userReg) NewUserCtrl(cfg *config.Config, pgDB *sqlx.DB, loggingUC logger.LoggerUC) (ctrl controller.UserCtrl, err error) {
	repo, err := c.NewUserRepo(cfg, pgDB)
	if err != nil {
		return
	}
	uc, err := c.NewUserUC(repo)
	if err != nil {
		return
	}
	ctrl, err = controller.NewUserCtrl(uc, loggingUC)
	if err != nil {
		return
	}
	return
}

func (c *userReg) NewUserUC(repo repository.UserRepo) (uc usecase.UserUC, err error) {
	uc, err = usecase.NewUserUC(repo)
	if err != nil {
		return
	}
	return
}

func (c *userReg) NewUserRepo(cfg *config.Config, pgDB *sqlx.DB) (repo repository.UserRepo, err error) {
	repo, err = repository.NewUserRepo(cfg, pgDB)
	if err != nil {
		return
	}
	return
}
