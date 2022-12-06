package controller

import (
	"github.com/gofiber/fiber/v2"
	"probabilisticTimeSeriesModeling/internal/middleware"
)

func CreditRoutesGroup(mw middleware.MDWManager, creditRout fiber.Router, h CreditCtrl) {
	creditRout.Get("/:code/retrieve_two_columns", mw.APIMiddleware(), h.RetrieveTwoColumns())
	creditRout.Get("/:code/:years/bank_forecast", mw.APIMiddleware(), h.ForecastingBankData())
	creditRout.Get("/bank_forecast", mw.APIMiddleware(), h.GetCodesList())
	creditRout.Get("/:id/:code/get_code_data", mw.APIMiddleware(), h.GetCodeDataByID())
	creditRout.Delete("/:id/:code/code", mw.APIMiddleware(), h.DeleteCodeDataByID())
	creditRout.Patch("/code", mw.APIMiddleware(), h.UpdateCodeDataByID())
	creditRout.Post("/code_data", mw.APIMiddleware(), h.AddCodeData())
	creditRout.Post("/create_table", mw.APIMiddleware(), h.CreateCustomUserDataTable())
	creditRout.Post("/code_data_list", mw.APIMiddleware(), h.AddListCodeData())
}
