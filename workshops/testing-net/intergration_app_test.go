package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// section: testHello
func Test_I_HelloHandler(t *testing.T) {
	// section: newServer
	ts := httptest.NewServer(App())
	defer ts.Close()
	// section: newServer

	// section: get
	res, err := http.Get(ts.URL + "/hello")
	if err != nil {
		t.Fatal(err)
	}

	if got, exp := res.StatusCode, http.StatusOK; got != exp {
		t.Errorf("unexpected status code: got %d, exp %d\n", got, exp)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if got, exp := string(b), "Hello, World!"; got != exp {
		t.Errorf("unexpected body: got %s, exp %s\n", got, exp)
	}
	// section: get
}

// section: testHello

func Test_I_HelloHandler_WithQuery(t *testing.T) {
	r := require.New(t)

	ts := httptest.NewServer(App())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/hello?name=Ringo")
	r.NoError(err)

	r.Equal(200, res.StatusCode)
	b, err := ioutil.ReadAll(res.Body)
	r.NoError(err)
	r.Equal("Hello, Ringo!", string(b))
}

// section: formHandler
func Test_I_FormHandler(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	form := url.Values{}
	form.Add("name", "Ringo")

	res, err := http.PostForm(ts.URL+"/form", form)
	if err != nil {
		t.Fatal(err)
	}

	if got, exp := res.StatusCode, http.StatusOK; got != exp {
		t.Errorf("unexpected status code: got %d, exp %d\n", got, exp)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if got, exp := string(b), "Posted Hello, Ringo!"; got != exp {
		t.Errorf("unexpected body: got %s, exp %s\n", got, exp)
	}
}

// section: formHandler

// section: exercise
func Test_I_FormHandler_Error_Template(t *testing.T) {
	// Create a new httptest.NewServer with our App
	// Defer the test server Close

	// Post `%zzzz` to the `/form` endpoint

	// Test the status code is http.StatusInternalServerError

	// test the body is `Oops!`

}

// section: exercise

// section: solution
func Test_I_FormHandler_Error(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	res, err := http.Post(ts.URL+"/form",
		"application/x-www-form-urlencoded",
		strings.NewReader("%zzzzz"))
	if err != nil {
		t.Fatal(err)
	}

	if got, exp := res.StatusCode, http.StatusInternalServerError; got != exp {
		t.Errorf("unexpected status code: got %d, exp %d\n", got, exp)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if got, exp := string(b), "Oops!"; got != exp {
		t.Errorf("unexpected body: got %s, exp %s\n", got, exp)
	}

}

// section: solution
