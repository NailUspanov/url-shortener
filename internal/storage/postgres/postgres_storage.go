package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
	"url-shortener/internal/model"
)

type StoragePostgres struct {
	db *sqlx.DB
}

func NewStoragePostgres(db *sqlx.DB) *StoragePostgres {
	return &StoragePostgres{db: db}
}

func (s *StoragePostgres) Create(shortUrl, longUrl string) error {
	createUrlQuery := fmt.Sprintf("INSERT INTO %s (long_url, short_url, expiration_date) VALUES ($1, $2, $3)"+
		" ON CONFLICT (short_url) DO NOTHING", urlsTable)
	_, err := s.db.Exec(createUrlQuery, longUrl, shortUrl, time.Now().AddDate(0, 0, 1))
	if err != nil {
		return err
	}
	return nil
}

func (s *StoragePostgres) Find(shortUrl string) (string, error) {
	var url model.URL
	findURLQuery := fmt.Sprintf("SELECT u.* FROM %s u WHERE u.short_url = $1", urlsTable)
	err := s.db.Get(&url, findURLQuery, shortUrl)
	return url.LongURL, err
}

func (s *StoragePostgres) Flush(period time.Duration) {

}
