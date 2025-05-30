package database

import (
	// "database/sql"

	"spo-data/internal/models"

	"gorm.io/gorm"

	// _ "github.com/mattn/go-sqlite3" // SQLite driver
	"github.com/glebarez/sqlite"
)

var db *gorm.DB

// Initialize opens a database connection and sets up the schema
func Initialize(dbPath string) error {
	var err error
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&models.Price{})

	return nil

}

// StoreValue inserts a value into the database
func StorePrice(price models.Price) (models.Price, error) {
	result := db.Create(&price)

	if result.Error != nil {
		return models.Price{}, result.Error
	}

	return price, nil
}

// GetValueByName retrieves a value from the database by its name
func GetLastPrice() (models.Price, error) {
	var price models.Price
	result := db.Last(&price)
	if result.Error != nil {
		return models.Price{}, result.Error
	}

	return price, nil
}
