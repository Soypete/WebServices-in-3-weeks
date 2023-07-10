//go:build ignore

//
// test with curl -I -X POST localhost:8080/ -H "Content-Type: text/plain"
// expected output: Hello
// Follow up: What are some ways to improve the code?
// Follow up: what are some safety concerns with the code?
// Follow up: what are some ways to handle errors?

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		contentType := r.Header.Get("Content-Type")
		if contentType != "text/plain" {
			http.Error(w, "Content-Type header is incorrect", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Hello %s", b)

	})

	http.ListenAndServe(":8080", nil)
}
