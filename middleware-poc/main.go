package main

import (
	"fmt"
	"log"
	"net/http"
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

func main() {

	mux := http.NewServeMux()

	mux.Handle("/", loggingMiddleware(headerMiddleware(http.HandlerFunc(homeHandler))))

	mux.Handle("/about", loggingMiddleware(headerMiddleware(http.HandlerFunc(aboutHandler))))

	log.Println("Starting server on port 8080 ...")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		// The logic should be executed incase the listenandserve returns error
		log.Fatal("Server failed!", err)
	}
}
