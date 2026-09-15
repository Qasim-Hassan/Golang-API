package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}

type application struct {
	cfg config
}

// mount
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// Base middleware stack
	r.Use(middleware.RequestID) // For Rate-Limiting
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // recover from crashes

	r.Use(middleware.Timeout(60 * time.Second)) // stops processing a request once the request timed out

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good"))
	})

	return r
}

// run server
func (app *application) run(h http.Handler) error {
	server := &http.Server{
		Addr:         app.cfg.addr,
		Handler:      h,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server is running on %v", app.cfg.addr)

	return server.ListenAndServe()
}
