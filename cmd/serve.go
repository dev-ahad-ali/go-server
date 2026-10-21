package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manger := middleware.NewManger()

	// last one runs first
	manger.Use(
		middleware.Testing,
		middleware.Logger,
		middleware.CorsWithPreflight,
	)

	mux := http.NewServeMux()

	initRoutes(mux, manger)

	fmt.Println("Sever listening on :3000")

	wrappedMax := manger.With(mux)

	err := http.ListenAndServe(":3000", wrappedMax)
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
