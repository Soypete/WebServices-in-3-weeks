//go:build ignore

//
// test with curl localhost:8080/hello\?name=World
// expected output: Hello World! status 200

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
