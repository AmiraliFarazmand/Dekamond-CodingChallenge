package db

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func SetOTP(phoneNumber, otpCode string) error {
	shortSetStatus := REDIS.Set(context.Background(), "OTP:"+phoneNumber, otpCode, time.Minute*2)
	longSetStatus := REDIS.Set(context.Background(), "OTPHISTORY:"+phoneNumber+":"+time.Now().String(), otpCode, time.Minute*10)
	if shortSetStatus.Err() != nil || longSetStatus.Err() != nil {
		return fmt.Errorf(">REDIS ERROR")
	}
	return nil
}

func IsValidToCreateOTP(phoneNumber string) bool {
	keys, _ := REDIS.Keys(context.Background(), "OTPHISTORY:"+phoneNumber+":*").Result()
	if len(keys) > 2 {
		return false
	} else {
		return true
	}
}

func CheckOTP(phoneNumber, otpCode string) bool {
	otp, err := REDIS.Get(context.Background(), "OTP:"+phoneNumber).Result()
	if err == redis.Nil {
		return false
	}
	if otp == otpCode {
		REDIS.Del(context.Background(), "OTP:"+phoneNumber)
		return true
	}
	return false
}
