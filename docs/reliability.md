# Reliable Go Services

In the final part of building production-ready Go web services, we focus on **reliability** — making sure our APIs work under stress, are easy to debug, and can gracefully handle failures. We'll cover:

* Security fundamentals (AuthN & AuthZ)
* Adding middleware for observability and safety
* Metrics and monitoring tools in Go
* Practices for building reliable web systems

---

## Security for Web Services

In any web service, understanding who is using your API and what they can access is critical. This breaks down into two core concepts:

### Authentication (AuthN)

**Who are you?**  
This happens *before* authorization. Common techniques include:

- Basic Auth
- Bearer Token (JWT)
- OAuth 2.0
- API Key headers
- VPN-based identity for internal services

Example using Basic Auth:

```go
user, pass, ok := r.BasicAuth()
if !ok || user != "admin" || pass != "password" {
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	return
}
```

For more production-ready strategies, look at:
- [Go Guardian](https://github.com/shaj13/go-guardian)
- [OAuth2 lib](https://pkg.go.dev/golang.org/x/oauth2)
- [JWT middleware](https://github.com/golang-jwt/jwt)

### Authorization (AuthZ)

**What are you allowed to do?**  
This typically uses:
- Role-Based Access Control (RBAC)
- Claims from a JWT token
- OAuth 2.0 scopes

---

## Middleware in Go

Middleware is software that sits between the request and your endpoint logic. It’s where you can:

- Log requests and responses
- Enforce authentication
- Set timeouts and rate limits

Here’s a simple example of middleware that logs incoming requests:

```go
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)

	loggedMux := loggingMiddleware(mux)
	http.ListenAndServe(":8080", loggedMux)
}
```

To explore production-ready middleware, check out:
- [chi middleware](https://github.com/go-chi/chi#middleware)
- [gin middleware](https://github.com/gin-gonic/gin#middleware)
- [fiber middleware](https://docs.gofiber.io/api/middleware)

---

## Monitoring and Observability

You can't manage what you can't measure. Metrics allow you to answer key questions:

- Are users experiencing errors?
- Is latency increasing?
- What’s the memory or CPU usage?

### The Four Golden Signals

From SRE practices at Google, every service should measure:

1. **Latency** – how long it takes to respond
2. **Traffic** – how many requests it handles
3. **Errors** – how often requests fail
4. **Saturation** – how close it is to resource limits

### Go's Built-in Monitoring Tools

#### `net/http/pprof` – CPU, memory, goroutine analysis

```go
import _ "net/http/pprof"

func main() {
	go http.ListenAndServe(":6060", nil) // debug server
	log.Fatal(http.ListenAndServe(":8080", yourHandler))
}
```

Visit `http://localhost:6060/debug/pprof/` for a full profile.

#### `expvar` – Expose runtime metrics over HTTP

```go
import "expvar"

var requestCount = expvar.NewInt("request_count")

func helloHandler(w http.ResponseWriter, r *http.Request) {
	requestCount.Add(1)
	w.Write([]byte("Hello"))
}
```

Visit `http://localhost:8080/debug/vars` to view output.

#### `log` – Built-in logging tool

Use structured logs with context:

```go
log.Printf("method=%s path=%s status=%d", r.Method, r.URL.Path, http.StatusOK)
```

---

## Metrics with Prometheus

[Prometheus](https://prometheus.io) is the most popular open-source metrics tool. With a Go library like [`promhttp`](https://pkg.go.dev/github.com/prometheus/client_golang/prometheus/promhttp), you can expose metrics easily.

### Example: Basic Prometheus Setup

```go
import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var httpRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total HTTP requests",
	},
	[]string{"path"},
)

func main() {
	prometheus.MustRegister(httpRequests)

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		httpRequests.WithLabelValues("/hello").Inc()
		w.Write([]byte("Hello, Prometheus"))
	})

	http.ListenAndServe(":8080", nil)
}
```

Now you can point Prometheus at `/metrics` and start collecting stats.

---

## Alerting

Once metrics are collected, you should set up alerts for things like:

- High 500 error rates
- Spike in latency
- Low memory availability

Popular tools:
- [Grafana](https://grafana.com/)
- [Slack Alerts](https://cloud.google.com/blog/products/devops-sre/use-slack-and-webhooks-for-notifications)
- [PagerDuty](https://www.pagerduty.com/)
- [Prometheus AlertManager](https://prometheus.io/docs/alerting/latest/alertmanager/)

---

## Exercises

### [Exercise 1](/reliable-webservice-go/ex-1-auth/auth.go) Add auth to your server endpoints

Auth tooling is sometimes the first or last measure of security for your endpoints. There are various methods for adding auth to your server endpoint. When building production services the methods you choose for authentication and authorization will be determined by security professionals, but how you implement them is up to you as a developer.

Add an auth method to your server. You can use any method you like such as a middle-ware, a helper functions or, by manually adding the logic to a single function.

Here are some examples of how to add different kinds of auth in your apps. You can pick one to use as a reference for you code.

* [chi middleware](https://github.com/go-chi/chi/blob/master/middleware/basic_auth.go)
* [manual basic token](https://github.com/Soypete/golang-cli-game/blob/main/server/helpers.go#L36)
* [JWT](https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-Parse-Hmac)
* [go-JWT example package](https://pkg.go.dev/github.com/golang-jwt/jwt/v5)
* [Go-guardian](https://github.com/shaj13/go-guardian/tree/master/_examples)
* [Oauth twitter](https://github.com/forgeutah/tweet_automated_bot/blob/main/client/setup.go)
* [Oauth2 golang.com/x](https://github.com/Soypete/Meetup-chat-server/blob/main/twitch/auth.go)

### [Exercise 2](/reliable-webservice-go/ex-2-middleware/middleware.go) Add middleware to your go server

_Middleware_: _Middleware also refers to the software that separates two or more APIs and provides services such as rate-limiting, authentication, and logging._[wikipedia](https://en.wikipedia.org/wiki/Middleware) The implementation is typically “built-in” functions. In Go, this tends to be platform style tooling shared across the organization. It allows you to add complex functionality to your endpoints.

Using the same web frameworks you used for your web server or the go standard library, add a middleware function to your server. You can use middleware to add metrics, auth, profiling or custom logic to your programs. In this exercise add logging, retry, rate limiting or replace the auth from exercise 1 with a middleware.

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

Pprof is an incredible profiling tool. It is the only tool currently provided to in the standard library what will let you follow memory hot path.
If you plan on using pprof as part of your monitoring suit you will need to install [graphviz](https://graphviz.org/download/) first.

[Pprof YouTube video](https://youtu.be/KzivSSjnBls)

For more information check out this talk, [Pprof for beginners](https://www.youtube.com/watch?v=HjzJ5r2D8ZM)

### [Exercise 3](/reliable-webservice-go/ex-3-monitoring/monitoring.go) Add some monitoring endpoints to your server

Monitoring is often setup as part of the middleware for commonly used metrics like db calls and http status codes. Often there are other metrics that should be added to track specific business logic and functionality. [Expvars](https://pkg.go.dev/expvar) are provided by the go standard library as a method for exposing metrics to an endpoint where they can be read via a web browser or consumed by a tracking service.

[Prometheus](https://prometheus.io/docs/guides/go-application/) is a very common opensource solution for adding metrics to your web services. It adds metrics to end points that can be scraped into a prometheus instance.

_NOTE:_ In this exercise it is not intended to have a prometheus instance up and running, just to set up the endpoint where you can manually view the metrics.

Using Expvars and/or Prometheus SDK add some custom metrics.

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

	// add prometheus endpoint at /metrics. The above collectors will be shown
	// in the reverse order they are registered.
	r.Mount("/metrics", promhttp.HandlerFor(
		reg,
		promhttp.HandlerOpts{
			// Opt into OpenMetrics to support exemplars.
			EnableOpenMetrics: true,
		},
	))
```

#### Bonus exercise: Add Pprof

Add pprof to your service to see how it uses memory when handling API calls. Run pprof and see what insights are available to you.

First add the pprof driver to your app.

```go
import _ "net/http/pprof"
```

_*NOTE*: the "\_" means that the import is added globally as a backend system. This is common for servers, db drivers, etc_

Add a pprof server as its own goroutine in your main function.

```go
// run pprof
go func() {
	http.ListenAndServe("localhost:6060", nil)
}()
```

Install [graphviz](https://graphviz.org/download/) on your machine to get the visual insights.

_Mac:_

```bash
brew install graphviz
```

run pprof while your worker-pool is executing

```bash
go tool pprof -http=:18080 http://localhost:6060/debug/pprof/profile?seconds=30
```

In the default graph each node is a function that your program is running. Size and color indicate how much CPU and time each function is taking.

to access the command-line tool run

```bash
go tool pprof http://localhost:6060/debug/pprof/allocs
```

---

## Conclusion

Reliability isn’t just about keeping your server "up"—it’s about observability, security, and scalability. By using tools like `pprof`, `expvar`, Prometheus, and middleware, you can build web services that don’t just work—but are maintainable, measurable, and secure.

