package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

// Structures
type Websites struct {
	Data map[string]string `json:"data"`
	sync.RWMutex
}

func main() {
	var websites = &Websites{
		Data: make(map[string]string),
	}

	go RunCronJobs(websites)

	router := createNewRouter(websites)

	setUpServer(router)
}

// Server Setup
func setUpServer(router *http.ServeMux) {
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	slog.Info("Server listening at", slog.String("address", "localhost:8080"))

	if err := server.ListenAndServe(); err != nil {
		slog.Error("Server error", slog.String("error", err.Error()))
	}
}

// Routing Setup
func createNewRouter(websites *Websites) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("POST /websites", AddWebsiteToWatchList(websites))
	router.HandleFunc("GET /websites", GetWebsitesStatus(websites))

	return router
}

// Utils
type Response struct {
	Message string      `json:"message"`
	Result  interface{} `json:"result,omitempty"`
}

func WriteJson(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	response := Response{
		Message: message,
		Result:  data,
	}

	marshaledResponse, err := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message" : "Failed to marshal response" }`))
		return
	}

	w.WriteHeader(statusCode)
	w.Write(marshaledResponse)

}

// Handlers
func AddWebsiteToWatchList(websites *Websites) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestBody := &struct {
			Data []string `json:"data"`
		}{Data: nil}

		err := json.NewDecoder(r.Body).Decode(requestBody)
		if err != nil {
			WriteJson(w, http.StatusInternalServerError, "Failed to parse request body", nil)
			return
		}

		if len(requestBody.Data) == 0 {
			WriteJson(w, http.StatusBadRequest, "Empty request body", nil)
			return
		}

		websites.RLock()
		defer websites.RUnlock()

		for _, val := range requestBody.Data {
			if _, ok := websites.Data[val]; !ok {
				websites.Data[val] = "DOWN"
			}
		}

		WriteJson(w, 200, "Websites added successfully", nil)
	}
}

func GetWebsitesStatus(websites *Websites) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if name == "" {
			WriteJson(w, 200, "Websites status fetched successfully", websites.Data)
			return
		}

		websiteStatus, ok := websites.Data[name]
		if !ok {
			WriteJson(w, http.StatusBadRequest, fmt.Sprintf("Website with url: %s is not being tracked", name), nil)
			return
		}

		WriteJson(w, 200, "Website status fetched successfully", map[string]string{name: websiteStatus})
	}
}

// Status checker
type StatusChecker interface {
	Check(ctx context.Context, website string) (bool, error)
}

type HttpStatusChecker struct {
}

func (h HttpStatusChecker) Check(ctx context.Context, website string) (bool, error) {
	response, err := http.Get(website)

	if err != nil {
		return false, err
	}

	return response.StatusCode >= 200 && response.StatusCode < 300, nil
}

func updateWebsiteStatus(wg *sync.WaitGroup, statusChecker StatusChecker, website string, websites *Websites) {
	defer wg.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	status, err := statusChecker.Check(ctx, website)

	websites.RLock()
	defer websites.RUnlock()

	if err != nil || !status {
		websites.Data[website] = "DOWN"
		return
	}

	websites.Data[website] = "UP"
}

// Cron job
func StatusCheckerCronJob(websites *Websites) {
	var wg sync.WaitGroup

	statusChecker := HttpStatusChecker{}

	for website := range websites.Data {
		wg.Add(1)
		go updateWebsiteStatus(&wg, statusChecker, website, websites)
	}

	wg.Wait()
}

func RunCronJobs(websites *Websites) {
	for {
		time.Sleep(time.Minute)
		StatusCheckerCronJob(websites)
	}
}
