package usecase

import (
	"encoding/json"
	"probabilisticTimeSeriesModeling/internal/credit"
	"probabilisticTimeSeriesModeling/internal/credit/repository"
)

type creditUC struct {
	repo repository.CreditRepo
}

type CreditUC interface {
	RetrieveTwoColumns() (credit.Dataset, error)
}

func NewCreditUC(repo repository.CreditRepo) (obj CreditUC, err error) {
	return &creditUC{
		repo: repo,
	}, err
}

func (uc *creditUC) RetrieveTwoColumns() (response credit.Dataset, err error) {
	response, err = uc.repo.RetrieveTwoColumns()
	if err != nil {
		return
	}
	return
}

func (uc *creditUC) ForecastingBankData() (response credit.BankForecast, err error) {
	var bankData credit.BankElement
	rawData, err := uc.RetrieveTwoColumns()
	if err != nil {
		return
	}
	for _, obj := range rawData.Datasett.Data {
		for _, el := range obj {
			bankData =
		}
	}
	return
}
