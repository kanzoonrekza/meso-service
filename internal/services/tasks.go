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

	jsonData, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("Error marshaling tasks to JSON:", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
