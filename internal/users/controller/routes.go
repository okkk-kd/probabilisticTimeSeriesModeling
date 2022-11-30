package controller

import (
	"github.com/gofiber/fiber/v2"
	"probabilisticTimeSeriesModeling/internal/middleware"
)

func UserRoutesGroup(mw middleware.MDWManager, userRout fiber.Router, h UserCtrl) {
	userRout.Post("registration", mw.APIMiddleware(), h.UserRegistration())
	userRout.Post("authorization", mw.APIMiddleware(), h.UserAuthorization())
}
