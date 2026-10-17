package main

import (
	"ecommerce/handlers"
	"ecommerce/lib"
	"ecommerce/utils"
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

func main() {

	mux := http.NewServeMux() // router

	mux.Handle("GET /", http.HandlerFunc(helloHandler))      // route
	mux.Handle("GET /about", http.HandlerFunc(aboutHandler)) // route

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /create-products", http.HandlerFunc(handlers.CreateProducts))

	fmt.Println("Sever listening on :3000")

	globalRouter := utils.GlobalRouter(mux)

	err := http.ListenAndServe(":3000", globalRouter) // failed to start the server
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}

func init() {
	prd1 := lib.Product{
		ID:          1,
		Title:       "Orange",
		Description: "lorem diej eifie e iifieh id feuhiojgn oe oe oiejoij oe o",
		Price:       10.99,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSNYH1D8zGWCajbNbCmeD16reW0UBuZwyVu-LeFLsrsRA&s=10",
	}

	prd2 := lib.Product{
		ID:          2,
		Title:       "Apple",
		Description: "lorem diej eifie e iifieh id feuhiojgn oe oe oiejoij oe o oaojdfodof",
		Price:       20.99,
		ImgUrl:      "https://cdn.britannica.com/22/187222-050-07B17FB6/apples-on-a-tree-branch.jpg",
	}
	prd3 := lib.Product{
		ID:          3,
		Title:       "Banana",
		Description: "lorem diej eifie e iifieh id feuhiojgn oe oe oiejoij oe o",
		Price:       15.49,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/8/8a/Banana-Single.jpg",
	}

	prd4 := lib.Product{
		ID:          4,
		Title:       "Strawberry",
		Description: "lorem diej eifie e iifieh id feuhiojgn oe oe oiejoij oe o oaojdfodof",
		Price:       12.99,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/e/e1/Strawberries.jpg",
	}

	prd5 := lib.Product{
		ID:          5,
		Title:       "Mango",
		Description: "lorem diej eifie e iifieh id feuhiojgn oe oe oiejoij oe o",
		Price:       18.50,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/9/90/Haden_mango_aa.jpg",
	}

	prd6 := lib.Product{
		ID:          6,
		Title:       "Pineapple",
		Description: "lorem diej eifie e iifieh id feuhiojgn oe oe oiejoij oe o oaojdfodof",
		Price:       25.00,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/c/cb/Pineapple_and_cross_section.jpg",
	}

	lib.ProductList = append(lib.ProductList, prd1, prd2, prd3, prd4, prd5, prd6)
}
