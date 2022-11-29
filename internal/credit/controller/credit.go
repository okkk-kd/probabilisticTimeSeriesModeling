package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"probabilisticTimeSeriesModeling/internal/credit"
	"probabilisticTimeSeriesModeling/internal/credit/usecase"
	"probabilisticTimeSeriesModeling/pkg/logger"
	"time"
)

type creditCtrl struct {
	creditUC  usecase.CreditUC
	loggingUC logger.LoggerUC
}

type CreditCtrl interface {
	RetrieveTwoColumns() fiber.Handler
	ForecastingBankData() fiber.Handler
}

func NewCreditCtrl(creditUC usecase.CreditUC, loggingUC logger.LoggerUC) (obj CreditCtrl, err error) {
	return &creditCtrl{
		creditUC:  creditUC,
		loggingUC: loggingUC,
	}, err
}

func (ctrl *creditCtrl) RetrieveTwoColumns() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		result, err := ctrl.creditUC.RetrieveTwoColumns()
		if err != nil {
			err = errors.Wrapf(err, "creditCtrl.RetrieveTwoColumns()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (ctrl *creditCtrl) ForecastingBankData() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params credit.ForecastingBankDataRequest
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		params.Years = ctx.Params("years", params.Years)
		result, err := ctrl.creditUC.ForecastingBankData(params)
		if err != nil {
			err = errors.Wrapf(err, "creditCtrl.RetrieveTwoColumns()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}
