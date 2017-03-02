# g..

g.. is a wrapper around the go command for running go sub-commands
(e.g. test, vet, fmt, ...) on paths ending with `...` with excluded
vendor packages.

It is meant as a cross-platform workaround for the
[golang/go#19090](https://github.com/golang/go/issues/19090)
issue.

## Installation

```bash
go get -u github.com/mibk/g..
```

## Example of usage

```bash
# list all packages of the current application
g.. # it's basically equivalent to `go list |grep -v vendor`

# test them
g.. test

# gofmt them
g.. fmt

# list all Docker packages (compare it with go list github.com/docker/docker/...)
g.. github.com/docker/docker/...

# test all packages in Docker
g.. github.com/docker/docker/... test
```

For additional information run `g.. -h`.
