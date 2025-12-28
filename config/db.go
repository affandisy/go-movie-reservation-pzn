package config

import (
	"fmt"
	"go-movie-reservation/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(Config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Config.DBHost, Config.DBPort, Config.DBUser, Config.DBPassword, Config.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// enable UUID extension
	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		return nil, fmt.Errorf("failed to create UUID Extension: %w", err)
	}

	// Auto-Migrate all models
	if err := db.AutoMigrate(
		&model.Movie{},
		&model.User{},
		&model.Showtime{},
		&model.Reservation{},
	); err != nil {
		return nil, fmt.Errorf("failed to auto-migrate: %w", err)
	}

	return db, nil
}
