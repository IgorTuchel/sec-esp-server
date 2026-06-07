package main

import (
	"log"
	"net/http"
	"time"

	"github.com/igortuchel/sec-esp-server/internal/config"
	"github.com/igortuchel/sec-esp-server/internal/handlers"
)

func main() {
	cfg := &handlers.Config{}
	config.Startup(cfg)
	defer cfg.Db.Close()
	defer cfg.Redis.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/status", cfg.Status)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Println("server listening on :8080")
	log.Fatal(srv.ListenAndServe())
}
