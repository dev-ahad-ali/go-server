package main

import "net/http"

func getProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Bad request", 400)
		return
	}

	sendData(w, productList, 200)
}
