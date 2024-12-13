package services

import (
	"encoding/json"
	"fmt"
	"meso/db"
	"meso/internal/utils"
	"net/http"
	"strconv"
)

func GetAllListsHandler(w http.ResponseWriter, r *http.Request) {
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

func GetListByIDHandler(w http.ResponseWriter, r *http.Request) {
	queries, ctx := utils.GetDBCtx(w, r)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		fmt.Println("Error getting ID from request:", err)
		return
	}

	list, err := queries.GetListByID(ctx, int64(id))
	if err != nil {
		fmt.Println("Error getting list:", err)
		return
	}

	jsonData, err := json.Marshal(list)
	if err != nil {
		fmt.Println("Error marshaling list to JSON:", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func CreateListHandler(w http.ResponseWriter, r *http.Request) {
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

func DeleteListHandler(w http.ResponseWriter, r *http.Request) {
	queries, ctx := utils.GetDBCtx(w, r)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		fmt.Println("Error getting ID from request:", err)
		return
	}

	// TODO: Find the list with that id first before executing the deletion

	err = queries.DeleteList(ctx, int64(id))
	if err != nil {
		fmt.Println("Error deleting list:", err)
		return
	}

	response := map[string]interface{}{
		"message": "Delete success",
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error marshaling list to JSON:", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}