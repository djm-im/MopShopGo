package admin

import (
    "MopShopGo/exceptions"
    "MopShopGo/products"
    "MopShopGo/users"
    "encoding/json"
    "net/http"
)

func GetAllUsers(response http.ResponseWriter, request *http.Request) {
    username, password, ok := request.BasicAuth()

    if !ok {
        exceptions.UnauthorizedResponse(&response)
        return
    }

    if username != "admin" || password != "admin123" {
        exceptions.UnauthorizedResponse(&response)
        return
    }

    allUsers, err := users.GetAllUsers()

    if err != nil {
        exceptions.BuildErrorResponse(&response, http.StatusNotFound, err.Error())
        return
    }

    response.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(response).Encode(allUsers)

}

func AddProduct(response http.ResponseWriter, request *http.Request) {
    username, password, ok := request.BasicAuth()

    if !ok {
        exceptions.UnauthorizedResponse(&response)
        return
    }

    if username != "admin" || password != "admin123" {
        exceptions.UnauthorizedResponse(&response)
        return
    }

    decoder := json.NewDecoder(request.Body)
    var newProduct products.Product
    err := decoder.Decode(&newProduct)

    acceptable := http.StatusNotAcceptable
    if err != nil {
        exceptions.BuildErrorResponse(&response, acceptable, err.Error())
        return
    }

    product, err := products.AddProduct(newProduct)

    if err != nil {
        exceptions.BuildErrorResponse(&response, acceptable, err.Error())
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
