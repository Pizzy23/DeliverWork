package main

import (
	"deliver/db"
	"deliver/middleware"
)

func main() {
	r := middleware.SetupRouter()

	db.ConnectDatabase()
	r.Run(":8080")
}
