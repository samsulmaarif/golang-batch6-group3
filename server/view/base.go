package view

import "net/http"

type Response struct {
	Status      int         `json:"status"`
	Message     interface{} `json:"message"`
	Payload     interface{} `json:"payload"`
	GeneralInfo interface{} `json:"general_info"`
	Error       interface{} `json:"error"`
}

func SuccessCreated(message interface{}) *Response {
	return &Response{
		Status:      http.StatusCreated,
		Message:     message,
		GeneralInfo: "NooBee-Shop",
	}
}

func SuccessLogin(message interface{}, payload interface{}) *Response {
	return &Response{
		Status:  http.StatusCreated,
		Message: message,
		Payload: payload,
	}
}

func SuccessDeleted(message interface{}) *Response {
	return &Response{
		Status:  http.StatusCreated,
		Message: message,
	}
}

func SuccessFindAll(payload interface{}) *Response {
	return &Response{
		Status:  http.StatusOK,
		Payload: payload,
	}
}

func ErrBadRequest(err interface{}) *Response {
	return &Response{
		Status: http.StatusBadRequest,
		Error:  err,
	}
}

func ErrInternalServer(err interface{}) *Response {
	return &Response{
		Status: http.StatusInternalServerError,
		Error:  err,
	}
}
func ErrNotFound() *Response {
	return &Response{
		Status: http.StatusNotFound,
		Error:  "NO_DATA",
	}
}
func ErrUnauthorized() *Response {
	return &Response{
		Status: http.StatusUnauthorized,
		Error:  "UNAUTHORIZED",
	}
}
