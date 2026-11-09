package product

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	pId, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "Please provide a valid product Id", http.StatusBadRequest)
		return
	}

	product, err := h.productRepo.Get(pId)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if product == nil {
		utils.SendError(w, 404, "Product not found")
		return
	}

	utils.SendData(w, product, 200)

}
