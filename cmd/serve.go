package cmd

import (
	"ecommerce/handlers"
	"ecommerce/utils"
	"fmt"
	"net/http"
)

func Serve() {

	mux := http.NewServeMux() // router

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProducts))
	mux.Handle("GET /products/{productId}", http.HandlerFunc(handlers.GetProductById))

	fmt.Println("Sever listening on :3000")

	globalRouter := utils.GlobalRouter(mux)

	err := http.ListenAndServe(":3000", globalRouter) // failed to start the server
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
