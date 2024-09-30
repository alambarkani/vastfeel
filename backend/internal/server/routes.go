package server

import (
	"vastfeel-backend/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, user *controllers.UserController) {
	api := app.Group("/api")
	v1 := api.Group("v1")

	v1.Get("/users", user.GetAllUser())
	v1.Get("/users/:id", user.GetUserById())
	v1.Post("/users", user.CreateUser())
	v1.Put("/users/:id", user.UpdateUser())
	v1.Delete("/users/:id", user.DeleteUser())
}