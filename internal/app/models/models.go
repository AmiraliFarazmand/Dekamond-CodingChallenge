package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"size:64;not null;unique"`
	Password string `gorm:"size:64;not null"`
}

type Table struct {
	ID          uint `gorm:"primaryKey"`
	SeatsNumber int  `gorm:"not null;check:seats_number >= 4 AND seats_number <= 10;index:,sort:desc"`
	IsReserved  bool `gorm:"not null,default:false"`
}

type Reservation struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint `gorm:"not null"`
	TableID uint `gorm:"not null"`
	SeatsReserved int     `gorm:"not null"`
	TotalPrice    float64 `gorm:"not null"`
	IsActive      bool    `gorm:"not null,default:true"`
}
