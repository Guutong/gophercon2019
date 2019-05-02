package main

import (
	"fmt"
	"net/http"
)

// section: hello
func HelloHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(200)
	name := req.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(res, "Hello, %s!", name)
}

// section: hello

// section: formHandler
func FormHandler(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		res.WriteHeader(500)
		return
	}
	name := req.Form.Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(res, "Posted Hello, %s!", name)
}

// section: formHandler

type Responder struct {
	http.ResponseWriter
	status int
}

func (r *Responder) WriteHeader(i int) {
	r.status = i
	r.ResponseWriter.WriteHeader(i)
}

// section: app
func App() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", HelloHandler) // GET
	mux.HandleFunc("/form", FormHandler)   // POST

	h := func(res http.ResponseWriter, req *http.Request) {
		rs := &Responder{ResponseWriter: res, status: 200}
		mux.ServeHTTP(rs, req)
		if rs.status == 500 {
			rs.Write([]byte("Oops!"))
		}
	}
	return http.HandlerFunc(h)
}

// section: app

func main() {
	http.ListenAndServe(":3000", App())
}
