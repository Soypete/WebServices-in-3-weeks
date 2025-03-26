# Databases in Go

In the previous tutorial, we learned how to build and test a RESTful API server using Go. But most APIs are not just logic—they’re data access layers. In this section, we’ll cover how to interact with databases from Go. Specifically, we’ll cover:

* Why SQL is still the standard
* How to connect Go servers to a database
* Tools for managing queries and migrations
* How to test database-integrated code

## Why SQL?

Structured Query Language (SQL) is the most popular language for working with relational databases. It’s declarative—meaning you tell the database *what* you want, not *how* to do it. This makes your data layer flexible and readable.

> "SQL works with a variety of data stores… and allows you to define schemas that serve as contracts." — [Wikipedia](https://en.wikipedia.org/wiki/SQL)

Some popular SQL dialects you’ll encounter:
- PostgreSQL
- MySQL
- SQLite
- DuckDB
- and warehouse SQL variants like BigQuery SQL or Redshift SQL

## Mapping RESTful Calls to SQL

When building APIs, we can think of HTTP methods as shorthand for SQL operations:

| HTTP Method | SQL Operation |
|-------------|----------------|
| GET         | SELECT         |
| POST        | INSERT         |
| PUT         | UPDATE         |
| DELETE      | DELETE         |

So when you write an API that fetches a user profile, you might write a `GET /users/{id}` endpoint that maps to:

```sql
SELECT * FROM users WHERE id = $1;
```

## Using Go’s `database/sql` Package

Go’s built-in database library, `database/sql`, provides a generic interface to many SQL databases. It’s extensible by choosing a driver, such as:

- PostgreSQL: [`github.com/lib/pq`](https://github.com/lib/pq)
- MySQL: [`github.com/go-sql-driver/mysql`](https://github.com/go-sql-driver/mysql/)
- SQLite: [`github.com/mattn/go-sqlite3`](https://github.com/mattn/go-sqlite3)
- DuckDB: [`github.com/marcboeker/go-duckdb`](https://github.com/marcboeker/go-duckdb)

To get a drive into your module, you can use `go get`:

```bash
go mod init github.com/{username}/repo
go get "github.com/mattn/go-sqlite3"
go mod tidy
```

### Example: Connecting to Postgres

```go
import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

func main() {
	connectionString := "postgresql://user:secret@localhost/mydb?sslmode=disable"
	db, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")
	db.Close()
}
```

### Example: mysql

```go
import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	connectionString := "server=127.0.0.1;uid=root;pwd=12345;database=test"
	db, err := sqlx.Connect("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")
	db.Close()
}
```

### Example sqlite3

```go
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fileName := "/database.db"
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")
	db.Close()
}
```

### Best Practices

- Always defer `rows.Close()` when querying.
- Use `context.Context` in production apps to support timeouts and cancellations.
- Check for errors from both the query and the rows object.

## Making Queries

Here’s how to insert and fetch data:

```go
// Insert
_, err := db.ExecContext(ctx,
  "INSERT INTO users(name, email) VALUES($1, $2)", "Miriah", "miriah@example.com")

// Select
row := db.QueryRowContext(ctx, "SELECT id FROM users WHERE email = $1", "miriah@example.com")
var id int
if err := row.Scan(&id); err != nil {
	log.Fatal(err)
}
```

for postgres you use the numbered placeholders `$1`, `$2`, etc. For MySQL you use `?`.


here is an example of a select query to get many columns from a table:

```go
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate over the rows
	for rows.Next() {
		var id int
		var name, email string
		err = rows.Scan(&id, &name, &email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name, email)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
```

## Managing Schema and Migrations

Rather than hand-writing `CREATE TABLE` statements, you should manage your schema with tools. Go has some great ones:

- [`sqlc`](https://sqlc.dev): write SQL, generate Go code
- [`gorm`](https://gorm.io): ORM with structs as models
- [`sqlx`](https://github.com/jmoiron/sqlx): lightweight extensions over `database/sql`
- [`goose`](https://github.com/pressly/goose): schema migration tool

### Example: Goose Migration

Goose is a database migration tool that lets you write SQL migrations and run them from the command line. You can also use it programmatically by running the migrations in your Go code. I prefer this approach because it keeps your migrations in source control. It also allows you to embed the migrations in your binary.

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest

# Create a migration
goose create create_users_table sql
```

This creates a file with up/down SQL sections. You can then run:

```bash
goose -dir db/migrations postgres "your-connection-string" up
```

here is an video of the [goose migration demo]()

### Example: SQLC

Sqlc is a tool that generates Go code from SQL queries. You write your queries in a `.sql` file and run `sqlc generate` to create Go code. This is great for keeping your queries in one place and generating type-safe Go code.

```bash

```

herre is a videro of the of the [sqlc demo]()

## Organizing Your DB Code

It’s common in Go to wrap your DB in a struct and expose methods that encapsulate queries. For example:

```go
type DBClient struct {
	db *sql.DB
}

func (c *DBClient) GetUserByEmail(email string) (*User, error) {
	row := c.db.QueryRow("SELECT id, name FROM users WHERE email = $1", email)
	u := &User{}
	err := row.Scan(&u.ID, &u.Name)
	return u, err
}
```

This gives you:
- Better testability with mockable interfaces
- A clear “database API” surface
- Separation of concerns between logic and data

## Testing with Mocks

You should avoid writing integration tests that depend on a running database in every unit test. Instead, create interfaces and use mock implementations.

```go
type UserRepository interface {
	GetUserByEmail(email string) (*User, error)
}
```

Then in your test:

```go
type mockUserRepo struct{}

func (m *mockUserRepo) GetUserByEmail(email string) (*User, error) {
	return &User{ID: 1, Name: "Mock User"}, nil
}
```

Now your handlers can depend on the interface rather than a concrete DB connection.

## Exercises

### Exercise 1

In your server project, add your preferred database driver and connect to the database in the main function (if you missed day one's exercises or have them in a different location using the [ex-1-connection/main.go](ex-1-connection/main.go)). After you have connected and verified your connection, explore the database. Make sure to query the database's users table and handle the error. Try running `SELECT`, `INSERT`, and `UPDATE` statements

### Follow-up questions:

* What kind of package organization would make sense for organizing your database logic?
* What database driver did you pick?
* Did the data persist?

_NOTE_: If you are not using postgres or are completing this independently. You can run many databases locally using docker. Below is an example of running postgres locally in a docker container.

```
docker pull postgres
docker run -e POSTGRES_PASSWORD=postgres -e POSTGRES_USERNAME=postgres -p 5431:5432 postgres
```

After you get docker running in your local environment set up your database. You will need to `CREATE` your tables and `INSERT` data into the table. You can do this in your Go app or via a sql script editor. [psql](https://www.postgresql.org/docs/current/app-psql.html) is postgres's command line tool.

An example of a go app that connect to a local postgres instance is in [database/ex-1-connection/solution](/database/ex-1-connection/solution/postgres.go).


### Exercise 2

Build a client and interface around your database connection. You can use your existing main.go file and build new database package for your abstraction, or you can use the template files found in [database/ex-2-abstraction](/database/ex-2-abstraction/main.go). Make sure to create an interface, a User struct, a database client, and the methods to create, update, and query the User data. If you have any questions please put them in the chat.

Follow-up Questions:

* what are the differences between manually creating a database object vs generating one with sqlc?

### Exercise 3 mock database

Using your new database interface mock the database functions into your [tests from last week](../restful-go/ex-4-tests/solution/framework_test.go). The goal is to imitate db interactions without connecting to the db. You will need to add the DB package to the same repo that your server lives in.

[Here](https://github.com/Soypete/golang-cli-game/blob/24dc57852dee27bb17120555d3d390bd17a78d13/server/api_test.go#L14) are some working tests that use `passBD{}` and `failDB{}` to mock database functionality in an API test.
