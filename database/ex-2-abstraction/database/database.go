package database

import (
	"database/sql"
	"fmt"
)

// Connection struct holds the database connection
type Connection struct {
	db *sql.DB
}

// User struct holds the user data
type User struct {
	//TODO: add user fields
}

type Connector interface {
	QueryUser() (*User, error)
	UpdateUser() error
	AddUser() error
}

// Connect to database and return a Connection struct.
// This return value needs to be a pointer because we are
// using it to implement the Connector interface.
func Connect() (*Connection, error) {
	fileName := "database/ex-1-connection/sqlite.db"
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		return nil, fmt.Errorf("cannot open sqlite connection: %w", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("cannot ping database: %w", err)
	}
	return &Connection{db: db}, nil
}

func (c *Connection) QueryUser() (*User, error) {
	// TODO: query database
}

func (c *Connection) UpdateUser() error {
	// TODO: update database
}

func (c *Connection) AddUser() error {
	// TODO: add user to database
}
