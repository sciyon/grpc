package main

import (
	"log"
	"net/http"
)

const (
	httpAddr = ":8923"
)

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)

	log.Printf("Starting HTTP server on %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start HTTP server")
	}
}
