package main

import (
	"log"

	"github.com/Netflix-Clone-MicFlix/User-service/config"
	"github.com/Netflix-Clone-MicFlix/User-service/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
