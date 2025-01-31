package exceptions

import (
	"github.com/gofiber/fiber/v2"
)

var UserErrors = struct {
	ErrorInGeneratingToken *CustomError
	UserWithEmailExists    *CustomError
}{
	ErrorInGeneratingToken: NewCustomError(fiber.StatusInternalServerError, "Token Generation Error", "An error occurred while generating the token. Please try again later."),
	UserWithEmailExists:    NewCustomError(fiber.StatusConflict, "User Already Exists, Email", "A user with the provided email already exists. Please try again with a different email."),
}
