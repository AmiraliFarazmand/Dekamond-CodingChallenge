package controllers

import (
	"net/http"
	"resturant-task/internal/app/db"
	"resturant-task/internal/app/models"
	"resturant-task/internal/app/utils"

	"github.com/gin-gonic/gin"
)

const pricePerSeat float64 = 10.0

type bookTableRequest struct {
	Seats int `json:"seats" binding:"required,min=1"`
}
type cancelReservationRequest struct {
	ReservationID uint `json:"reservation_id" binding:"required,min=1"`
}

func calculatePrice(seats, tableSeats int) float64 {
	if seats == tableSeats {
		return float64((tableSeats - 1)) * pricePerSeat
	}
	return float64(tableSeats) * pricePerSeat
}

func BookTable(c *gin.Context) {

	// Get authenticated user from context
	user, _ := c.Get("user") // in middleware we already check if user exists!
	authUser := user.(models.User)

	// Parse request body
	var req bookTableRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	// find an available table with minimum enough seats
	var table models.Table
	if err := db.DB.
		Where("seats_number >= ? AND is_reserved = ?", req.Seats, false).
		Order("seats_number ASC").First(&table).Error; err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "No available table found")
		return
	}

	// update table, change to reserved
	table.IsReserved = true
	db.DB.Save(&table)

	// Calculate price
	totalPrice := calculatePrice(req.Seats, table.SeatsNumber)

	// Create reservation
	reservation := models.Reservation{
		UserID:        authUser.ID,
		TableID:       table.ID,
		SeatsReserved: req.Seats,
		TotalPrice:    totalPrice,
		IsActive:      true,
	}
	db.DB.Create(&reservation)

	// Return success response
	c.JSON(http.StatusCreated, gin.H{
		"message":        "Reservation successful",
		"reservation_id": reservation.ID,
		"total_price":    totalPrice,
		"table_id":       table.ID,
		"seats_booked":   req.Seats,
	})
}

func CancelReservation(c *gin.Context) {

	// Get authenticated user from context
	user, _ := c.Get("user") // in middleware we already check if user exists!
	authUser := user.(models.User)

	// Parse request body
	var req cancelReservationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	var reservation models.Reservation
	if err := db.DB.
		Where("id = ? AND user_id = ?", req.ReservationID, authUser.ID).
		First(&reservation).Error; err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Reservation not found")
		return
	}

	// Check if the reservation is already inactive
	if !reservation.IsActive {
		utils.RespondWithError(c, http.StatusBadRequest, "Reservation already canceled")
		return
	}

	// Mark reservation as inactive and free up the table
	reservation.IsActive = false
	db.DB.Save(&reservation)
	var table models.Table
	if err := db.DB.First(&table, reservation.TableID).Error; err == nil {
		table.IsReserved = false
		db.DB.Save(&table)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reservation canceled successfully"})
}
