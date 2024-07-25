package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	CategoryID  uint
	AuthorID    uint
	Category    Category `gorm:"foreignKey:CategoryID"`
	Author      Author   `gorm:"foreignKey:AuthorID"`
	Image       string
}
