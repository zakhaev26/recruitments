package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/zakhaev26/recruitments/schemas"
)

func (c *Controller) GetApplicantInformation(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("applicant_id")
	var user schemas.User

	if err := c.db.Preload("Profile").Select("id", "created_at", "updated_at", "deleted_at", "name", "email", "address", "user_type", "profile_headline", "profile_id").Where("id = ?", userID).First(&user).Error; err != nil {
		http.Error(w, "Failed to fetch Applicant Details", http.StatusInternalServerError)
		return
	}

	responseData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseData)
}
