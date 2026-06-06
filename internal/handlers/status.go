package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

type Handler struct {
	DB *sql.DB
}

func (h *Handler) Status(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	dbStatus := "ok"
	statusCode := http.StatusOK

	if err := h.DB.Ping(); err != nil {
		dbStatus = "down"
		statusCode = http.StatusServiceUnavailable
	}

	w.WriteHeader(statusCode)

	_ = json.NewEncoder(w).Encode(map[string]any{
		"status": "ok",
		"db":     dbStatus,
		"time":   time.Now().UTC().Format(time.RFC3339),
	})
}
