package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/zakhaev26/recruitments/schemas"
	"gorm.io/gorm"
)

func (c *Controller) GetJobsCotroller(w http.ResponseWriter, r *http.Request) {
	var jobs []schemas.Job
	if err := c.db.Preload("PostedBy", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name", "email", "user_type", "profile_headline", "profile_id") // Select specific fields from PostedBy
	}).Preload("PostedBy.Profile", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "education", "skills")
	}).Find(&jobs).Error; err != nil {
		http.Error(w, "Failed to fetch jobs", http.StatusInternalServerError)
		return
	}

	// Return the fetched job data as a response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jobs)
}
