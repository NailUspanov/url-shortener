package _map

import (
	"errors"
	"sync"
	"time"
	"url-shortener/internal/model"
)

type MapStorage struct {
	storage *sync.Map
}

func NewMapStorage(storage *sync.Map) *MapStorage {
	return &MapStorage{storage: storage}
}

func (m *MapStorage) Create(shortUrl, longUrl string) (string, error) {
	url := model.URL{
		LongURL:        longUrl,
		ShortURL:       shortUrl,
		ExpirationDate: time.Now().AddDate(0, 0, 1),
	}
	val, loaded := m.storage.LoadOrStore(shortUrl, &url)
	if loaded {
		val.(*model.URL).ExpirationDate = url.ExpirationDate
	}
	return url.ShortURL, nil
}

func (m *MapStorage) Find(shortUrl string) (string, error) {
	value, ok := m.storage.Load(shortUrl)
	if !ok {
		return "", errors.New("error occurred during find operation in map storage")
	}

	if value.(*model.URL).ExpirationDate.Before(time.Now()) {
		return "", errors.New("the link has expired")
	}

	return value.(*model.URL).LongURL, nil
}

func (m *MapStorage) Flush(period time.Duration) {
	t := time.NewTicker(period * time.Second)
	defer t.Stop()
	for {
		select {
		case <-t.C: // Activate periodically
			m.storage.Range(func(key, value any) bool {
				if value.(*model.URL).ExpirationDate.Before(time.Now()) {
					m.storage.Delete(key)
				}
				return true
			})
		}
	}
}
