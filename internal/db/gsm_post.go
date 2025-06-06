package db

import (
	"time"

	"app/dbconfig"
	"app/internal/logger"

	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func InsertGSMData(t time.Time) (int, string, error) {
	log := logger.GetLogger()
	db := dbconfig.GetDB()
	if db == nil {
		log.Error("База данных не инициализирована")
		return 1, "Ошибка соединения с базой данных", ErrDBNotInitialized
	}

	_, err := db.Exec("INSERT INTO gsm_data (recorded_at) VALUES (?)", t)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				log.Warn("Данные уже существуют", zap.Time("time", t))
				return 2, "Данные уже существуют", nil
			}
		}
		log.Error("Ошибка вставки записи", zap.Time("time", t), zap.Error(err))
		return 1, "Ошибка при добавлении записи", err
	}

	log.Info("Успешная вставка", zap.Time("time", t))
	return 0, "Данные успешно добавлены", nil
}
