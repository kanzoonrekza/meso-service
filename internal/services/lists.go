package services

import (
	"encoding/json"
	"fmt"
	"meso/db"
	"meso/internal/utils"
	"net/http"
)

func GetAllLists(w http.ResponseWriter, r *http.Request) {
	queries, ctx := utils.GetDBCtx(w, r)

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

func CreateList(w http.ResponseWriter, r *http.Request) {
	queries, ctx := utils.GetDBCtx(w, r)

	var list db.CreateListParams

	err := json.NewDecoder(r.Body).Decode(&list)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	createdList, err := queries.CreateList(ctx, list)
	if err != nil {
		fmt.Println("Error creating list:", err)
		return
	}

	jsonData, err := json.Marshal(createdList)
	if err != nil {
		fmt.Println("Error marshaling list to JSON:", err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
