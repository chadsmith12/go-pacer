package main

import (
	"log"

	"github.com/chadsmith12/pacer/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to startup server - %v", err)
	}
	app := app.New()

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

