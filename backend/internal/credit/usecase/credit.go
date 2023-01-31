package usecase

import (
	"context"
	"probabilisticTimeSeriesModeling/internal/credit"
	"probabilisticTimeSeriesModeling/internal/credit/repository"
	"probabilisticTimeSeriesModeling/pkg/forecast"
	"probabilisticTimeSeriesModeling/pkg/utils"
	"sort"
)

type creditUC struct {
	repo repository.CreditRepo
}

type CreditUC interface {
	RetrieveTwoColumns(code string) (credit.Dataset, error)
	ForecastingBankData(params credit.ForecastingBankDataRequest) (response *forecast.BankForecast, err error)
	GetCodesList(ctx context.Context) (result []credit.GetCodesListResponse, err error)
	GetCodeDataByID(ctx context.Context, params credit.GetCodeDataByID) (result credit.GetCodeDataByIDResponse, err error)
	DeleteCodeDataByID(ctx context.Context, params credit.DeleteCodeDataByID) (err error)
	UpdateCodeDataByID(ctx context.Context, params credit.UpdateCodeDataByID) (err error)
	AddCodeData(ctx context.Context, params credit.AddCodeData) (err error)
	CreateCustomUserDataTable(dbName string) (err error)
	AddListCodeData(ctx context.Context, params credit.AddListCodeData) (err error)
	GetDataTablesFromDBByCode(code string) (response []forecast.ForecastEl, err error)
}

func NewCreditUC(repo repository.CreditRepo) (obj CreditUC, err error) {
	return &creditUC{
		repo: repo,
	}, err
}

func (uc *creditUC) RetrieveTwoColumns(code string) (response credit.Dataset, err error) {
	response, err = uc.repo.RetrieveTwoColumns(code)
	if err != nil {
		return
	}
	return
}

func (uc *creditUC) GetDataTablesFromDBByCode(code string) (response []forecast.ForecastEl, err error) {
	response, err = uc.repo.GetDataFromTable(code)
	if err != nil {
		return
	}
	return
}

func (uc *creditUC) ForecastingBankData(params credit.ForecastingBankDataRequest) (response *forecast.BankForecast, err error) {
	var bankData []forecast.ForecastEl

	rawData, err := uc.RetrieveTwoColumns(params.Code)
	if err != nil {
		return
	}
	if rawData.Datasett.Data == nil {
		bankData, err = uc.repo.GetDataFromTableByCode(params)
		if err != nil {
			return
		}
	}
	for _, obj := range rawData.Datasett.Data {
		bank, err := utils.ConvertToBankData(obj)
		if err != nil {
			return response, err
		}
		bankData = append(bankData, bank)
	}
	sort.SliceStable(bankData, func(i, j int) bool {
		return bankData[i].Date.Before(bankData[j].Date)
	})
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

func (uc *creditUC) GetCodesList(ctx context.Context) (result []credit.GetCodesListResponse, err error) {
	result, err = uc.repo.GetCodesList(ctx)
	if err != nil {
		return
	}
	return
}

func (uc *creditUC) GetCodeDataByID(ctx context.Context, params credit.GetCodeDataByID) (result credit.GetCodeDataByIDResponse, err error) {
	result, err = uc.repo.GetCodeDataByID(ctx, params)
	if err != nil {
		return
	}
	return
}

func (uc *creditUC) DeleteCodeDataByID(ctx context.Context, params credit.DeleteCodeDataByID) (err error) {
	err = uc.repo.DeleteCodeDataByID(ctx, params)
	if err != nil {
		return
	}
	return
}

func (uc *creditUC) UpdateCodeDataByID(ctx context.Context, params credit.UpdateCodeDataByID) (err error) {
	err = uc.repo.UpdateCodeDataByID(ctx, params)
	if err != nil {
		return
	}
	return
}

func (uc *creditUC) AddCodeData(ctx context.Context, params credit.AddCodeData) (err error) {
	err = uc.repo.AddCodeData(ctx, params)
	if err != nil {
		return
	}
	return
}

func (uc *creditUC) CreateCustomUserDataTable(dbName string) (err error) {
	err = uc.repo.CreateCustomUserDataTable(dbName)
	if err != nil {
		return
	}
	return
}
func (uc *creditUC) AddListCodeData(ctx context.Context, params credit.AddListCodeData) (err error) {
	err = uc.repo.AddListCodeData(ctx, params)
	if err != nil {
		return
	}
	return
}
