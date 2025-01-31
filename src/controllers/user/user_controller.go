package user

import (
	"app/src/dtos"
	"app/src/exceptions"
	"app/src/utils"

	"github.com/gofiber/fiber/v2"
)

func HandleCreateUser(c *fiber.Ctx) error {
	dto := new(dtos.CreateUserDto)
	if err := c.BodyParser(dto); err != nil {
		return exceptions.ErrorHandler(c, exceptions.CommonErrors.BodyParsingFailed)
	}
	if validationMessage, err := utils.ValidateStruct(dto, dtos.CreateUserDtoValidationMessages); err != nil {
		return exceptions.ErrorHandler(c, exceptions.CommonErrors.BodyValidationFailed, validationMessage)
	}

	err := CreateUser(c, dto)
	if err != nil {
		return exceptions.ErrorHandler(c, err)
	}

	return utils.SendResponse(c, fiber.StatusCreated, "", "User created successfully", nil)
}
