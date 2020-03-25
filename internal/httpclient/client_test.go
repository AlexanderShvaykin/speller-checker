package httpclient

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	response := `OK!`
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Send response to be tested
		rw.Write([]byte(response))
	}))

	result := Get(server.URL)
	if result != response {
		t.Errorf("unexpected result: got %s, want %s", result, response)
	}
}
