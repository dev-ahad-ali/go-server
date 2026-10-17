package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	fmt.Fprintln(w, "Hello, World!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Text about information....")
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

var productList []Product

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Bad request", 400)
		return
	}

	sendData(w, productList, 200)
}

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

func main() {

	mux := http.NewServeMux() // router

	mux.Handle("GET /", http.HandlerFunc(helloHandler))      // route
	mux.Handle("GET /about", http.HandlerFunc(aboutHandler)) // route

	mux.Handle("GET /products", corsMiddleware(http.HandlerFunc(getProducts)))
	mux.Handle("POST /create-products", corsMiddleware(http.HandlerFunc(createProducts)))

	fmt.Println("Sever listening on :3000")

	globalRouter := globalRouter(mux)

	err := http.ListenAndServe(":3000", globalRouter) // failed to start the server
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	handleCors := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	}

	handler := http.HandlerFunc(handleCors)

	return handler
}

func globalRouter(mux *http.ServeMux) http.Handler {
	handleAllReq := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Content-Type", "application/json")

			w.WriteHeader(200)
		} else {
			mux.ServeHTTP(w, r)
		}
	}

	return http.HandlerFunc(handleAllReq)
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "lorem diej eifie e iifieh id feuhiojgn oe oe oiejoij oe o",
		Price:       10.99,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSNYH1D8zGWCajbNbCmeD16reW0UBuZwyVu-LeFLsrsRA&s=10",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "lorem diej eifie e iifieh id feuhiojgn oe oe oiejoij oe o oaojdfodof",
		Price:       20.99,
		ImgUrl:      "https://cdn.britannica.com/22/187222-050-07B17FB6/apples-on-a-tree-branch.jpg",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "lorem diej eifie e iifieh id feuhiojgn oe oe oiejoij oe o",
		Price:       15.49,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/8/8a/Banana-Single.jpg",
	}

	prd4 := Product{
		ID:          4,
		Title:       "Strawberry",
		Description: "lorem diej eifie e iifieh id feuhiojgn oe oe oiejoij oe o oaojdfodof",
		Price:       12.99,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/e/e1/Strawberries.jpg",
	}

	prd5 := Product{
		ID:          5,
		Title:       "Mango",
		Description: "lorem diej eifie e iifieh id feuhiojgn oe oe oiejoij oe o",
		Price:       18.50,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/9/90/Haden_mango_aa.jpg",
	}

	prd6 := Product{
		ID:          6,
		Title:       "Pineapple",
		Description: "lorem diej eifie e iifieh id feuhiojgn oe oe oiejoij oe o oaojdfodof",
		Price:       25.00,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/c/cb/Pineapple_and_cross_section.jpg",
	}

	productList = append(productList, prd1, prd2, prd3, prd4, prd5, prd6)
}
