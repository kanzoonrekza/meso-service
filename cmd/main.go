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

	mux.HandleFunc("/status", services.Status)
	mux.HandleFunc("GET /tasks", services.GetAllTasks)
	mux.HandleFunc("GET /lists", services.GetAllLists)
	mux.HandleFunc("POST /lists", services.CreateList)

	server.ListenAndServe()
}
