package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetResumeSummary(buffer bytes.Buffer) (string, error) {
	// Create a new HTTP request
	req, err := http.NewRequest("POST", "https://api.apilayer.com/resume_parser/upload", &buffer)
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("apikey", "gNiXyflsFu3WNYCz1ZCxdWDb7oQg1Nl1")
	log.Printf("idhar agaya hun")
	// Create a HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	// Return summary string and nil error
	return string(respBody), nil
}
