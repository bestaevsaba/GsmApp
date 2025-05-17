package dbconfig

import (
	"app/cmd/logger"
	"encoding/json"
	"os"

	"go.uber.org/zap"
)

type Config struct {
	Database struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		DBName   string `json:"dbname"`
	} `json:"database"`
}

func LoadConfig() *Config {
	log := logger.GetLogger()
	path := "../../dbconfig/config.json"
	log.Info("Чтение конфига", zap.String("path", path))

	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Не удалось открыть config.json", zap.String("path", path), zap.Error(err))
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		log.Fatal("Ошибка парсинга config.json", zap.Error(err))
	}

	log.Info("Конфигурация БД загружена", zap.String("host", config.Database.Host),
		zap.String("dbname", config.Database.DBName))
	return &config
}
