package http_request

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// DoRequest performs an HTTP request for the specified method.
func HttpRequest(method, url string, body []byte, headers map[string]string) (string, error) {
	// Create a new request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers if provided
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(responseBody), nil
}
