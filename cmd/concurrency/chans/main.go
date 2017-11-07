package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	ch := make(chan string)

	go startServer(":8080", ch)

	// range ch loops over ch with read operations
	// range continues until ch is closed
	for msg := range ch {
		fmt.Println(msg)
	}

	fmt.Println("chan closed")
}

func startServer(port string, stream chan<- string) {
	msgHandler := http.HandlerFunc(messenger)
	closerHandler := http.HandlerFunc(closer)

	http.Handle("/", streamMiddleware(stream, msgHandler))
	http.Handle("/close", closerMiddleware(stream, closerHandler))

	fmt.Println("listening on port: ", port)
	http.ListenAndServe(port, nil)
}

func messenger(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello chan world"))
}

func closer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("goodbye chan world"))
}

func streamMiddleware(stream chan<- string, next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var msg message
		json.NewDecoder(r.Body).Decode(&msg)

		stream <- msg.String() // sends msg to stream chan

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func closerMiddleware(stream chan<- string, next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var msg message
		json.NewDecoder(r.Body).Decode(&msg)

		stream <- msg.String() // sends msg to stream chan

		close(stream) // signals that the channel is closed
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

type message struct {
	Greeting string `json:"greeting"`
	Name     string `json:"name"`
}

func (m message) String() string {
	return fmt.Sprintf("%s %s", m.Greeting, m.Name)
}
