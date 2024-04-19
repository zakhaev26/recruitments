package controllers

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/charmbracelet/log"
	"github.com/zakhaev26/recruitments/schemas"
	"github.com/zakhaev26/recruitments/utils"
)

func (c *Controller) ResumeUploadController(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Info("hereeeeee")


	file, _, err := r.FormFile("cv")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Check file format
	fileType, err := utils.DetectFileType(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if fileType != ".pdf" && fileType != ".zip" {
		http.Error(w, "Unsupported file format. Only PDF and DOCX are allowed.", http.StatusBadRequest)
		return
	}

	fileData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
	fmt.Println(claims.UserID)
	fileRecord := schemas.File{
		UserID:   claims.UserID,
		FileName: "uploaded_file_" + time.Now().Local().Format("2006-01-02_15-04-05") + fileType,
		FileType: fileType,
		FileData: fileData,
	}

	c.db.Create(&fileRecord)
	log.Info("hihiihihi")
	fmt.Fprintf(w, "File uploaded successfully. File saved in the database.")
}
