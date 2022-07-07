package service

import "url-shortener/internal/storage"

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type UrlService interface {
	Create(longUrl string) (string, error)
	Find(shortUrl string) (string, error)
}

type Service struct {
	UrlService
}

func NewService(storage *storage.Storage) *Service {
	return &Service{UrlService: NewURL(storage.URLStorage)}
}
