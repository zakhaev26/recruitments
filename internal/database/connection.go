package database

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"github.com/zakhaev26/recruitments/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitAuth() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DATABASE_CONNECTION_STRING := os.Getenv("DB_DEV")

	Db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  DATABASE_CONNECTION_STRING,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Error("Failed to connect to Database", "err", err)
		os.Exit(1)
	}
	log.Info("Connected to PostgreSQL")
	if err := Db.AutoMigrate(&schemas.User{}, &schemas.Profile{}, &schemas.Job{}, &schemas.File{}, &schemas.Applications{}); err != nil {
		log.Error("Schema Migration Failed", "err", err)
		return
	}

	log.Info("Migrations Successful")
}
