package routes

import (
	c "app/src/controllers/user"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	// User Routes
	router.Post("/", c.HandleCreateUser)
}
