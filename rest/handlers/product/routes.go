package product

import (
	"ecommerce/rest/handlers"
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manger *middleware.Manger) {
	mux.Handle("GET /test", manger.With(
		http.HandlerFunc(handlers.Test),
		middleware.Testing,
	))

	mux.Handle("GET /products", manger.With(
		http.HandlerFunc(h.GetProducts),
	))

	mux.Handle("POST /products", manger.With(
		http.HandlerFunc(h.CreateProducts),
		middleware.AuthenticateJwt,
	))

	mux.Handle("GET /products/{id}", manger.With(
		http.HandlerFunc(h.GetProductById),
	))

	mux.Handle("PUT /products/{id}", manger.With(
		http.HandlerFunc(h.UpdateProduct),
	))

	mux.Handle("DELETE /products/{id}", manger.With(
		http.HandlerFunc(h.DeleteProduct),
	))

}
