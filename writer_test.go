package iorpc

import (
	"bytes"
	"io"
	"net/rpc"
	"testing"
)

func TestWriter_Impl(t *testing.T) {
	var raw interface{}
	raw = new(Writer)
	if _, ok := raw.(io.Writer); !ok {
		t.Fatal("Writer must be an io.Writer")
	}
}

func TestWriter(t *testing.T) {
	buf := new(bytes.Buffer)

	server := rpc.NewServer()
	RegisterWriter(server, buf)
	addr := serveSingleConn(server)

	rpcClient, err := rpc.Dial("tcp", addr)
	if err != nil {
		t.Fatalf("Couldn't connect: %s", err)
	}

	writer := NewWriter(rpcClient)
	n, err := writer.Write([]byte("hello!"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if n != 6 {
		t.Fatalf("n != 4: %d", n)
	}

	if buf.String() != "hello!" {
		t.Fatalf("bad value: %s", buf.String())
	}
}
