package handlers

import (
	"ecommerce/lib"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProducts(w http.ResponseWriter, r *http.Request) {
	var newProduct lib.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid JSON", 400)
	}

	newProduct.ID = len(lib.ProductList) + 1

	lib.ProductList = append(lib.ProductList, newProduct)

	utils.SendData(w, newProduct, 201)
}
