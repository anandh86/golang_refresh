# Goroutines

Goroutines are functions that run in the background.

To start a goroutine, add the `go` keyword when calling a function.

```go
func main() {
	// Run in the background
	go Greet("World")
	
	Greet("Mike")
}

func Greet(who string) {
	fmt.Println("Hello,", who)
}
```

The first `Greet` function is executed *concurrently* to the `main` function. It means that `main` will continue to run, not waiting for it to finish.

If you run the snippet above, the result will most likely be:

```bash
Hello, Mike
```

This is because the application ends after the `main` function has finished, and the goroutine we started might not have a chance to execute.

To see both results, we can add `time.Sleep()`. It will block the current function for the given time.

```go
func main() {
	// Run in the background
	go Greet("World")
	
	Greet("Mike")
	
	time.Sleep(time.Millisecond)
}
```

Now, the results will likely be:

```bash
Hello, Mike
Hello, World
```

"Likely", because the results may vary, depending on which goroutine is executed first.

When working with goroutines, keep in mind they are *concurrent*.
You can't determine the order in which they are executed (at least without using more advanced constructs from the next exercises).

## Exercise

File: `15-concurrency/01-goroutines/main.go`

`SignUp` is a function that saves the user and sends them an email notification.
However, both actions take quite long to complete, and they're executed in a sequence.

Make `SaveUser` and `SendNotification` run concurrently by using the `go` keyword.
It will make `SignUp` run faster.