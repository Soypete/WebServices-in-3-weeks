// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package database

import (
	"database/sql"
	"time"
)

type Game struct {
	ID        int32
	Host      string
	Players   []string
	Answer    string
	Questions []string
	Guesses   []string
	StartTime time.Time
	EndTime   sql.NullTime
	Ended     sql.NullBool
}

type Guess struct {
	ID      int32
	Guess   string
	UserID  int32
	GameID  int32
	Correct sql.NullBool
}

type Question struct {
	ID       int32
	Question string
	UserID   int32
	GameID   int32
}

type User struct {
	ID        int32
	Username  string
	CreatedAt time.Time
}
