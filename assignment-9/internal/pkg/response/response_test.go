package response

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteJson(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		message    string
		result     interface{}
	}{
		{"Request with empty body", http.StatusOK, "Success", nil},
		{"Request with body of type number", http.StatusOK, "Success", 1000},
		{"Request with body of type map", http.StatusOK, "Success", map[string]string{"key": "value"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()

			WriteJson(recorder, tt.statusCode, tt.message, tt.result)

			result := recorder.Result()
			defer result.Body.Close()

			contentType := result.Header.Get("Content-Type")
			if contentType != "application/json" {
				t.Errorf("expected Content-Type 'application/json', got '%s'", contentType)
			}

			var responseBody Response
			err := json.NewDecoder(result.Body).Decode(&responseBody)
			if err != nil {
				t.Fatalf("failed to decode response body: %v", err)
			}

			if result.StatusCode != tt.statusCode {
				t.Errorf("expected status code %d, got %d", tt.statusCode, result.StatusCode)
			}

			if responseBody.Message != tt.message {
				t.Errorf("expected message '%s', got '%s'", tt.message, responseBody.Message)
			}
		})
	}
}
