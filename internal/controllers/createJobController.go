package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/zakhaev26/recruitments/schemas"
	"github.com/zakhaev26/recruitments/utils"
)

func (c *Controller) CreateJobController(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		CompanyName string `json:"companyName"`
		Email       string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		log.Error(err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	token, err := utils.ExtractTokenFromHeader(r)
	if err != nil {
		log.Error(err)
		http.Error(w, "Failed to extract token from headers", http.StatusBadRequest)
		return
	}

	claims, err := utils.ParseAccessToken(token)
	if err != nil {
		log.Error(err)
		http.Error(w, "Failed to parse token from headers", http.StatusBadRequest)

		return
	}
	job := schemas.Job{
		Title:             reqBody.Title,
		Description:       reqBody.Description,
		PostedOn:          time.Now(),
		TotalApplications: 0,
		CompanyName:       reqBody.CompanyName,
		PostedByID:        claims.UserID,
	}
	job.ID = uuid.New()

	if err := c.db.Create(&job).Error; err != nil {
		http.Error(w, "Failed to create job", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(job)
}
