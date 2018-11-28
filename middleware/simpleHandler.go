package main

import (
	"fmt"
	"net/http"
)

type helloHandler struct{}

func (h helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, you have hit %s", r.URL.Path)
}

func main() {
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8081", helloHandler{})
}
