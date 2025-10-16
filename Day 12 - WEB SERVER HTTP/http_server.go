package main

import (
	"fmt"
	"net/http"
)

func main() {
    // Register a handler for the "/" route
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, you've reached the home page!")
    })

    fmt.Println("Starting server on port 8080...")
    // Start listening for requests
    // This is a blocking call
    http.ListenAndServe(":8080", nil)
}
