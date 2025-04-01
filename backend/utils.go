package main

import (
	"encoding/json"
	"net/http"
)

// Función para responder con un JSON
func respondWithJSON(w http.ResponseWriter, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}

// Función para responder con un mensaje de error en formato JSON
func respondWithError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ApiResponse{
		Success: false,
		Message: message,
	})
}
