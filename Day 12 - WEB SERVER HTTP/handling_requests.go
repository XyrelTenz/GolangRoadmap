package main

import (
	"fmt"
	"net/http"
)

// A handler function must have this signature
func helloHandler(w http.ResponseWriter, r *http.Request) {
    // r contains information about the request
    fmt.Printf("Handling request for %s from %s\n", r.URL.Path, r.RemoteAddr)

    // w is where you write your response to
    fmt.Fprintln(w, "Hello from the handler function!")
}

func main() {
    http.HandleFunc("/hello", helloHandler)

    fmt.Println("Starting server on port 8080...")
    http.ListenAndServe(":8080", nil)
}
