# Interfaces

Passing structs as arguments is useful, but sometimes you don't need all methods the struct exposes.
You might need to call only one or two methods, and they can be shared among different types. 

Consider a snippet that sends email notifications:

```go
type EmailSender struct {}

func (e EmailSender) SendNotification(user User, content string) error { 
	// ...
}

func SignUp(user User, sender EmailSender) error {
	// ...
	
	return sender.SendNotification(user, "Welcome to our website!")
}
```

If the user prefers text messages over email, we'd need to introduce another sender.

```go
type SMSSender struct {}

func (s SMSSender) SendNotification(user User, content string) error { 
	// ...
}
```

But notice that the `SignUp` function doesn't care how the messages are sent.
It trusts `SendNotification` to do the job. We just need a way to pass one of the senders to the function.

Interfaces let you specify what methods you need, without relying on the exact struct.

```go
type UserNotifier interface {
	SendNotification(user User, content string) error	
}

func SignUp(user User, notifier UserNotifier) error {
	// ... 
	
	return notifier.SendNotification(user, "Welcome to our website!")
}
```

You can now pass both `EmailSender` and `SMSSender` to the `SignUp` function, depending on what you need.

## Exercise

File: `12-interfaces/01-interfaces/main.go`

Introduce an interface and modify the `NewUser` function, so that both `MapStorage` and `SliceStorage` can be passed to it as `storage`.