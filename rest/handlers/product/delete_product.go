package product

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	pId, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "Please provide a valid product Id", http.StatusBadRequest)
		return
	}

	err = h.productRepo.Delete(pId)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	utils.SendData(w, "Product deleted successfully", 201)
}
