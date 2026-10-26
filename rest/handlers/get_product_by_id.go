package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	pId, err := strconv.Atoi(r.PathValue("productId"))

	if err != nil {
		http.Error(w, "Please provide a valid product Id", 400)
		return
	}

	for _, product := range database.ProductList {
		if product.ID == pId {
			utils.SendData(w, product, 200)
			return
		}
	}

	utils.SendData(w, "Data not found", 404)

}
