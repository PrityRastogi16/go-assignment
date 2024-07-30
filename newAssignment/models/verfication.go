package models

import "time"

type VerificationToken struct {
	ID        uint      `json:"id" gorm:"primary_key;auto_increment"`
	UserID    uint      `json:"userId" gorm:"not null"`
	Token     string    `json:"token" gorm:"type:varchar(255);not null"`
	ExpiresAt time.Time `json:"expiresAt" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:current_timestamp"`
}
