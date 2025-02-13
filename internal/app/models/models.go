package models

type User struct {
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"size:64;not null;unique"`
	Password string `gorm:"size:64;not null"`
}

type Table struct {
	ID          uint `gorm:"primaryKey"`
	SeatsNumber int  `gorm:"not null;check:seats_number >= 4 AND seats_number <= 10"`
	IsReserved  bool `gorm:"not null,default:true"`
}

type Reservation struct {
	ID            uint    `gorm:"primaryKey"`
	UserID        uint    `gorm:"not null;index"`
	User          User    `gorm:"foreignKey:UserID"`
	TableID       uint    `gorm:"not null;index"`
	Table         Table   `gorm:"foreignKey:TableID"`
	SeatsReserved int     `gorm:"not null"`
	TotalPrice    float64 `gorm:"not null"`
	IsActive      bool    `gorm:"not null,default:true"`
}
