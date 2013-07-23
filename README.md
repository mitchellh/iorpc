# iorpc

iorpc is a Go library for serving and accessing `io` interfaces across
`net/rpc` connections.

## Installation

Standard `go get`:

```
$ go get github.com/mitchellh/iorpc
```

## Usage & Example

For the most up-to-date examples and usage information, you should always
access the documentation. However, for a quick look at what using
this library looks like, this basic example is shown below:

```go
// On the server side, you just register your reader onto an RPC
// server, and serve it as usual.
server := rpc.NewServer()
reader := new(bytes.Buffer) // This can be any io.Reader
iorpc.RegisterReader(server, reader)

// On the client side, you connect to an RPC server like usual,
// the use the iorpc client to read it.
client, _ := rpc.Dial("tcp", "127.0.0.1:1234")
reader := iorpc.NewReader(client)

// "reader" is now a valid io.Reader to use that will read data from
// the remote io.Reader
```