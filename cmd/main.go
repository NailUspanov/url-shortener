package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
	"url-shortener/internal/handlers"
	"url-shortener/internal/service"
	"url-shortener/internal/storage"
	_map "url-shortener/internal/storage/map"
)

func main() {
	//db, err := postgres.NewPostgresDB(postgres.Config{
	//	Host:     os.Getenv("DB_HOST"),
	//	Port:     os.Getenv("DB_PORT"),
	//	Username: os.Getenv("DB_USERNAME"),
	//	Password: os.Getenv("DB_PASSWORD"),
	//	DBName:   os.Getenv("DB_NAME"),
	//	SSLMode:  os.Getenv("DB_SSLMODE"),
	//})
	//
	//if err != nil {
	//	logrus.Fatalf("failed to initialize db %s", err.Error())
	//}

	//прокидываю инстанс бд и создаю репозитории
	storages := storage.NewStorage(_map.NewMapStorage(_map.NewMap()))
	//прокидываю репозитории в сервисы
	services := service.NewService(storages)
	//сервисы в хендлеры
	handler := handlers.NewHandler(services)

	//запускаю сервер на порту 8000
	srv := new(Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while runnning http server: %s", err.Error())
	}
}

func periodicGreet(msg string, period time.Duration) {
	t := time.NewTicker(period * time.Second)
	defer t.Stop()
	for {
		select {
		case <-t.C: // Activate periodically
			fmt.Printf(msg)
		}
	}
}
