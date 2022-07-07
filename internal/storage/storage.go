package storage

import (
	"time"
)

type URLStorage interface {
	Create(shortUrl, longUrl string) error
	Find(shortUrl string) (string, error)
	Flush(period time.Duration)
}

type Storage struct {
	URLStorage
}

func NewStorage(URLStorage URLStorage) *Storage {
	return &Storage{URLStorage: URLStorage}
}
