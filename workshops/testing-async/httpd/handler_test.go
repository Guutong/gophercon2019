package httpd_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/gopherguides/learn/_training/testing/async/src/httpd"
	"github.com/gopherguides/learn/_training/testing/async/src/keys"
)

// section: nosleep
func TestSetNoSleep(t *testing.T) {
	handler := httpd.New()
	store := keys.NewStore()
	handler.Store = store

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/key", nil)
	r.Form = url.Values{"key": []string{"foo"}, "value": []string{"bar"}}

	handler.ServeHTTP(w, r)
	if exp, got := w.Code, http.StatusAccepted; exp != got {
		t.Errorf("unexpected error code. exp: %d, got %d", exp, got)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/key?key=foo", nil)

	handler.ServeHTTP(w, r)
	if exp, got := w.Code, http.StatusOK; exp != got {
		t.Log(w.Body)
		t.Errorf("unexpected error code. exp: %d, got %d", exp, got)
	}
}

// section: nosleep
/*
// section: nosleep-output
$ go test . -run TestSetNoSleep
2019/04/25 09:39:27 set...
2019/04/25 09:39:27 get...
--- FAIL: TestSetNoSleep (0.00s)
    handler_test.go:36: foo not found

    handler_test.go:37: unexpected error code. exp: 400, got 200
FAIL
FAIL    github.com/gopherguides/learn/_training/testing/async/src/httpd 0.014s
// section: nosleep-output
*/

// section: sleep
func TestSetSleep(t *testing.T) {
	handler := httpd.New()
	store := keys.NewStore()
	handler.Store = store

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/key", nil)
	r.Form = url.Values{"key": []string{"foo"}, "value": []string{"bar"}}

	handler.ServeHTTP(w, r)
	if exp, got := w.Code, http.StatusAccepted; exp != got {
		t.Errorf("unexpected error code. exp: %d, got %d", exp, got)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/key?key=foo", nil)

	// Magic test time bomb! Just increase this if the test times out... AHHHHH!!!
	time.Sleep(4 * time.Second)

	handler.ServeHTTP(w, r)
	if exp, got := w.Code, http.StatusOK; exp != got {
		t.Log(w.Body)
		t.Errorf("unexpected error code. exp: %d, got %d", exp, got)
	}
}

// section: sleep

/*
// section: sleep-output
$ go test . -v -run TestSetSleep
=== RUN   TestSetSleep
2019/04/25 09:42:47 set...
2019/04/25 09:42:48 inserted:  foo  with value of  bar
2019/04/25 09:42:51 get...
2019/04/25 09:42:51 took 239.544µs
--- PASS: TestSetSleep (4.00s)
PASS
ok      github.com/gopherguides/learn/_training/testing/async/src/httpd 4.020s

// section: sleep-output
*/

// section: channels
func TestSetChannels(t *testing.T) {
	handler := httpd.New()
	store := keys.NewStore()
	handler.Store = store

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/key", nil)
	r.Form = url.Values{"key": []string{"foo"}, "value": []string{"bar"}}

	handler.ServeHTTP(w, r)
	if exp, got := w.Code, http.StatusAccepted; exp != got {
		t.Errorf("unexpected error code. exp: %d, got %d", exp, got)
	}

	// Create a function that we can fire off in a go routine
	test := func() error {
		w = httptest.NewRecorder()
		r = httptest.NewRequest("GET", "/key?key=foo", nil)
		handler.ServeHTTP(w, r)

		if got, exp := w.Code, http.StatusOK; got != exp {
			return fmt.Errorf("unexpected status code.  got %d, expected %d", got, exp)
		}
		data := map[string]interface{}{}
		if err := json.Unmarshal(w.Body.Bytes(), &data); err != nil {
			t.Fatal(err)
		}
		if got, exp := data["foo"], "bar"; got != exp {
			return fmt.Errorf("unexpected value.  got: %v, exp %v", got, exp)
		}
		// test successful
		return nil
	}

	// Use a ticker to periodically check if the test has completed successfully
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	// Set a timeout for the entire test.
	timeout := time.NewTimer(5 * time.Second)
	defer timeout.Stop()

	// Create a variable to keep track of the last error
	var testErr error

	for {
		testErr = test()
		// If we don't have an error, we succeeded
		if testErr == nil {
			// test successful
			return
		}
		// There was an error in the test, try it again

		select {
		case <-timeout.C:
			t.Fatalf("test timed out waiting for success.  last error: %s", testErr)
			return
		case <-ticker.C:
			continue
		}
	}
}

// section: channels

/*
// section: channels-output
$ go test . -v -run TestSetChannels
=== RUN   TestSetChannels
2019/04/25 09:49:43 set...
2019/04/25 09:49:43 inserted:  foo  with value of  bar
2019/04/25 09:49:43 get...
2019/04/25 09:49:43 took 293.541µs
--- PASS: TestSetChannels (0.11s)
PASS
ok      github.com/gopherguides/learn/_training/testing/async/src/httpd 0.121s
// section: channels-output
*/

// section: timeout-function
func TestSetTimeout(t *testing.T) {
	handler := httpd.New()
	store := keys.NewStore()
	handler.Store = store

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/key", nil)
	r.Form = url.Values{"key": []string{"foo"}, "value": []string{"bar"}}

	handler.ServeHTTP(w, r)
	if exp, got := w.Code, http.StatusAccepted; exp != got {
		t.Errorf("unexpected error code. exp: %d, got %d", exp, got)
	}

	// Give this 5 seconds to retry before it fails
	TimeoutRetry(t, 5*time.Second, func() error {
		w = httptest.NewRecorder()
		r = httptest.NewRequest("GET", "/key?key=foo", nil)
		handler.ServeHTTP(w, r)

		if got, exp := w.Code, http.StatusOK; got != exp {
			return fmt.Errorf("unexpected status code.  got %d, expected %d", got, exp)
		}
		data := map[string]interface{}{}
		if err := json.Unmarshal(w.Body.Bytes(), &data); err != nil {
			t.Fatal(err)
		}
		if got, exp := data["foo"], "bar"; got != exp {
			return fmt.Errorf("unexpected value.  got: %v, exp %v", got, exp)
		}
		// test successful
		return nil
	})
}

// section: timeout-function

/*
// section: timeout-function-output
$ go test . -v -run TestSetTimeout -count=1
=== RUN   TestSetTimeout
2019/04/25 09:59:05 set...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:05 get...
2019/04/25 09:59:06 get...
2019/04/25 09:59:06 get...
2019/04/25 09:59:06 get...
2019/04/25 09:59:06 get...
2019/04/25 09:59:06 inserted:  foo  with value of  bar
2019/04/25 09:59:06 get...
2019/04/25 09:59:06 took 158.599µs
--- PASS: TestSetTimeout (1.01s)
PASS
ok      github.com/gopherguides/learn/_training/testing/async/src/httpd 1.021s
// section: timeout-function-output
*/

// section: timeout-retry
// TimeoutRetry returns failes if fn doesn't return a nil error within the timeout duration.
// It will continue to retry even if it receives an error as long as it has not timed out
func TimeoutRetry(t *testing.T, timeout time.Duration, fn func() error) {
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	ticker := time.NewTicker(25 * time.Millisecond)
	defer ticker.Stop()

	var err error
	for {
		// Run the function and save the last error, if any. Exit if no error.
		if err = fn(); err == nil {
			return
		}

		select {
		case <-timer.C:
			// Time is up, we've exceeded the deadline.
			// Include the last known error as well for debugging purposes
			t.Fatalf("%s (%s timeout)", err, timeout)
		case <-ticker.C:
			// interval is up, time to continue and retry our test
			continue
		}
	}
}
