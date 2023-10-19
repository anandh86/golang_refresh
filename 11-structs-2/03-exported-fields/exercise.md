# Exported fields

Only fields starting with an uppercase letter can be used outside the package they're in. We call them *exported*.

Fields starting with a lowercase letter (*unexported*) aren't visible for other packages.

We'll dive into writing your own packages soon, but for now consider the `json.Marshal` function you've used.
The `encoding/json` package won't be able to access unexported fields on any struct you pass to its functions.

```go
type Payload struct {
	Source          string // exported
	internalDetails string // unexported
}
```

## Exercise

File: `11-structs-2/03-exported-fields/main.go`

Change the `Account` struct so the `Password` isn't included after marshaling to JSON.