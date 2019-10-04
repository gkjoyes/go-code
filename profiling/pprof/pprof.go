// Sample program that implements a simple web service that will allow us to explore
// how to use the http/pprof tooling.

package main

import (
	"encoding/json"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/sendjson", sendJSON)
	log.Println("Listening on: http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}

// sendJSON returns a simple JSON document.
func sendJSON(w http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "george",
		Email: "george@email.com",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}
