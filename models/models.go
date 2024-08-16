package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `swaggerignore:"true"`
	ID         uint   `gorm:"primaryKey"`
	Username   string `gorm:"unique"`
	Email      string
}
