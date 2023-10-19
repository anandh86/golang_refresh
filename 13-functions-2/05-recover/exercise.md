# Recover

A way to keep `panic` from stopping your application is to use `recover`.
You have to call it inside `defer`, otherwise it won't work.

`recover` prevents panics only in the function where it's called.
It won't prevent panics happening in other functions or goroutines.

```go
func Example() {
	defer func() {
		r := recover()	
		if r != nil {
			fmt.Println("Recovered panic:", r)
		}
	}()
	
	panic("boom")
}
```

You shouldn't overuse `recover`. It's useful for handling specific scenarios, like a single
goroutine panicking and stopping your program. It doesn't mean you should use `panic` over
returning errors.

## Exercise

File: `13-functions-2/05-recover/main.go`

Fill in the `RunSafely` function.

It should execute the given function `f`, prevent any panics, and return an error if a panic happened.


<div class="accordion" id="hints-accordion">

<div class="accordion-item">
	<h3 class="accordion-header" id="hints-accordion-header-1">
	<button class="accordion-button fs-4 fw-semibold collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#hints-accordion-body-1" aria-expanded="false" aria-controls="hints-accordion">
		Hint #1
	</button>
	</h3>
	<div id="hints-accordion-body-1" class="accordion-collapse collapse" aria-labelledby="hints-accordion-header-1" data-bs-parent="#hints-accordion">
	<div class="accordion-body">

`recover()` needs to be called **after** executing the `f` function. Use `defer` to guarantee this.

</div>
	</div>
	</div>

<div class="accordion-item">
	<h3 class="accordion-header" id="hints-accordion-header-2">
	<button class="accordion-button fs-4 fw-semibold collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#hints-accordion-body-2" aria-expanded="false" aria-controls="hints-accordion">
		Hint #2
	</button>
	</h3>
	<div id="hints-accordion-body-2" class="accordion-collapse collapse" aria-labelledby="hints-accordion-header-2" data-bs-parent="#hints-accordion">
	<div class="accordion-body">

Make sure you call `f` outside the `defer` block.

</div>
	</div>
	</div>

</div>
