package models

import "time"

type User struct {
	PhoneNumber string    `gorm:"primaryKey,index:idx_phonenumber,unique"`
	CreatedAt   time.Time `gorm:"default:current_timestamp"`
}
