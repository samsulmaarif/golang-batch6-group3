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
		Message:     "REGISTER_SUCCESS",
		GeneralInfo: "NooBee-Shop",
	}
}

func SuccessLogin(message interface{}, payload interface{}) *Response {
	return &Response{
		Status:      http.StatusCreated,
		Message:     message,
		Payload:     payload,
		GeneralInfo: "NooBee-Shop",
	}
}

func SuccessDeleted(message interface{}) *Response {
	return &Response{
		Status:      http.StatusCreated,
		Message:     message,
		GeneralInfo: "NooBee-Shop",
	}
}

func SuccessFindAll(payload interface{}) *Response {
	return &Response{
		Status:      http.StatusOK,
		Payload:     payload,
		GeneralInfo: "NooBee-Shop",
	}
}

func ErrBadRequest(err interface{}) *Response {
	return &Response{
		Status:      http.StatusBadRequest,
		GeneralInfo: "NooBee-Shop",
		Error:       err,
	}
}

func ErrInternalServer(err interface{}) *Response {
	return &Response{
		Status:      http.StatusInternalServerError,
		Error:       err,
		GeneralInfo: "NooBee-Shop",
	}
}
func ErrNotFound() *Response {
	return &Response{
		Status:      http.StatusNotFound,
		Error:       "NO_DATA",
		GeneralInfo: "NooBee-Shop",
	}
}
func ErrUnauthorized() *Response {
	return &Response{
		Status:      http.StatusUnauthorized,
		Error:       "UNAUTHORIZED",
		GeneralInfo: "NooBee-Shop",
	}
}
