# Restful Go 

In order to build a foundational web server, we need to understand some basics. Here we will cover the following:

* What is RESTful?
* Golang RESTful basics
* How to build Golang servers and clients
* Golang testing tools for APIs


## What is RESTful?

An API is defined this way:

>An application programming interface (API) defines the rules that you must follow to communicate with other software systems - [source](https://aws.amazon.com/what-is/restful-api/)

RESTful is a set of rules that are used for create APIs that communicate over the internet between different systems.
When you are building a RESTful API, you need to follow these rules:

* The client sends a request to the server. The client follows the API documentation to format the request in a way that the server understands.
* The server authenticates the client and confirms that the client has the right to make that request.
* The server receives the request and processes it internally.
* The server returns a response to the client. The response contains information that tells the client whether the request was successful. The response also includes any information that the client requested.

There are many use cases for RESTful APIs, such as:
* Public APIs where you can access data from a third-party service.
* Public APIs where you can update data in a third-party service.
* Internal APIs where you can communicate between different services in your organization.
* Create, Update, Delete, and Read (CRUD) APIs that interact with a webservice on a database.

We are mostly interested in the last use case, where we can create a RESTful API that interacts with a database since this is the most common use case for web servers. When we can to interact with a database we create a [Request]() to contain all of the necessary information. The Requests can be called CRUD operations if they correstond to actions in a database, and they have specific HTTP methods that we use to interact with the server.

### HTTP Methods
* GET - Read data from the server
* POST - Create data on the server
* PUT - Update data on the server
* DELETE - Delete data from the server

Requests are used to communicate between Clients and Server. In addition to the HTTP methods, we also have the following components of a Request:

* Header- This stores information like Roles, Identity, and Authentication
* Payload- This stores any information needed be manipulated by or stored by the service
* URL argument- This is qualifying information that is used to make decisions
* Query Parameter - This is filtering information that can be used by service or stored

I am a data practitioner by trade, so when I think about making API calls, I tend to focus on different types of data that are being to compose the Request. Each of the previous components can be used to store different types of data. For example, the Header can store information about the user, the Payload can store information about the data being manipulated, the URL argument can store information about the service, and the Query Parameter can store information about the data being filtered. Each has a specific use case and so it is important to understand how and when to use each f of these components and to make sure that you are using them correctly.

An example of this is the Header. The Header is used to store information about the user, such as the user's identity, roles, and authentication. This information is used by the server to determine whether the user has the right to make the request. If the user does not have the right to make the request, the server will return an error message to the client. The Header is also used to store information about the request itself, such as the content type of the request and the content type of the response. This information is used by the server to determine how to process the request and how to format the response. 

### curl

[Curl]() is a command-line tool that is used to make HTTP requests. It is a very powerful tool that can be used to test RESTful APIs. Here is an example of how to use curl to make a GET request to a RESTful API:

```bash
curl -X GET https://cat-fact.herokuapp.com/facts
```     

This command will make a GET request to the cat-fact API and return a list of facts about cats. You can use curl to make other types of requests as well, such as POST, PUT, and DELETE requests. I always recommend testing apis with curl before you start writing code. This will help you understand how the API works and what kind of data you can expect to receive from the server.

_NOTE: For extra practice download the [jq]() tool and use it to parse the JSON response from the cat-fact API. This will help you understand how to work with JSON data in your code._ 

```bash
curl -X GET https://cat-fact.herokuapp.com/facts | jq
```

### Postman

[Postman]) is a popular tool for testing Restful APIs. The strength of Postman is saving and sharing requests. You can create environments and collections to organize your requests. You can also use Postman to create tests for your APIs. This is a great way to automate testing and make sure that your APIs are working correctly. Postman is a great tool for testing APIs. If you are not as comfortable with the command line, I recommend using Postman to test your APIs.


## Golang RESTful basics

To start our journey into understanding how to build webservers with Go, we will start with the standard library.

### Writing a simple server

All the tools needed to create and maintain a RESTful API server can be found in the Standard Library. It supports a variety of protocols including TCP, UDP, HTTP/1.0, HTTP/2.0. The Standard Library allows you to create a variety of endpoints, custom middleware, and contexts like timeouts. 

Here is an example of a simple server that listens on port 8080 and returns "Hello, World!" when you make a GET request to the root endpoint:

```go
func main() {
	helloHandler := func(
		w http.ResponseWriter,
		req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

```

This curl command will make a GET request to the server and return "Hello, World!":

```bash
curl -X GET http://localhost:8080/hello
```

Go by default uses multiplexing to handle multiple requests at the same time. This means the classic `http.ListenAndServe` is a Mux server.

_MUX is a session management protocol separating the underlying transport from the upper level application protocols. It provides a lightweight communication channel to the application layer by multiplexing data streams on top of a reliable stream oriented transport._  [source](https://www.w3.org/Protocols/MUX/)

For each endpoint you define on a server you will need to create a handler function. The handler function takes two arguments, a `http.ResponseWriter` and a `*http.Request`. The `http.ResponseWriter` is used to write the response to the client. The `*http.Request` is used to read the request from the client. You can use the `http.ResponseWriter` to write the response to the client. You can use the `*http.Request` to read the request from the client. If you want additional data to be provided to the handler functions you should make sure that the functions are method on a struct that contains the data you need, or you can use a closure to pass the data to the handler function.

#### Status Codes

Go has a built-in package called `http` that contains a list of status codes that you can use to return to the client. Since these are constant you can use them in you code to make sure that you are returning the correct status code to the client. I highly recommend using the harcoded constants for status codes because they are self-documenting and make your code easier to read.

```go
func main() {
    helloHandler := func(
        w http.ResponseWriter,
        req *http.Request) {
        w.WriteHeader(http.StatusOK)
        io.WriteString(w, "Hello, world!\n")
    }

    http.HandleFunc("/hello", helloHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

This is a good point to stop and test your skills. here is an exercise for you to try:

## Using a Web Framework

The standard library contains all the tools you need to create a RESTful API server, but sometimes for team conseensuus and speed you make choose a webframe work that abstracts away a little but of the boilerplate code. There are many web frameworks available for Go, such as [Echo](https://echo.labstack.com/), [Chi](https://github.com/go-chi/chi), [Gin](https://github.com/gin-gonic/gin), and [Fiber](https://github.com/gofiber/fiber). Each of these frameworks has its own strengths and weaknesses, so you should choose the one that best fits your needs.

When you are choosing a web framework, you should consider the following factors:
 - Adding complex middleware
 - Adding error handling
 - Add url parsing tools

 _*NOTE*: url parsining for ids was added in Go (1.22)[https://go.dev/blog/routing-enhancements] so you may not need a web framework for this feature, but it is still a consideration because every web framework has its own way of handling url parsing._

Echo on the newest and in my opinons is the most perscriptive. It has a lot of boilerplate and it think it is hard to get used to. Gin is the oldest and post popular by stars(and a good friend of mine is an original contributor), but it was written pre Go context package so it has it's own concept of context. Chi is very lightweight and has a lot of features that are not in the standard library, I thought it was the easiest to implement and chose it for my own project. Fiber was the one most recommeneded by my twitch chat. They say it is the fastest and it also is very lightweight.

Here is an example in Chi:

Here is an example in Fiber:
