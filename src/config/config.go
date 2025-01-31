package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	AppName            string `env:"APP_NAME" envDefault:"gofiber-jumpstart"`
	CORSOrigin         string `env:"CORS_ORIGIN" envDefault:"*"`
	CORSHeader         string `env:"CORS_HEADER" envDefault:"*"`
	CORSMethod         string `env:"CORS_METHOD" envDefault:"*"`
	DBDSN              string `env:"DB_DSN,required"`
	JwtSecret          string `env:"JWT_SECRET,required"`
	Host               string `env:"HOST" envDefault:"0.0.0.0"`
	AppPort            int    `env:"APP_PORT" envDefault:"3000"`
	ShutDownTimeOut    int    `env:"SHUTDOWN_TIMEOUT" envDefault:"70"`
	RunMigrations      bool   `env:"RUN_MIGRATIONS" envDefault:"true"`
}

var config Config

func GetConfig() (Config, error) {
	// If already parsed, return
	if config != (Config{}) {
		return config, nil
	}

	// Load the environment variables from .env
	_ = godotenv.Load()

	// Parse the environment variables
	if err := env.Parse(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}