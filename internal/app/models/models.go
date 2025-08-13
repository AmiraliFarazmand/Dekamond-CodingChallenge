package models

import "time"

type User struct {
	PhoneNumber string    `gorm:"primaryKey,index:idx_phonenumber,unique" json:"phone_number"`
	CreatedAt   time.Time `gorm:"default:current_timestamp" json:"created_at"`
}
