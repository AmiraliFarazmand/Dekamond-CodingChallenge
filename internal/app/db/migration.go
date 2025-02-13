package db

import (
	"log"
	"resturant-task/internal/app/models"

	"gorm.io/gorm"
)

func runMigrations(db *gorm.DB) {
	if err := db.AutoMigrate(
		&models.User{},
		&models.Table{},
		&models.Reservation{},
	); err != nil {
		log.Fatalf(">ERR db.RunMigraitons(). Failed to run migrations: %v", err)
	}
	log.Println("Database migrations completed successfully!")
}
