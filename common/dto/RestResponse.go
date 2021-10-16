package dto

import "net/http"

type RestResponse struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewNotFoundError(message string) RestResponse {
	return RestResponse{
		Code:    http.StatusNotFound,
		Message: message,
		Data:    nil,
	}
}

func NewBadRequest(message string) RestResponse {
	return RestResponse{
		Code:    http.StatusBadRequest,
		Message: message,
		Data:    nil,
	}
}
