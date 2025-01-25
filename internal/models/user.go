package models

import "gorm.io/gorm"

// User Модель користувача.
type User struct {
	gorm.Model
	Email    string
	Password string
}
