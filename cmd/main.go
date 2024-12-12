package main

import (
	"context"
	"encoding/json"
	"fmt"
	"meso/db"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	fmt.Println("☄️☄️☄️")

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, os.Getenv("DB_URL"))
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer conn.Close(context.Background())

	queries := db.New(conn)

	tasks, err := queries.GetAllTasks(ctx)
	if err != nil {
		fmt.Println("Error getting tasks:", err)
		return
	}

	jsonData, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("Error marshaling tasks to JSON:", err)
	}

	fmt.Println("JSON Output Tasks:", string(jsonData))

}
