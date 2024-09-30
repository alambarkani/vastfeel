package main

import (
	"vastfeel-backend/internal/controllers"
	"vastfeel-backend/internal/database"
	"vastfeel-backend/internal/repositories"
	"vastfeel-backend/internal/server"
	"vastfeel-backend/internal/services"
	"vastfeel-backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
)

func main() {
	myValidator := &utils.XValidator{
		Validator: utils.Validate,
	}

	app := fiber.New(fiber.Config{
		AppName: "VastFeel Backend",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(utils.GlobalErrorHandlerResponse{
				Success: false,
				Message: err.Error(),
			})
		},
	})

	app.Use(csrf.New(csrf.Config{
		CookieSecure:   true,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(utils.GlobalErrorHandlerResponse{
				Success: false,
				Message: err.Error(),
			})
		},
	}))

	app.Use(cors.New())

	db, err := database.DBConnect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userService := services.NewUserService(repositories.NewUserRepository(db))
	userController := controllers.NewUserController(userService, myValidator)
	server.UserRoutes(app, userController)
	server.Server(app)
}