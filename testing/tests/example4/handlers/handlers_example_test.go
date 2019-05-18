// Sample to show how to write a basic example.
// godoc -http :3000

package handlers_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

func ExampleSendJSON() {
	r := httptest.NewRequest("GET", "/sendjson", nil)
	w := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, r)

	var u struct {
		Name  string
		Email string
	}

	if err := json.NewDecoder(w.Body).Decode(&u); err != nil {
		log.Println("Error: ", err)
	}
	fmt.Println(u)
	// Output:
	// {George george@email.com}
}
