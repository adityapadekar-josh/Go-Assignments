package main

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/adityapadekar-josh/assignment-8/internal/api"
	"github.com/adityapadekar-josh/assignment-8/internal/repository"
)

func main() {
	ctx := context.Background()

	dataStore := repository.InitDataStore()

	go repository.RunCronJobs(ctx, dataStore)

	router := api.NewRouter(dataStore)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	slog.Info("Server listening at", slog.String("address", "localhost:8080"))

	if err := server.ListenAndServe(); err != nil {
		slog.Error("Server error", slog.String("error", err.Error()))
	}
}
