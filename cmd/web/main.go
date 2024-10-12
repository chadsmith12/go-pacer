package main

import (
	"log"

	"github.com/chadsmith12/pacer/internal/app"
)

func main() {
	app := app.New()

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

