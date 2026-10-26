package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	pId, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "Please provide a valid product Id", 400)
		return
	}

	product := database.Get(pId)

	if product == nil {
		utils.SendError(w, 404, "Product not found")
		return
	}

	utils.SendData(w, product, 200)

}
