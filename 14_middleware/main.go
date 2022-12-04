package main

import (
	"fmt"
	"log"
	"net/http"
)

// https://go.dev/blog/defer-panic-and-recover
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/panic/", panicDemo)
	// handleWare is a handler wrapper to defer panics
	mux.HandleFunc("/panic-after/", handleWare(panicDemo))
	mux.HandleFunc("/", hello)

	log.Fatal(http.ListenAndServe(":3000", mux))
}

func handleWare(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer handlePanic()
		h(w, r)
	}
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Something bad happened, logging:", r)
	}
}

func panicDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")
	funcThatPanics()
}

func funcThatPanics() {
	panic("bad")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")
}
