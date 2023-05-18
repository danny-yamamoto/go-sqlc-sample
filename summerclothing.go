package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/danny-personal/go-sqlc-sample/tutorial"
	_ "github.com/lib/pq"
)

func run() error {
	ctx := context.Background()
	db, err := sql.Open("postgres", "postgres://postgres:password@192.168.0.56/pagila")
	if err != nil {
		return err
	}
	queries := tutorial.New(db)
	payments, err := queries.ListPayment(ctx)
	if err != nil {
		return err
	}
	log.Println(payments)
	return nil
}

func main() {
	fmt.Println("hello")
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
