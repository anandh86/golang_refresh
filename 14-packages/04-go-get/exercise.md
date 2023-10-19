# go get

To import an external module, you have to add it as a dependency.

To do so, run the `go get` command followed by the module's full path.

For example:

```bash
go get github.com/google/uuid
```

This will change your go.mod file to something like:

```text
module tax

go 1.18

require github.com/google/uuid v1.3.0 // indirect
```

It will also generate a `go.sum` file with checksums of the dependencies.
You won't need to modify this file.
If you use version control (like a git repository), commit both `go.mod` and `go.sum` files.

You can now import the module in your code.

## Exercise

File: `14-packages/04-go-get/main.go`

The `CalculateTax` function uses an external `decimal` library for calculations.

Add a Go module (with `go mod init`).

Run `go get` to add the missing dependency.
