package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/zakhaev26/recruitments/schemas"
	"github.com/zakhaev26/recruitments/utils"
	"gorm.io/gorm"
)

func (c *Controller) SignUpController(w http.ResponseWriter, r *http.Request) {
	log.Info("aya")
	var (
		reqBody struct {
			User    schemas.User    `json:"user"`
			Profile schemas.Profile `json:"profile"`
		}
	)

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		log.Error(err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	user := reqBody.User
	profile := reqBody.Profile

	if user.Name == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "Name, Email, and Password are required fields", http.StatusBadRequest)
		return
	}

	user.ID = uuid.New()
	hashedPass, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	user.Password = string(hashedPass)
	profile.ID = uuid.New()

	if reqBody.Profile.Skills == "" || reqBody.Profile.Education == "" || reqBody.Profile.Experience == "" {
		http.Error(w, "ResumeFileAddress, Skills, Education, and Experience are required fields for profile", http.StatusBadRequest)
		return
	}

	if err := c.db.Create(&profile).Error; err != nil {
		http.Error(w, "Failed to create profile", http.StatusInternalServerError)
		return
	}

	user.ProfileID = profile.ID
	user.Profile = profile

	if err := c.db.Create(&user).Error; err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(reqBody.User)
}

func (c *Controller) LoginController(w http.ResponseWriter, r *http.Request) {

	var (
		reqBody struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
	)

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		log.Error(err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	if reqBody.Email == "" || reqBody.Password == "" {
		http.Error(w, "Required fields missing", http.StatusBadRequest)
		return
	}

	var user schemas.User
	if err := c.db.Where("email = ?", reqBody.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}
		log.Error(err)
		http.Error(w, "Failed to find user", http.StatusInternalServerError)
		return
	}

	if utils.ComparePassword(user.Password, reqBody.Password) != nil {
		http.Error(w, "Invalid password", http.StatusBadRequest)
		return
	}

	token, err := utils.GenerateAccessToken(user.ID, user.UserType)
	if err != nil {
		http.Error(w, "Invalid password", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"accessToken": token,
	})
}
