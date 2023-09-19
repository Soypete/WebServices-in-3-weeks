# Day 1: Writing Restful Golang Servers and Client

These are the code-samples, exercises, quizzes, and links needed for this course. To complete these exercises follow the instructions in this doc. In the subdirectories of this section are example solutions. Your soluutions do not need to match the provided solutions. The goal of the exercises is to learn and understand the why things are implemented so that a participant apply principles to any framework or problem set.

## Exercises

### [Exercise 1](/ex-1-servers/server.go) Build a Standard Library Server

In a main function, create a webserver using the `ListenAndServe()` function. This server should have at least one endpoint that accepts a parameter and returns it in a response. An example of a finished exercise in in [ex-1-server](ex-1-servers/server.go)

* reference: [go-by-example Https server](https://gobyexample.com/http-servers)

\_ _NOTE_: `ListenAndServe()` is great for prototyping and non-production development. When you are building out your services, make sure you have security in mind. Use a custom `mux`. We will learn more aabout security best practices later on in the course. Check out [Reliable Webservices](../reliable-webservice-go/README.md).

## [Exercise 2](/ex-2-web-frameworks/framework.go) Build a Server Using an opensource Framework

The [Standard Library](https://pkg.go.dev/net/http) provides all the functionality needed to build a robust web server, but sometimes, instead of building out your own infrastructre, there is a benefit to adopting an opensource frameworks at the backbone of your server infrastructure. This exercise is you chance to experiment with some popular Go web frameworks.

_Instructions:_ Pick a web framework and implement a server that can accept a username, validate it, and print the username to standard out. Add logic to handle duplicate usernames, empty username parameters, or any additional error cases. When a username is accepted and printed to standard out return a 200 class status code. If the parameter is empty or invalid return a 400 class error code. Handle any other errors. _We will connect this server to a database in the Day 2 exercises_

Examples of the finished exercise:

* using the [Chi web framework](/ex-2-webframeworks/framework.go).
* using the [Fiber web framework](/ex-2-fiber/fiber.go).

_*Note:* you do not have to do this with a web framework. Feel free to use the standard library `net/http` tooling if you are more comfortable. The opstins for using a web framework are just to try you hand at evaluating popular opensource tools. But just like adding any tool to your software start there are risks and rewards._

Here are some examples of opensource frameworks:

* [Chi](https://github.com/go-chi/chi)
* [Gin](https://github.com/gin-gonic/gin) <!-- uses it own context that predates context.Context-->
* [Fiber](https://github.com/gofiber/fiber) <!-- uses fasthhtp -->

### [Exercise 3](/ex-3-clients/client.go) Build a Go app that calls your new endpoint.

Using the standard library to create a client that calls your new web endpoint. To connect to you server, you will need to run your server in one different and the client app in a different terminal window.

### [Exercise 4](/ex-4-test/framework_tests.go)

Using the [httptest](https://pkg.go.dev/net/http/httptest#example-ResponseRecorder) package add some tests for at least one of our newly created endpoints. You should have one test that checks for a 200s class error and one that tests for a 400s class error. Add additional test validations as time allows.

Here are some examples of tests for the above opensource frameworks.

* [Chi tests](https://go-chi.io/#/pages/testing)
* [Gin tests](https://gin-gonic.com/docs/testing/)
* [Fiber tests](https://docs.gofiber.io/api/app#test)

## Helpful links:

We will be using `curl` to test the functionality of our endpoints and servers. If you are unfamiliar with the `curl` syntax check out these resources:

* [Free code camp](https://www.freecodecamp.org/news/how-to-start-using-curl-and-why-a-hands-on-introduction-ea1c913caaaa/)
* [Curl docs](https://curl.se/docs/manual.html)
