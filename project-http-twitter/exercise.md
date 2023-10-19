# The "Twitter" Project: Rate Limiting

*This exercise continues your previous HTTP project. You will work on the same files.*

Let's now protect our HTTP server from evil hackers (like us).

### Middlewares
 
Middlewares are HTTP handlers executed for all incoming requests in a sequence.
You've already used a `Logger` middleware before that logs a request details.

Chi comes with many [middlewares that are ready to use](https://github.com/go-chi/chi/tree/master/middleware).
You can also write your own.

To apply the middleware to one endpoint only you need to call the `With` method:

```go
r.With(authMiddleware).Post("/comments", addComment)
```

You can also enable the middleware for a group of endpoints:

```go
r.With(authMiddleware).Group(func(r chi.Router) {
	r.Post("/comments", addComment)
	r.Post("/follows", addFollow)
})
```

## Exercise

File: `project-http-twitter/main.go`

* Use the [`go-chi/httprate`](https://github.com/go-chi/httprate) library and add a rate limiting middleware to your router.
* Add limits per IP using `httprate.LimitByIP`.
* Allow adding up to 10 tweets per minute.
* Limit requests only for the POST endpoint. We don't care about limiting reads for now.
