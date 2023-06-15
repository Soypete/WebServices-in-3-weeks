package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	connectionString := "server=127.0.0.1;uid=root;pwd=12345;database=test"
	db, err := sqlx.Connect("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")
	db.Close()
}
