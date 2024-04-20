package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/zakhaev26/recruitments/schemas"
)

func (c *Controller) ApplicantsController(w http.ResponseWriter, r *http.Request) {
	var systemUsers []schemas.User

	// Select fields excluding password
	if err := c.db.Preload("Profile").Select("id", "created_at", "updated_at", "deleted_at", "name", "email", "address", "user_type", "profile_headline", "profile_id").Find(&systemUsers).Error; err != nil {
		http.Error(w, "Failed to fetch jobs", http.StatusInternalServerError)
		return
	}

	// Marshal the systemUsers slice to JSON
	responseData, err := json.Marshal(systemUsers)
	if err != nil {
		http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header and write the response data
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseData)
}
