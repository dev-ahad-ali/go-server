package product

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	pId, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Please provide a valid product Id")
		return
	}

	product, err := h.productRepo.Get(pId)

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if product == nil {
		utils.SendError(w, http.StatusNotFound, "Product not found")
		return
	}

	utils.SendData(w, http.StatusOK, product)

}
