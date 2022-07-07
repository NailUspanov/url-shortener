package model

import "time"

type URL struct {
	Id             int       `json:"id,omitempty" db:"id"`
	LongURL        string    `json:"long_url" db:"long_url"`
	ShortURL       string    `json:"short_url" db:"short_url"`
	ExpirationDate time.Time `json:"expiration_date" db:"expiration_date"`
}
