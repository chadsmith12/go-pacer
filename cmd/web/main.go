package main

import (
	"log"
	"net/http"

	"github.com/chadsmith12/pacer/pkgs/pulse"
)

func main() {
	pulseApp := pulse.Pulse(":4500")

	pulseApp.Get("/hello", hello)

	if err := pulseApp.Start(); err != nil {
		log.Fatal(err)
	}
}

func hello(r *http.Request) pulse.PuleHttpWriter {
	var result = struct { Ok bool } { Ok: true }

	return pulse.Json(result)
}

