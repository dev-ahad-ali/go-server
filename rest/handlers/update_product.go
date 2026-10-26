package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	pId, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "Please provide a valid product Id", 400)
		return
	}

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		utils.SendError(w, 400, "Give valid json data")
		return
	}

	newProduct.ID = pId

	database.Update(newProduct)

	utils.SendData(w, "Product successfully created", 201)
}
