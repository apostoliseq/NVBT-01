package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func read(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func save(w http.ResponseWriter, req *http.Request) {

	// INSERT INTO user_input (text) VALUES ($1)

	fmt.Fprintf(w, "Saved!")
}

func main() {

	// Authenticate with postgres

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/read", read)
	http.HandleFunc("/save", save)

	http.ListenAndServe(":8090", nil)
}
