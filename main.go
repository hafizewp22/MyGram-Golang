package main

import (
	"fmt"
	"os"
	"project_final/database"
	"project_final/router"
)

func main() {
	port := os.Getenv("PORT")
	database.StartDB()

	router.New().Run(fmt.Sprintf(":%s", port))
}
