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

func ErrorResponse(response *http.ResponseWriter, statusCode int, errorMessage string) {
	(*response).Header().Set("Content-Type", "application/json")

	(*response).WriteHeader(statusCode)

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
