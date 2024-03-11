package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	connStr := "postgres://postgres:1234@localhost:5432/db?sslmode=require"
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close(ctx)
}