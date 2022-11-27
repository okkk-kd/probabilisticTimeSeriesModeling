package repository

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"probabilisticTimeSeriesModeling/config"
	"probabilisticTimeSeriesModeling/internal/credit"
	"probabilisticTimeSeriesModeling/pkg/fhttp"
)

type creditRepo struct {
	cfg         *config.Config
	fhttpClient *fhttp.Client
}

type CreditRepo interface {
	RetrieveTwoColumns() (credit.Dataset, error)
}

func NewCreditRepo(cfg *config.Config, fhttpClient *fhttp.Client) (obj CreditRepo, err error) {
	return &creditRepo{
		cfg:         cfg,
		fhttpClient: fhttpClient,
	}, err
}

func (repo *creditRepo) RetrieveTwoColumns() (response credit.Dataset, err error) {
	url := repo.cfg.Server.NasdaqURL + retrieveTwoColumns + nasdaqDataLinkCode
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
