# Day 3: Building Reliable Web Services

These are the code-samples, exercises, quizzes, and links needed for this course. To complete these exercises follow the instructions in this doc. In the subdirectories of this section are example solutions. Your solutions do not need to match the provided solutions. The goal of the exercises is to learn and understand the why things are implemented so that a participant apply principles to any framework or problem set.

## Exercises

### [Exercise 1](/ex-1-auth/auth.go) Add auth to your server endpoints

Auth tooling is sometimes the first or last measure of security for your endpoints. There are various methods for adding auth to your server endpoint. When building production services the methods you choose for authentiction and authorization will be determined by security professionals, but how you implement them is up to you as a developer.

Add an auth method to your server. You can use any method you like such as a middle-ware, a helper functions or, by manually adding the logic to a single function.

Here are some examples of how to add different kinds of auth in your apps. You can pick one to use as a reference for you code.

* [chi middleware](https://github.com/go-chi/chi/blob/master/middleware/basic_auth.go)
* [manual basic token](https://github.com/Soypete/golang-cli-game/blob/main/server/helpers.go#L36)
* [JWT](https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-Parse-Hmac)
* [go-JWT example package]()
* [Go-guardian]
* [Oauth twitter]
* [Oauth2 golang.com/x]

### [Exercise 2](/ex-2-middleware/middleware.go) Add middleware to your go server

_Middleware_: _Middleware also refers to the software that separates two or more APIs and provides services such as rate-limiting, authentication, and logging._[wikipedia](https://en.wikipedia.org/wiki/Middleware) The implementation is typically “built-in” functions. In Go, this tends to be platform style tooling shared across the organization. It allows you add complex funcitonality to your endpoints.

Using the same web frameworks you used for your webserver or the go standard library, add a middleware function to your server. You can use middleware to add metrics, auth, profiling or custom logic to your programs. In this exercise add logging, retry, rate limiting or replace the auth from exercise 1 with a middleware.

Below are framework docs, they will contain examples of build in middleware that you can add with single line functions. They also show you ways of adding custom middleware to your services.

Web frameworks:

* [Chi](https://github.com/go-chi/chi)
* [Gin](https://github.com/gin-gonic/gin) <!-- uses it own context that predates context.Context-->
* [Fiber](https://github.com/gofiber/fiber) <!-- uses fasthhtp -->

Here is an [example](https://github.com/Soypete/golang-cli-game/blob/main/server/setup.go) of setting it up using the [chi](https://pkg.go.dev/github.com/go-chi/chi) framework

```go
	r := chi.NewRouter()

	// add prebuilt middleware for all requests
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
```

### Pprof live Demo

Pprof is an incredible builting profiling tool. It is the only tool currently provided to in the standard library what will let you follow memory hot path.
If you plad on Using pprof are part of your monitoring suit you will need to install [graphviz](https://graphviz.org/download/) first.

[Pprof youtube video](https://youtu.be/KzivSSjnBls)

For more information check out this talk, [Pprof for beginners](https://www.youtube.com/watch?v=HjzJ5r2D8ZM)

### [Exercise 3](/ex-3-monitoring/monitoring.go) Add some monitoring endpoints to your server

Monitoring is often setup as part of the middleware for commonly used metrics like db calls and http status codes. Often there are other metrics that should be added to track specific business logic and functionality. [Expvars](https://pkg.go.dev/expvar) are provided by the go standard library as a method for exposing metrics to an endpoint where they can be read via a web browser or consumed by a tracking service.

[Prometheus](https://prometheus.io/docs/guides/go-application/) is a very common opensource solution for adding metrics to your web services. It adds metrics to end points that can be scraped into a prometheus instance.

_NOTE:_ In this exercise it is not intended to have a prometheus instance up and running, just to setup the endpoint where you can manually view the metrics.

Using Expvars and/or Prometheus sdk add some custom metrics.

[Example](https://github.com/Soypete/golang-cli-game/blob/main/server/setup.go#L53)

```go
    reg := prometheus.NewRegistry()
	reg.MustRegister(collectors.NewBuildInfoCollector())
	reg.MustRegister(collectors.NewDBStatsCollector(db.GetSqlDB(), "postgres"))
	reg.MustRegister(collectors.NewExpvarCollector(
		map[string]*prometheus.Desc{
			"counter200Code": prometheus.NewDesc("expvar_200Status", "number of status 200 api calls", nil, nil),
			"counter400Code": prometheus.NewDesc("expvar_400status", "number of status 400 api calls", nil, nil),
			"counter500Code": prometheus.NewDesc("expvar_500status", "number of status 500 api calls", nil, nil),
		},
	))

	// add prometheus endpoint at /metrics. The above collectors will be show
	// in the reverse order the are registered.
	r.Mount("/metrics", promhttp.HandlerFor(
		reg,
		promhttp.HandlerOpts{
			// Opt into OpenMetrics to support exemplars.
			EnableOpenMetrics: true,
		},
	))
```

#### Bonus exercise: Add Pprof

Add pprof to your service to see how it uses memory when handling api calls. Run pprof and see what insights are available to you.

First add the pprof driver to your app.

```go
import _ "net/http/pprof"
```

_*NOTE*: the "\_" means that the import is added globally as a backend system. This is common for servers, db drivers, etc_

add a pprof server as it's own goroutine in your main function.

```go
// run pprof
go func() {
	http.ListenAndServe("localhost:6060", nil)
}()
```

install [graphviz](https://graphviz.org/download/) on your machine to get the visual insights.

_Mac:_

```bash
brew install graphviz
```

run pprof while your worker-pool is executing

```bash
go tool pprof -http=:18080 http://localhost:6060/debug/pprof/profile?seconds=30
```

In the default graph each node is a function that your program is running. Size and color indicate how much cpu and time each function is taking.

to access the commandline tool tool run

```bash
go tool pprof http://localhost:6060/debug/pprof/allocs
```
