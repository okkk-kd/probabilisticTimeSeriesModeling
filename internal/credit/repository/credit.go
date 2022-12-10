package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"probabilisticTimeSeriesModeling/config"
	"probabilisticTimeSeriesModeling/internal/credit"
	"probabilisticTimeSeriesModeling/pkg/fhttp"
	"probabilisticTimeSeriesModeling/pkg/forecast"
	"probabilisticTimeSeriesModeling/pkg/utils"
)

type creditRepo struct {
	cfg         *config.Config
	fhttpClient *fhttp.Client
	db          *sqlx.DB
}

type CreditRepo interface {
	RetrieveTwoColumns(string) (credit.Dataset, error)
	InitCodeLists() (err error)
	GetCodesList(ctx context.Context) (result []credit.GetCodesListResponse, err error)
	GetCodeDataByID(ctx context.Context, params credit.GetCodeDataByID) (result credit.GetCodeDataByIDResponse, err error)
	DeleteCodeDataByID(ctx context.Context, params credit.DeleteCodeDataByID) (err error)
	UpdateCodeDataByID(ctx context.Context, params credit.UpdateCodeDataByID) (err error)
	AddCodeData(ctx context.Context, params credit.AddCodeData) (err error)

	GetDataFromTableByCode(params credit.ForecastingBankDataRequest) (data []forecast.ForecastEl, err error)
	CreateCustomUserDataTable(dbName string) (err error)
	AddListCodeData(ctx context.Context, params credit.AddListCodeData) (err error)
}

func NewCreditRepo(cfg *config.Config, fhttpClient *fhttp.Client, pgDB *sqlx.DB) (obj CreditRepo, err error) {
	return &creditRepo{
		cfg:         cfg,
		fhttpClient: fhttpClient,
		db:          pgDB,
	}, err
}

func (repo *creditRepo) RetrieveTwoColumns(code string) (response credit.Dataset, err error) {
	url := repo.cfg.Server.NasdaqURL + retrieveTwoColumns + nasdaqDataLinkCode + code
	method := fiber.MethodGet

	var headers = make(map[string]string)
	headers[fiber.HeaderContentType] = fiber.MIMEApplicationJSON
	headers[fiber.HeaderAccept] = fiber.MIMEApplicationJSON

	queryParams := make(map[string]string)
	queryParams["column_index"] = "1"
	queryParams["api_key"] = repo.cfg.Server.APIKey

	responseBody, statusCode, err := repo.fhttpClient.Request(method, url, nil, queryParams, headers)
	if err != nil {
		return response, err
	}

	if statusCode != 200 {
		return response, errors.Wrapf(err, "Wrong status code = %d | creditRepo.RetrieveTwoColumns()",
			statusCode)
	}

	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return response, err
	}

	return response, err
}

func (repo *creditRepo) GetCodesList(ctx context.Context) (result []credit.GetCodesListResponse, err error) {
	err = repo.db.SelectContext(ctx, &result, queryGetCodesList)
	if err != nil {
		return
	}
	return
}

func (repo *creditRepo) InitCodeLists() (err error) {
	codeList, err := repo.GetCodesList(context.Background())
	if err != nil {
		return
	}
	tx, err := repo.db.Begin()
	for _, el := range codeList {
		dataSet, err := repo.RetrieveTwoColumns(el.Code)
		if err != nil {
			errR := tx.Rollback()
			if errR != nil {
				return errors.Wrapf(err, errR.Error())
			}
			return err
		}
		if _, err := repo.db.Exec(fmt.Sprintf(queryCreateCodesDataTable, el.Code, el.Code, el.Code, el.Code, el.Code)); err != nil {
			errR := tx.Rollback()
			if errR != nil {
				return errors.Wrapf(err, errR.Error())
			}
			return err
		}
		for _, data := range dataSet.Datasett.Data {
			convertedData, err := utils.ConvertToBankData(data)
			if err != nil {
				errR := tx.Rollback()
				if errR != nil {
					return errors.Wrapf(err, errR.Error())
				}
				return err
			}
			if _, err := tx.Exec(fmt.Sprintf(queryInsertDataIntoCodesDataTable, el.Code), convertedData.Price, convertedData.Date); err != nil {
				errR := tx.Rollback()
				if errR != nil {
					return errors.Wrapf(err, errR.Error())
				}
				return err
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		return
	}
	return
}

func (repo *creditRepo) GetCodeDataByID(ctx context.Context, params credit.GetCodeDataByID) (result credit.GetCodeDataByIDResponse, err error) {
	err = repo.db.GetContext(ctx, &result, fmt.Sprintf(queryGetCodeDataByID, params.Code), params.ID)
	if err != nil {
		return
	}
	return
}

func (repo *creditRepo) DeleteCodeDataByID(ctx context.Context, params credit.DeleteCodeDataByID) (err error) {
	_, err = repo.db.ExecContext(ctx, fmt.Sprintf(queryDeleteCodeDataByID, params.Code), params.ID)
	if err != nil {
		return
	}
	return
}

func (repo *creditRepo) UpdateCodeDataByID(ctx context.Context, params credit.UpdateCodeDataByID) (err error) {
	_, err = repo.db.ExecContext(ctx, fmt.Sprintf(queryUpdateCodeDataByID, params.Code), params.Amount, params.Date, params.ID)
	if err != nil {
		return
	}
	return
}

func (repo *creditRepo) AddCodeData(ctx context.Context, params credit.AddCodeData) (err error) {
	_, err = repo.db.ExecContext(ctx, fmt.Sprintf(queryAddCodeData, params.Code), params.Amount, params.Date)
	if err != nil {
		return
	}
	return
}

func (repo *creditRepo) AddListCodeData(ctx context.Context, params credit.AddListCodeData) (err error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return
	}
	for _, el := range params.Data {
		_, err = tx.ExecContext(ctx, fmt.Sprintf(queryAddCodeData, params.Code), el.Amount, el.Date)
		if err != nil {
			errR := tx.Rollback()
			if errR != nil {
				return errors.Wrapf(err, errR.Error())
			}
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return
	}
	return
}

func (repo *creditRepo) CreateCustomUserDataTable(dbName string) (err error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return
	}
	if _, err := tx.Exec(fmt.Sprintf(queryCreateCodesDataTable, dbName, dbName, dbName, dbName, dbName)); err != nil {
		errR := tx.Rollback()
		if errR != nil {
			return errors.Wrapf(err, errR.Error())
		}
		return err
	}
	if _, err := tx.Exec(queryInsertTableIntoTableList, dbName, dbName); err != nil {
		errR := tx.Rollback()
		if errR != nil {
			return errors.Wrapf(err, errR.Error())
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		return
	}
	return
}

func (repo *creditRepo) GetDataFromTableByCode(params credit.ForecastingBankDataRequest) (data []forecast.ForecastEl, err error) {
	err = repo.db.Select(&data, fmt.Sprintf(querySelectDataByCode, params.Code))
	if err != nil {
		return
	}
	return
}
