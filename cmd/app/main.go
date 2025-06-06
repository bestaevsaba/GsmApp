package main

import (
	"net/http"

	"app/internal/logger"
	"app/dbconfig"
	api "app/internal/handlers"

	"go.uber.org/zap"
)

func main() {
	log := logger.GetLogger()
	log.Info("Запуск...")
	log.Info("Инициализация базы данных")
	dbconfig.InitDB()
	defer func() {
		log.Info("Закрытие соединения с БД")
		dbconfig.CloseDB()
	}()

	http.HandleFunc("/gsm", api.GSMHandler)
	http.HandleFunc("/gsm/get", api.GSMGetHandler)

	addr := "0.0.0.0:8080"
	log.Info("Сервер запускается", zap.String("address", addr))
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("Ошибка запуска HTTP‑сервера", zap.Error(err))
	}
}
