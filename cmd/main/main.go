package main

import (
	"net/http"
	"os"

	"github.com/zakhaev26/recruitments/internal/database"
	"github.com/zakhaev26/recruitments/internal/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}


	database.InitAuth()
	r := routes.Router()
	http.ListenAndServe(":"+port, r)
}
