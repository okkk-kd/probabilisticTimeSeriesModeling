package controller

import (
	"github.com/gofiber/fiber/v2"
	"probabilisticTimeSeriesModeling/internal/middleware"
)

func CreditRoutesGroup(mw middleware.MDWManager, creditRout fiber.Router, h CreditCtrl) {
	creditRout.Get("/retrieve_two_columns", mw.APIMiddleware(), h.RetrieveTwoColumns())
}
