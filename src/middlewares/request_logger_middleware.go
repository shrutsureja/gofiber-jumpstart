package middlewares

import (
	"app/src/utils/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
)

var requestLogger = logger.GetLogger()

func RequestLoggerMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestId := c.Locals("requestId").(string)
		l := requestLogger.With().Str("requestId", requestId).Logger()
		return fiberzerolog.New(fiberzerolog.Config{
			Logger: &l,
		})(c)
	}
}
