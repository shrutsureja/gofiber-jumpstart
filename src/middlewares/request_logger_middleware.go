package middlewares

import (
	"app/src/utils/custom_logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
)

var requestLogger = custom_logger.GetLogger()

func RequestLoggerMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestId := c.Locals("requestId").(string)
		l := requestLogger.With().Str("requestId", requestId).Logger()
		return fiberzerolog.New(fiberzerolog.Config{
			Logger: &l,
		})(c)
	}
}
