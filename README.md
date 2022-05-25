idservice
====

[![GoDoc](https://godoc.org/github.com/JaneliaSciComp/idservice?status.png)](https://godoc.org/github.com/JaneliaSciComp/idservice) 

Returns unique version ids via various protocols.

- [x] HTTP
- [ ] gRPC
- [ ] zeromq

### HTTP example

    % idservice http -p :8000 -w /path/to/workdir

To use, the client sends a POST request with an optional count:

    POST /v1/id

        Returns application/json:
        {"id":1}

    POST /v1/id?count=10

        Returns application/json:
        {"ids":[1,10]}

### Installation

If you have Go installed, you can install `idservice` with:

    go install github.com/JaneliaSciComp/idservice@latest

This will install the `idservice` executable in `$GOPATH/bin` or `$GOBIN`
or by default in `~/go/bin` if neither is set.

Or you can build from source by cloning the repo, changing into its directory,
and running a `go build`.
