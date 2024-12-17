package services

import (
	"database/sql"
	"errors"
	"fmt"
	"meso/db"
	"meso/internal/utils"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	queries, ctx := utils.GetDBCtx(w, r)

	var task db.CreateTaskParams
	if !utils.GetDecodedJSON(w, r.Body, &task) {
		return
	}

	pTask, err := queries.GetParentTaskData(ctx, task.ParentTaskID.Int64)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("Error parent task not found")
			utils.CreateJSONResponse(w, http.StatusNotFound, "Parent task not found")
			return
		}
		fmt.Println("Error checking parent task:", err)
		utils.CreateJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	if pTask.ParentTaskID.Valid {
		fmt.Println("Error parent task is a subtask")
		utils.CreateJSONResponse(w, http.StatusBadRequest, "Parent task is a subtask")
		return
	}

	createdTask, err := queries.CreateTask(ctx, task)
	if err != nil {
		fmt.Println("Error creating task:", err)
		utils.CreateJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.CreateJSONResponse(w, http.StatusCreated, createdTask)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	queries, ctx := utils.GetDBCtx(w, r)

	var id int
	if !utils.GetPathParamAsNumber(w, r, "id", &id) {
		return
	}

	var task db.UpdateTaskParams
	if !utils.GetDecodedJSON(w, r.Body, &task) {
		return
	}
	task.ID = int64(id)

	updatedTask, err := queries.UpdateTask(ctx, task)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("Error task not found:", err)
			utils.CreateJSONResponse(w, http.StatusNotFound, "Task not found")
			return
		}
		fmt.Println("Error updating task:", err)
		utils.CreateJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.CreateJSONResponse(w, http.StatusOK, updatedTask)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	queries, ctx := utils.GetDBCtx(w, r)

	var id int
	if !utils.GetPathParamAsNumber(w, r, "id", &id) {
		return
	}

	rows, err := queries.DeleteTask(ctx, int64(id))
	if err != nil {
		fmt.Println("Error deleting task:", err)
		utils.CreateJSONResponse(w, http.StatusInternalServerError, err)
		return
	}
	if rows == 0 {
		utils.CreateJSONResponse(w, http.StatusNotFound, "Task not found")
		return
	}

	utils.CreateJSONResponse(w, http.StatusOK, "Delete successful")
}

func GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	queries, ctx := utils.GetDBCtx(w, r)

	tasks, err := queries.GetAllTasks(ctx)
	if err != nil {
		fmt.Println("Error getting tasks:", err)
		return
	}

	utils.CreateJSONResponse(w, http.StatusOK, tasks)
}

// TODO
func GetTaskDetailsHandler(w http.ResponseWriter, r *http.Request) {
	queries, ctx := utils.GetDBCtx(w, r)

	var id int
	if !utils.GetPathParamAsNumber(w, r, "id", &id) {
		return
	}

	pTask, err := queries.GetTaskByID(ctx, int64(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.CreateJSONResponse(w, http.StatusNotFound, "Task not found")
			return
		}
		fmt.Println("Error getting task:", err)
		utils.CreateJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	sTasks, err := queries.GetTasksByParentTaskID(ctx, pgtype.Int8{Int64: int64(id), Valid: true})
	if err != nil {
		fmt.Println("Error getting subtasks:", err)
		utils.CreateJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	response := db.TaskWithSubtasks{
		Task:     pTask,
		Subtasks: sTasks,
	}

	utils.CreateJSONResponse(w, http.StatusOK, response)
}

func GetTasksByParentListIDHandler(w http.ResponseWriter, r *http.Request) {
	queries, ctx := utils.GetDBCtx(w, r)

	var id int
	if !utils.GetPathParamAsNumber(w, r, "id", &id) {
		return
	}

	list, err := queries.GetListExistanceByID(ctx, int64(id))
	if err != nil {
		fmt.Println("Error getting list:", err)
		utils.CreateJSONResponse(w, http.StatusInternalServerError, err)
		return
	}
	if !list {
		utils.CreateJSONResponse(w, http.StatusNotFound, "List not found")
		return
	}

	tasks, err := queries.GetTasksByParentListID(ctx, pgtype.Int8{Int64: int64(id), Valid: true})
	if err != nil {
		fmt.Println("Error getting tasks:", err)
		utils.CreateJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.CreateJSONResponse(w, http.StatusOK, tasks)
}
