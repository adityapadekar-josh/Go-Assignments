package repository

import (
	"context"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/adityapadekar-josh/assignment-8/internal/models"
)

type StatusChecker interface {
	Check(ctx context.Context, website string) (bool, error)
}

type HttpStatusChecker struct {
	Client *http.Client
}

func (h HttpStatusChecker) Check(ctx context.Context, website string) (bool, error) {

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, website, nil)
	if err != nil {
		slog.Error("Request creation failed", slog.String("Error:", err.Error()))
		return false, err
	}

	response, err := h.Client.Do(request)
	if err != nil {
		slog.Error("API Get request failed", slog.String("Error:", err.Error()))
		return false, err
	}

	defer response.Body.Close()

	return response.StatusCode >= 200 && response.StatusCode < 300, nil
}

func updateWebsiteStatus(wg *sync.WaitGroup, statusChecker StatusChecker, website string, dataStore *models.DataStore) {
	defer wg.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	status, err := statusChecker.Check(ctx, website)

	dataStore.RLock()
	defer dataStore.RUnlock()

	if err != nil || !status {
		dataStore.Data[website] = "DOWN"
		return
	}

	dataStore.Data[website] = "UP"
}

func StatusCheckerCronJob(ctx context.Context, statusChecker StatusChecker, dataStore *models.DataStore) {
	var wg sync.WaitGroup

	for website := range dataStore.Data {
		wg.Add(1)
		go updateWebsiteStatus(&wg, statusChecker, website, dataStore)
	}

	wg.Wait()
}

func RunCronJobs(ctx context.Context, websites *models.DataStore) {
	statusChecker := HttpStatusChecker{
		Client: &http.Client{},
	}
	for {
		time.Sleep(time.Minute)
		StatusCheckerCronJob(ctx, statusChecker, websites)
	}
}
