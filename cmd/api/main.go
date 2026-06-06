package main

import (
	"log"
	"net/http"
	"time"

	"github.com/igortuchel/sec-esp-server/internal/config"
	"github.com/igortuchel/sec-esp-server/internal/handlers"
)

func main() {
	db, err := config.OpenSQLite("data/app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	h := &handlers.Handler{DB: db}

	mux := http.NewServeMux()
	mux.HandleFunc("/status", h.Status)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Println("server listening on :8080")
	log.Fatal(srv.ListenAndServe())
}
