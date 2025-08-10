package models

import "time"

type User struct {
	PhoneNumber string    `gorm:"primaryKey"`
	CreatedAt   time.Time `gorm:"default:current_timestamp"`
}
