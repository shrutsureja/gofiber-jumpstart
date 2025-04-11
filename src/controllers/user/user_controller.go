package user

import (
	"app/src/custom_errors"
	"app/src/dtos"
	"app/src/utils"

	"github.com/gofiber/fiber/v2"
)

func HandleCreateUser(c *fiber.Ctx) error {
	dto := new(dtos.CreateUserDto)
	if err := c.BodyParser(dto); err != nil {
		return custom_errors.BadRequestError(custom_errors.CommonErrorsCodes.BodyParsingFailed, custom_errors.CommonErrors.BodyParsingFailed, nil)
	}
	if validationMessage, err := utils.ValidateStruct(dto, dtos.CreateUserDtoValidationMessages); err != nil {
		return custom_errors.BadRequestError(custom_errors.CommonErrorsCodes.BodyValidationFailed, custom_errors.CommonErrors.BodyValidationFailed, validationMessage)
	}

	err := CreateUser(c, dto)
	if err != nil {
		return err
	}

	return utils.SendResponse(c, fiber.StatusCreated, "User created successfully", nil)
}
