package testing

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTestServer(t *testing.T) {
	server := testServer(200, "testing")
	defer server.Close()

	if server == nil {
		t.Errorf("Something wrong with Test-Server")
	}
}

// Create a Test HTTP Server that will return a response with HTTP code and body.
func testServer(code int, body string) *httptest.Server {
	serv := testServerForQuery("", code, body)
	return serv
}

// testServerForQuery returns a mock server that only responds to a particular query string.
func testServerForQuery(query string, code int, body string) *httptest.Server {
	server := &httptest.Server{}

	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if query != "" && r.URL.RawQuery != query {
			log.Printf("Query != Expected Query: %s != %s", query, r.URL.RawQuery)
			http.Error(w, "fail", 999)
			return
		}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, body)
	}))

	return server
}
