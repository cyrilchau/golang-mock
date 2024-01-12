package datasource

import (
	"fmt"
	"myapp/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(cfg config.DatabaseConfig) (*gorm.DB, error) {
	// Create the connection string.
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=%s", cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port, cfg.TimeZone)

	// Connect to the Postgres database.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
