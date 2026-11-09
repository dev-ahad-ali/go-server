package user

import (
	"ecommerce/config"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req ReqLogin

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		fmt.Println(err)
		utils.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	usr, err := h.userRepo.Find(req.Email, req.Password)

	if err != nil {
		utils.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	cnf := config.GetConfig()

	accessToken, err := utils.CreateJwt(cnf.JwtSecretKey, utils.Payload{
		Sub:       usr.ID,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
	})

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	utils.SendData(w, http.StatusCreated, accessToken)
}
