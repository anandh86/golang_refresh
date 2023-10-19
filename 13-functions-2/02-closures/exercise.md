# Closures

## Returning functions

If anonymous functions are values, they can also be returned from other functions.

```go
func NewGreeter() func() {
	return func() {
		fmt.Println("Hello!")
	}
}

func main() {
	greeter := NewGreeter()
	greeter()
}
```

## Closures

Closures are anonymous functions that use local variables outside their scope.

This is useful for keeping a variable that's shared across the function's calls.

```go
func NewCounter() func() int {
	counter := 0
	return func() int {
		counter += 1	
		return counter
	}	
}

func main() {
	counter := NewCounter()
	
	a := counter() // 1 
	b := counter() // 2 
	c := counter() // 3
}
```

## Exercise

File: `13-functions-2/02-closures/main.go`

Fill in the `WordGenerator` function.

The function takes a slice of words and returns a function that generates the next word from the slice. 
After returning the last slice element, it should start from the beginning.

For example, if the `words` slice is `[]string{"A", "B", "C"}`, the 
function should return a function that returns `"A"`, `"B"`, `"C"`, `"A"`, `"B"`, (...) in a sequence.
