package exceptions

import (
	"github.com/gofiber/fiber/v2"
)

var CommonErrors = struct {
	BodyParsingFailed     *CustomError
	BodyValidationFailed  *CustomError
	QueryParsingFailed    *CustomError
	QueryValidationFailed *CustomError
	UserUnauthorized      *CustomError
	UserForbiddenAccess   *CustomError
	DatabaseError         *CustomError
}{
	BodyParsingFailed:     NewCustomError(fiber.StatusBadRequest, "Parsing Failed", "Failed to parse the request body. Please check the data format and types."),
	BodyValidationFailed:  NewCustomError(fiber.StatusBadRequest, "Validation Failed", "Request body validation failed. Ensure all required fields are correctly filled."),
	QueryParsingFailed:    NewCustomError(fiber.StatusBadRequest, "Parsing Failed", "Failed to parse the query parameters. Please check the data format."),
	QueryValidationFailed: NewCustomError(fiber.StatusBadRequest, "Validation Failed", "Query parameter validation failed. Ensure all required parameters are correctly provided."),
	UserUnauthorized:      NewCustomError(fiber.StatusUnauthorized, "Unauthorized", "User is not authorized. Please log in to continue."),
	UserForbiddenAccess:   NewCustomError(fiber.StatusForbidden, "Forbidden", "User does not have permission to access this resource."),
	DatabaseError:         NewCustomError(fiber.StatusInternalServerError, "Database Error", "An error occurred while accessing the database. Please try again later."),
}
