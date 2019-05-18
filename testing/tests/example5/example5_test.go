// go test -v
// go test -run TestDownload/statusok -v
// go test -run TestDownload/statusnotfound -v
// go test -run TestParallelize -v

// Sample test to show how to write a basic sub unit table test.
package example5

import (
	"net/http"
	"testing"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

// TestDownload validates the http Get function can download content and handles
// different status conditions properly.
func TestDownload(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		statusCode int
	}{
		{"statusok", "https://www.ardanlabs.com/blog/index.xml", http.StatusOK},
		{"statusnotfound", "http://rss.cnn.com/rss/cnn_topstorie.rss", http.StatusNotFound},
	}

	t.Log("Given the need to test downloading different content.")
	{
		for i, tt := range tests {
			tf := func(t *testing.T) {
				t.Logf("\tTest: %d\tWhen checking %q for status code %d", i, tt.url, tt.statusCode)
				{
					resp, err := http.Get(tt.url)
					if err != nil {
						t.Fatalf("\t%s\tShould be able to make the Get call: %v", failed, err)
					}
					t.Logf("\t%s\tShould be able to make the Get call.", succeed)

					defer resp.Body.Close()

					if resp.StatusCode == tt.statusCode {
						t.Logf("\t%s\tShould receive a %d status code.", succeed, tt.statusCode)
					} else {
						t.Errorf("\t%s\tShould receive a %d status code: %v", failed, tt.statusCode, resp.StatusCode)
					}
				}
			}
			t.Run(tt.name, tf)
		}
	}
}

// TestParallelize validate the http Get function can download content and handles
// different status conditions properly but runs the tests in parallel.
func TestParallelize(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		statusCode int
	}{
		{"statusok", "http://rss.cnn.com/rss/cnn_topstorie.rss", http.StatusOK},
		{"statusnotfound", "http://rss.cnn.com/rss/cnn_topstorie.rss", http.StatusNotFound},
	}

	t.Log("Given the need to test downloading different content.")
	{
		for i, tt := range tests {
			tf := func(t *testing.T) {
				t.Parallel()

				t.Logf("\tTest: %d\tWhen checking %q for status code %d", i, tt.url, tt.statusCode)
				{
					resp, err := http.Get(tt.url)
					if err != nil {
						t.Fatalf("\t%s\tShould be able to make the Get call: %v", failed, err)
					}
					t.Logf("\t%s\tShould be able to make the Get call.", succeed)

					defer resp.Body.Close()

					if resp.StatusCode == tt.statusCode {
						t.Logf("\t%s\tShould receive a %d status code.", succeed, tt.statusCode)
					} else {
						t.Errorf("\t%s\tShould receive a %d status code: %v", failed, tt.statusCode, resp.StatusCode)
					}
				}
			}
			t.Run(tt.name, tf)
		}
	}
}
