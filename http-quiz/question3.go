//go:build ignore

//
// test with curl localhost:8080/hello
// expected output: Hello World! status 400
// follow up question: should it return a different status code or error message?
// What other status codes does this service return?

package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			http.Error(w, "missing name", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Hello %s!", name)
	})
	http.ListenAndServe(":8080", nil)
}
