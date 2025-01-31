package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"app/src/config"
	"app/src/database"
	"app/src/middlewares"
	"app/src/routes"
	"app/src/utils"
	"app/src/utils/custom_logger"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

var log = custom_logger.GetLogger()

func main() {
	// Load configuration
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}

	// Initialize Database and run migrations
	database.InitDB(cfg)
	log.Info().Msg("Database connected successfully")

	// App Config
	app := fiber.New(fiber.Config{
		AppName:               cfg.AppName,
		ServerHeader:          cfg.AppName,
		Immutable:             true,
		ETag:                  false,
		DisableKeepalive:      true,
		DisableDefaultDate:    true,
		DisableStartupMessage: true,
		BodyLimit:             20 * 1024 * 1024, // 20 MB body limit
		JSONEncoder:           sonic.Marshal,
		JSONDecoder:           sonic.Unmarshal,
	})

	// Middlewares
	app.Use(recover.New())
	app.Use(requestid.New(
		requestid.Config{
			Next:       nil,
			Header:     fiber.HeaderXRequestID,
			ContextKey: "requestId",
			Generator: func() string {
				return utils.GenerateUUID().String()
			},
		},
	))
	app.Use(middlewares.RequestLoggerMiddleware())
	app.Use(helmet.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.CORSOrigin,
		AllowHeaders: cfg.CORSHeader,
		AllowMethods: cfg.CORSMethod,
	}))
	app.Use(compress.New())
	app.Use(healthcheck.New())

	// Routes
	routes.SetupRoutes(app)

	// Start server in goroutine and handle shutdown
	log.Info().Msgf("Listening on %s:%d", cfg.Host, cfg.AppPort)
	go func() {
		if err := app.Listen(fmt.Sprintf("%s:%d", cfg.Host, cfg.AppPort)); err != nil {
			log.Error().Err(err).Msg("Failed to start listener")
			panic(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutting down server")

	if err := app.ShutdownWithTimeout(time.Duration(cfg.ShutDownTimeOut) * time.Millisecond); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	log.Info().Msg("Server exiting")
}
