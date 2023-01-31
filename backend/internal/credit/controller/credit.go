package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"probabilisticTimeSeriesModeling/internal/credit"
	"probabilisticTimeSeriesModeling/internal/credit/usecase"
	"probabilisticTimeSeriesModeling/pkg/logger"
	"probabilisticTimeSeriesModeling/pkg/utils/reqvalidator"
	"strings"
	"time"
)

type creditCtrl struct {
	creditUC  usecase.CreditUC
	loggingUC logger.LoggerUC
}

type CreditCtrl interface {
	RetrieveTwoColumns() fiber.Handler
	ForecastingBankData() fiber.Handler
	GetCodesList() fiber.Handler
	GetCodeDataByID() fiber.Handler
	DeleteCodeDataByID() fiber.Handler
	UpdateCodeDataByID() fiber.Handler
	AddCodeData() fiber.Handler
	CreateCustomUserDataTable() fiber.Handler
	AddListCodeData() fiber.Handler
	GetDataFromTableByCode() fiber.Handler
}

func NewCreditCtrl(creditUC usecase.CreditUC, loggingUC logger.LoggerUC) (obj CreditCtrl, err error) {
	return &creditCtrl{
		creditUC:  creditUC,
		loggingUC: loggingUC,
	}, err
}

func (ctrl *creditCtrl) RetrieveTwoColumns() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params credit.ForecastingBankDataRequest
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		params.Code = strings.ToLower(ctx.Params("code", params.Code))
		result, err := ctrl.creditUC.RetrieveTwoColumns(params.Code)
		if err != nil {
			err = errors.Wrapf(err, "creditCtrl.RetrieveTwoColumns()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (ctrl *creditCtrl) GetDataFromTableByCode() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var code string
		code = strings.ToLower(ctx.Params("code", code))
		result, err := ctrl.creditUC.GetDataTablesFromDBByCode(code)
		if err != nil {
			err = errors.Wrapf(err, "creditCtrl.GetDataFromTableByCode()")
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
		params.Code = strings.ToLower(ctx.Params("code", params.Code))
		result, err := ctrl.creditUC.ForecastingBankData(params)
		if err != nil {
			err = errors.Wrapf(err, "creditCtrl.ForecastingBankData()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (ctrl *creditCtrl) GetCodesList() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		result, err := ctrl.creditUC.GetCodesList(ctx.Context())
		if err != nil {
			err = errors.Wrapf(err, "creditCtrl.GetCodesList()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (ctrl *creditCtrl) GetCodeDataByID() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params credit.GetCodeDataByID
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		params.ID = ctx.Params("id", params.ID)
		params.Code = strings.ToLower(ctx.Params("code", params.Code))
		result, err := ctrl.creditUC.GetCodeDataByID(ctx.Context(), params)
		if err != nil {
			err = errors.Wrapf(err, "creditCtrl.GetCodeDataByID()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (ctrl *creditCtrl) DeleteCodeDataByID() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params credit.DeleteCodeDataByID
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		params.ID = ctx.Params("id", params.ID)
		params.Code = strings.ToLower(ctx.Params("code", params.Code))
		err := ctrl.creditUC.DeleteCodeDataByID(ctx.Context(), params)
		if err != nil {
			err = errors.Wrapf(err, "creditCtrl.DeleteCodeDataByID()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.SendStatus(fiber.StatusOK)
	}
}

func (ctrl *creditCtrl) UpdateCodeDataByID() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params credit.UpdateCodeDataByID
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		err := reqvalidator.ReadRequest(ctx, &params)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		err = ctrl.creditUC.UpdateCodeDataByID(ctx.Context(), params)
		if err != nil {
			err = errors.Wrapf(err, "creditCtrl.UpdateCodeDataByID()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.SendStatus(fiber.StatusOK)
	}
}

func (ctrl *creditCtrl) AddCodeData() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params credit.AddCodeData
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		err := reqvalidator.ReadRequest(ctx, &params)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		err = ctrl.creditUC.AddCodeData(ctx.Context(), params)
		if err != nil {
			err = errors.Wrapf(err, "creditCtrl.AddCodeData()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.SendStatus(fiber.StatusOK)
	}
}

func (ctrl *creditCtrl) CreateCustomUserDataTable() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params credit.CreateCustomUserDataTable
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		err := reqvalidator.ReadRequest(ctx, &params)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		err = ctrl.creditUC.CreateCustomUserDataTable(params.DBName)
		if err != nil {
			err = errors.Wrapf(err, "creditCtrl.CreateCustomUserDataTable()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.SendStatus(fiber.StatusOK)
	}
}

func (ctrl *creditCtrl) AddListCodeData() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params credit.AddListCodeData
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		err := reqvalidator.ReadRequest(ctx, &params)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		params.Code = strings.ToLower(ctx.Params("code", params.Code))
		err = ctrl.creditUC.AddListCodeData(ctx.Context(), params)
		if err != nil {
			err = errors.Wrapf(err, "creditCtrl.AddListCodeData()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.SendStatus(fiber.StatusOK)
	}
}
