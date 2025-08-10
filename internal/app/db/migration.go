package db

import (
	"Dakomond/internal/app/models"
	"log"

	"gorm.io/gorm"
)

func runMigrations(db *gorm.DB) {
	if err := db.AutoMigrate(
		&models.User{},
	); err != nil {
		log.Fatalf(">ERR db.RunMigraitons(). Failed to run migrations: %v", err)
	}
	log.Println("Database migrations completed successfully!")
}
