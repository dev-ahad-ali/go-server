package product

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	pId, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "Please provide a valid product Id", 400)
		return
	}

	database.Delete(pId)

	utils.SendData(w, "Product deleted successfully", 201)
}
