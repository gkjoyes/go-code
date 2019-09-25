// Sample program that implement a simple web service that will allow us to explore
// how to use the GODEBUG variable.
package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var leak bool

func main() {
	http.HandleFunc("/sendjson", sendJSON)

	// Leak goroutine if we have any argument.
	if len(os.Args) == 2 {
		leak = true
	}

	log.Printf("listener: Started: Listening on: http://localhost:4000: Leak[%v]\n", leak)
	http.ListenAndServe(":4000", nil)
}

// sendJSON returns a simple JSON document.
func sendJSON(w http.ResponseWriter, r *http.Request) {

	// Lead a goroutine every so often.
	if leak {
		if rand.Intn(100) == 5 {
			go func() {
				for {
					time.Sleep(time.Millisecond * 10)
				}
			}()
		}
	}

	u := struct {
		Name  string
		Email string
	}{
		Name:  "george",
		Email: "george@mail.com",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&u)
}
