package main

import (
	"log/slog"
	"os"
)

func main() {
	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}

	api := application{
		cfg: cfg,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil)) // structured logging for INFO or ERRORS
	slog.SetDefault(logger)

	h := api.mount()

	if err := api.run(h); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(500)
	}
}
