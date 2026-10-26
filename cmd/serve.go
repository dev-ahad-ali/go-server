package cmd

import (
	"ecommerce/config"
	"ecommerce/middleware"
	"fmt"
	"net/http"
	"strconv"
)

func Serve() {
	cnf := config.GetConfig()

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
