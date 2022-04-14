package utils

import (
	"encoding/json"
	"net/http"
)

//JSON success response model
type SuccessResponse struct {
	Status  int         `json:"status_code"`
	Message string      `json:"message"`
	Results interface{} `json:"results"`
}
type InsertSuccessResponse struct {
	Status     int         `json:"status_code"`
	Message    string      `json:"message"`
	Created_id interface{} `json:"created_id"`
}
type UpdateSuccessResponse struct {
	Status        int         `json:"status_code"`
	Message       string      `json:"message"`
	Updated_value interface{} `json:"updated_value"`
}

func BuildResponse(r http.ResponseWriter, status int, message string, data interface{}) {
	result := SuccessResponse{
		Status:  status,
		Message: message,
		Results: data,
	}
	r.Header().Set("Content-Type", "application/json")
	r.WriteHeader(status)
	json.NewEncoder(r).Encode(result)
}

func BuildInsertResponse(r http.ResponseWriter, status int, message string, data interface{}) {
	result := InsertSuccessResponse{
		Status:     status,
		Message:    message,
		Created_id: data,
	}
	r.Header().Set("Content-Type", "application/json")
	r.WriteHeader(status)
	json.NewEncoder(r).Encode(result)
}

func BuildUpdateResponse(r http.ResponseWriter, status int, message string, data interface{}) {
	result := UpdateSuccessResponse{
		Status:        status,
		Message:       message,
		Updated_value: data,
	}
	r.Header().Set("Content-Type", "application/json")
	r.WriteHeader(status)
	json.NewEncoder(r).Encode(result)
}
