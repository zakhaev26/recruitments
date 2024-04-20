package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/zakhaev26/recruitments/internal/database"
	"github.com/zakhaev26/recruitments/internal/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	database.InitAuth()
	r := routes.Router()
	http.ListenAndServe(PORT, r)
}
