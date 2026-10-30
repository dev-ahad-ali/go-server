package product

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Bad request", 400)
		return
	}

	utils.SendData(w, database.List(), 200)
}
