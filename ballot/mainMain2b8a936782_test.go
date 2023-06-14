First, modify the main function to make it testable by extracting the HTTP server creation into a separate function:

```go
package main

import (
	"log"
	"net"
	"net/http"
)

const port = "8080"

func main() {
	log.Println("ballot is ready to store votes")
	startServer(port)
}

func startServer(port string) {
	http.HandleFunc("/", serveRoot)
	http.HandleFunc("/tests/run", runTest)
	log.Println(http.ListenAndServe(net.JoinHostPort("", port), nil))
}

func serveRoot(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement your logic here
}

func runTest(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement your logic here
}
```

Next, create a test file named `main_test.go` with the following integration tests:

```go
package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServeRoot(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(serveRoot))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	// TODO: Check the response body for expected content
	if !strings.Contains(string(body), "expected content") {
		t.Errorf("Expected 'expected content' in the response body, got %s", body)
	}
}

func TestRunTest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(runTest))
	defer ts.Close()

	res, err := http.Get(ts.URL + "/tests/run")
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	// TODO: Check the response body for expected content
	if !strings.Contains(string(body), "expected content") {
		t.Errorf("Expected 'expected content' in the response body, got %s", body)
	}
}