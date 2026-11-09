package cmd

import (
	"ecommerce/config"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	userRepo := repo.NewUserRepo()
	productRepo := repo.NewProductRepo()

	middlewares := middleware.NewMiddleware(cnf)

	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start()
}
