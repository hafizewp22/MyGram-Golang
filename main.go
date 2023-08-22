package main

import (
	"project_final/database"
	"project_final/router"
)

func main() {
	database.StartDB()

	router.New().Run(":5000")
}
