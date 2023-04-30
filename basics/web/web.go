package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hello", hello)
	fmt.Println("Starting Server...")
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, rp *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func hello(w http.ResponseWriter, rp *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World!</h1>")
}
