package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/zakhaev26/recruitments/schemas"
	"gorm.io/gorm"
)

type JobResponse struct {
	schemas.Job `json:"jobDetails"`
	Applicants  []schemas.User `json:"applicants"`
}

func (c *Controller) JobDetailController(w http.ResponseWriter, r *http.Request) {
	jobID := r.URL.Query().Get("job_id")
	var job schemas.Job
	if err := c.db.Preload("PostedBy", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name", "email", "address", "user_type", "profile_headline", "profile_id")
	}).First(&job, "id = ?", jobID).Error; err != nil {
		http.Error(w, "Failed to fetch job details", http.StatusInternalServerError)
		return
	}

	var applications []schemas.Applications
	if err := c.db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name", "email", "address", "user_type", "profile_headline", "profile_id")
	}).Preload("User.Profile").Where("job_id = ?", jobID).Find(&applications).Error; err != nil {
		http.Error(w, "Failed to fetch applications", http.StatusInternalServerError)
		return
	}

	response := JobResponse{
		Job:        job,
		Applicants: getUsersFromApplications(applications),
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func getUsersFromApplications(applications []schemas.Applications) []schemas.User {
	var users []schemas.User
	for _, app := range applications {
		users = append(users, app.User)
	}
	return users
}
