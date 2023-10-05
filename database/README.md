# Day 2 Databases and your Go app

This section of the course contains examples and exercises for effectively leveraging a database in the backend of your Web services.. The following are examples of finished exercises for day 2 of the [O'reilly Media Online Learning Course](https://www.oreilly.com/live-events/go-for-web-development-in-3-weeks/0636920091015/). They can be completed by following along with the instruction or independently.

## Exercise 1

In your server project, add your preferred database driver and connect to the database in the main function (if you missed day one's exercises or have them in a different location using the [ex-1-connection/main.go](ex-1-connection/main.go)). After you have connected and verified your connection, explore the database. Make sure to query the database's users table and handle the error. Try running `SELECT`, `INSERT`, and `UPDATE` statements

### Follow up questions:

* what kind of package organization would make sense for organizing your database logic?
* what data base driver did you pick?
* did the data persist?

_NOTE_: If you are not using postgres or are completeling this indepently. You can run many datbases locally using docker. Below is an example of running postgres locally in a docker container.

```
docker pull postgres
docker run -e POSTGRES_PASSWORD=my_password -e POSTGRES_USERNAME=user1 -p 5431:5432 postgres
```

After you get docker running in your local environment setup your database. You will need to `CREATE` your tables and `INSERT` data into the table. You can do this in your Go app or via a sql script editor. [psql](https://www.postgresql.org/docs/current/app-psql.html) is postgres's command line tool.

An example of a go app that connect to a local postgres instance is in [database/ex-1-connection/solution](database/ex-1-connection/solution/postgres.go).

## Live Demo

Using tools to manage a database. Tools like goose and sqlc can be used to easliy abstract database management into your software stack. The following page is the resulting [code from the live demo](database/demo/main.go)

Recording to come

## Exercise 2

Build a client and interface around your database connection. You can use your existing main.go file and build new database package for your abstraction, or you can use the template files found in [database/ex-2-abstraction](database/ex-2-abstraction/main.go). Make sure to create an interface, a User struct, a database client, and the methods to create, update, and query the User data. If you have any questions please put them in the chat.

Follow-up Questions:

* what are the differences between manually creating a database object vs generating one with sqlc?

## Exercise 3 mock database

Using your new database interface mock the database functions into your [tests from last week.](../restful-go/ex-4-tests/framework_test.go)

TODO: copy tests from cli-game. and edit the interface to be inside the test suite.
