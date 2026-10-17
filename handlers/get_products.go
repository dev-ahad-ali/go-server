package handlers

import (
	"ecommerce/lib"
	"ecommerce/utils"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Bad request", 400)
		return
	}

	utils.SendData(w, lib.ProductList, 200)
}
