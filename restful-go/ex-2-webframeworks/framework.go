// This is exercise uses the following web frameworks:
// 1. chi - https://github.com/go-chi/chi

// to run this exercise:
// go test -run TestUserEndpoints
package framework

import (
	"net/http"

	"github.com/go-chi/chi"
)

func getUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	// check if username is empty -> return 400
	if username == "" {
		http.Error(w, http.StatusText(400)+", username parameter cannot be empty", 400)
		return
	}

	// return username and 200
	w.Write([]byte(username))
}

func updateUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	// check if username is empty -> return 400
	if username == "" {
		http.Error(w, http.StatusText(400)+", username parameter cannot be empty", 400)
		return
	}

	// return response and 200
	w.Write([]byte("user registered"))
}

func deleteUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	// check if username is empty -> return 400
	if username == "" {
		http.Error(w, http.StatusText(400)+", username parameter cannot be empty", 400)
		return
	}

	// return response and 200
	w.Write([]byte("user deleted")) // TODO: return deleted username
}
