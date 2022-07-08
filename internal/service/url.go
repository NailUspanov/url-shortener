package service

import (
	"sync/atomic"
	"url-shortener/internal/helpers"
	"url-shortener/internal/storage"
)

type URL struct {
	storage storage.URLStorage
}

type count uint64

var c count = 100000000000000000

func NewURL(storage storage.URLStorage) *URL {
	return &URL{storage: storage}
}

func (u *URL) Create(longUrl string) (string, error) {
	c.inc()
	shortUrl := helpers.ToBase63(c.get())
	shortUrl, err := u.storage.Create(shortUrl, longUrl)
	if err != nil {
		return "", err
	}
	return shortUrl, err
}

func (u *URL) Find(shortUrl string) (string, error) {
	return u.storage.Find(shortUrl)
}

func (c *count) inc() uint64 {
	return atomic.AddUint64((*uint64)(c), 1)
}

func (c *count) get() uint64 {
	return atomic.LoadUint64((*uint64)(c))
}
