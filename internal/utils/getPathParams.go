package utils

import (
	"fmt"
	"net/http"
	"strconv"
)

func GetPathParamAsString(w http.ResponseWriter, r *http.Request, value string, v *string) bool {
	pathParams := r.PathValue(value)
	if pathParams == "" {
		fmt.Println("Path parameter not found:", value)
		return false
	}

	*v = pathParams
	return true
}
func GetPathParamAsNumber(w http.ResponseWriter, r *http.Request, value string, v *int) bool {
	var str string
	if !GetPathParamAsString(w, r, value, &str) {
		return false
	}

	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error converting params as number:", err)
		return false
	}

	*v = num
	return true
}
