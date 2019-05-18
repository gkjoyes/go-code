// Sample test to show how to write a basic unit test.

package example1

import (
	"net/http"
	"testing"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

// TestDownload validates the http Get function can download content.
func TestDownload(t *testing.T) {
	url := "https://www.ardanlabs.com/blog/index.xml"
	statusCode := 200

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tTest 0:\tWhen checking %q for status code %d", url, statusCode)
		{
			res, err := http.Get(url)
			if err != nil {
				t.Fatalf("\t%s\tShould able to make the Get call: %v", failed, err)
			}
			t.Logf("\t%s\tShould be able to make the Get call.", succeed)

			defer res.Body.Close()

			if res.StatusCode == statusCode {
				t.Logf("\t%s\tShould receive a %d status code.", succeed, statusCode)
			} else {
				t.Errorf("\t%s\tShould receive a %d status code: %d", failed, statusCode, res.StatusCode)
			}
		}
	}
}
