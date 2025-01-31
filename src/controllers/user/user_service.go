package user

import (
	"app/src/database/models"
	"app/src/database/repository"
	"app/src/dtos"
	"app/src/exceptions"
	"app/src/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx, dto *dtos.CreateUserDto) error {
	// Create a user
	user := new(models.User)
	user.Username = dto.Username
	user.Email = dto.Email

	// Hash the password
	hashedPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	// Save the user
	err = repository.UserRepository.CreateUser(user)
	if exceptions.DuplicateKeyError(err) {
		return exceptions.UserErrors.UserWithEmailExists
	}
	if err != nil {
		return err
	}
	return nil
}
