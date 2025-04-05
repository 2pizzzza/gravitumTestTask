package main

import (
	"log"
	"testTaskGravitum/internal/app"
	"testTaskGravitum/internal/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
	app.New(cfg)
}
