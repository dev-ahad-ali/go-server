package cmd

import (
	"ecommerce/handlers"
	"ecommerce/middleware"
	"ecommerce/utils"
	"fmt"
	"net/http"
)

func Serve() {
	manger := middleware.NewManger()

	manger.Use(
		middleware.Logger,
		middleware.Testing,
	)

	mux := http.NewServeMux() // router

	mux.Handle("GET /test", manger.With(
		http.HandlerFunc(handlers.Test),
	))

	mux.Handle("GET /products", manger.With(
		http.HandlerFunc(handlers.GetProducts),
	))

	mux.Handle("POST /products", manger.With(
		http.HandlerFunc(handlers.CreateProducts),
	))

	mux.Handle("GET /products/{productId}", manger.With(
		http.HandlerFunc(handlers.GetProductById),
	))

	fmt.Println("Sever listening on :3000")

	globalRouter := utils.GlobalRouter(mux)

	err := http.ListenAndServe(":3000", globalRouter) // failed to start the server
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
