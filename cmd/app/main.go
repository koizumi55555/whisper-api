package main

import (
	"koizumi55555/whisper-api/config"
	"koizumi55555/whisper-api/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	if err := app.Run(cfg); err != nil {
		log.Fatalf("runtime error %s", err)
	}
}
