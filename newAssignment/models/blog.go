package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	ID          uint     `gorm:"primaryKey"`
	Title       string   `gorm:"not null"form:"title"`
	Description string   `form:"description"`
	CategoryID  uint     `form:"categoryId"`
	AuthorID    uint     `form:"authorId"`
	Category    Category `gorm:"foreignKey:CategoryID"`
	Author      Author   `gorm:"foreignKey:AuthorID"`
	Image       string
}
