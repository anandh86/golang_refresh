# Channels

Goroutines don't know about each other.
The functions running as goroutines can't return values in the traditional way, as the caller doesn't wait for them to finish.

The standard way to communicate between goroutines is using **channels**. 

A channel is a type that lets you send and receive values of a specific type.
You need to initialize it with `make`, similarly to maps.

```go
var users chan string
users = make(chan string)

// Same thing
users := make(chan string)
```

To send and receive values, use the `<-` operator.
When sending, the channel goes on the left side. When receiving, on the right.

```go
// Send to the users channel
users <- "Alice"
users <- "Bob"

// Retrieve from the users channel
nextUser := <-users
```

The `chan` keyword is always followed by a type. If you pass channels as arguments, remember to specify it:

```go
func AddUser(users chan string) {
	users <- "Alice"
}
```

## Exercise

File: `15-concurrency/02-channels/main.go`

While channels can pass values between goroutines, they can also be used to signal when something happened.

`SendNewsletter` concurrently sends newsletters to all given emails.
It uses the `done` channel, so the `SendNewsletterToEmail` goroutine can signal when it's ready.

Implement this part. Send a `true` value to the `done` channel after the function is done processing.
