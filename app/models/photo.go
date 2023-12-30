package models

import (
	"github.com/jinzhu/gorm"
)

// Photo model represents the photo entity in the database
type Photo struct {
	gorm.Model
	UserID    uint
	FileName  string `gorm:"not null"`
}
