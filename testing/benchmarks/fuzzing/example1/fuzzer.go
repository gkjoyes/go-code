// +build gofuzz

package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
)

func init() {
	Routes()
}

// Fuzz is executed by the go-fuzz tool. Input data modifications are provided and
// used to validate API call.
func Fuzz(data []byte) int {
	r := httptest.NewRequest("POST", "/process", bytes.NewBuffer(data))
	w := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, r)

	if w.Code != http.StatusOK {

		// Report the data that produced this error as not interesting.
		return 0
	}

	// Report the data that did not cause an error as interesting.
	return 1
}
