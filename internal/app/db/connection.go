package db

import (
	"Dakomond/internal/app/utils"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var REDIS *redis.Client
func connect() error {
	dsn, err := utils.ReadEnv("DSN")
	if err != nil {
		return fmt.Errorf(">ERR db.Connect().%w", err)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.New(">ERR db.Connect(). Failed to connect to database")
	}

	redisAddress, err := utils.ReadEnv("REDIS_ADDR")
	if err != nil {
		return fmt.Errorf(">ERR db.Connect().%w", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:         redisAddress, // or container DNS if running inside Docker
		Password:     "",           // set if you configured requirepass
		DB:           0,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	})
	
	// Assign the global variables
	REDIS = rdb
	DB = db
	return nil
}
