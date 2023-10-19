# Functions as Values

Functions can be values, too.

A function stored as a value is called an *anonymous function*.

You can use them in function arguments like any other type. 

```go
func main() {
	add := func(x int, y int) int {
		return x + y
	}
	
	multiply := func(x int, y int) int {
		return x * y
	}
	
	Calculate(add)
	Calculate(multiply)
}

func Calculate(f func(int, int) int) {
	result := f(10, 50)
	fmt.Println(result)
}
```

## Exercise

File: `13-functions-2/01-anonymous-functions/main.go`

Add the missing `uppercase` anonymous function in `main`.

The function should return the uppercase version of the given string.
You can use [`strings.ToUpper()`](https://pkg.go.dev/strings#ToUpper) to do this.