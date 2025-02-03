package models

import "sync"

type DataStore struct {
	Data map[string]string `json:"data"`
	sync.RWMutex
}
