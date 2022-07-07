package storage

type URLStorage interface {
	Create(shortUrl, longUrl string) error
	Find(shortUrl string) (string, error)
}

type Storage struct {
	URLStorage
}

func NewStorage(URLStorage URLStorage) *Storage {
	return &Storage{URLStorage: URLStorage}
}
