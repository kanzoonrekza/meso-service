package services

import (
	"encoding/json"
	"fmt"
	"meso/db"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetAllLists(w http.ResponseWriter, r *http.Request) {
	pool, ok := r.Context().Value("pool").(*pgxpool.Pool)
	if !ok {
		http.Error(w, "Database connection not available", http.StatusInternalServerError)
		return
	}

	queries := db.New(pool)
	ctx := r.Context()

	lists, err := queries.GetAllLists(ctx)
	if err != nil {
		fmt.Println("Error getting lists:", err)
		return
	}

	jsonData, err := json.Marshal(lists)
	if err != nil {
		fmt.Println("Error marshaling lists to JSON:", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
