package main

import (
	"net/http"
)

// endpoint: /getUsername/{username}
func getUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Path[len("/getUsername/"):]
	// check if username is empty -> return 400
	if username == "" {
		http.Error(w, http.StatusText(400)+", username parameter cannot be empty", 400)
		return
	}
	// return username and 200
	w.Write([]byte(username))
}

// endpoint: /updateUsername/{username}
func updateUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Path[len("/updateUsername/"):]
	// check if username is empty -> return 400
	if username == "" {
		http.Error(w, http.StatusText(400)+", username parameter cannot be empty", 400)
		return
	}

	// return response and 200
	w.Write([]byte("user registered"))
}

// endpoint: /deleteUsername/{username}
func deleteUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Path[len("/deleteUsername/"):]
	// check if username is empty -> return 400
	if username == "" {
		http.Error(w, http.StatusText(400)+", username parameter cannot be empty", 400)
		return
	}

	// return response and 200
	w.Write([]byte("user deleted")) // TODO: return deleted username
}

func main() {
	http.HandleFunc("/getUsername", getUsername)
	http.HandleFunc("/updateUsername", updateUsername)
	http.HandleFunc("/deleteUsername", deleteUsername)

	http.ListenAndServe(":8090", nil)
}
