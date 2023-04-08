package main

import (
	"log"
	"os"
	"project_final/database"
	"project_final/router"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

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
