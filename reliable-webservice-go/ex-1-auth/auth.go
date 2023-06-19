// Auth middleware for basic auth
// to test: curl -H "Authorization:Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==" http://localhost:8080

package main

import (
	"fmt"
	"net/http"
)

// check if the that username and password are provided
func authMiddleware(w http.ResponseWriter, r *http.Request) {
	//https://pkg.go.dev/net/http#Request.BasicAuth
	username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, http.StatusText(http.StatusBadRequest)+", Authorization header must be in the form username:password", http.StatusBadRequest)
		return
	}
	// Add logic to check if the username and password are correct or not
	// This is often done by checking the username and password against a database
	// For the sake of simplicity, we will hardcode the username and password
	if username != "Aladdin" || password != "open sesame" {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	authMiddleware(w, r)
	fmt.Fprintf(w, "Hello")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
