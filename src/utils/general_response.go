package utils

import "github.com/gofiber/fiber/v2"

type GeneralResponse struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Title   string      `json:"title,omitempty"` // Include title for error responses, optional
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // Omit if no data
}

// SendResponse is a utility function to send structured JSON responses, with optional title
func SendResponse(c *fiber.Ctx, code int, title, message string, data interface{}) error {
	success := code < 400
	response := GeneralResponse{
		Success: success,
		Code:    code,
		Title:   title,
		Message: message,
		Data:    data,
	}

	return c.Status(code).JSON(response)
}