package iorpc

import (
	"io"
	"net/rpc"
)

// Writer is an implementation of io.Writer that performs its operations
// across an RPC connection.
type Writer struct {
	client   *rpc.Client
	typeName string
}

// WriterServer wraps a io.Writer and serves it across an RPC connection.
// The WriterServer can be accessed using an RPC client and Writer.
// Use RegisterWriter and RegisterWriterName to serve a writer from your
// RPC server.
type WriterServer struct {
	writer io.Writer
}

// NewWriter returns a new Writer that communicates with the given
// RPC client.
func NewWriter(client *rpc.Client) *Writer {
	return NewWriterName(client, "WriterServer")
}

// NewWriterName returns a new Writer that communicates with the given
// RPC client using name as the type name for the remote WriterServer in
// RPC calls.
func NewWriterName(client *rpc.Client, name string) *Writer {
	return &Writer{
		client:   client,
		typeName: name,
	}
}

// RegisterWriter registers the io.Writer on the RPC server.
func RegisterWriter(server *rpc.Server, r io.Writer) error {
	return server.Register(&WriterServer{writer: r})
}

// RegisterWriterName registers the io.Writer on the RPC server with
// the given type name instead of the default inferred type name.
func RegisterWriterName(s *rpc.Server, name string, r io.Writer) error {
	return s.RegisterName(name, &WriterServer{writer: r})
}

func (w *Writer) Write(p []byte) (n int, err error) {
	err = w.client.Call(w.typeName+".Write", p, &n)
	if err != nil && err.Error() == io.EOF.Error() {
		err = io.EOF
	}

	return
}

func (w *WriterServer) Write(p []byte, n *int) (err error) {
	*n, err = w.writer.Write(p)
	return
}
