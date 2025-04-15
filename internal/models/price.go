package models

import "gorm.io/gorm"

// Price represents a key-value pair stored in the database
type Price struct {
	gorm.Model
	Price string `gorm:"price" json:"price"` // Price field

}
