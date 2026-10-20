package cmd

import (
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

	mux := http.NewServeMux()

	initRoutes(mux, manger)

	fmt.Println("Sever listening on :3000")

	globalRouter := utils.GlobalRouter(mux)

	err := http.ListenAndServe(":3000", globalRouter) // failed to start the server
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
