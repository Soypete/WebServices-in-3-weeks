package main

import (
	"fmt"
	"net/url"

	"github.com/jmoiron/sqlx"
)

type connection interface {
	GetUserData(string) (string, error)
	UpsertUsername(string) error
	DeleteUsername(string) error
	CreateGame(string) (int64, error)
	AddUserToGame(string, int64) error
	GetGameData(int64) (string, error)
	StopGame(int64) error
}

type client struct {
	db *sqlx.DB
}

func setup() *client {
	// connect to db
	params := url.Values{}
	params.Set("sslmode", "disable")

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

	return &client{
		db: db,
	}
}
func main() {
	dbClient := setup()

	fmt.Println(dbClient.db.Ping())

}
