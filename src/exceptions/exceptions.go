package exceptions

import (
	"app/src/utils"
	"app/src/utils/custom_logger"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var log = custom_logger.GetLogger()

// CustomError is a struct that holds a message, title and error code
type CustomError struct {
	Code    int
	Title   string
	Message string
}

// Implement the `Error` method so CustomError satisfies the `error` interface
func (e *CustomError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

// NewCustomError creates a new instance of CustomError.
// It is commonly used in the service or repository layer to generate custom errors with a specific code, title, and message.
// The generated CustomError is then caught in the centralized ErrorHandler to produce a structured JSON response.
//
// Example
// NewCustomError(fiber.StatusNotFound, "Not Found", "User not found.")
func NewCustomError(code int, title string, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Title:   title,
		Message: message,
	}
}

// ErrorHandler is a centralized error handler for handling and responding to errors in a structured JSON format.
// It checks if the error is a CustomError and responds with the corresponding code, title, and message.
// If the error is not a CustomError, it logs and sends a generic Internal Server Error.
func ErrorHandler(c *fiber.Ctx, err error, data ...interface{}) error {
	// This is a safety check in case `nil` is passed.
	if err == nil {
		return nil
	}

	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	} else {
		responseData = nil
	}

	// Check if the error is a CustomError
	if customErr, ok := err.(*CustomError); ok {
		log.Info().Msgf("RequestId %s, Error: %s, Code: %d, Data: %v", c.Locals("requestId"), customErr.Message, customErr.Code, data)
		return utils.SendResponse(c, customErr.Code, customErr.Title, customErr.Message, responseData)
	}

	// Internal Server Error
	log.Error().Msgf("RequestId %s, Internal Server Error: %v", c.Locals("requestId"), err)
	return utils.SendResponse(c, fiber.StatusInternalServerError, "Error", "Internal Server Error", nil)
}

// DuplicateKeyError checks if the given error corresponds to a duplicate key error in the database.
// Returns true if the error is a duplicate key error, false otherwise.
func DuplicateKeyError(err error) bool {
	return errors.Is(err, gorm.ErrDuplicatedKey)
}

// NotFoundError checks if the given error corresponds to a record not being found in the database.
// Returns true if the error is a record not found error, false otherwise.
func NotFoundError(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
