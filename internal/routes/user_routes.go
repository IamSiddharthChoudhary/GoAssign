package routes

import (
	"github.com/gofiber/fiber/v2"

	"assignment/internal/handler"
)

func RegisterUserRoutes(app *fiber.App, h *handler.UserHandler) {
	app.Post("/users", h.CreateUser)
	app.Get("/users", h.ListUsers)
	app.Get("/users/:id", h.GetUserByID)
	app.Put("/users/:id", h.UpdateUser)
	app.Delete("/users/:id", h.DeleteUser)
}
