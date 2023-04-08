package main

import (
	"os"
	"project_final/database"
	"project_final/router"
	"strconv"
)

func main() {
	dbPort, _ := strconv.Atoi(os.Getenv("DB_USER"))
	dbConf := database.Database{
		Host:      os.Getenv("DB_HOST"),
		Username:  os.Getenv("DB_USER"),
		Password:  os.Getenv("DB_PASSWORD"),
		Port:      dbPort,
		Name:      os.Getenv("DB_NAME"),
		DebugMode: os.Getenv("DEBUG_MODE"),
	}

	database.StartDB(&dbConf)

	router.New().Run(":5000")
}
