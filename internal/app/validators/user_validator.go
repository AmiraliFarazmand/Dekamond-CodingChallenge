package validators

import (
	"Dakomond/internal/app/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func checkUniquenessPhoneNumber(db *gorm.DB, phoneNumber string) error {
	var user models.User
	if db.Where("phone_number = ?", phoneNumber).First(&user).Error == nil {
		return fmt.Errorf("user with phone number:%s already exists", phoneNumber)
	}
	return nil
}

func validatePhoneNumber(db *gorm.DB, phoneNumber string) error {
	fmt.Println(phoneNumber)
	if len(phoneNumber) < 3 || len(phoneNumber) > 64 {
		return errors.New("phone number must be between 3 and 64 characters")
	}
	if err := checkUniquenessPhoneNumber(db, phoneNumber); err != nil {
		return err
	}
	return nil
}

func ValidateUsernameAndOTP(db *gorm.DB, phoneNumber string, otp string) error {
	if err := validatePhoneNumber(db, phoneNumber); err != nil {
		return err
	}
	return nil
}
