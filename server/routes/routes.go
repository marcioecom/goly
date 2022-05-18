package routes

import (
	"goly/server/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/r/:redirect", handler.Redirect)
	app.Get("/goly", handler.GetAllGolies)
	app.Get("/goly/:id", handler.GetGoly)
	app.Post("/goly", handler.CreateGoly)
	app.Put("/goly", handler.UpdateGoly)
	app.Delete("/goly/:id", handler.DeleteGoly)
}
