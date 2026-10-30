package user

import (
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manger) {
	mux.Handle("POST /users", manager.With(
		http.HandlerFunc(h.CreateUser),
	))

	mux.Handle("Post /users/login", manager.With(
		http.HandlerFunc(h.Login),
	))

}
