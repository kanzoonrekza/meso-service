package services

import (
	"encoding/json"
	"net/http"
)

func Status(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"service": "Meso",
		"status":  "OK",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
