package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string `gorm:"unique"`
	Username  string
	FirstName string
	LastName  string
	Password  string
	Phone     string
}
