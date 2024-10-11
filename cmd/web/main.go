package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("wwwroot")))
	mux.HandleFunc("GET /api/hello", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("{\"hello\": \"world\"}")
	})

	err := http.ListenAndServe(":6969", mux)
	if err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
