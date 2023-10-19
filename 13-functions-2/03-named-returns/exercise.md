# Named returns

Go allows naming the return values.

```go
func Add(x int, y int) (sum int) {
	sum = x + y	
	return
}
```

They act as normal local variables.
You still need to write the empty `return` statement to exit from the function.

Since they're already declared, you use `=`, not `:=` to assign values to them.

## Exercise

File: `13-functions-2/03-named-returns/main.go`

Update the `DirectionFromNumber` function to use a named return value called `direction`.
Add the missing `return` statement.