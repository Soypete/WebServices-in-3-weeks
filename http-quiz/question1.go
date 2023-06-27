//go:build ignore

//
// test with curl localhost:8080/
// expected output: Hello World! status 200

package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	http.ListenAndServe(":8080", nil)
}
