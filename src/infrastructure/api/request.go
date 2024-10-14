package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// SendDataToAPI sends JSON data to the specified API endpoint
func SendDataToAPI(url string, data interface{}) error {
	// Encode the data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshaling data to JSON: %v", err)
	}

	// Create a new POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %v", err)
	}

	// Set the content type to application/json
	req.Header.Set("Content-Type", "application/json")

	// Send the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending HTTP request: %v", err)
	}
	defer resp.Body.Close() // Ensure the response body is closed

	// Check if the request was successful
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("received non-201 response status: %s", resp.Status)
	}

	return nil
}
