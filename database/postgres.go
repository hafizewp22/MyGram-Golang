package database

import (
	"fmt"
	"project_final/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Host      string
	Username  string
	Password  string
	Port      int
	Name      string
	DebugMode string
}

const (
	DB_HOST     = "localhost"
	DB_USER     = "postgres"
	DB_PASSWORD = "root"
	DB_PORT     = 5432
	DB_NAME     = "project_final"
	DEBUG_MODE  = true // true/false
)

var (
	db  *gorm.DB
	err error
)

func StartDB(conf *Database) *gorm.DB {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)

	if conf.Host != "" {
		config = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			conf.Host, conf.Username, conf.Password, conf.Name, conf.Port)
	}

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if conf.DebugMode == "true" || DEBUG_MODE {
		fmt.Println(conf.DebugMode, DEBUG_MODE)
		db.Debug().AutoMigrate(models.User{}, models.SocialMedia{}, models.Photo{}, models.Comment{})
		return db
	}

	db.AutoMigrate(models.User{}, models.SocialMedia{}, models.Photo{}, models.Comment{})
	return db
}

func GetDB() *gorm.DB {
	if DEBUG_MODE {
		return db.Debug()
	}

	return db
}
