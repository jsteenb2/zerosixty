package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Write([]byte("Hello world"))
	end := time.Now()
	log.Printf("[%s] %q %v", r.Method, r.URL.String(), end.Sub(start))
}

func about(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Write([]byte("About page"))
	end := time.Now()
	log.Printf("[%s] %q %v", r.Method, r.URL.String(), end.Sub(start))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	fmt.Println("listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
