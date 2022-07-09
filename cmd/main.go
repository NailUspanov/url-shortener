package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"url-shortener/internal/handlers"
	"url-shortener/internal/service"
	"url-shortener/internal/storage"
	_map "url-shortener/internal/storage/map"
	"url-shortener/internal/storage/postgres"
)

func main() {

	var urlStorage storage.URLStorage

	isRdbms, err := strconv.ParseBool(os.Getenv("RDBMS"))
	if err != nil {
		logrus.Fatalf("error occured while parsing RDBMS variable: %s", err.Error())
	}

	if isRdbms {
		db, err := postgres.NewPostgresDB(postgres.Config{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		})

		if err != nil {
			logrus.Fatalf("failed to initialize db %s", err.Error())
		}
		urlStorage = postgres.NewStoragePostgres(db)
		logrus.Println("RDBMS started")
	} else {
		urlStorage = _map.NewMapStorage(_map.NewMap())
		logrus.Println("Map storage started")
	}

	storages := storage.NewStorage(urlStorage)
	services := service.NewService(storages)
	handler := handlers.NewHandler(services)

	// каждые 24 часа запускается очистка неактивных ссылок
	go storages.URLStorage.Flush(86400) // 1 day

	srv := new(Server)
	if err := srv.Run(os.Getenv("PORT"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while runnning http server: %s", err.Error())
	}
}
