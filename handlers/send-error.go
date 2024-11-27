package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func SendError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResponse := ErrorResponse{
		Message: message,
	}

	if err := json.NewEncoder(w).Encode(errorResponse); err != nil {
		log.Println("Failed to send error response:", err)
	}
}
