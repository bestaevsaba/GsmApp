package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"app/internal/db"
	"app/internal/logger"
	"app/internal/models"

	"go.uber.org/zap"
)

func GSMHandler(w http.ResponseWriter, r *http.Request) {
	log := logger.GetLogger()

	if r.Method != http.MethodPost {
		http.Error(w, "Только POST поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var data models.GSMData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Невалидный JSON", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(data.RecordedAt) == "" {
		http.Error(w, "Поле recorded_at обязательно", http.StatusBadRequest)
		return
	}

	formats := []string{
		"02.01.2006 15:04:05",
		"02.01.2006 15:04",
		time.RFC3339,
	}

	var parsedTime time.Time
	var parseErr error
	for _, f := range formats {
		parsedTime, parseErr = time.Parse(f, data.RecordedAt)
		if parseErr == nil {
			break
		}
	}

	if parseErr != nil {
		http.Error(w, "Неверный формат времени. Допустимые форматы: DD.MM.YYYY HH:MM[:SS] или RFC3339", http.StatusBadRequest)
		return
	}

	log.Info("Получен POST запрос с recorded_at", zap.Time("parsedTime", parsedTime))

	code, message, err := db.InsertGSMData(parsedTime)
	if err != nil {
		log.Error("Ошибка записи в БД", zap.Error(err))
		http.Error(w, fmt.Sprintf("Ошибка записи в БД: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    code,
		"message": message,
	})
}
