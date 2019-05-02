package main

import (
	"errors"
	"io/ioutil"
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
	res string
	err error
}

func (j jack) RoundTrip(req *http.Request) (*http.Response, error) {
	res := httptest.NewRecorder()
	res.Body.WriteString(j.res)
	return res.Result(), j.err
}

// section: tripper

// section: success
func Test_Ping_Success(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/", nil)
	if err != nil {
		t.Fatal(err)
	}

	exp := "hello"

	c := http.Client{
		Transport: jack{res: exp},
	}

	res, err := Ping(req, c)
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	got := string(b)
	if got != exp {
		t.Fatalf("expected %q got %q", exp, got)
	}
}

// section: success

// section: error
func Test_Ping_Error(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/", nil)
	if err != nil {
		t.Fatal(err)
	}

	c := http.Client{
		Transport: jack{err: errors.New("oops!")},
	}

	_, err = Ping(req, c)
	if err == nil {
		t.Fatal("expected error got none")
	}
}

// section: error
