package main

import (
	"fmt"
	"net/http"
)

type HomeHandler struct{}

func (h HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to go server!")
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", HomeHandler{})

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world! You are accessing %s and using User Agent %s\n", r.URL.Path, r.Header.Get("User-Agent"))
	})

	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"status": "ok"}`)
	})

	fmt.Println("server is starting on port 8080")
	http.ListenAndServe(":8080", mux)
}
