package services

import (
	"database/sql"
	"errors"
	"fmt"
	"meso/db"
	"meso/internal/utils"
	"net/http"
)

func CreateListHandler(w http.ResponseWriter, r *http.Request) {
	queries, ctx := utils.GetDBCtx(w, r)

	var list db.CreateListParams
	if !utils.GetDecodedJSON(w, r.Body, &list) {
		return
	}

	createdList, err := queries.CreateList(ctx, list)
	if err != nil {
		fmt.Println("Error creating list:", err)
		utils.CreateJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.CreateJSONResponse(w, http.StatusCreated, createdList)
}

func UpdateListHandler(w http.ResponseWriter, r *http.Request) {
	queries, ctx := utils.GetDBCtx(w, r)

	var id int
	if !utils.GetPathParamAsNumber(w, r, "id", &id) {
		return
	}

	var list db.UpdateListParams
	if !utils.GetDecodedJSON(w, r.Body, &list) {
		return
	}
	list.ID = int64(id)

	updatedList, err := queries.UpdateList(ctx, list)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("Error list not found:", err)
			utils.CreateJSONResponse(w, http.StatusNotFound, "List not found")
			return
		}
		fmt.Println("Error updating list:", err)
		utils.CreateJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.CreateJSONResponse(w, http.StatusOK, updatedList)
}

func DeleteListHandler(w http.ResponseWriter, r *http.Request) {
	queries, ctx := utils.GetDBCtx(w, r)

	var id int
	if !utils.GetPathParamAsNumber(w, r, "id", &id) {
		return
	}

	rows, err := queries.DeleteList(ctx, int64(id))
	if err != nil {
		fmt.Println("Error deleting list:", err)
		utils.CreateJSONResponse(w, http.StatusInternalServerError, err)
		return
	}
	if rows == 0 {
		utils.CreateJSONResponse(w, http.StatusNotFound, "List not found")
		return
	}

	utils.CreateJSONResponse(w, http.StatusOK, "Delete successful")
}

func GetAllListsHandler(w http.ResponseWriter, r *http.Request) {
	queries, ctx := utils.GetDBCtx(w, r)

	lists, err := queries.GetAllLists(ctx)
	if err != nil {
		fmt.Println("Error getting lists:", err)
		utils.CreateJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.CreateJSONResponse(w, http.StatusOK, lists)
}

func GetListByIDHandler(w http.ResponseWriter, r *http.Request) {
	queries, ctx := utils.GetDBCtx(w, r)

	var id int
	if !utils.GetPathParamAsNumber(w, r, "id", &id) {
		return
	}

	list, err := queries.GetListByID(ctx, int64(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("Error list not found:", err)
			utils.CreateJSONResponse(w, http.StatusNotFound, "List not found")
			return
		}
		fmt.Println("Error getting list:", err.Error())
		utils.CreateJSONResponse(w, http.StatusNotFound, err)
		return
	}

	utils.CreateJSONResponse(w, http.StatusOK, list)
}
