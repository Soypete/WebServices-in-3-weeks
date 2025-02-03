# Why Go is the Best Modern Programming Language for Web Services  
by: Miriah Peterson
---

## What is a Web Service?

Web Service is an overloaded term in the current software ecosystem. I am ok with that, but for argument sake let me explain why I think a webservice is more that just something that connects to a website.

A web service, as my colleague [Tod Hansmann](https://www.linkedin.com/in/thansmann/) aptly defines, is:

"any piece of software that makes and receives communication over a network protocol." 

This broad definition captures a variety of applications. Some of them we call [Sass services](https://en.wikipedia.org/wiki/Software_as_a_service), some we call [CRUD](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete) services, and others we use just for serving data onto websites like [Back end as a Service](https://www.cloudflare.com/learning/serverless/glossary/backend-as-a-service-baas/). To me the use case doesn't matter all that much. If you are build a piece of software from the ground up intented to handle request you are building a webservice. Whether you're handling [HTTP requests](https://developer.mozilla.org/en-US/docs/Web/HTTP/Overview), [gRPC communication](https://grpc.io/), or [WebSocket connections](https://en.wikipedia.org/wiki/WebSocket), that software is a web service and needs to be built with the right tools.

Go, in my humble opinion, is the best modern programming language, period. I basically use Go as my swiss army knife for all my programming needs, but when I am evaluating languages and tools for production grade, scalable, and reliable Web Services, even with my bias, Go is the best choice. Why is that? It starts with the language itself.

### The Power of Go’s Standard Library and Ecosystem

Go's standard library is a treasure trove for web service development. Unlike many languages that require third-party libraries in order to effectively build endpoints and make efficient network requests (yes I am looking at [Rust](https://www.rust-lang.org/) and [Python](https://www.python.org/)), Go provides built-in, high-performance tools to handle critical tasks. For instance:  

- The **`net` package** simplifies network communication.  
- The **`database/sql` package** streamlines database interactions.  
- Built-in **testing tools** support reproducibility, making it easier to write, run, and maintain reliable code.  

This comprehensive standard library means you can start building robust web services without searching for or vetting dependencies. If you have ever had to deal with dependency hell, you know how valuable this is. If you have ever had to worry about compliance and security, you know how valuable this is. If you have ever had to worry licensing and support, you know how valuable this is.

Go excels in supporting modern web service protocols. You can accomplish any protocol with just standard library. Its native support for RESTful APIs is particularly noteworthy. It does request parsing and writing very easily. The native `encode/json` package provides 1st class support for you're public-facing APIs or internal microservices, Go offers scalability and ease of use for both backend and frontend needs.


If you need to lean into speed or want to use a different protocol like HTTP, gRPC, WebSockets, or GraphQL, Go’s ecosystem delivers. The community have build great tools and libraries to help you build your services. Most of these libraries 
are built with great licenses and are well maintained. It is easy to find a library that fits your needs and if you can't find one, it is easy to build one. If you want to go the extra mile you can even replace JSON with [Protocol Buffers](https://developers.google.com/protocol-buffers) for faster and more efficient data serialization.

### Readable, Scalable, and Maintainable  

Go’s simplicity is intentional and by design. Its syntax is clean and minimalistic, making it easy to learn for newcomers and straightforward to debug for seasoned developers. This readability translates directly to scalability in organizations, as teams can onboard quickly and collaborate effectively. In their most recent blog post, the Go team lead recommited to _"stability, safety, and supporting software engineering and production at scale."_ -[Go Turns 15](https://go.dev/blog/15years). The maintainers are aware that Go is used for many services fundamental for the high software we use everyday.

Since Go is a compiled language, and in my opinion this gives a ton of benefits to reliability and scalability of software. First runtime errors are minimized. Section you can run compiler check in CI/CD pipelines in addition to meaningful unit and integration testing which leads to faster deployments and a smoother development lifecycle. 

Moreover, Go was created with cloud-native development in mind. Its lightweight binaries and cross-platform support mean your web services can run on any major operating system and thrive in containerized environments like Docker and Kubernetes.  

### Learn to Build Web Services with Go  

I while ago, I gave several 3 way workshops on Go for WebServices. I have turned them into a series of Youtube videos. You can watch them here:

- [Go for Web Services: Part 1](https://www.youtube.com/watch?v=V9SvDd3e1eM)
- [Go for Web Services: Part 2](https://www.youtube.com/watch?v=V9SvDd3e1eM)
- [Go for Web Services: Part 3](https://www.youtube.com/watch?v=V9SvDd3e1eM)


for Exercices please check out this [repo](https://github.com/Soypete/WebServices-in-3-weeks)

---

### Why Go Beats the Competition  

Currently, Python frameworks like Django/FastAPI or Java's Spring/Boot framework are popular for building web services, but they oversimplify the process of writing software making it easy to loose sight of programming fundamentals. Go is a modern language that prioritizes performance, reliability, and simplicity. It doesn't shy away from the need for intentional code design by enforcing strict typing and error handling. The `httptest` package allows for easy testting of failures cases and intentional creations of [API contracts](https://www.geeksforgeeks.org/api-contracts-system-design/). This is why Go is the best choice, it is quick and easy to build with so we can focus on the important things like experince and predictability with tests, contacts, and solid design.

### Secure by Design

Many webservice developers get caught in their frameworks like [Java/Spring Boot](Spring Boot) or [Django](https://www.djangoproject.com/), and leave the security of their services to a diffent part of the organization. I have worked at a lot of small companies where that could not be left to someone else. It is a fundamental part of building any kind of software and Go has you covered.

Security is non-negotiable for Web Services, and Go delivers here as well. Their [Fuzzing test](https://en.wikipedia.org/wiki/Fuzzing) suite is a prime example of how Go prioritizes security and reliability. This allows you to create unpredictable inputs to your code, ensuring it can handle unexpected data gracefully and to prevent security vulnerabilities.

It is also easy to built authentication and authorization mechanisms, ensuring your services are secure by design from anything like oauth2 to SAML. I know that not everyone is a security expert, but Go makes it easy to live within the frameworks of security best practices and compliance without having to be an expert.

### Learn to Build Web Services with Go  

If you’re ready to explore Go’s capabilities, the following resources will help you get started. There is text and assignments to help you learn the basics of Go and build your first web service.

## Conclusion

Go is the best modern programming language for web services. Its standard library, protocol support, security features, readability, scalability, and maintainability make it the ideal choice for building robust, reliable, and secure web services. Whether you're building RESTful APIs, gRPC services, or WebSocket connections, Go has you covered. If you're new to Go, I encourage you to explore its capabilities and see for yourself why it's the best choice for web service development.

---

Resources:

