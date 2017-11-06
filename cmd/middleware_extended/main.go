package main

import (
	"log"
	"net/http"
	"time"

	"github.com/justinas/alice"
)

func logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		total := time.Now().Sub(start)
		log.Printf("[%s] %q %v", r.Method, r.URL.String(), total.String())
	}

	return http.HandlerFunc(fn)
}

func recovery(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About page"))
}

func main() {
	common := alice.New(logger, recovery)

	http.Handle("/", common.ThenFunc(index))
	http.Handle("/about", common.ThenFunc(about))
	log.Println("listening on port: 3000")
	http.ListenAndServe(":3000", nil)
}
