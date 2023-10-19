# Pointer receivers

Just as with other types, you can pass a pointer to a struct to modify it inside a function.

```go
func ChangeToZero(m *Money) {
	m.Amount = 0	
}

func main() {
	m := &Money{}
	ChangeToZero(m)
}
```

Similarly, methods allow changing structures in place, like this:

```go
func main() {
	m := &Money{}
	m.ChangeToZero()
}
```

To be able to do this, you need to add an asterisk (`*`) in the receiver type.
It lets you modify the struct fields inside methods.

```go
func (m *Money) ChangeToZero() {
	m.Amount = 0
}
```

## What to use?

As a rule of thumb:

* If your methods don't modify the struct, use plain struct receivers.
* If any of your methods modifies the struct, use pointer receivers everywhere for consistency.

You might also use pointer receivers if your struct is huge and copying it on each method call would be expensive.
You don't have to concern about it for now.

## Exercise

Add the missing `Move` method to the `Position` struct.

It should take two integer arguments, `dx` and `dy`. Add them to the `X` and `Y` fields.