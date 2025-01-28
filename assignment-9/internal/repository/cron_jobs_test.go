package repository

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/adityapadekar-josh/assignment-9/internal/models"
	stcMock "github.com/adityapadekar-josh/assignment-9/internal/repository/mocks"

	"github.com/undefinedlabs/go-mpatch"
)

const (
	instagram = "https://instagram.com"
	facebook  = "https://facebook.com"
	x         = "https://x.com"
)

var currentTime = time.Now()

func TestUpdateWebsiteStatus(t *testing.T) {

	mpatch.PatchMethod(time.Now, func() time.Time {
		return currentTime
	})

	stc := stcMock.NewStatusChecker(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tests := []struct {
		name     string
		setup    func(context.Context)
		website  string
		websites models.DataStore
	}{
		{
			"Empty string",
			func(ctx context.Context) {
				stc.On("Check", ctx, instagram).Return(true, nil)
			},
			instagram,
			models.DataStore{Data: map[string]string{instagram: "DOWN"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(ctx)
			var wg sync.WaitGroup
			wg.Add(1)
			go updateWebsiteStatus(context.Background(), &wg, stc, tt.website, &tt.websites)
			wg.Wait()
			fmt.Println(tt.websites.Data[tt.website])
		})
	}
}
