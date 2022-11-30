package controller

import (
	"github.com/gofiber/fiber/v2"
	"probabilisticTimeSeriesModeling/internal/middleware"
)

func UserRoutesGroup(mw middleware.MDWManager, userRout fiber.Router, h UserCtrl) {
	userRout.Post("/registration", mw.NonAuthed(), h.CreateUser())
	userRout.Post("/authorization", mw.NonAuthed(), h.Authorization())
	userRout.Post("/update_password", mw.APIMiddleware(), h.UpdateUserPassword())
}
