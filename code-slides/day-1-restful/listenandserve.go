//go:build ignore

package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(
		w http.ResponseWriter,
		req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
