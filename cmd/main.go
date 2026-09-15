package main

import (
	"log"
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

	h := api.mount()
	if err := api.run(h); err != nil {
		log.Printf("Server failed %v", err)
		os.Exit(500)
	}
}
