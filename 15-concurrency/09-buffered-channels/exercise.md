# Buffered Channels

By default, channels don't "hold" any values, they just pass it from the sender to the receiver.
Reading and writing blocks until the other side is ready.

You can add a buffer to any channel to make it keep values until they're consumed.

```go
usernames := make(chan string, 10)
```

Writing to a buffered channel doesn't block as long as there's still free space in the buffer.
Once the buffer is full, writing blocks again.

Reading works the same as with unbuffered channels. It blocks until there's any value in the buffer.

## Exercise

File: `15-concurrency/09-buffered-channels/main.go`

The file shows an over-engineered payments aggregation pipeline.

`SendPayments` sends chosen amounts on the channel that `Aggregate` picks up. 
`Aggregate` sums up the amounts and replies on another channel. However, it is quite slow.

`SendPayments` has a 25-millisecond timeout on sending *all* payments.
It worked good enough until we started receiving more than five payments at once. Now it times out.

Introduce a buffer to a channel so that all payments fit it, `SendPayments` doesn't fail, and `Aggregate` can take whatever time it needs.