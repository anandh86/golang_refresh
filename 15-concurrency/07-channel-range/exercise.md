# Channel range

As with slices and maps, you can use `for` and `range` to go over values in a channel.

```go
luckyNumbers := make(chan int)
for number := range luckyNumbers {
	fmt.Println("Your next lucky number is", number)	
}
```

Such loop blocks until there's a value to be received from the channel.

You can use `break` and `continue` like in other `for` loops.

### Closing channels

A channel can be *closed* using `close`.

```go
close(luckyNumbers)
```

Closing a channel interrupts any loops waiting for more messages.

Reading from a closed channel returns immediately, returning the zero-value of the channel's type.

```go
// Always zero after the channel is closed
number := <-luckyNumbers
```

You can also check if a channel is closed using the two-value syntax:

```go
number, ok := <-luckyNumbers
if ok {
	fmt.Println("Your lucky number is", number)	
} else {
	fmt.Println("Out of lucky numbers")
}
```

Don't write to closed channels — it panics.

### Channel directions

When using channels as function arguments, you can decide if the channel should allow reads or writes.

```go
func readAndWrite(ch chan string) {
	
}

func readOnly(ch <-chan string) {
	
}

func writeOnly(ch chan<- string) {
	
}
```

This has impact on the channel only within the function.

## Exercise

File: `15-concurrency/07-channel-range/main.go`

Fill in the `ReadMessages` function so that it reads all messages from the given channel until it's closed. 
It should return all captured messages.

Update `SendMessages` so that it closes the channel after it's done writing to it.