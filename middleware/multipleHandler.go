package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is home page")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is aout page")
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request received for :", r)
		next.ServeHTTP(w, r)
		fmt.Println("Request handled successfully")
	})
}

func main() {

	mux := http.NewServeMux()

	mux.Handle("/index", loggerMiddleware(http.HandlerFunc(indexHandler)))
	mux.Handle("/about", loggerMiddleware(http.HandlerFunc(aboutHandler)))
	mux.HandleFunc("/about", aboutHandler)
	h := loggerMiddleware(mux)

	http.ListenAndServe(":8081", h)
}
