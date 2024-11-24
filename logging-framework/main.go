package main

import (
	"log"

	"logging-framework/config"
	"logging-framework/domain"
	"logging-framework/logger"
)

func main() {
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	logr, err := logger.ConfigureLogger(cfg)
	if err != nil {
		log.Fatalf("could not configure logger: %v", err)
	}
	logr.Log("This is an error message", domain.ERROR, "main")
	logr.Log("This is an info message", domain.INFO, "main")
	logr.Log("This is a debug message", domain.DEBUG, "main")

}
