package custom_errors

import (
	"app/src/utils"
	"app/src/utils/custom_logger"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

var log = custom_logger.GetLogger()

// CustomError is a struct that holds a message, title and error code
type CustomError struct {
	StatusCode int
	ErrorCode  string
	Message    string
	Details    any
}

// Implement the `Error` method so CustomError satisfies the `error` interface
func (e *CustomError) Error() string {
	return fmt.Sprintf("StatusCode: %d, ErrorCode: %s, Message: %s", e.StatusCode, e.ErrorCode, e.Message)
}

// Helper Functions for Different Error Types
func BadRequestError(errorCode, message string, details any) *CustomError {
	return &CustomError{
		StatusCode: fiber.StatusBadRequest,
		ErrorCode:  errorCode,
		Message:    message,
		Details:    details,
	}
}

func BadGatewayError(errorCode, message string, details any) *CustomError {
	return &CustomError{
		StatusCode: fiber.StatusBadGateway,
		ErrorCode:  errorCode,
		Message:    message,
		Details:    details,
	}
}

func NotFoundError(errorCode, message string, details any) *CustomError {
	return &CustomError{
		StatusCode: fiber.StatusNotFound,
		ErrorCode:  errorCode,
		Message:    message,
		Details:    details,
	}
}

func UnauthorizedError(errorCode, message string, details any) *CustomError {
	return &CustomError{
		StatusCode: fiber.StatusUnauthorized,
		ErrorCode:  errorCode,
		Message:    message,
		Details:    details,
	}
}

func ForbiddenError(errorCode, message string, details any) *CustomError {
	return &CustomError{
		StatusCode: fiber.StatusForbidden,
		ErrorCode:  errorCode,
		Message:    message,
		Details:    details,
	}
}

func ConflictError(errorCode, message string, details any) *CustomError {
	return &CustomError{
		StatusCode: fiber.StatusConflict,
		ErrorCode:  errorCode,
		Message:    message,
		Details:    details,
	}
}

func InternalServerError(errorCode, message string, details any) *CustomError {
	return &CustomError{
		StatusCode: fiber.StatusInternalServerError,
		ErrorCode:  errorCode,
		Message:    message,
		Details:    details,
	}
}

// DB_DuplicateKeyError checks if the given error corresponds to a duplicate key error in the database.
// Returns true if the error is a duplicate key error, false otherwise.
func DB_DuplicateKeyError(err error) bool {
	return errors.Is(err, gorm.ErrDuplicatedKey)
}

// DB_NotFoundError checks if the given error corresponds to a record not being found in the database.
// Returns true if the error is a record not found error, false otherwise.
func DB_NotFoundError(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	requestId, ok := c.Locals("requestId").(string)
	if !ok {
		requestId = "unknown"
	}

	logger, ok := c.Locals("logger").(zerolog.Logger)
	if !ok {
		logger = custom_logger.GetLogger().With().Str("requestId", requestId).Logger()
	}

	// Default to Internal Server Error
	response := InternalServerError(CommonErrorsCodes.InternalServerError, CommonErrors.InternalServerError, nil)

	// Check if error is our CustomError type
	var customErr *CustomError
	if errors.As(err, &customErr) {
		response = customErr
	}

	// Check if error is a Fiber error (e.g., "Cannot GET /unknown")
	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		response = &CustomError{
			StatusCode: fiberErr.Code,
			Message:    fiberErr.Message,
			Details: map[string]any{
				"path":   c.OriginalURL(),
				"method": c.Method(),
			},
		}
	}

	// Handle specific Fiber errors like request body size exceeded
	if errors.Is(err, fiber.ErrRequestEntityTooLarge) {
		response = BadRequestError(CommonErrorsCodes.BodySizeExceeded, CommonErrors.BodySizeExceeded, nil)
	}

	// Structured logging based on error type
	logEvent := logger.Error().Err(err).
		Int("statusCode", response.StatusCode).
		Str("message", response.Message).
		Any("details", response.Details).
		Str("method", c.Method()).
		Str("path", c.OriginalURL())

	logEvent.Msg("Handled error")

	return utils.SendErrorResponse(c, response.StatusCode, response.ErrorCode, response.Message, response.Details)
}
