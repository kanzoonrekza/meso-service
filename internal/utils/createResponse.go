package utils

import (
	"encoding/json"
	"net/http"
)

func CreateJSONResponse(w http.ResponseWriter, status int, payload interface{}) {
	response := map[string]interface{}{
		"status_code":    status,
		"status_message": http.StatusText(status),
	}

	switch t := payload.(type) {
	case nil:
		// Do nothing
	case error:
		response["error_details"] = t.Error()
	case string:
		response["message"] = t
	default:
		response["payload"] = t
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
