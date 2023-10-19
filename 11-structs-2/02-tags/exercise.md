# Tags

A common use case for structs is transforming them into text (called *marshaling* or *serializing*).
For example, using the JSON format, which Go provides with the `encoding/json` package.

```go
type User struct {
	Name        string
	YearJoined  int
	Premium     bool
}

alice := User{
	Name:       "Alice", 
	YearJoined: 2022, 
	Premium:    true,
}

marshaledUser, err := json.Marshal(alice)
```

The output of marshaling this struct would be:

```json
{
  "Name": "Alice",
  "YearJoined": 2022,
  "Premium": true
}
```

By default, `encoding/json` bases key names on the struct fields. You can use *tags* to control this.

Tags are string literals used as metadata on fields. Define them within backticks (\`) so you can use the standard double-quotes inside.


<div class="alert alert-dismissible bg-light-primary d-flex flex-column flex-sm-row p-7 mb-10">
    <div class="d-flex flex-column">
        <h3 class="mb-5 text-dark">
			<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-lightbulb text-primary" viewBox="0 0 16 16">
			  <path d="M2 6a6 6 0 1 1 10.174 4.31c-.203.196-.359.4-.453.619l-.762 1.769A.5.5 0 0 1 10.5 13a.5.5 0 0 1 0 1 .5.5 0 0 1 0 1l-.224.447a1 1 0 0 1-.894.553H6.618a1 1 0 0 1-.894-.553L5.5 15a.5.5 0 0 1 0-1 .5.5 0 0 1 0-1 .5.5 0 0 1-.46-.302l-.761-1.77a1.964 1.964 0 0 0-.453-.618A5.984 5.984 0 0 1 2 6zm6-5a5 5 0 0 0-3.479 8.592c.263.254.514.564.676.941L5.83 12h4.342l.632-1.467c.162-.377.413-.687.676-.941A5 5 0 0 0 8 1z"/>
			</svg>
			Tip
		</h3>
        <span>

Strings within backticks can span multiple lines.

```go
var beautifulPoem = `
Lorem ipsum dolor sit amet,
consectetur adipiscing elit,
sed do eiusmod tempor incididunt
ut labore et dolore magna aliqua
`
```

</span>
	</div>
	</div>

```go
type User struct {
	Name        string `json:"name"`
	YearJoined  int	   `json:"year_joined"`
	Premium     bool   `json:"premium"`
}
```

The output of marshaling this struct would be:

```json
{
  "name": "Alice",
  "year_joined": 2022,
  "premium": true
}
```

## Exercise

File: `11-structs-2/02-tags/main.go`

Fill the `Event` struct with `json` tags, so that the output keys are in `snake_case`.


<div class="alert alert-dismissible bg-light-primary d-flex flex-column flex-sm-row p-7 mb-10">
    <div class="d-flex flex-column">
        <h3 class="mb-5 text-dark">
			<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-lightbulb text-primary" viewBox="0 0 16 16">
			  <path d="M2 6a6 6 0 1 1 10.174 4.31c-.203.196-.359.4-.453.619l-.762 1.769A.5.5 0 0 1 10.5 13a.5.5 0 0 1 0 1 .5.5 0 0 1 0 1l-.224.447a1 1 0 0 1-.894.553H6.618a1 1 0 0 1-.894-.553L5.5 15a.5.5 0 0 1 0-1 .5.5 0 0 1 0-1 .5.5 0 0 1-.46-.302l-.761-1.77a1.964 1.964 0 0 0-.453-.618A5.984 5.984 0 0 1 2 6zm6-5a5 5 0 0 0-3.479 8.592c.263.254.514.564.676.941L5.83 12h4.342l.632-1.467c.162-.377.413-.687.676-.941A5 5 0 0 0 8 1z"/>
			</svg>
			Tip
		</h3>
        <span>

Notice how we're *casting* the `marshaled` variable (`[]byte`) into a `string` before printing it.

```go
fmt.Println(string(marshaled))
```

A slice of `byte`s and a `string` can be casted either way.

</span>
	</div>
	</div>