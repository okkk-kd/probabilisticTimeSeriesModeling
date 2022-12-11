package controller

import (
	"github.com/gofiber/fiber/v2"
	"probabilisticTimeSeriesModeling/internal/middleware"
)

func CreditRoutesGroup(mw middleware.MDWManager, creditRout fiber.Router, h CreditCtrl) {
	creditRout.Get("/:code/retrieve_two_columns", mw.NonAuthed(), h.RetrieveTwoColumns())
	creditRout.Get("/:code/:years/bank_forecast", mw.NonAuthed(), h.ForecastingBankData())
	creditRout.Get("/codes_list", mw.NonAuthed(), h.GetCodesList())
	creditRout.Get("/:id/:code/get_code_data", mw.NonAuthed(), h.GetCodeDataByID())
	creditRout.Delete("/:id/:code/code", mw.NonAuthed(), h.DeleteCodeDataByID())
	creditRout.Patch("/code", mw.NonAuthed(), h.UpdateCodeDataByID())
	creditRout.Post("/code_data", mw.NonAuthed(), h.AddCodeData())
	creditRout.Post("/create_table", mw.NonAuthed(), h.CreateCustomUserDataTable())
	creditRout.Post("/add_code_data_list", mw.NonAuthed(), h.AddListCodeData())
}
