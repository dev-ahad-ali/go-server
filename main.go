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

	mux := http.NewServeMux() // router

	mux.HandleFunc("/hello", helloHandler) // route
	mux.HandleFunc("/about", aboutHandler) // route

	fmt.Println("Sever listening on :3000")

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
