package postgres

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
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

	timeNow := now().AddDate(0, 0, 1)

	createUrlQuery := fmt.Sprintf("INSERT INTO %s (long_url, short_url, expiration_date) VALUES ($1, $2, $3)"+
		" ON CONFLICT (short_url) DO UPDATE SET expiration_date = $4", urlsTable)
	_, err := s.db.Exec(createUrlQuery, longUrl, shortUrl, timeNow, timeNow)

	if err != nil {
		return "", err
	}
	return shortUrl, nil
}

func (s *StoragePostgres) Find(shortUrl string) (string, error) {
	var url model.URL
	findURLQuery := fmt.Sprintf("SELECT u.* FROM %s u WHERE u.short_url = $1", urlsTable)
	err := s.db.Get(&url, findURLQuery, shortUrl)

	if url.ExpirationDate.Before(now()) {
		logrus.Println(url)
		return "", errors.New("the link has expired")
	}

	return url.LongURL, err
}

func (s *StoragePostgres) Flush(period time.Duration) {
	t := time.NewTicker(period * time.Second)
	defer t.Stop()
	for {
		select {
		case <-t.C: // Activate periodically
			deleteQuery := fmt.Sprintf("DELETE FROM %s u WHERE u.expiration_date <= $1::date", urlsTable)
			_, err := s.db.Exec(deleteQuery, now())
			if err != nil {
				logrus.Println("ERROR DURING FLUSH EXECUTION")
			}
		}
	}
}
