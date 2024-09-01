package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Howdy!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/howdy", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Howdy!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
