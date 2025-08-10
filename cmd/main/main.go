package main

import (
	"Dakomond/internal/app/db"
	"Dakomond/internal/app/routes"
)

func main() {
	db.InitDB()
	r := routes.SetupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
