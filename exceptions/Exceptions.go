package exceptions

import (
    "encoding/json"
    "log"
    "net/http"
)

type ErrorMessage struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

func UnauthorizedResponse(response *http.ResponseWriter) {
    (*response).Header().Set("WWW-Authenticate", `Basic realm="MyRealm", charset="UTF-8"`)
    (*response).Header().Set("Content-Type", "application/json")
    (*response).WriteHeader(http.StatusUnauthorized)

    BuildErrorResponse(response, http.StatusUnauthorized, "Unauthorized")
}

func BuildErrorResponse(response *http.ResponseWriter, status int, message string) {
    (*response).Header().Set("Content-Type", "application/json")
    (*response).WriteHeader(status)

    errorResponse(response, status, message)
}

func errorResponse(response *http.ResponseWriter, statusCode int, errorMessage string) {
    _ = json.NewEncoder(*response).
        Encode(ErrorMessage{
            Code:    statusCode,
            Message: errorMessage,
        })
}

func FuncNotImplemented(response http.ResponseWriter) {
    response.Header().Set("Content-Type", "application/json")

    response.WriteHeader(http.StatusNotImplemented)

    err := json.NewEncoder(response).
        Encode(ErrorMessage{
            Code:    http.StatusNotImplemented,
            Message: "Method is not implemented, yet.",
        })

    if err != nil {
        log.Printf("Error %s ", err)
        return
    }
}
