package app

import (
	"fmt"
	"github.com/VyacheslavKuzharov/gophermart/config"
	"github.com/VyacheslavKuzharov/gophermart/pkg/logger"
	"html"
	"log"
	"net/http"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/howdy", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Howdy!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
