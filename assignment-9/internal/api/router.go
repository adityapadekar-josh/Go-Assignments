package api

import (
	"net/http"

	"github.com/adityapadekar-josh/assignment-9/internal/models"
)

func NewRouter(dataStore *models.DataStore) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("POST /websites", AddWebsiteToWatchList(dataStore))
	router.HandleFunc("GET /websites", GetWebsitesStatus(dataStore))

	return router
}
