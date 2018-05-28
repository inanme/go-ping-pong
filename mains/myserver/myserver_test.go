package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// $./bin/go-wrk -d 5 http://localhost:8090/mert.inan
// $go tool pprof --seconds=5 localhost:8090/debug/pprof/profile
func Test_handler(t *testing.T) {
	test_cases := []struct {
		in, out string
	}{
		{"http://localhost:34/mert.inan", "mert inan\n"},
		{"http://localhost:34/x.y", "x y\n"},
	}

	for _, test_case := range test_cases {
		req, err := http.NewRequest(http.MethodGet, test_case.in, nil)

		if err != nil {
			t.Fatalf("Could not create request : %v", err)
		}

		rec := httptest.NewRecorder()
		handler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status 200: got %d", rec.Code)
		}
		if rec.Body.String() != "hello there "+test_case.out {
			t.Errorf("Unexpected body reponse %q", rec.Body.String())
		}
	}

}
