package main

import (
	"log"
	"testTaskGravitum/internal/app"
	"testTaskGravitum/internal/config"
)

// @title User API
// @version 1.0
// @description This is a simple User management API
// @host localhost:8080
// @BasePath /
func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
	app.New(cfg)
}
