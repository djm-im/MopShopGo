package carts

import (
    "MopShopGo/exceptions"
    "MopShopGo/sessions"
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
    "strconv"
)

func GetCart(response http.ResponseWriter, request *http.Request) {
    exceptions.FuncNotImplemented(response)
}

func AddProduct(response http.ResponseWriter, request *http.Request) {
    cookie, err1 := request.Cookie("session_token")
    if err1 != nil {
        if err1 == http.ErrNoCookie {
            exceptions.BuildErrorResponse(
                &response,
                http.StatusUnauthorized,
                fmt.Sprintf("User is not authorized."))
            return
        }

        // other errors
        exceptions.BuildErrorResponse(
            &response,
            http.StatusBadRequest,
            fmt.Sprintf("Not proper requesrt."))
        return
    }

    token := cookie.Value

    sessionToken, err2 := sessions.GetSessionToken(token)
    if err2 != nil {
        exceptions.BuildErrorResponse(
            &response,
            http.StatusUnauthorized,
            fmt.Sprintf(err2.Error()))
        return
    }

    if !sessions.IsValidToken(sessionToken) {
        exceptions.BuildErrorResponse(
            &response,
            http.StatusUnauthorized,
            fmt.Sprintf("Session token is not valid"))
        return
    }

    params := mux.Vars(request)
    userId, _ := strconv.Atoi(params["userId"])
    productId, _ := strconv.Atoi(params["productId"])
    quantity, _ := strconv.Atoi(params["quantity"])

    err3 := insertCartItem(CartItems{
        UserId:    int64(userId),
        ProductId: int64(productId),
        Quantity:  quantity,
    })

    if err3 != nil {
        exceptions.BuildErrorResponse(
            &response,
            http.StatusConflict,
            fmt.Sprintf(err3.Error()))
    }

    response.Header().Set("Content-Type", "application/json")
    response.WriteHeader(http.StatusNoContent)
}

func RemoveProduct(response http.ResponseWriter, request *http.Request) {
    exceptions.FuncNotImplemented(response)
}

func DeleteCart(response http.ResponseWriter, request *http.Request) {
    exceptions.FuncNotImplemented(response)
}