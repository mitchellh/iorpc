package iorpc

import (
	"bytes"
	"io"
	"net/rpc"
	"testing"
)

func TestReader_Impl(t *testing.T) {
	var raw interface{}
	raw = new(Reader)
	if _, ok := raw.(io.Reader); !ok {
		t.Fatal("Reader must be an io.Reader")
	}
}

func TestReader(t *testing.T) {
	buf := new(bytes.Buffer)
	buf.WriteString("hello\n")

	server := rpc.NewServer()
	RegisterReader(server, buf)
	addr := serveSingleConn(server)

	rpcClient, err := rpc.Dial("tcp", addr)
	if err != nil {
		t.Fatalf("Couldn't connect: %s", err)
	}

	data := make([]byte, 4)
	reader := NewReader(rpcClient)
	n, err := reader.Read(data)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if n != 4 {
		t.Fatalf("n != 4: %d", n)
	}

	if string(data) != "hell" {
		t.Fatalf("bad value: %s", string(data))
	}

	n, err = reader.Read(data)
	if n != 2 {
		t.Fatalf("n != 2: %d", n)
	}

	if string(data[0:n]) != "o\n" {
		t.Fatalf("bad value: %s", string(data))
	}

	n, err = reader.Read(data)
	if err != io.EOF {
		t.Fatalf("non-EOF: %s", err)
	}

	if n != 0 {
		t.Fatalf("n != 0: %d", n)
	}
}
