# Day 2 Databases and your Go app

This section of the course contains examples and exercises for effectively leveraging a database in the backend of your Web services.. The following are examples of finished exercises for day 2 of the [O'reilly Media Online Learning Course](https://www.oreilly.com/live-events/go-for-web-development-in-3-weeks/0636920091015/). They can be completed by following along with the instruction or independently.

## Exercise 1

Use the appropriate database driver and go module to connect to the provided postgres database. The username and password are provided by the instructor (See below if completing independently). Once you have connected run a query based on the following schema...

TODO: include an ERD or other schema docs.

_NOTE_: If you are completeling this indepently. You can run postgres locally in a docker container using the following command.

```
docker pull postgres
docker run -e POSTGRES_PASSWORD=my_password -e POSTGRES_USERNAME=user1 -p 5431:5432 postgres
```

## Live Demo

Using tools to manage a database. Tools like goose and sqlc can be used to easliy abstract database management into your software stack. The following page is the resulting [code from the live demo](database/demo/main.go)

Recording to come

## Exercise 2

Build a client and interface around your database connection.

Questions:

* what are the differences between manually creating a database object vs generating one with sqlc?

## Exercise 3 mock database

Using your new database interface mock the database functions into your [tests from last week.](../restful-go/ex-4-tests/framework_test.go)

TODO: copy tests from cli-game. and edit the interface to be inside the test suite.
