package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Redis *redis.Client
	Db    *sql.DB
}

func (cfg *Config) Status(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	dbStatus := "ok"
	redisStatus := "ok"
	statusCode := http.StatusOK

	if err := cfg.Db.Ping(); err != nil {
		dbStatus = "down"
		statusCode = http.StatusServiceUnavailable
	}

	ctx := context.Background()
	if err := cfg.Redis.Ping(ctx).Err(); err != nil {
		redisStatus = "down"
		statusCode = http.StatusServiceUnavailable
	}

	w.WriteHeader(statusCode)

	_ = json.NewEncoder(w).Encode(map[string]any{
		"status": "ok",
		"db":     dbStatus,
		"redis":  redisStatus,
		"time":   time.Now().UTC().Format(time.RFC3339),
	})
}
