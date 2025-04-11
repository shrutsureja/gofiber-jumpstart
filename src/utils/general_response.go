package utils

import "github.com/gofiber/fiber/v2"

type GeneralResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"` // Omit if no data
}

type ErrorResponse struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
	Data      any    `json:"data,omitempty"` // Omit if no data
}

// SendResponse is a utility function to send structured JSON responses, with optional title
func SendResponse(c *fiber.Ctx, statusCode int, message string, data any) error {
	success := statusCode < 400
	response := GeneralResponse{
		Success: success,
		Message: message,
		Data:    data,
	}
	return c.Status(statusCode).JSON(response)
}

func SendErrorResponse(c *fiber.Ctx, statusCode int, errorCode string, message string, data any) error {
	response := ErrorResponse{
		Success:   false,
		ErrorCode: errorCode,
		Message:   message,
	}

	// Only include data if status code is not a server error (5xx)
	// This is to prevent leaking sensitive information in case of server errors
	if data != nil && statusCode < 500 {
		response.Data = data
	}

	return c.Status(statusCode).JSON(response)
}
