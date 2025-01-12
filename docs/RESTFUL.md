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

We are mostly interested in the last use case, where we can create a RESTful API that interacts with a database since this is the most common use case for web servers. We call these CRUD operations, and they have specific HTTP methods that we use to interact with the server.

### HTTP Methods
* GET - Read data from the server
* POST - Create data on the server
* PUT - Update data on the server
* DELETE - Delete data from the server
