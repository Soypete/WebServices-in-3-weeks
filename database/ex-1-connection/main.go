// to run postgres locally, you can use docker
// docker pull postgres
// docker run --name postgres -e POSTGRES_PASSWORD=postgres -p 5431:5431 -d postgres
//
// otherwise use the sqlite provided in the repo

package main

// add driver here

func main() {
	connectionString := "/database.db"
	// TODO: connect to sqlite database here

	// TODO: query the database here
	db.Close()
}
