# Introduction

Here there are just some references to how use external dependencies.

To define a module, just run:

```shell
go mod init <module name>
```

In go.mod, whenever there is a new dependency declaration and a `go` command is launched (build, run test, ...) a new line with the additional dependency is added.

To downgrade to a specific dependency version use (for instance to install v0.8.1 of `gituhb.com/pkg/errors`):

```shell
go get gituhb.com/pkg/errors@v0.8.1
```

The go.sum file contains checksums to validate the dependency itself.

To remove old references from the go.sum fil simply run:

```shell
go mod tidy
```

To update recursively all the dependencies just run:

```shell
go get -u ./...
```

# Credits

Heavily based on [this video](https://www.youtube.com/watch?v=AJxKKmzRTUY)