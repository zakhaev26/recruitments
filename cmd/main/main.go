package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/charmbracelet/log"
	"github.com/zakhaev26/recruitments/internal/database"
	"github.com/zakhaev26/recruitments/internal/routes"
	"github.com/zakhaev26/recruitments/utils"
)

func main() {
	database.InitAuth()
	fmt.Println("Yay")
	userId := "gonjigreatUWU"
	tok, err := utils.GenerateAccessToken(userId)
	if err != nil {
		log.Info(err)
		return
	}
	for {
		log.Info(tok)
		cl, err := utils.ParseAccessToken(tok)
		if err != nil {
			log.Info(err)
			break
		}
		log.Info("isExpired", utils.IsTokenExpired(tok))
		log.Info("userId", cl.UserID)
		time.Sleep(time.Second * 1)
	}

	r := routes.Router()

	http.ListenAndServe(":3000", r)
}
