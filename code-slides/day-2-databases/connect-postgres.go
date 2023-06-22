//go:build ignore

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

func main() {
	connectionString := "postgresql://user:secret@localhost/mydb?sslmode=disable"
	db, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")
	db.Close()
}
