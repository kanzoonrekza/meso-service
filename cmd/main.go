package main

import (
	"context"
	"fmt"
	"meso/internal/middleware"
	"meso/internal/services"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	fmt.Println("☄️☄️☄️")

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, os.Getenv("DB_URL"))
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer pool.Close()

	mux := http.NewServeMux()

	server := http.Server{
		Addr:    ":8000",
		Handler: middleware.DBPoolHandler(pool)(mux),
	}

	mux.HandleFunc("/status", services.StatusHandler)

	mux.HandleFunc("GET /lists", services.GetAllListsHandler)
	mux.HandleFunc("GET /lists/{id}", services.GetListByIDHandler)
	mux.HandleFunc("POST /lists", services.CreateListHandler)
	mux.HandleFunc("DELETE /lists/{id}", services.DeleteListHandler)

	mux.HandleFunc("GET /tasks", services.GetAllTasksHandler)

	server.ListenAndServe()
}
