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

func (m *MapStorage) Create(shortUrl, longUrl string) error {
	url := model.URL{
		LongURL:        longUrl,
		ShortURL:       shortUrl,
		ExpirationDate: time.Now().AddDate(0, 0, 1),
	}
	_, loaded := m.storage.LoadOrStore(shortUrl, &url)
	if loaded {
		return errors.New("error occurred during create operation in map storage")
	}
	return nil
}

func (m *MapStorage) Find(shortUrl string) (string, error) {
	value, ok := m.storage.Load(shortUrl)
	if !ok {
		return "", errors.New("error occurred during find operation in map storage")
	}
	return value.(*model.URL).LongURL, nil
}
