package registry

import (
	"github.com/jmoiron/sqlx"
	"probabilisticTimeSeriesModeling/config"
	"probabilisticTimeSeriesModeling/internal/credit/controller"
	"probabilisticTimeSeriesModeling/internal/credit/repository"
	"probabilisticTimeSeriesModeling/internal/credit/usecase"
	"probabilisticTimeSeriesModeling/pkg/fhttp"
	"probabilisticTimeSeriesModeling/pkg/logger"
)

type creditReg struct {
}

type CreditReg interface {
	NewCreditCtrl(*config.Config, *fhttp.Client, logger.LoggerUC, *sqlx.DB) (controller.CreditCtrl, error)
}

func NewCreditReg() (obj CreditReg, err error) {
	return &creditReg{}, err
}

func (c *creditReg) NewCreditCtrl(cfg *config.Config, fhttpClient *fhttp.Client, log logger.LoggerUC, pgDB *sqlx.DB) (ctrl controller.CreditCtrl, err error) {
	repo, err := c.NewCreditRepo(cfg, fhttpClient, pgDB)
	if err != nil {
		return
	}
	uc, err := c.NewCreditUC(repo)
	if err != nil {
		return
	}
	ctrl, err = controller.NewCreditCtrl(uc, log)
	if err != nil {
		return
	}
	return
}

func (c *creditReg) NewCreditUC(repo repository.CreditRepo) (uc usecase.CreditUC, err error) {
	uc, err = usecase.NewCreditUC(repo)
	if err != nil {
		return
	}
	return
}

func (c *creditReg) NewCreditRepo(cfg *config.Config, fhttpClient *fhttp.Client, pgDB *sqlx.DB) (repo repository.CreditRepo, err error) {
	repo, err = repository.NewCreditRepo(cfg, fhttpClient, pgDB)
	if err != nil {
		return
	}
	err = repo.InitCodeLists()
	if err != nil {
		return
	}
	return
}
