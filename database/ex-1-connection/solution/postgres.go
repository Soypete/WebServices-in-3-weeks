package main

import (
	"fmt"
	"net/url"

	"github.com/jmoiron/sqlx"
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

	db, err := sqlx.Connect("postgres", connectionString.String())
	if err != nil {
		panic(err)
	}
	fmt.Println(db.Ping())
}
