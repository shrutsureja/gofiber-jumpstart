package database

import (
	"app/src/config"
	"app/src/utils/logger"
	"embed"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//go:embed migrations/*.sql
var EmbedMigrations embed.FS

var log = logger.GetLogger()
var DB *gorm.DB

func InitDB(cfg config.Config) {
	var err error
	DB, err = gorm.Open(postgres.Open(cfg.DBDSN), &gorm.Config{TranslateError: true})

	if err != nil {
		panic("failed to connect database")
	}

	// Running migrations
	if cfg.RunMigrations {
		RunMigration()
	}
}

func RunMigration() {
	// Set the base filesystem to the embedded migrations
	goose.SetBaseFS(EmbedMigrations)

	// Set the database dialect to PostgreSQL
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal().Err(err).Msg("Failed to database dialect to PostgreSQL")
	}

	// Get the database handle from your DB object to type *sql.DB
	dbHandle, err := DB.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("Error in getting DB handle")
	}

	// Run the migration using the embedded files
	if err := goose.Up(dbHandle, "migrations"); err != nil {
		log.Fatal().Err(err).Msg("Error in running migrations:")
	}
}