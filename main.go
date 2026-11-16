package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const version = "v0.1.0"

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "unset"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "ok")
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "svc-hello %s (env=%s)\n", version, env)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("svc-hello %s listening on :%s", version, port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
