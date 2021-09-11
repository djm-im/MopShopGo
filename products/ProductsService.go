package products

import (
	"MopShopGo/exceptions"
	json "encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetAll(response http.ResponseWriter, request *http.Request) {
	products := getAllProducts()

	response.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(response).Encode(products)
}

func GetProduct(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	productId, _ := strconv.Atoi(params["productId"])

	product, err := getProduct(productId)

	if err != nil {
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusNotFound)

		exceptions.ErrorResponse(
			&response,
			http.StatusNotFound,
			"Cannot fina a product with ID: "+strconv.Itoa(productId))

		return
	}

	response.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(response).Encode(product)
}
