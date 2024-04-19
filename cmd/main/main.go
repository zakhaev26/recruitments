package main

import (
	"net/http"

	"github.com/zakhaev26/recruitments/internal/database"
	"github.com/zakhaev26/recruitments/internal/routes"
)

func main() {
	database.InitAuth()
	r := routes.Router()
	http.ListenAndServe(":3000", r)
}
