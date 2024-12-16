package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetDecodedJSON(w http.ResponseWriter, r io.Reader, decoded interface{}) bool {
	if err := json.NewDecoder(r).Decode(&decoded); err != nil {
		fmt.Println("Error decoding JSON:", err)
		CreateJSONResponse(w, http.StatusBadRequest, err)
		return false
	}

	return true
}
