package controllers

import (
	"lotus-task/internal/app/db"
	"lotus-task/internal/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

const pricePerSeat float64 = 10.0

func BookTable(c *gin.Context) {
	type RequestBody struct {
		Seats int `json:"seats" binding:"required,min=1"`
	}
	// Get authenticated user from context
	user, _ := c.Get("user") // in middleware we already check if user exists!
	authUser := user.(models.User)

	// Parse request body
	var req RequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// find an available table with minimum enough seats
	var table models.Table
	if err := db.DB.
		Where("seats_number >= ? AND is_reserved = ?", req.Seats, false).
		Order("seats_number ASC").First(&table).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No available table found"})
		return
	}

	// update table, change to reserved
	table.IsReserved = true
	db.DB.Save(&table)

	// Calculate price
	var totalPrice float64
	if req.Seats == table.SeatsNumber {
		totalPrice = float64((table.SeatsNumber - 1)) * pricePerSeat
		} else {
			totalPrice = float64(table.SeatsNumber) * pricePerSeat
		}

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
		"message":      "Reservation successful",
		"reservation_id":  reservation.ID,
		"total_price":  totalPrice,
		"table_id":     table.ID,
		"seats_booked": req.Seats,
	})
}
