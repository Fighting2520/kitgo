package jsonresponse

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	ErrorResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	Response struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)

func ErrorEncoder(c context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(ErrorResponse{Code: 9999, Message: err.Error()})
}
