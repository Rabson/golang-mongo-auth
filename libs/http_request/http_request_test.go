package http_request

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHttpRequest(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		url            string
		body           []byte
		headers        map[string]string
		mockStatusCode int
		mockResponse   string
		expectError    bool
	}{
		{
			name:           "GET request success",
			method:         http.MethodGet,
			url:            "/test",
			body:           nil,
			headers:        map[string]string{"Content-Type": "application/json"},
			mockStatusCode: http.StatusOK,
			mockResponse:   `{"message": "success"}`,
			expectError:    false,
		},
		{
			name:           "POST request success",
			method:         http.MethodPost,
			url:            "/test",
			body:           []byte(`{"key": "value"}`),
			headers:        map[string]string{"Content-Type": "application/json"},
			mockStatusCode: http.StatusCreated,
			mockResponse:   `{"message": "created"}`,
			expectError:    false,
		},
		{
			name:           "Request failure",
			method:         http.MethodGet,
			url:            "/test",
			body:           nil,
			headers:        nil,
			mockStatusCode: http.StatusInternalServerError,
			mockResponse:   `{"error": "internal server error"}`,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock server
			mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != tt.url {
					t.Errorf("expected URL %s, got %s", tt.url, r.URL.Path)
				}
				for key, value := range tt.headers {
					if r.Header.Get(key) != value {
						t.Errorf("expected header %s to be %s, got %s", key, value, r.Header.Get(key))
					}
				}
				w.WriteHeader(tt.mockStatusCode)
				w.Write([]byte(tt.mockResponse))
			}))
			defer mockServer.Close()

			// Perform the HTTP request
			response, err := HttpRequest(tt.method, mockServer.URL+tt.url, tt.body, tt.headers)

			// Validate the results
			if tt.expectError {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("did not expect an error but got: %v", err)
				}
				if strings.TrimSpace(response) != tt.mockResponse {
					t.Errorf("expected response %s, got %s", tt.mockResponse, response)
				}
			}
		})
	}
}
