package services

import (
	"meso/internal/utils"
	"net/http"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	utils.CreateJSONResponse(w, http.StatusOK, "Welcome to Meso!")
}
