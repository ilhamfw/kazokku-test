package models

import (
	"github.com/jinzhu/gorm"
)

// User model represents the user entity in the database
type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null;unique"`
	Address  string `gorm:"not null"`
	Password string `gorm:"not null"`
	Photos []Photo `gorm:"foreignKey:UserID"`
	CreditCards []CreditCard
}
