package controllers

import (
	"Dakomond/internal/app/db"
	"Dakomond/internal/app/models"
	"Dakomond/internal/app/utils"
	"Dakomond/internal/app/validators"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const tokenExpireTime int = 72 // expire time in hour

type authRequest struct {
	PhoneNumber string `json:"phone number"`
	OTP         string
}

type otpRequest struct {
	PhoneNumber string `json:"phone number"`
}

func CreateOTP(c *gin.Context) {
	var body otpRequest
	// Validate format of request
	if c.ShouldBindJSON(&body) != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	// Validate Username and Password of request body
	if err := validators.ValidatePhoneNumber(db.DB, body.PhoneNumber); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	otp := utils.RandString(4)
	if ok := db.IsValidToCreateOTP(body.PhoneNumber); !ok {
		utils.RespondWithError(c, http.StatusTooManyRequests, "Too many attempts")
		return
	}
	if err := db.SetOTP(body.PhoneNumber, otp); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Internal server errro")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"OTP code": otp})
}

func createToken(userID string) (string, error) {

	// Create Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * time.Duration(tokenExpireTime)).Unix(),
	})
	secretKey, err := utils.ReadEnv("SECRET_KEY")
	if err != nil {
		return "", err
	}
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Signup(c *gin.Context) {
	var body authRequest
	// Validate format of request
	if c.ShouldBindJSON(&body) != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	if err := validators.ValidatePhoneNumber(db.DB, body.PhoneNumber); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := validators.CheckUniquenessPhoneNumber(db.DB, body.PhoneNumber); err!=nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	// validate if the OTP code was correct
	if ok := db.CheckOTP(body.PhoneNumber, body.OTP); !ok {
		utils.RespondWithError(c, http.StatusNotAcceptable, "OTP code wasn't correct or even doesn't exist!")
		return
	}

	// Insert into DB
	user := models.User{PhoneNumber: body.PhoneNumber}
	result := db.DB.Create(&user)
	if result.Error != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create user")
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func Login(c *gin.Context) {
	var body authRequest
	// Validate format of request
	if c.ShouldBindJSON(&body) != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	var user models.User
	db.DB.First(&user, "phone_number = ?", body.PhoneNumber)
	if user.PhoneNumber == "" {
		utils.RespondWithError(c, http.StatusBadRequest, "User not found")
		return
	}
	// validate if the OTP code was correct
	if ok := db.CheckOTP(body.PhoneNumber, body.OTP); !ok {
		utils.RespondWithError(c, http.StatusNotAcceptable, "OTP code wasn't correct or even don't exist!")
		return
	}

	tokenString, err := createToken(user.PhoneNumber)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create token")
		return
	}

	// Set the token in a cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*tokenExpireTime, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func ValidateIsAuthenticated(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.RespondWithError(c, http.StatusUnauthorized, "UnAuthorized user")
		return
	}

	phoneNumber := user.(models.User).PhoneNumber
	c.JSON(http.StatusOK, gin.H{
		"message":      "I am Authenticated",
		"phone number": phoneNumber,
	})
}
