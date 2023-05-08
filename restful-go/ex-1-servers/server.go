package main

import (
	"fmt"
	"net/http"
)

// curl http://localhost:8090/hello?name=world
func hello(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if name := req.FormValue("name"); name != "" {
		fmt.Fprintf(w, "hello %s\n", name)
	}

}

func main() {
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":8090", nil)
}
