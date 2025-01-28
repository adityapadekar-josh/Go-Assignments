package repository

import "github.com/adityapadekar-josh/assignment-8/internal/models"

func InitDataStore() *models.DataStore {
	return &models.DataStore{
		Data: make(map[string]string),
	}
}
