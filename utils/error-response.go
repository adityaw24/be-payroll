package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status_code    int    `json:"status_code"`
	Status_message string `json:"status_message"`
}

func BuildErrorResponse(r http.ResponseWriter, status int, message string) {
	result := ErrorResponse{
		Status_code:    status,
		Status_message: message,
	}
	r.Header().Set("Content-Type", "application/json")
	r.WriteHeader(status)
	json.NewEncoder(r).Encode(result)
}
