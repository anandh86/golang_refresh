# Select

When you need to read from multiple channels at the same time, use `select`.

`select` is similar to `switch` but it works with channels. It blocks until one of the `case` statements can be executed.

## One of a few

When reading from multiple channels, `select` executes the `case` that has values ready to be received.
If more than one `case` can be executed, `select` picks one randomly.

```go
select {
case command := <-commands:
	// Execute a command
case err := <-errors:
	// Handle an error 
}
```

A common use-case is using `time.After()` to introduce a timeout.

```go
select {
case command := <-commands:
	// Execute a command
case <-time.After(time.Second):
	// Timed out waiting for a command
}
```

## Non-blocking wait

A `default` case is executed if no other actions are possible.

An empty `default` case is a way to continue immediately if there are no values to be received from a channel.

```go
ready := make(chan bool)

// Blocks until ready
<-ready

// Checks if ready and continues otherwise
select {
case <-ready:
	fmt.Println("ready")
default:
	fmt.Println("not ready yet")
}
```

## Looping over select

`select` is often useful in combination with `for`.

For example, the snippet below infinitely checks until a value can be received from a channel, with 1-second sleep between checks.

```go
func WaitUntilReady(ready chan bool) {
	for {
		select {
		case <-ready:
			fmt.Println("ready")
			return 
		default:
			fmt.Println("not ready yet")
			time.Sleep(time.Second)
		}
	}
}
```

Warning: using `break` in a select statement interrupts the `select`, not the `for`.
In this scenario, it's best to use a separate function and a `return` instead.

## An empty select

An empty `select` waits forever. Similarly to an empty `for`, but is much easier on the CPU.

```go
select {}
```

## Exercise

File: `15-concurrency/08-select/main.go`

`RunWithTimeout` runs the provided function and returns a timeout error if it runs over 1 second.

Fill in the missing code using `select`:

* If an `error` can be received from the channel, return it.
* If it doesn't happen after 1 second, return a new error with a `timeout` message.