package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {
	fmt.Println("Starting server on port 8080")
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":8080", nil)
}
