package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// RespondWithJSON is a utility function to respond with JSON data
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}

// HandleError is a utility function to handle errors and respond with an appropriate message
func HandleError(w http.ResponseWriter, err error, code int) {
	log.Printf("Error: %v", err)
	RespondWithJSON(w, code, map[string]string{"error": err.Error()})
}
