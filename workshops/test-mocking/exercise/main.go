package main

import (
	"log"
	"net/http"

	"github.com/gopherguides/learn/_training/testing/mocking/src/exercise/httpd"
	"github.com/gopherguides/learn/_training/testing/mocking/src/exercise/keys"
)

func main() {
	// section: main
	handler := httpd.NewHandler()
	handler.Store = keys.NewStore()

	log.Println("starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
	// section: main
}
