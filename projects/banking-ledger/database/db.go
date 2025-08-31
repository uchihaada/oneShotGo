// Package database handles database connection and setup
// Learning focus: GORM connection, auto-migration, connection management
package database

import (
	"banking-ledger/models"
	"errors"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect initializes database connection and runs migrations
func Connect() error {
	// open database connection
	db, err := gorm.Open(sqlite.Open("banking.db"), &gorm.Config{})
	DB = db
	if err != nil {
		log.Println("failed to connect to database", err)
		return err
	}

	log.Println("Database connection successful")

	// Auto-migrate all models (User, Account, Transaction)
	err = DB.AutoMigrate(&models.Account{}, &models.User{}, &models.Transaction{})

	if err != nil {
		return errors.New("migrating models failed")
	}

	// Test connection with a simple ping/query
	sqlDB, _ := DB.DB()
	err = sqlDB.Ping()
	if err != nil {
		return errors.New("ping failed")
	}

	log.Println("Database connected and migrated successfully")
	return nil
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}

// CloseDB closes the database connection
func CloseDB() error {
	sqlDB, err := DB.DB()
	if err == nil {
		return sqlDB.Close()
	}
	return err
}
