package middlewares

import (
	"app/src/utils/custom_logger"

	"github.com/gofiber/fiber/v2"
)

// The only purpose of this middleware is to add logger to the context with the request id
func LoggerMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestId := c.Locals("requestId").(string)
		logger := custom_logger.GetLogger().With().Str("requestId", requestId).Logger()
		c.Locals("logger", logger)
		return c.Next()
	}
}
