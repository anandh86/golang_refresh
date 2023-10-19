# Packages

You can group `.go` files in directories. Go treats each directory as a **package**.
All files in the directory need to have the same `package` at the top.

```go
package storage
```

To import a package, add an import path using the full module's path.

```go
import "my-app/storage"
```

To call functions or access variables from the package, use its name.

```go
users, err := storage.AllUsers()
```

The package name doesn't need to be the same as the directory name.
But it's confusing otherwise, so name packages the same as the directories that keep them.

### Modules and Packages

Remember the difference:

* One **module** can contain many **packages** (directories).
* You need only one `go.mod` file in one module (in the top directory).
* Modules are defined by `go.mod`, packages are defined by directories.

## Exercise

File: `14-packages/02-packages/main.go`

Add a new `money` package (and directory).

Write a `money.New` function that returns new `money.Money` struct with the given amount (an integer) and currency (a string).

See the usage in `main.go` for an example.


<div class="accordion" id="hints-accordion">

<div class="accordion-item">
	<h3 class="accordion-header" id="hints-accordion-header-1">
	<button class="accordion-button fs-4 fw-semibold collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#hints-accordion-body-1" aria-expanded="false" aria-controls="hints-accordion">
		Hint #1
	</button>
	</h3>
	<div id="hints-accordion-body-1" class="accordion-collapse collapse" aria-labelledby="hints-accordion-header-1" data-bs-parent="#hints-accordion">
	<div class="accordion-body">

You don't need to add another `go.mod` file inside the `money` directory.
It's a package, not a module.

</div>
	</div>
	</div>

</div>
