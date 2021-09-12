package admin

import (
    "MopShopGo/exceptions"
    "MopShopGo/products"
    "MopShopGo/users"
    "encoding/json"
    "net/http"
)

const (
    basicAuthUser = "admin"
    basicAuthPass = "admin123"
)

func GetAllUsers(response http.ResponseWriter, request *http.Request) {
    if !isAuthorized(request) {
        exceptions.UnauthorizedResponse(&response)
        return
    }

    allUsers, err := users.GetAllUsers()

    if err != nil {
        exceptions.BuildErrorResponse(&response, http.StatusNotFound, err.Error())
        return
    }

    response.Header().Set("Content-Type", "application/json")
    response.WriteHeader(http.StatusOK)

    _ = json.NewEncoder(response).Encode(allUsers)
}

func AddProduct(response http.ResponseWriter, request *http.Request) {
    if !isAuthorized(request) {
        exceptions.UnauthorizedResponse(&response)
        return
    }

    decoder := json.NewDecoder(request.Body)
    var newProduct products.Product
    err := decoder.Decode(&newProduct)

    if err != nil {
        exceptions.BuildErrorResponse(&response, http.StatusNotAcceptable, err.Error())
        return
    }

    product, err := products.AddProduct(newProduct)

    if err != nil {
        exceptions.BuildErrorResponse(&response, http.StatusNotAcceptable, err.Error())
        return
    }

    response.Header().Set("Content-Type", "application/json")
    response.WriteHeader(http.StatusCreated)

    _ = json.NewEncoder(response).Encode(product)
}

func isAuthorized(request *http.Request) bool {
    username, password, ok := request.BasicAuth()

    if !ok {
        return false
    }

    if username != basicAuthUser || password != basicAuthPass {
        return false
    }

    return true
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
