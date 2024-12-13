package utils

import (
	"context"
	"meso/db"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetDBCtx(w http.ResponseWriter, r *http.Request) (*db.Queries, context.Context) {
	pool, ok := r.Context().Value("pool").(*pgxpool.Pool)
	if !ok {
		http.Error(w, "Database connection not available", http.StatusInternalServerError)
		return db.New(nil), context.Background()
	}

	queries := db.New(pool)
	ctx := r.Context()

	return queries, ctx
}
