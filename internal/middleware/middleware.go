package middleware

import (
	"github.com/gofiber/fiber/v2"
	"probabilisticTimeSeriesModeling/config"
	"probabilisticTimeSeriesModeling/pkg/logger"
)

type mdwManager struct {
	cfg    *config.Config
	logger logger.Logger
}

type MDWManager interface {
	APIMiddleware() fiber.Handler
}

func NewMDWManager(
	cfg *config.Config,
	logger logger.Logger,
) MDWManager {
	return &mdwManager{
		cfg:    cfg,
		logger: logger,
	}
}

func (mw *mdwManager) APIMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Get("API-Key") != mw.cfg.Server.APIKey {
			mw.logger.Warn(" APIMiddleware wrong API Key")
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		return c.Next()
	}
}
