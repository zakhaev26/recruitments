package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetResumeSummary(buffer bytes.Buffer) (string, error) {
	req, err := http.NewRequest("POST", "https://api.apilayer.com/resume_parser/upload", &buffer)
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("apikey", "gNiXyflsFu3WNYCz1ZCxdWDb7oQg1Nl1")
	log.Printf("idhar agaya hun")
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
