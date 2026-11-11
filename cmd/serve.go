package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
	"fmt"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection()

	if err != nil {
		fmt.Println(err)
		return
	}

	userRepo := repo.NewUserRepo(dbCon)
	productRepo := repo.NewProductRepo()

	middlewares := middleware.NewMiddleware(cnf)

	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start()
}
