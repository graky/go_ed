package main

import (
    "log"

    "go_ed/internal/server"
    "go_ed/pkg/config"
    "go_ed/pkg/database"
)

func main() {
	cfg := config.Load()
	db, err := database.InitDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	srv := server.New(cfg, db)
	log.Printf("Starting server on %s", cfg.ServerAddress)
	log.Fatal(srv.Run())
}