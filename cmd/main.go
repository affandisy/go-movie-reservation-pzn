package main

import (
	"go-movie-reservation/config"
	"go-movie-reservation/internal/routes"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := config.NewDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	router := routes.SetupRouter(db)
	router.Run(":8080")
}
