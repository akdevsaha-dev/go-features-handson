package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//? Implement logic
		w.Header().Set("X-Custom-Header", "Pokemon")

		//end of middleware logic
		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
		next.ServeHTTP(w, r)
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the home page!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the about page!")

}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Anon"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")
	if len(pathSegments) >= 2 && pathSegments[1] == "user" {
		userId := pathSegments[2]
		fmt.Fprintf(w, "Hi user: %s!", userId)
	} else {
		http.NotFound(w, r)
	}
}
func main() {

	mux := http.NewServeMux()

	mux.Handle("/", loggingMiddleware(headerMiddleware(http.HandlerFunc(homeHandler))))

	mux.Handle("/about", loggingMiddleware(headerMiddleware(http.HandlerFunc(aboutHandler))))

	mux.Handle("/name", http.HandlerFunc(queryHandler))

	mux.Handle("/user/", http.HandlerFunc(pathHandler))

	log.Println("Starting server on port 8080 ...")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		// The logic should be executed incase the listenandserve returns error
		log.Fatal("Server failed!", err)
	}
}
