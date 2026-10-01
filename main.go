package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	fmt.Fprintln(w, "Hello, World!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Text about information....")
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Bad request", 400)
		return
	}
}

func main() {

	mux := http.NewServeMux() // router

	mux.HandleFunc("/", helloHandler)      // route
	mux.HandleFunc("/about", aboutHandler) // route
	mux.HandleFunc("/products", getProducts)

	fmt.Println("Sever listening on :3000")

	err := http.ListenAndServe(":3000", mux) // failed to start the server
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
