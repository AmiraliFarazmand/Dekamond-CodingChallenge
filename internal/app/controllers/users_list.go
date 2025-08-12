package controllers

import (
	"Dakomond/internal/app/db"
	"Dakomond/internal/app/models"
	"fmt"
	"math"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ListQuery struct {
	Page    int    `form:"page,default=1" binding:"gte=1"`
	PerPage int    `form:"per_page,default=10" binding:"gte=1,lte=100"`
	Search  string `form:"search"`
}

func UsersPagination(c *gin.Context) {
	var q ListQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	offset := (q.Page - 1) * q.PerPage

	tx := db.DB.Model(&models.User{})

	// search (case-insensitive contains on name/description)
	if s := strings.TrimSpace(q.Search); s != "" {
		like := "%" + s + "%"
		tx = tx.Where("phone_number ILIKE ?", like)
	}

	// total count (for meta)
	var total int64
	if err := tx.Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "count failed"})
		return
	}

	// page data (stable order!)
	var items []models.User
	if err := tx.
		Order("created_at DESC").
		Order(`"phone_number" DESC`).
		Limit(q.PerPage).
		Offset(offset).
		Find(&items).Error; err != nil {
			fmt.Println("KIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIR",err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       items,
		"page":       q.Page,
		"per_page":   q.PerPage,
		"total":      total,
		"total_page": int(math.Ceil(float64(total) / float64(q.PerPage))),
		"search":     q.Search,
	})
}
