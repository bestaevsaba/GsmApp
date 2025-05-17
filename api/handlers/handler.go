package api

import (
	"app/api/db"
	"app/api/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func GSMHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Только POST поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var data models.GSMData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Невалидный JSON", http.StatusBadRequest)
		return
	}

	code, message, err := db.InsertGSMData(data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка записи в БД: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if code == 0 {
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    code,
		"message": message,
	})
}
