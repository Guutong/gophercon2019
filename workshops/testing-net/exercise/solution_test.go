package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// section: ping
func Ping(req *http.Request, c http.Client) (*http.Response, error) {
	return c.Do(req)
}

// section: ping

// section: tripper
type jack struct {
	code int
	res  string
	err  error
}

func (j jack) RoundTrip(req *http.Request) (*http.Response, error) {
	res := httptest.NewRecorder()
	if j.code != 0 {
		res.WriteHeader(j.code)
	}
	res.Body.WriteString(j.res)
	return res.Result(), j.err
}

// section: tripper

// section: status
func Test_Ping_Status(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/", nil)
	if err != nil {
		t.Fatal(err)
	}

	exp := 500

	c := http.Client{
		Transport: jack{code: exp},
	}

	res, err := Ping(req, c)
	if err != nil {
		t.Fatal(err)
	}

	got := res.StatusCode

	if exp != got {
		t.Fatalf("expected %q got %q", exp, got)
	}

}

// section: status
