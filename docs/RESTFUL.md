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

## Golang RESTful basics


