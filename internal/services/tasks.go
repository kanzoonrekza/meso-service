package services

import (
	"encoding/json"
	"fmt"
	"meso/internal/utils"
	"net/http"
)

func GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	queries, ctx := utils.GetDBCtx(w, r)

	tasks, err := queries.GetAllTasks(ctx)
	if err != nil {
		fmt.Println("Error getting tasks:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}
