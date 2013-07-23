package iorpc

import (
	"io"
	"net/rpc"
)

// Reader is an implementation of io.Reader that performs its operations
// across an RPC connection.
type Reader struct {
	client *rpc.Client
	typeName string
}

// ReaderServer wraps an io.Reader and serves it across an RPC connection.
// The ReaderServer can be accessed by a Reader.
type ReaderServer struct {
	reader io.Reader
}

// NewReader returns a new Reader that communicates with the given
// RPC client.
func NewReader(client *rpc.Client) *Reader {
	return NewReaderName(client, "ReaderServer")
}

// NewReaderName returns a new Reader that communicates with the given
// RPC client using name as the type name for the remote ReaderServer in
// RPC calls.
func NewReaderName(client *rpc.Client, name string) *Reader {
	return &Reader{
		client: client,
		typeName: name,
	}
}

// RegisterReader registers the io.Reader on the RPC server.
func RegisterReader(server *rpc.Server, r io.Reader) error {
	return server.Register(&ReaderServer{reader: r})
}

func (r *Reader) Read(p []byte) (n int, err error) {
	var result []byte
	err = r.client.Call(r.typeName + ".Read", len(p), &result)
	if err != nil {
		if result != nil {
			n = len(result)
		}

		if err.Error() == io.EOF.Error() {
			err = io.EOF
		}

		return
	}

	copy(p, result)
	return len(result), nil
}

func (r *ReaderServer) Read(n int, data *[]byte) error {
	*data = make([]byte, n)
	n, err := r.reader.Read(*data)
	*data = (*data)[0:n]
	return err
}
