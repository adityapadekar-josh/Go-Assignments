package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/adityapadekar-josh/assignment-9/internal/models"
	"github.com/adityapadekar-josh/assignment-9/internal/pkg/response"
)

func AddWebsiteToWatchList(dataStore *models.DataStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestBody := &AddWebsiteToWatchListRequestBody{Data: nil}

		err := json.NewDecoder(r.Body).Decode(requestBody)
		if err != nil {
			slog.Error("Failed to parse request body", slog.String("Error:", err.Error()))
			response.WriteJson(w, http.StatusInternalServerError, "Failed to parse request body", nil)
			return
		}

		if len(requestBody.Data) == 0 {
			slog.Error("Empty request body")
			response.WriteJson(w, http.StatusBadRequest, "Empty request body", nil)
			return
		}

		dataStore.RLock()
		defer dataStore.RUnlock()

		for _, val := range requestBody.Data {
			if _, ok := dataStore.Data[val]; !ok {
				dataStore.Data[val] = "DOWN"
			}
		}

		response.WriteJson(w, 200, "Websites added successfully", nil)
	}
}

func GetWebsitesStatus(dataStore *models.DataStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if name == "" {
			response.WriteJson(w, 200, "Websites status fetched successfully", dataStore.Data)
			return
		}

		websiteStatus, ok := dataStore.Data[name]
		if !ok {
			slog.Error(fmt.Sprintf("Website with url: %s is not being tracked", name))
			response.WriteJson(w, http.StatusBadRequest, fmt.Sprintf("Website with url: %s is not being tracked", name), nil)
			return
		}

		response.WriteJson(w, 200, "Website status fetched successfully", map[string]string{name: websiteStatus})
	}
}
