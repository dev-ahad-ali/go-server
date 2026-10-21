package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manger := middleware.NewManger()
	manger.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrappedMax := manger.WrapMux(mux)

	initRoutes(mux, manger)

	fmt.Println("Sever listening on :3000")

	err := http.ListenAndServe(":3000", wrappedMax)

	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
