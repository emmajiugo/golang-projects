package api

import (
	"encoding/json"
	"net/http"
)

// Coin Balance Params
type CoinBalanceParams struct {
	Username string
}

// Coin Balance Response
type CoinBalanceResponse struct {
	// Success code, usally 200
	Code int

	// Account balance
	Balance int64
}

// Error Response
type ErrorResponse struct {
	// Error code, usally 400 or 500
	Code int
	// Error message
	Message string
}

func writeError(w http.ResponseWriter, code int, message string) {
	response := ErrorResponse{
		Code:    code,
		Message: message,
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(response)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, http.StatusBadRequest, err.Error())
	}
	InternalServerErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, http.StatusInternalServerError, "Internal server error")
	}
)