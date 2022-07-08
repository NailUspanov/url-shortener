package service

import (
	"url-shortener/internal/helpers"
	"url-shortener/internal/storage"
)

type URL struct {
	storage storage.URLStorage
}

func NewURL(storage storage.URLStorage) *URL {
	return &URL{storage: storage}
}

func (u *URL) Create(longUrl string) (string, error) {
	shortUrl := helpers.Encode(10)
	shortUrl, err := u.storage.Create(shortUrl, longUrl)
	if err != nil {
		return "", err
	}
	return shortUrl, err
}

func (u *URL) Find(shortUrl string) (string, error) {
	return u.storage.Find(shortUrl)
}
