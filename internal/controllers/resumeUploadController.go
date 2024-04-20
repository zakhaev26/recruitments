package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/charmbracelet/log"
	"github.com/zakhaev26/recruitments/schemas"
	"github.com/zakhaev26/recruitments/utils"
)

func (c *Controller) ResumeUploadController(w http.ResponseWriter, r *http.Request) {
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
	var buffer bytes.Buffer

	// Copy the request body to the buffer
	_, err = io.Copy(&buffer, r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	fileBytes := buffer.Bytes()
	fileRecord := schemas.File{
		UserID:   claims.UserID,
		FileName: "uploaded_file_" + time.Now().Local().Format("2006-01-02_15-04-05") + ".pdf",
		FileType: ".pdf",
		FileData: &fileBytes,
	}

	c.db.Create(&fileRecord)
	summaryJSON, err := utils.GetResumeSummary(buffer)
	if err != nil {
		log.Error(err)
		http.Error(w, "Error parsing CV summary", http.StatusInternalServerError)
		return
	}

	var summary schemas.Summary
	err = json.Unmarshal([]byte(summaryJSON), &summary)
	if err != nil {
		log.Error(err)
		http.Error(w, "Error parsing CV summary", http.StatusInternalServerError)
		return
	}

	var profile schemas.Profile
	result := c.db.First(&profile, "id = ?", claims.ProfileID)
	if result.Error != nil {
		log.Error(result.Error)
		http.Error(w, "Error patching Profiles", http.StatusInternalServerError)
		return
	}
	utils.PatchProfileFromSummary(&profile, summary)

	c.db.Save(profile)
	fmt.Fprintf(w, "File uploaded successfully. File saved in the database.")
}
