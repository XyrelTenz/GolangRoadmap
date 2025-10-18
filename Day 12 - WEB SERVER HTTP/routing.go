package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is the root page.")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is the about page.")
	})

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		// Example of getting a query parameter: /user?name=Xyrel
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "Guest"
		}
		fmt.Fprintf(w, "Hello, %s!", name)
	})

	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}
