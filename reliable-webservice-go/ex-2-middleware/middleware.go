// to test: curl -H "Authorization:Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==" http://localhost:8080/hello/Aladdin
package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func middlewareHandler(next http.Handler) http.Handler {
	return authHandler(next)
}
func authHandler(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//https://pkg.go.dev/net/http#Request.BasicAuth
		user, pass, ok := r.BasicAuth()
		if !ok {
			fmt.Println("authHandler: !ok")
			http.Error(w, http.StatusText(http.StatusBadRequest)+", Authorization header must be in the form username:password", http.StatusBadRequest)
			return
		}
		if user != "Aladdin" || pass != "open sesame" {
			fmt.Println("authHandler: user != Aladdin || pass != open sesame")
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World!\n")
}

func main() {
	r := chi.NewRouter()

	// add prebuild middleware for all requests
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Get("/hello", handler)
	r.With(middlewareHandler).Route("/middleware", func(r chi.Router) {
		r.Get("/hello", handler)
	})

	http.ListenAndServe(":8080", r)

}
