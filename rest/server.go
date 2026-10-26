package rest

import (
	"ecommerce/config"
	middleware "ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"strconv"
)

func Start(cnf config.Config) {
	manger := middleware.NewManger()
	manger.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrappedMax := manger.WrapMux(mux)

	initRoutes(mux, manger)

	addr := ":" + strconv.Itoa(cnf.HttpPort)

	fmt.Println("Server listening on port", addr)

	err := http.ListenAndServe(addr, wrappedMax)

	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
