// to test: curl -H "Authorization:Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==" http://localhost:8080/hello/Aladdin
package main

import (
	"expvar"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var (
	// expvar counters for the number of status codes returned
	counter200Code expvar.Int
	counter400Code expvar.Int
	counter500Code expvar.Int
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
			counter400Code.Add(1)
			http.Error(w, http.StatusText(http.StatusBadRequest)+", Authorization header must be in the form username:password", http.StatusBadRequest)
			return
		}
		if user != "Aladdin" || pass != "open sesame" {
			fmt.Println("authHandler: user != Aladdin || pass != open sesame")
			counter400Code.Add(1)
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	counter200Code.Add(1)
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
	// add expvars at /debug/vars
	expvar.Publish("counter200Code", &counter200Code)
	expvar.Publish("counter400Code", &counter400Code)
	// add pprof at /debug/
	r.Mount("/debug", middleware.Profiler())

	http.ListenAndServe(":8080", r)

}
