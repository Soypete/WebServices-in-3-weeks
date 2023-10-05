package main

import (
	"log"

	"github.com/soypete/WebServices-in-3-weeks/database/ex-2-abstraction/database"
)

func main() {

	// this will enforge the interface
	var db database.Connector
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	// call your client methods here
}
