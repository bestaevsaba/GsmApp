package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"app/internal/db"
	"app/internal/logger"

	"go.uber.org/zap"
)

func GSMGetHandler(w http.ResponseWriter, r *http.Request) {
	log := logger.GetLogger()

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()
	fromStr := query.Get("from")
	toStr := query.Get("to")
	pageStr := query.Get("page")
	sizeStr := query.Get("size")

	from, err := time.Parse("2006-01-02", fromStr)
	if err != nil {
		log.Warn("Неверный формат даты 'from'", zap.String("from", fromStr), zap.Error(err))
		http.Error(w, "Неверный формат даты from (ожидается YYYY-MM-DD)", http.StatusBadRequest)
		return
	}
	to, err := time.Parse("2006-01-02", toStr)
	if err != nil {
		log.Warn("Неверный формат даты 'to'", zap.String("to", toStr), zap.Error(err))
		http.Error(w, "Неверный формат даты to (ожидается YYYY-MM-DD)", http.StatusBadRequest)
		return
	}

	page, _ := strconv.Atoi(pageStr)
	size, _ := strconv.Atoi(sizeStr)
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	offset := (page - 1) * size

	log.Info("Обработка GET запроса к /gsm",
		zap.Time("from", from),
		zap.Time("to", to),
		zap.Int("page", page),
		zap.Int("size", size),
	)

	results, err := db.GetGSMData(from, to, size, offset)
	if err != nil {
		log.Error("Ошибка при получении данных из БД", zap.Error(err))
		http.Error(w, "Ошибка запроса к БД", http.StatusInternalServerError)
		return
	}

	log.Info("Данные успешно получены", zap.Int("кол-во записей", len(results)))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":  results,
		"page":  page,
		"size":  size,
		"total": len(results),
	})
}
