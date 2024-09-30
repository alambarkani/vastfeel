package server

import "github.com/gofiber/fiber/v2"

func Server(app *fiber.App) {
	app.Listen(":3000")
}