package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fileName := "database/ex-1-connection/sqlite.db"
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, email TEXT UNIQUE, password TEXT NOT NULL)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ('John Doe', ' john.doe@email.com', 'opensesame')")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ('Jane Smith', 'smithjane@email.com', 'opensesame')")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ('Robert Johnson', 'me@robertjohnson.com', 'opensesame')")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ('Emily Davis', 'emily_davis@email.com', 'opensesame')")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ('Michael Wilson', 'mwilson@email.com', 'opensesame')")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ('Sarah Brown', 'sbbrown@email.com', 'opensesame')")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ('William Jones', 'whjones@email.com', 'opensesame')")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ('Olivia Taylor', 'olivia.taylord@email.com', 'opensesame')")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ('David Evans', 'david@evans.com', 'opensesame')")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ('Sophia Miller', 'sophiamiller@sophiamiller.com', 'opensesame')")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ('James Wilson', 'jamesthewil@email.com', 'opensesame')")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ('Ava Martin', 'avamart@email.com', 'opensesame')")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ('Charles Anderson', 'charles@anderson', 'opensesame')")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ('Mia Martinez', 'miacasaessucasa@martinez.com', 'opensesame')")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ('Joseph Thomas', 'joe@thomas.com', 'opensesame')")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ('Chloe White', 'chloewhite@email.com', 'opensesame')")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ('Daniel Harris', 'test@user.com', 'opensesame')")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ('Emma Moore', 'emma_more@test.com', 'opensesame')")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ('Matthew Clark', 'matt_clark@email.com', 'opensesame')")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ('Harper Lewis', 'harp@lewis.com', 'opensesame')")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var email string
		var password string
		rows.Scan(&id, &name, &email, &password)
		fmt.Println(id, name, email, password)
	}

}
