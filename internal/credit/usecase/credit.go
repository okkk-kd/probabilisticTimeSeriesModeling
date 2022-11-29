package usecase

import (
	"fmt"
	"probabilisticTimeSeriesModeling/internal/credit"
	"probabilisticTimeSeriesModeling/internal/credit/repository"
	"probabilisticTimeSeriesModeling/pkg/forecast"
	"strconv"
)

type creditUC struct {
	repo repository.CreditRepo
}

type CreditUC interface {
	RetrieveTwoColumns() (credit.Dataset, error)
	ForecastingBankData(params credit.ForecastingBankDataRequest) (response *forecast.BankForecast, err error)
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

func (uc *creditUC) ConvertToBankData(before []interface{}) (result forecast.ForecastEl, err error) {
	result.Data = fmt.Sprintf("%v", before[0])
	price, err := strconv.ParseFloat(fmt.Sprintf("%v", before[1]), 64)
	if err != nil {
		return
	}
	result.Price = price
	return
}

func (uc *creditUC) ForecastingBankData(params credit.ForecastingBankDataRequest) (response *forecast.BankForecast, err error) {
	var bankData []forecast.ForecastEl

	rawData, err := uc.RetrieveTwoColumns()
	if err != nil {
		return
	}
	for _, obj := range rawData.Datasett.Data {
		bank, err := uc.ConvertToBankData(obj)
		if err != nil {
			return response, err
		}
		bankData = append(bankData, bank)
	}
	fore, err := forecast.NewForecast()
	if err != nil {
		return
	}
	response, err = fore.ForecastingBankData(bankData, params)
	if err != nil {
		return
	}
	return
}
