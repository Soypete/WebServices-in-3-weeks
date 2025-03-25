Here’s a fully structured and styled tutorial on **Building Reliable Web Services in Go**, based on Day 3 from your PDF and using the same tone and format as your "RESTful Go" example:

---

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

Try these exercises to solidify your understanding:

1. **Add Logging Middleware**  
   Wrap your existing routes with a logger that prints HTTP method and path.

2. **Use `pprof` to inspect CPU usage**  
   Add `net/http/pprof`, hammer your server with requests, and view flame graphs.

3. **Add Prometheus HTTP request counters**  
   Count requests per endpoint using a `CounterVec`.

4. **Secure an endpoint**  
   Add basic auth to a `/secure` route and verify that only authenticated users can access it.

5. **Set up a `/metrics` endpoint**  
   Use Prometheus to collect and view metrics.

> [Exercise repo](https://github.com/Soypete/WebServices-in-3-weeks/tree/main/reliable-webservice-go)

---

## Conclusion

Reliability isn’t just about keeping your server "up"—it’s about observability, security, and scalability. By using tools like `pprof`, `expvar`, Prometheus, and middleware, you can build web services that don’t just work—but are maintainable, measurable, and secure.

---

Would you like this pushed as a Markdown doc, or posted as the third article on your Substack/blog?
