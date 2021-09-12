package admin

import (
    "MopShopGo/exceptions"
    "MopShopGo/products"
    "encoding/json"
    "net/http"
)

func GetAllUsers(response http.ResponseWriter, request *http.Request) {
    exceptions.FuncNotImplemented(response)
}

func AddProduct(response http.ResponseWriter, request *http.Request) {
    username, password, ok := request.BasicAuth()

    if !ok {
        response.Header().Set("Content-Type", "application/json")
        response.WriteHeader(http.StatusForbidden)

        exceptions.ErrorResponse(
            &response,
            http.StatusForbidden,
            "No proper authorization")

        return
    }

    if username != "admin" || password != "admin123" {
        response.Header().Set("Content-Type", "application/json")
        response.WriteHeader(http.StatusForbidden)

        exceptions.ErrorResponse(
            &response,
            http.StatusForbidden,
            "Not proper combination of username/password.")

        return
    }

    decoder := json.NewDecoder(request.Body)
    var newProduct products.Product
    err := decoder.Decode(&newProduct)

    if err != nil {
        response.Header().Set("Content-Type", "application/json")
        response.WriteHeader(http.StatusNotAcceptable)

        exceptions.ErrorResponse(
            &response,
            http.StatusNotAcceptable,
            err.Error())

        return
    }

    product, err := products.AddProduct(newProduct)

    if err != nil {
        response.Header().Set("Content-Type", "application/json")
        response.WriteHeader(http.StatusNotAcceptable)

        exceptions.ErrorResponse(
            &response,
            http.StatusNotAcceptable,
            err.Error())

        return
    }

    response.Header().Set("Content-Type", "application/json")
    response.WriteHeader(http.StatusCreated)

    _ = json.NewEncoder(response).Encode(product)
}

func UpdateProduct(response http.ResponseWriter, request *http.Request) {
    exceptions.FuncNotImplemented(response)
}

func UploadImage(response http.ResponseWriter, request *http.Request) {
    exceptions.FuncNotImplemented(response)
}

func DeleteProduct(response http.ResponseWriter, request *http.Request) {
    exceptions.FuncNotImplemented(response)
}

func ProductStatistics(response http.ResponseWriter, request *http.Request) {
    exceptions.FuncNotImplemented(response)
}
