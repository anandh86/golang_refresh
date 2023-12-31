# Defer

The `defer` keyword makes a function execute at the end of the current function.

Think of it like of a cleanup.

```go
func Execute() {
	defer cleanup()
	
	// ...
}

func cleanup() { 
	// ...
}
```

## Accessing named returns

If you use named returns, you can access them in `defer` after the function returns.

```go
func Execute() (err error) {
	defer func() {
		if err == nil {
			fmt.Println("Completed successfully")
		} else {
			fmt.Println("Failed:", err)
		}
	}()
	
	// ...
}
```


<div class="alert alert-dismissible bg-light-primary d-flex flex-column flex-sm-row p-7 mb-10">
    <div class="d-flex flex-column">
        <h3 class="mb-5 text-dark">
			<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-lightbulb text-primary" viewBox="0 0 16 16">
			  <path d="M2 6a6 6 0 1 1 10.174 4.31c-.203.196-.359.4-.453.619l-.762 1.769A.5.5 0 0 1 10.5 13a.5.5 0 0 1 0 1 .5.5 0 0 1 0 1l-.224.447a1 1 0 0 1-.894.553H6.618a1 1 0 0 1-.894-.553L5.5 15a.5.5 0 0 1 0-1 .5.5 0 0 1 0-1 .5.5 0 0 1-.46-.302l-.761-1.77a1.964 1.964 0 0 0-.453-.618A5.984 5.984 0 0 1 2 6zm6-5a5 5 0 0 0-3.479 8.592c.263.254.514.564.676.941L5.83 12h4.342l.632-1.467c.162-.377.413-.687.676-.941A5 5 0 0 0 8 1z"/>
			</svg>
			Tip
		</h3>
        <span>

Notice the brackets (`()`) after the anonymous function's body.

</span>
	</div>
	</div>

## Order

You can call `defer` many times. The last deferred function is the first that's going to be called.

```go
func Execute() {
	defer third()
	defer second()
	defer first()
}
```

## Exercise

File: `13-functions-2/04-defer/main.go`

Fill in the `Execute` function.

The function should execute the given `f` function, and call proper methods on `metrics`:

* Call `StoreExecution()` before calling the function.
* Call `StoreSuccess()` if the function `f` returns without an error.
* Call `StoreFailure()` if the function `f` returns with an error.

Use `defer` to handle the `err` return value.
