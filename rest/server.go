package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"strconv"
)

type Server struct {
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(productHandler *product.Handler, userHandler *user.Handler) *Server {
	return &Server{
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (server *Server) Start(cnf config.Config) {
	manager := middleware.NewManger()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrappedMax := manager.WrapMux(mux)

	addr := ":" + strconv.Itoa(cnf.HttpPort)

	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	fmt.Println("Server listening on port", addr)

	err := http.ListenAndServe(addr, wrappedMax)

	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
