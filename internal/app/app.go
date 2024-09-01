package app

import (
	"fmt"
	"github.com/VyacheslavKuzharov/gophermart/config"
	"html"
	"log"
	"net/http"
)

func Run(cfg *config.Config) {
	fmt.Println("Howdy!")
	fmt.Println("cfg.App.Version", cfg.App.Version)
	fmt.Println("cfg.PG.DSN", cfg.PG.DatabaseUri)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/howdy", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Howdy!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
