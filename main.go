package main

import (
	"fmt"
	"net/http"
)

func index(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Hello World!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
