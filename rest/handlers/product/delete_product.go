package product

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	pId, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Please provide a valid product Id")
		return
	}

	err = h.productRepo.Delete(pId)

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	utils.SendData(w, http.StatusOK, "Product deleted successfully")
}
