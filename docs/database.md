Great — I’ve reviewed the contents of the pull request. Based on that and your example format, here’s the draft of your **Go + SQL** tutorial, styled and structured to match your "RESTful Go" example:

---

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

### Example: Connecting to Postgres

```go
import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	connStr := "postgres://user:pass@localhost/dbname?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("could not connect to db:", err)
	}

	log.Println("Connected to database!")
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

## Managing Schema and Migrations

Rather than hand-writing `CREATE TABLE` statements, you should manage your schema with tools. Go has some great ones:

- [`sqlc`](https://sqlc.dev): write SQL, generate Go code
- [`gorm`](https://gorm.io): ORM with structs as models
- [`sqlx`](https://github.com/jmoiron/sqlx): lightweight extensions over `database/sql`
- [`goose`](https://github.com/pressly/goose): schema migration tool

### Example: Goose Migration

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest

# Create a migration
goose create create_users_table sql
```

This creates a file with up/down SQL sections. You can then run:

```bash
goose -dir db/migrations postgres "your-connection-string" up
```

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

Here are some hands-on activities to try:

1. **Connect to a Postgres database**  
   Use the `sql.Open` example to connect to a database running locally or in Docker.
   [Exercise](https://github.com/Soypete/WebServices-in-3-weeks/tree/main/database/ex-1-connection)

2. **Wrap your database logic**  
   Create a `DBClient` struct with methods like `CreateUser` or `FindUserByID`.
   [Example](https://github.com/Soypete/WebServices-in-3-weeks/blob/main/database/demo/database/db.go)

3. **Write SQL with `sqlc` and auto-generate Go code**  
   Write a `.sql` file with your query and use sqlc to generate Go code.
   [Example user.sql](https://github.com/Soypete/WebServices-in-3-weeks/blob/main/database/demo/queries/user.sql)

4. **Write tests using mocks**  
   Use interfaces to test your handler logic without spinning up a real database.
   [Example](https://github.com/Soypete/golang-cli-game/blob/main/server/api_test.go)

---

Would you like this turned into a markdown file for your repo, or published to your blog as Article 2?
