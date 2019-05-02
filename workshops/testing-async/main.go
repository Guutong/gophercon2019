package main

import (
	"log"
	"net/http"

	"github.com/gopherguides/learn/_training/testing/async/src/httpd"
	"github.com/gopherguides/learn/_training/testing/async/src/keys"
)

func main() {
	// Get a new instance of our API service
	handler := httpd.New()
	// Back the API service with a store implementation
	handler.Store = keys.NewStore()

	log.Println("starting server on http://localhost:8080")
	// Start up the API service on port 8080
	log.Fatal(http.ListenAndServe(":8080", handler))
}
