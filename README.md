# Golang Webservices in 3 weeks
This repo contains the examples and exercises for The O'reilly online learning course[Go Web Development in 3 weeks](https://www.oreilly.com/live-events/go-for-web-development-in-3-weeks/0636920091015/).

[![wakatime](https://wakatime.com/badge/user/953eeb5a-d347-44af-9d8b-a5b8a918cecf/project/815add1c-01f3-412e-b6cd-730805338e0e.svg)](https://wakatime.com/badge/user/953eeb5a-d347-44af-9d8b-a5b8a918cecf/project/815add1c-01f3-412e-b6cd-730805338e0e)
---

In this course you will learn all the steps to build a webservice in [Go](https://go.dev/). From starting the service to monitoring your service sit is meant to give you a comprehensive guide for building production level service. The first section of the course will handle building restful bestpractices in go. Connumication is key for designing  and bulding yoru webservices and is th efoundation on qhich your functionality will be built. The second section is all about data bases. Each webservice needs a layer to store, fetch and manipulate the data communicated with it. WE need to make sure our data foudnations are strong so we maintain the state of our services. The last section is an overview of reliability. This section is jsut goes over rilaiabity basics, but they are vital things that avery engineer should include when building a web service. This course does not go over [Go](https://go.dev/) basics. 

## pre-requisites
- [Go](https://go.dev/) installed and running
- Working knowledge of go

## New to go?

If you are new to go, work through these exercises first
- [Golang Zero to Hero](https://github.com/Soypete/Golang_tutorial_zero_to_hero)
- [Tour of Go](https://go.dev/tour/welcome/1)
- [Gophercises](https://gophercises.com/)

---

# Day 1 - Rest API protocols

* [Exercise 1](restful-go/README.md): std lib listenAndServe
* [Quiz](http-quiz/): Status Codes
* [Exercise 2](restful-go/README.md): Using a web framework
* [Exercise 3](restful-go/README.md): Client to query hosted server 
* [Exercise 4](restful-go/README/md): HttpTests
 
# Day 2 - Databases for webservices

* [Exercise 1](database/README.md): Connect to a live database
* [live coding](database/demo/): sqlc
* [Exercise 2](database/README.md): Add client and interface
* [Exercise 3](database/README.md): Unit tests with mock client
* [SQL quiz](sql-quiz)

# Day 3 - Metrics and Monitoring 

* [Exercise 1](reliable-webservice-go/README.md): API Auth
* [Exercise 2](reliable-webservice-go/README.md): Middleware
* [Exercise 3](reliable-webservice-go/README.md): add monitoring for endpoints

---

# Companion Service

This is a live production service that implements the game 20 questions. Code from this service is pulled out for the course exercises. The service code can be found [here](https://github.com/Soypete/golang-cli-game/).
---

## Explore More
- [Echo](https://echo.labstack.com/)
- [Chi](https://github.com/go-chi/chi)
- [Gin](https://github.com/gin-gonic/gin)
- [Fiber](https://github.com/gofiber/fiber)
- [Expvar](https://pkg.go.dev/expvar)
- [pgx](https://github.com/jackc/pgx)
- [pg](https://github.com/lib/pq)
- [sqlx](https://github.com/jmoiron/sqlx)
- [sqlc](https://sqlc.dev/)
