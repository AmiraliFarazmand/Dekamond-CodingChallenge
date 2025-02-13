package main

import (
	"resturant-task/internal/app/db"
	"resturant-task/internal/app/routes"
)

func main() {
	db.InitDB()
	r := routes.SetupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080

}
