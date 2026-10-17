package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func createProducts(w http.ResponseWriter, r *http.Request) {
	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid JSON", 400)
	}

	newProduct.ID = len(productList) + 1

	productList = append(productList, newProduct)

	sendData(w, newProduct, 201)
}
