package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex;not null" validate:"required"`
	Password string `json:"password" gorm:"not null" validate:"required"`
}
