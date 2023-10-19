# Modules

## Installing Go

Until now, you didn't need a local Go installation to run the code.
Now is a good time to do it, as you'll learn how to work with the `go` CLI tool.

Head to [go.dev](https://go.dev/dl/) and download the latest stable version of Go.

Once installed, confirm it works by running:

```bash
go version
```

## Modules

Modules are complete Go projects, like a library or an application.
Usually, a single repository keeps a single module.

A module is defined in a `go.mod` file. You don't need to modify this file.
The `go` CLI manages it for you.

To generate the `go.mod` file, run `go mod init` providing the module's name.

```bash
go mod init my-app
```

It generates content like this:

```text
module my-app

go 1.18
```

The module's name is usually the full path to the repository. For example, `github.com/ThreeDotsLabs/cli`.

It's possible to use any name you like, but it won't be seamless to import your module later.
As a rule of thumb, use the full path of your repository as the module's name.

## Exercise

File: `14-packages/01-modules/main.go`

Enter the directory with `main.go` and run `go mod init`, so the module name is `hello`.
