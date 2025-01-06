package main

import (
	"fmt"
	"net/url"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	params := url.Values{}
	params.Set("sslmode", "disable")

	// this is a personal preference to use url.URL to
	// build up the connection string. This works well for
	// postgres, but other drivers might have their own quirks.
	connectionString := url.URL{
		Scheme:   "postgresql",
		User:     url.UserPassword("postgres", "postgres"),
		Host:     "localhost:5431",
		Path:     "postgres",
		RawQuery: params.Encode(),
	}

	//"postgresql://postgres:postgres@localhost:5431/postgres?sslmode=disable"

	db, err := sqlx.Connect("postgres", connectionString.String())
	if err != nil {
		panic(err)
	}
	fmt.Println(db.Ping())

	// create a new user
	_, err = db.Exec(`INSERT INTO users (email, password) VALUES ($1, $2)`, "soypete@email.com", "password")
	if err != nil {
		panic(err)
	}

	// get the newly created user
	rows, err := db.Queryx("SELECT * FROM users")
	if err != nil {
		panic(err)
	}

	type User struct {
		ID       int    `db:"id"`
		Email    string `db:"email"`
		Password string `db:"password"`
	}

	for rows.Next() {
		var user User
		err = rows.StructScan(&user)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%#v\n", user)
	}

}
