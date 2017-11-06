package main

import (
	"log"
	"net/http"
	"time"
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

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing index")
	w.Write([]byte("OK"))
}

func about(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing about")
	w.Write([]byte("About page"))
}

func main() {
	http.Handle("/", logger(http.HandlerFunc(index)))
	http.Handle("/about", logger(http.HandlerFunc(about)))
	http.ListenAndServe(":3000", nil)
}
