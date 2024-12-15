package services

import (
	"encoding/json"
	"fmt"
	"meso/db"
	"meso/internal/utils"
	"net/http"
	"strconv"
)

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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdList)

}

func UpdateListHandler(w http.ResponseWriter, r *http.Request) {
	queries, ctx := utils.GetDBCtx(w, r)

	var list db.UpdateListParams

	err := json.NewDecoder(r.Body).Decode(&list)
	fmt.Println(list)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	updatedList, err := queries.UpdateList(ctx, list)
	if err != nil {
		fmt.Println("Error updating list:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedList)
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func GetAllListsHandler(w http.ResponseWriter, r *http.Request) {
	queries, ctx := utils.GetDBCtx(w, r)

	lists, err := queries.GetAllLists(ctx)
	if err != nil {
		fmt.Println("Error getting lists:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lists)

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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(list)
}
