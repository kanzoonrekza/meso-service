package middleware

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DBPoolHandler(pool *pgxpool.Pool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "pool", pool)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
