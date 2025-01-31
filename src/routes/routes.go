package routes

import (
	m "app/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Use(m.LoggerMiddleware())

	// User Routes
	userApi := app.Group("/user")
	SetupUserRoutes(userApi)
}
