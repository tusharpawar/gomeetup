package main

import (
	"fmt"
	"net/http"
)

type helloHandler struct{}

func (h helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello path: %v", r.URL.Path)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "you arein index page")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "you arein about page")
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//code preprocessing
		fmt.Println("request :", r)
		fmt.Fprintln(w, "request for", r.URL.Path)
		//next.ServeHTTP(w, r)
		fmt.Println("request processed")

	})
}

func authCheckMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//code preprocessing
		fmt.Println("auth check request :", r)
		fmt.Fprintln(w, "auth check request for", r.URL.Path)
		next.ServeHTTP(w, r)
		fmt.Println("authcheck processed")

	})
}
func main() {
	fmt.Println("Server started at :8080")

	mux := http.NewServeMux()

	// mux.Handle("/index", loggerMiddleware(http.HandlerFunc(indexHandler)))
	// mux.Handle("/about", loggerMiddleware(http.HandlerFunc(aboutHandler)))

	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/index", indexHandler)
	//mux.Handle("/", helloHandler{})
	//h := http.HandlerFunc(aboutHandler)
	h := loggerMiddleware(mux)
	http.ListenAndServe(":8080", h)
}
