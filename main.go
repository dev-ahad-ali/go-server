package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Text about information....")
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", aboutHandler)
}
