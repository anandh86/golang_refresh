# The error interface

You've been using one built-in interface for a long time, even if you weren't aware of it.

The `error` type is actually an interface defined as:

```go
type error interface {
	Error() string
}
```

## Exercise

File: `12-interfaces/02-error/main.go`

Implement the `error` interface for `UserNotFoundError`.

Make it return a message that says: `user not found: <id>`.
Use `fmt.Sprintf` or `fmt.Sprint` to format the message.