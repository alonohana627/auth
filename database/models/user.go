package models

import "gorm.io/gorm"

type Users struct {
	Nickname       string `gorm:"unique;not null"`
	Email          string `gorm:"unique;not null"`
	HashedPassword string `gorm:"not null"`
	gorm.Model
}
