idservice
====
[![Go Report Card](https://goreportcard.com/badge/github.com/JaneliaSciComp/idservice)](https://goreportcard.com/report/github.com/JaneliaSciComp/idservice)
[![GoDoc](https://godoc.org/github.com/JaneliaSciComp/idservice?status.png)](https://godoc.org/github.com/JaneliaSciComp/idservice) 

Returns unique version ids via various protocols.

- [x] HTTP
- [ ] gRPC
- [ ] zeromq

### HTTP example

Start HTTP service:

    idservice http

or specify port and working directly explicity

    idservice http -p :8000 -w /path/to/id-file-dir

To obtain ids, the client sends a POST request with an optional count:

    POST /v1/id

        Returns application/json:
        {"id":1}

    POST /v1/id?count=10

        Returns application/json:
        {"ids":[1,10]}

### Installation

Download the [latest binary executable releases](https://github.com/JaneliaSciComp/idservice/releases/latest)
for your platform.

Alternatively, if you have Go installed, you can install `idservice` with:

    go install github.com/JaneliaSciComp/idservice@latest

This will install the `idservice` executable in `$GOPATH/bin` or `$GOBIN`
or by default in `~/go/bin` if neither is set.

Or you can build from source by cloning the repo, changing into its directory,
and running a `go build`.
