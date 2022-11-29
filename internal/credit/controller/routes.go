package controller

import (
	"github.com/gofiber/fiber/v2"
	"probabilisticTimeSeriesModeling/internal/middleware"
)

func CreditRoutesGroup(mw middleware.MDWManager, creditRout fiber.Router, h CreditCtrl) {
	creditRout.Get("/:code/retrieve_two_columns", mw.APIMiddleware(), h.RetrieveTwoColumns())
	creditRout.Get("/:code/:years/bank_forecast", mw.APIMiddleware(), h.ForecastingBankData())
	creditRout.Get("/get_code_list", mw.APIMiddleware(), h.GetCodesList())
	creditRout.Get("/:id/:code/get_code_data", mw.APIMiddleware(), h.GetCodeDataByID())
	creditRout.Delete("/:id//:codeget_code_list", mw.APIMiddleware(), h.DeleteCodeDataByID())
	creditRout.Patch("/:id//:codeget_code_list", mw.APIMiddleware(), h.UpdateCodeDataByID())
	creditRout.Post("/:code/code_data", mw.APIMiddleware(), h.AddCodeData())
}
