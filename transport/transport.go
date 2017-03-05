package transport

import (
	"fmt"
	"log"
	"net/http"
)

// PingService sends a get request to service url and compares
// the status code to an expected value
func PingService(url string, expectedStatus int) (*http.Response, error) {
	response, err := http.Get(url)

	if err != nil {
		log.Println("Failure to get make request")
		return nil, err
	}

	if response.StatusCode != expectedStatus {
		err := fmt.Errorf("Expected Status %d Actual Status %d", expectedStatus, response.StatusCode)
		return nil, err
	}

	return response, nil
}
