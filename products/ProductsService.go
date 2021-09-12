package products

import (
    "MopShopGo/exceptions"
    "encoding/json"
    "github.com/gorilla/mux"
    "net/http"
    "strconv"
)

func GetAll(response http.ResponseWriter, request *http.Request) {
    products, err := getAllProducts()

    if err != nil {
        response.Header().Set("Content-Type", "application/json")
        response.WriteHeader(http.StatusInternalServerError)

        exceptions.ErrorResponse(
            &response,
            http.StatusInternalServerError,
            err.Error())

        return
    }

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
            err.Error())

        return
    }

    response.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(response).Encode(product)
}

func AddProduct(product Product) (Product, error) {
    // todo: validate that `product.price` is greater than zero
    return addProduct(product)
}
