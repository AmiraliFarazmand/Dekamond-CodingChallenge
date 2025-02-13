package main

import (
	"log"
	"lotus-task/internal/app/controllers"
	"lotus-task/internal/app/db"
	"lotus-task/internal/app/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	err := db.Connect()
	if err != nil {
		log.Println(err)
	}
	db.RunMigrations(db.DB)
	log.Println("successful to run migrations!")
}

func main() {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.ValidateIsAuthenticated)
	r.POST("/book", middleware.RequireAuth, controllers.BookTable)
	r.POST("/cancel", middleware.RequireAuth, controllers.CancelReservation)
	r.Run() // listen and serve on 0.0.0.0:8080

}
