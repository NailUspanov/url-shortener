package postgres

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"testing"
	"time"
)

func TestStoragePostgres_Create(t *testing.T) {

	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	clock := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	r := NewStoragePostgres(db)

	tests := []struct {
		name          string
		mock          func()
		inputShortUrl string
		inputLongUrl  string
		wantErr       bool
	}{
		{
			name: "Ok",
			mock: func() {
				mock.ExpectExec(fmt.Sprintf("INSERT INTO %s", urlsTable)).
					WithArgs("https://longurl.com", "sqwj2NSjka", clock.AddDate(0, 0, 1), clock.AddDate(0, 0, 1)).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			inputShortUrl: "sqwj2NSjka",
			inputLongUrl:  "https://longurl.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			now = func() time.Time {
				return clock
			}
			_, err := r.Create(tt.inputShortUrl, tt.inputLongUrl)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
