package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// section: testHello
func Test_U_HelloHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello", nil)
	res := httptest.NewRecorder()

	HelloHandler(res, req)

	if got, exp := res.Code, http.StatusOK; got != exp {
		t.Errorf("unexpected response code.  got: %d, exp %d\n", got, exp)
	}
	if got, exp := res.Body.String(), "Hello, World!"; got != exp {
		t.Errorf("unexpected body.  got: %s, exp %s\n", got, exp)
	}
}

// section: testHello

func Test_U_HelloHandler_WithQuery(t *testing.T) {
	r := require.New(t)

	req := httptest.NewRequest("GET", "/hello?name=Ringo", nil)
	res := httptest.NewRecorder()

	HelloHandler(res, req)

	r.Equal(http.StatusOK, res.Code)
	r.Equal("Hello, Ringo!", res.Body.String())
}

// section: formHandler
func Test_U_FormHandler(t *testing.T) {
	// The url.Values type allows us to assemble a "form" that we can send as part of the request.
	form := url.Values{}
	form.Add("name", "Ringo")

	// The `Encode` method on `url.Values` will properly encode the values we set into well formed `string` that can be read as the body of the request.
	req := httptest.NewRequest("POST", "/form", strings.NewReader(form.Encode()))

	// We must set the `Content-Type` correctly for `ParseForm` to work.
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res := httptest.NewRecorder()

	FormHandler(res, req)

	if got, exp := res.Code, http.StatusOK; got != exp {
		t.Errorf("unexpected response code.  got: %d, exp %d\n", got, exp)
	}
	if got, exp := res.Body.String(), "Posted Hello, Ringo!"; got != exp {
		t.Errorf("unexpected body.  got: %s, exp %s\n", got, exp)
	}
}

// section: formHandler

// section: exercise
func Test_U_FormHandler_Template_Error(t *testing.T) {

	// pass invalid hex strings
	//req := httptest.NewRequest("POST", "/form", strings.NewReader("%zzzzz"))

	// set the header `Content-Type` to `application/x-www-form-urlencoded`
	// create a new httptest.NewRecorder

	// Call the FormHandler
	// Test to see the the response code is 500

	// test the body is `Oops!`
}

// section: exercise

// section: solution
func Test_U_FormHandler_Error(t *testing.T) {

	// pass invalid hex strings
	req := httptest.NewRequest("POST", "/form", strings.NewReader("%zzzzz"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res := httptest.NewRecorder()

	FormHandler(res, req)
	if got, exp := res.Code, http.StatusInternalServerError; got != exp {
		t.Errorf("unexpected response code.  got: %d, exp %d\n", got, exp)
	}

	if got, exp := res.Body.String(), "Oops!"; got != exp {
		t.Errorf("unexpected body: got %s, exp %s\n", got, exp)
	}

}

// section: exercise
