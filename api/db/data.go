package db

import (
	"app/api/models"
	"app/dbconfig"
	"fmt"

	"app/cmd/logger"

	"go.uber.org/zap"
)

func InsertGSMData(data models.GSMData) (int, string, error) {
	log := logger.GetLogger()
	db := dbconfig.GetDB()
	if db == nil {
		log.Error("База данных не инициализирована")
		return 1, "Ошибка соединения с базой данных", fmt.Errorf("db is not initialized")
	}

	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM gsm_data WHERE name = ?)", data.Name).Scan(&exists)
	if err != nil {
		log.Error("Ошибка проверки существования записи", zap.Error(err))
		return 1, "Ошибка при проверке записи", err
	}

	if exists {
		log.Info("Запись уже существует", zap.String("name", data.Name))
		return 0, "Такая запись уже существует", nil 
	}

	_, err = db.Exec("INSERT INTO gsm_data (name) VALUES (?)", data.Name)
	if err != nil {
		log.Error("Ошибка вставки записи", zap.String("name", data.Name), zap.Error(err))
		return 1, "Ошибка при добавлении записи", err
	}

	log.Info("Успешная вставка", zap.String("name", data.Name))
	return 0, "Данные успешно добавлены", nil 
}
