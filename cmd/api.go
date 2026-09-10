package main

type application struct {
	cfg config
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
