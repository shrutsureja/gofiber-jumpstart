package user

import (
	"app/src/custom_errors"
	"app/src/database/models"
	"app/src/database/repository"
	"app/src/dtos"
	"app/src/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

func CreateUser(c *fiber.Ctx, dto *dtos.CreateUserDto) error {
	logger := c.Locals("logger").(zerolog.Logger)
	user := new(models.User)
	user.Username = dto.Username
	user.Email = dto.Email

	// Hash the password
	hashedPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		logger.Error().Err(err).Msg("Error hashing password")
		return err
	}
	logger.Trace().Msg("Password hashed successfully")
	user.Password = hashedPassword

	// Save the user
	err = repository.UserRepository.CreateUser(user)
	if custom_errors.DB_DuplicateKeyError(err) {
		logger.Error().Err(err).Msg("User with this email already exists")
		return custom_errors.ConflictError(custom_errors.UserErrorsCodes.UserWithEmailExists, custom_errors.UserErrors.UserWithEmailExists, nil)
	}
	if err != nil {
		logger.Error().Err(err).Msg("Error creating user")
		return err
	}
	logger.Trace().Msg("User created successfully")

	return nil
}
