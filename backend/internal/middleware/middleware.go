package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"probabilisticTimeSeriesModeling/config"
	"probabilisticTimeSeriesModeling/pkg/logger"
	"probabilisticTimeSeriesModeling/pkg/utils"
)

type mdwManager struct {
	cfg    *config.Config
	logger logger.Logger
	db     *sqlx.DB
}

type MDWManager interface {
	APIMiddleware() fiber.Handler
	NonAuthed() fiber.Handler
}

func NewMDWManager(
	cfg *config.Config,
	logger logger.Logger,
	pgDB *sqlx.DB,
) MDWManager {
	return &mdwManager{
		cfg:    cfg,
		logger: logger,
		db:     pgDB,
	}
}

func (mw *mdwManager) APIMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Get("API-Key") != mw.cfg.Server.APIKey {
			mw.logger.Warn(" APIMiddleware wrong API Key")
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		sessionKey := c.Get("Authorization")
		authed, err := utils.ValidateSession(mw.db, sessionKey)
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		if authed {
			return c.Next()
		} else {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
	}
}

func (mw *mdwManager) NonAuthed() fiber.Handler {
	return func(c *fiber.Ctx) error {
		//if c.Get("API-Key") != mw.cfg.Server.APIKey {
		//	mw.logger.Warn(" APIMiddleware wrong API Key")
		//	return c.SendStatus(fiber.StatusUnauthorized)
		//}
		return c.Next()
	}
}
