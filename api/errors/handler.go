package errors

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func ErrorResponseJSON(w http.ResponseWriter, err error, statusCode int) {
	response := ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: err.Error(),
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
