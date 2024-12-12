package services

import (
	"encoding/json"
	"fmt"
	"meso/db"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	pool, ok := r.Context().Value("pool").(*pgxpool.Pool)
	if !ok {
		http.Error(w, "Database connection not available", http.StatusInternalServerError)
		return
	}

	queries := db.New(pool)
	ctx := r.Context()

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
