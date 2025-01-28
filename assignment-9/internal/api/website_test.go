package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adityapadekar-josh/assignment-9/internal/models"
	"github.com/adityapadekar-josh/assignment-9/internal/pkg/response"
)

const (
	instagram = "https://instagram.com"
	facebook  = "https://facebook.com"
	x         = "https://x.com"
)

func TestAddWebsiteToWatchListHandler(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		body       []byte
		websites   models.DataStore
	}{
		{
			"Add website",
			http.StatusOK,
			[]byte(fmt.Sprintf(`{ "data" : ["%s"]}`, instagram)),
			models.DataStore{Data: make(map[string]string)},
		},
		{
			"Add multiple websites",
			http.StatusOK,
			[]byte(fmt.Sprintf(`{ "data" : ["%s", "%s"]}`, instagram, facebook)),
			models.DataStore{Data: make(map[string]string)},
		},
		{
			"Request with empty body",
			http.StatusBadRequest,
			[]byte(`{ "data" : []}`),
			models.DataStore{Data: make(map[string]string)},
		},
		{
			"Invalid json",
			http.StatusInternalServerError,
			[]byte(`{ data : false}`),
			models.DataStore{Data: make(map[string]string)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest("POST", "/websites", bytes.NewBuffer(tt.body))
			recorder := httptest.NewRecorder()

			AddWebsiteToWatchList(&tt.websites)(recorder, request)

			result := recorder.Result()
			defer result.Body.Close()

			var responseBody response.Response
			err := json.NewDecoder(result.Body).Decode(&responseBody)
			if err != nil {
				t.Fatalf("failed to decode response body: %v", err)
			}

			if result.StatusCode != tt.statusCode {
				t.Errorf("expected status code %d, got %d", tt.statusCode, result.StatusCode)
			}
		})
	}
}

func TestGetWebsiteStatus(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		query      string
		websites   models.DataStore
	}{
		{
			"No websites tracked",
			http.StatusOK,
			"",
			models.DataStore{Data: make(map[string]string)},
		},
		{
			"No websites tracked but fetching a website status",
			http.StatusBadRequest,
			instagram,
			models.DataStore{Data: make(map[string]string)},
		},
		{
			"Websites tracked",
			http.StatusOK,
			"",
			models.DataStore{Data: map[string]string{instagram: "UP", facebook: "DOWN"}},
		},
		{
			"Websites tracker and fetching a website status",
			http.StatusOK,
			instagram,
			models.DataStore{Data: map[string]string{instagram: "UP", facebook: "DOWN"}},
		},
		{
			"Websites tracker and fetching a website status that does not exists in data",
			http.StatusBadRequest,
			x,
			models.DataStore{Data: map[string]string{instagram: "UP", facebook: "DOWN"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest("GET", fmt.Sprintf("/websites?name=%s", tt.query), nil)
			recorder := httptest.NewRecorder()

			GetWebsitesStatus(&tt.websites)(recorder, request)

			result := recorder.Result()
			defer result.Body.Close()

			var responseBody response.Response
			err := json.NewDecoder(result.Body).Decode(&responseBody)
			if err != nil {
				t.Fatalf("failed to decode response body: %v", err)
			}

			if result.StatusCode != tt.statusCode {
				t.Errorf("expected status code %d, got %d", tt.statusCode, result.StatusCode)
			}
		})
	}
}
