package services

import (
	"encoding/json"
	"fmt"
	"meso/db"
	"meso/internal/utils"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	queries, ctx := utils.GetDBCtx(w, r)

	var task db.CreateTaskParams

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	createdTask, err := queries.CreateTask(ctx, task)
	if err != nil {
		fmt.Println("Error creating task:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTask)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	queries, ctx := utils.GetDBCtx(w, r)

	// var task map[string]interface{}
	var task db.UpdateTaskParams

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		fmt.Println("Error decoding JSON task:", err)
		return
	}
	fmt.Println(task)

	updatedTask, err := queries.UpdateTask(ctx, task)
	if err != nil {
		fmt.Println("Error updating task:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedTask)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	queries, ctx := utils.GetDBCtx(w, r)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		fmt.Println("Error getting ID from request:", err)
		return
	}

	// TODO: Find the task with that id first before executing the deletion

	err = queries.DeleteTask(ctx, int64(id))
	if err != nil {
		fmt.Println("Error deleting task:", err)
		return
	}

	response := map[string]interface{}{
		"message": "Delete success",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

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

func GetTaskDetailsHandler(w http.ResponseWriter, r *http.Request) {
	queries, ctx := utils.GetDBCtx(w, r)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		fmt.Println("Error getting ID from request:", err)
		return
	}

	task, err := queries.GetTaskByID(ctx, int64(id))
	if err != nil {
		fmt.Println("Error getting task:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func GetTasksByParentListIDHandler(w http.ResponseWriter, r *http.Request) {
	queries, ctx := utils.GetDBCtx(w, r)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		fmt.Println("Error getting ID from request:", err)
		return
	}

	tasks, err := queries.GetTasksByParentListID(ctx, pgtype.Int8{Int64: int64(id), Valid: true})
	if err != nil {
		fmt.Println("Error getting tasks:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}
