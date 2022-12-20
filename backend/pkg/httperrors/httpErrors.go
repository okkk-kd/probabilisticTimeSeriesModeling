package httperrors

import (
	"fmt"
	"probabilisticTimeSeriesModeling/config"
	"probabilisticTimeSeriesModeling/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type responseMsg struct {
	Message string `json:"message"`
}

func Init(cfg *config.Config, logger logger.Logger) func(c *fiber.Ctx, err error) error {
	return func(ctx *fiber.Ctx, err error) error {
		var (
			response   responseMsg
			statusCode int
		)

		if statusCode == 0 {
			statusCode = fiber.StatusInternalServerError
		}

		if response.Message == "" && cfg.Server.ShowUnknownErrorsInResponse {
			response.Message = fmt.Sprintf("hidden: %s", err.Error())
		} else if response.Message == "" {
			logger.Error(fmt.Errorf("%s %v", ctx.OriginalURL(), err))
			response.Message = "unknown error"
		}

		return ctx.Status(statusCode).JSON(response)
	}
}
