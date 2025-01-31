package user

import (
	"app/src/utils"

	"github.com/gofiber/fiber/v2"
)

func HandleCreateUser(c *fiber.Ctx) error {
		
	return utils.SendResponse(c, fiber.StatusCreated, "", "User created successfully", nil)
}