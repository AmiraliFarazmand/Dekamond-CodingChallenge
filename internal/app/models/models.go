package models

import "time"

type User struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	PhoneNumber string    `gorm:"primaryKey,index:idx_phonenumber,unique"`
	CreatedAt   time.Time `gorm:"default:current_timestamp"`

    // PhoneNumber string    `gorm:"primaryKey;column:phone_number;type:varchar(32);autoIncrement:false" json:"phone_number"`
    // CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}
