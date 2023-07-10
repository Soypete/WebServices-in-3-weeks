package main

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"net/url"

	"github.com/pressly/goose/v3"
	"github.com/soypete/WebServices-in-3-weeks/database/demo/database"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

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

	db, err := sql.Open("postgres", connectionString.String())
	if err != nil {
		panic(err)
	}

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}

	queires := database.New(db)

	user, err := queires.GetUser(context.Background(), "pete")
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
