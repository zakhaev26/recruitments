package controllers

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/zakhaev26/recruitments/schemas"
	"github.com/zakhaev26/recruitments/utils"
)

func (c *Controller) ApplyJobController(w http.ResponseWriter, r *http.Request) {
	jobID := r.URL.Query().Get("job_id")
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

	userId := claims.UserID
	jobId, err := uuid.Parse(jobID)
	if err != nil {
		log.Error(err)
		http.Error(w, "Failed to parse job id", http.StatusBadRequest)
		return
	}
	application := schemas.Applications{
		UserID: userId,
		JobID:  jobId,
	}

	application.ID = uuid.New()

	if err := c.db.Create(&application).Error; err != nil {
		log.Error(err)
		http.Error(w, "Failed to create job application", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Job application submitted successfully"))

}
