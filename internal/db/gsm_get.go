package db

import (
	"errors"
	"time"

	"app/dbconfig"
	"app/internal/logger"
	"app/internal/models"

	"go.uber.org/zap"
)

var ErrDBNotInitialized = errors.New("database is not initialized")

func GetGSMData(from, to time.Time, limit, offset int) ([]models.GSMData, error) {
	log := logger.GetLogger()
	db := dbconfig.GetDB()
	if db == nil {
		log.Error("База данных не инициализирована")
		return nil, ErrDBNotInitialized
	}

	log.Info("Выполняется запрос к базе данных",
		zap.Time("from", from),
		zap.Time("to", to),
		zap.Int("limit", limit),
		zap.Int("offset", offset),
	)

	rows, err := db.Query(`
		SELECT recorded_at 
		FROM gsm_data 
		WHERE recorded_at BETWEEN ? AND ?
		ORDER BY recorded_at
		LIMIT ? OFFSET ?`, from, to, limit, offset)
	if err != nil {
		log.Error("Ошибка при выполнении запроса SELECT", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var results []models.GSMData
	for rows.Next() {
		var data models.GSMData
		if err := rows.Scan(&data.RecordedAt); err != nil {
			log.Error("Ошибка при чтении строки результата", zap.Error(err))
			return nil, err
		}
		results = append(results, data)
	}

	log.Info("Запрос выполнен успешно", zap.Int("кол-во записей", len(results)))
	return results, nil
}
