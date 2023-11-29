package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" binding:"required" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password" binding:"required"`
	UserType string `json:"user_type" binding:"required"`
}
