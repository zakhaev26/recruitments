package main

import (
	"fmt"

	"github.com/zakhaev26/recruitments/internal/database"
)

func main() {
	database.InitAuth()
	fmt.Println("Yay")

}
