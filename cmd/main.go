package main

import "fmt"

func main() {
	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}

	api := application{
		cfg: cfg,
	}

	fmt.Println(api)
}
