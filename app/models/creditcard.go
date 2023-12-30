package models

import (
	"github.com/jinzhu/gorm"
)

// CreditCard model represents the credit card entity in the database
type CreditCard struct {
	gorm.Model
	UserID        uint
	Type          string `gorm:"not null"`
	Number        string `gorm:"not null"`
	Name          string `gorm:"not null"`
	Expired       string `gorm:"not null"`
	CVV           string `gorm:"not null"`
}
