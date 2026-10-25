package main

import (
	"ecommerce/config"
	"fmt"
)

func main() {
	cnf := config.GetConfig()

	fmt.Println(cnf.ServiceName)
	fmt.Println(cnf.Version)
	fmt.Println(cnf.HttpPort)

	// cmd.Serve()
}
