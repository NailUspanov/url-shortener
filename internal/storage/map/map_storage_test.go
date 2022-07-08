package _map

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"url-shortener/internal/model"
)

func TestMapStorage_Create(t *testing.T) {
	asserts := assert.New(t)

	var tests = []struct {
		name          string
		inputShortUrl string
		inputLongUrl  string
		storedData    *model.URL
		errorExpected bool
		err           error
		expected      string
	}{
		{
			name:          "OK",
			inputLongUrl:  "http:/somelongurl.com",
			inputShortUrl: "amw2MskMwq",
			errorExpected: false,
			storedData: &model.URL{
				Id:             0,
				LongURL:        "",
				ShortURL:       "",
				ExpirationDate: time.Time{},
			},
		},
		{
			name:          "collision",
			inputLongUrl:  "http:/somelongurl.com",
			inputShortUrl: "amw2MskMwq",
			errorExpected: true,
			storedData: &model.URL{
				Id:             0,
				LongURL:        "http:/website.com",
				ShortURL:       "amw2MskMwq",
				ExpirationDate: time.Now().AddDate(0, 0, 1),
			},
			expected: "amw2MskMwq",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			syncmap := NewMap()
			syncmap.Store(test.storedData.ShortURL, test.storedData)
			storage := NewMapStorage(syncmap)
			created, err := storage.Create(test.inputShortUrl, test.inputLongUrl)
			if test.errorExpected {
				asserts.Equal(test.expected, created)
			} else {
				asserts.NoError(err)
			}
		})
	}
}

func TestMapStorage_Find(t *testing.T) {
	asserts := assert.New(t)

	var tests = []struct {
		name          string
		inputShortUrl string
		storedData    *model.URL
		errorExpected bool
		err           error
	}{
		{
			name:          "OK",
			inputShortUrl: "amw2MskMwq",
			errorExpected: false,
			storedData: &model.URL{
				Id:             0,
				LongURL:        "http:/somelongurl.com",
				ShortURL:       "amw2MskMwq",
				ExpirationDate: time.Now().AddDate(0, 0, 1),
			},
		},
		{
			name:          "No such url in storage",
			inputShortUrl: "amw2MskMwq",
			errorExpected: true,
			storedData: &model.URL{
				Id:             0,
				LongURL:        "http:/somelongurl.com",
				ShortURL:       "aa",
				ExpirationDate: time.Time{},
			},
			err: errors.New("error occurred during find operation in map storage"),
		},
		{
			name:          "Link has expired",
			inputShortUrl: "amw2MskMwq",
			errorExpected: true,
			storedData: &model.URL{
				Id:             0,
				LongURL:        "http:/somelongurl.com",
				ShortURL:       "amw2MskMwq",
				ExpirationDate: time.Now(),
			},
			err: errors.New("the link has expired"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			syncmap := NewMap()
			syncmap.Store(test.storedData.ShortURL, test.storedData)
			storage := NewMapStorage(syncmap)

			time.Sleep(100 * time.Millisecond)
			v, err := storage.Find(test.inputShortUrl)
			if test.errorExpected {
				asserts.Equal(test.err, err)
			} else {
				asserts.NoError(err)
				asserts.Equal(test.storedData.LongURL, v)
			}
		})
	}
}
