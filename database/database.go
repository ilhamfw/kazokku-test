// database/database.go

package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"kazokku-app/config"
	"kazokku-app/app/models"
)

var DB *gorm.DB

// Init initializes the database connection
func Init(cfg *config.Config) error {
	dbURI := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword,
	)

	// Open a connection to the database
	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		return err
	}

	// Set database connection pool settings
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// Store the database connection globally
	DB = db

	// Run database migrations
	runMigrations()

	return nil
}

// runMigrations runs database migrations
func runMigrations() {
	// Run your database migrations here
	DB.AutoMigrate(&models.User{}, &models.Photo{}, &models.CreditCard{})
}
