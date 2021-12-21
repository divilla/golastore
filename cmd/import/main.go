package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

func main() {
	pool, err := pgxpool.Connect(context.Background(), "postgres://postgres:Kita1818@localhost:5432/ekupi")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	r := newRepository(context.Background(), pool)
	r.categoryNameToSlug()
}
