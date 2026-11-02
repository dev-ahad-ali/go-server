package product

import (
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manger *middleware.Manger) {
	mux.Handle("GET /products", manger.With(
		http.HandlerFunc(h.GetProducts),
	))

	mux.Handle("POST /products", manger.With(
		http.HandlerFunc(h.CreateProducts),
		h.middlewares.AuthenticateJwt,
	))

	mux.Handle("GET /products/{id}", manger.With(
		http.HandlerFunc(h.GetProductById),
	))

	mux.Handle("PUT /products/{id}", manger.With(
		http.HandlerFunc(h.UpdateProduct),
		h.middlewares.AuthenticateJwt,
	))

	mux.Handle("DELETE /products/{id}", manger.With(
		http.HandlerFunc(h.DeleteProduct),
		h.middlewares.AuthenticateJwt,
	))

}
