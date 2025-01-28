package repository

import "github.com/adityapadekar-josh/assignment-9/internal/models"

func InitDataStore() *models.DataStore {
	return &models.DataStore{
		Data: make(map[string]string),
	}
}
