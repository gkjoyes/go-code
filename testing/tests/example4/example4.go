package main

import (
	"log"
	"net/http"

	"github.com/george-kj/go-code/testing/tests/example4/handlers"
)

func main() {
	handlers.Routes()

	log.Println("Listener: Started: Listing on: http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}
