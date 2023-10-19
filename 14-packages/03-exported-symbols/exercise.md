# Exported symbols

You can only import symbols (variables, functions, types) from other packages if they start with an uppercase letter.

```go
func ExportedFunction() {
	
}

func unexportedFunction() {
	
}

const ExportedConst = 100

var unexportedVariable = "example"

type ExportedType struct {
    
}
```

Aim to export only what external packages would use. Make all internal details unexported.

## Struct fields

Similar rule applies to structs. Field names starting with lowercase won't be visible outside the package.

```go
type Money struct {
	amount int
	currency string
}

func (m Money) String() string {
	return fmt.Sprint(m.amount, " ", m.currency)
}
```

External packages are able to access the `String()` method, but not the `amount` or currency fields.


<div class="alert alert-dismissible bg-light-primary d-flex flex-column flex-sm-row p-7 mb-10">
    <div class="d-flex flex-column">
        <h3 class="mb-5 text-dark">
			<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-lightbulb text-primary" viewBox="0 0 16 16">
			  <path d="M2 6a6 6 0 1 1 10.174 4.31c-.203.196-.359.4-.453.619l-.762 1.769A.5.5 0 0 1 10.5 13a.5.5 0 0 1 0 1 .5.5 0 0 1 0 1l-.224.447a1 1 0 0 1-.894.553H6.618a1 1 0 0 1-.894-.553L5.5 15a.5.5 0 0 1 0-1 .5.5 0 0 1 0-1 .5.5 0 0 1-.46-.302l-.761-1.77a1.964 1.964 0 0 0-.453-.618A5.984 5.984 0 0 1 2 6zm6-5a5 5 0 0 0-3.479 8.592c.263.254.514.564.676.941L5.83 12h4.342l.632-1.467c.162-.377.413-.687.676-.941A5 5 0 0 0 8 1z"/>
			</svg>
			Tip
		</h3>
        <span>

If a struct has a `String() string` method, functions from the `fmt` package (like `fmt.Printf` and `fmt.Println`) use it to print out the structure.
See [`fmt.Stringer`](https://pkg.go.dev/fmt#Stringer) interface.

</span>
	</div>
	</div>

## Exercise

File: `14-packages/03-exported-symbols/account/account.go`

The `login` module shows a trivial login functionality. User accounts consist of an email address and password.

You can try logging in passing credentials as CLI arguments:

```bash
go run . email@example.com my-password
```

The `account.New()` constructor check whether both fields are not empty.
However, nothing stops you from changing the fields after an account is created.

```go
acc, err := account.New("kate@example.com", "t0ps3cr3t")
if err != nil {
	panic(err)
}

acc.Email = "joe@example.com"
acc.Password = ""

// Success!
ok := acc.Login("joe@example.com", "")
```

A better design would be to hide the details, so no one can access them outside the `account` package.
Introduce these changes.
