package httpd_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/guutong/gophercon2019/workshops/test-mocking/exercise/httpd"
)

// section: setsuccess
func TestSet_Success(t *testing.T) {
	// Create the handler
	handler := httpd.NewHandler()
	// create the mock store
	store := &MockStore{}

	// using the mock, create a successful call to `set` method
	store.getFn = func(string) (interface{}, error) {
		return nil, notFound{}
	}

	// assign the mock store the the handler's store
	handler.Store = store

	// uncomment the rest of this test to make it run
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/key", nil)

	r.Form = url.Values{"key": []string{"foo"}, "value": []string{"bar"}}

	handler.ServeHTTP(w, r)
	if exp, got := w.Code, http.StatusAccepted; exp != got {
		t.Errorf("unexpected error code. exp: %d, got %d", exp, got)
	}
}

// section: setsuccess

// section: seterror
func TestSet_Error(t *testing.T) {
	handler := httpd.NewHandler()
	store := &MockStore{}
	store.setFn = func(key string, value interface{}) error {
		return errors.New("boom")
	}
	handler.Store = store

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/key", nil)

	r.Form = url.Values{"key": []string{"foo"}, "value": []string{"bar"}}

	handler.ServeHTTP(w, r)
	if got, exp := w.Code, http.StatusInternalServerError; got != exp {
		t.Log(w.Body)
		t.Errorf("unexpected error code. got: %d, exp %d", got, exp)
	}
}

// section: seterror

// section: nokeys
func TestGet_NoKey(t *testing.T) {
	handler := httpd.NewHandler()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/key", nil)

	handler.ServeHTTP(w, r)
	if exp, got := w.Code, http.StatusBadRequest; exp != got {
		t.Log(w.Body)
		t.Errorf("unexpected error code. exp: %d, got %d", exp, got)
	}
}

// section: nokeys

// section: notfound
func TestGet_NotFound(t *testing.T) {
	handler := httpd.NewHandler()
	store := &MockStore{}
	// section: nfinject
	store.getFn = func(string) (interface{}, error) {
		return nil, notFound{}
	}
	// section: nfinject
	handler.Store = store

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/key?key=foo", nil)

	handler.ServeHTTP(w, r)
	if got, exp := w.Code, http.StatusNotFound; got != exp {
		t.Log(w.Body)
		t.Errorf("unexpected error code. got: %d, exp %d", got, exp)
	}
}

// section: notfound

// section: serverError
func TestGet_ServerError(t *testing.T) {
	handler := httpd.NewHandler()
	store := &MockStore{}
	// section: seinject
	store.getFn = func(string) (interface{}, error) {
		return nil, errors.New("boom")
	}
	// section: seinject
	handler.Store = store

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/key?key=foo", nil)

	handler.ServeHTTP(w, r)
	if got, exp := w.Code, http.StatusInternalServerError; got != exp {
		t.Log(w.Body)
		t.Errorf("unexpected error code. got: %d, exp %d", got, exp)
	}
}

// section: serverError

// section: success
func TestGet_Success(t *testing.T) {
	handler := httpd.NewHandler()
	store := &MockStore{}
	// section: successinject
	store.getFn = func(string) (interface{}, error) {
		return "bar", nil
	}
	// section: successinject
	handler.Store = store

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/key?key=foo", nil)

	handler.ServeHTTP(w, r)
	if got, exp := w.Code, http.StatusOK; got != exp {
		t.Fatalf("unexpected status code.  got %d, expected %d", got, exp)
	}
	data := map[string]interface{}{}
	if err := json.Unmarshal(w.Body.Bytes(), &data); err != nil {
		t.Fatal(err)
	}
	if got, exp := data["foo"], "bar"; got != exp {
		t.Fatalf("unexpected value.  got: %v, exp %v", got, exp)
	}
}

// section: success

// section: mockstore
type MockStore struct {
	setFn func(key string, value interface{}) error
	getFn func(key string) (interface{}, error)
}

func (ms *MockStore) Set(key string, value interface{}) error {
	if ms.setFn != nil {
		return ms.setFn(key, value)
	}
	return nil
}

func (ms *MockStore) Get(key string) (interface{}, error) {
	if ms.getFn != nil {
		return ms.getFn(key)
	}
	return nil, nil
}

// section: mockstore

// not found mock
type notFound struct{}

func (nf notFound) NotFound() {}

func (nf notFound) Error() string {
	return ""
}
