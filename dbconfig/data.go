package dbconfig

import (
	"database/sql"
	"fmt"

	"app/internal/logger"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

var db *sql.DB

func InitDB() {
	log := logger.GetLogger()
	cfg := LoadConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
	)
	log.Info("Формирование DSN", zap.String("dsn", dsn))

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Ошибка открытия соединения с БД", zap.Error(err))
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Ошибка пинга БД", zap.Error(err))
	}

	log.Info("Соединение с базой данных установлено")
}

func CloseDB() {
	log := logger.GetLogger()
	if db != nil {
		if err := db.Close(); err != nil {
			log.Error("Ошибка закрытия соединения с БД", zap.Error(err))
		} else {
			log.Info("Соединение с БД закрыто")
		}
	}
}

func GetDB() *sql.DB {
	return db
}
