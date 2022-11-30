package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"probabilisticTimeSeriesModeling/internal/users"
	"probabilisticTimeSeriesModeling/internal/users/usecase"
	"probabilisticTimeSeriesModeling/pkg/logger"
	"probabilisticTimeSeriesModeling/pkg/utils/reqvalidator"
	"time"
)

type userCtrl struct {
	uc        usecase.UserUC
	loggingUC logger.LoggerUC
}

type UserCtrl interface {
	CreateUser() fiber.Handler
	UpdateUserPassword() fiber.Handler
	Authorization() fiber.Handler
}

func NewUserCtrl(uc usecase.UserUC, loggingUC logger.LoggerUC) (obj UserCtrl, err error) {
	return &userCtrl{
		uc:        uc,
		loggingUC: loggingUC,
	}, err
}

func (ctrl *userCtrl) CreateUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params users.CreateUser
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		err := reqvalidator.ReadRequest(ctx, &params)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		err = ctrl.uc.CreateUser(params)
		if err != nil {
			err = errors.Wrapf(err, "userCtrl.CreateUser()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.SendStatus(fiber.StatusOK)
	}
}

func (ctrl *userCtrl) UpdateUserPassword() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params users.UpdateUserPassword
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		err := reqvalidator.ReadRequest(ctx, &params)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		err = ctrl.uc.UpdateUserPassword(params)
		if err != nil {
			err = errors.Wrapf(err, "userCtrl.UpdateUserPassword()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.SendStatus(fiber.StatusOK)
	}
}

func (ctrl *userCtrl) Authorization() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params users.Authorization
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		err := reqvalidator.ReadRequest(ctx, &params)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		result, err := ctrl.uc.Authorization(params)
		if err != nil {
			err = errors.Wrapf(err, "userCtrl.Authorization()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}
