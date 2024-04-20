package database

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/zakhaev26/recruitments/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitAuth() {

	var err error
	Db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=admin dbname=synergylabs host=localhost port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Error("Failed to connect to Database", "err", err)
		os.Exit(1)
	}
	log.Info("Connected to PostgreSQL")
	if err := Db.AutoMigrate(&schemas.User{}, &schemas.Profile{}, &schemas.Job{}, &schemas.File{}, &schemas.Applications{}	); err != nil {
		log.Error("Schema Migration Failed", "err", err)
		return
	}

	log.Info("Migrations Successful")
}
