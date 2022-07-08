package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"time"
	"url-shortener/internal/helpers"
	"url-shortener/internal/model"
)

type StoragePostgres struct {
	db *sqlx.DB
}

var now = time.Now

func NewStoragePostgres(db *sqlx.DB) *StoragePostgres {
	return &StoragePostgres{db: db}
}

func (s *StoragePostgres) Create(shortUrl, longUrl string) (string, error) {

	createUrlQuery := fmt.Sprintf("INSERT INTO %s (long_url, short_url, expiration_date) VALUES ($1, $2, $3)", urlsTable)
	_, err := s.db.Exec(createUrlQuery, longUrl, shortUrl, now().AddDate(0, 0, 1))

	pqerr, ok := err.(*pq.Error)
	if ok && pqerr.Code == "23505" { // checks for the collision
		shortUrl = helpers.Encode(10)
		shortUrl, err = s.Create(shortUrl, longUrl)
	}
	if err != nil {
		return "", err
	}
	return shortUrl, nil
}

func (s *StoragePostgres) Find(shortUrl string) (string, error) {
	var url model.URL
	findURLQuery := fmt.Sprintf("SELECT u.* FROM %s u WHERE u.short_url = $1", urlsTable)
	err := s.db.Get(&url, findURLQuery, shortUrl)
	return url.LongURL, err
}

func (s *StoragePostgres) Flush(period time.Duration) {

}
