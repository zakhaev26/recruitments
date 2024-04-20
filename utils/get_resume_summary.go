package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func GetResumeSummary(buffer bytes.Buffer) (string, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	RESUME_PARSER_API_KEY := os.Getenv("RESUME_PARSER_API_KEY")
	RESUME_PARSER_URI := os.Getenv("RESUME_PARSER_URL")
	req, err := http.NewRequest("POST", RESUME_PARSER_URI, &buffer)
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("apikey", RESUME_PARSER_API_KEY)
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	return string(respBody), nil
}
