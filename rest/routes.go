package rest

import (
	"ecommerce/rest/handlers"
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manger *middleware.Manger) {
	mux.Handle("GET /test",
		manger.With(
			http.HandlerFunc(handlers.Test),
			middleware.Testing,
		),
	)

	mux.Handle("GET /products", manger.With(
		http.HandlerFunc(handlers.GetProducts),
	))

	mux.Handle("POST /products", manger.With(
		http.HandlerFunc(handlers.CreateProducts),
	))

	mux.Handle("GET /products/{productId}", manger.With(
		http.HandlerFunc(handlers.GetProductById),
	))
}
