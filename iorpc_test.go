package iorpc

import (
	"fmt"
	"net"
	"net/rpc"
)

var testPort = 4000

func serveSingleConn(s *rpc.Server) string {
	l, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", testPort))
	if err != nil {
		panic(err)
	}
	testPort += 1

	// Accept a single connection in a goroutine and then exit
	go func() {
		defer l.Close()
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}

		s.ServeConn(conn)
	}()

	return l.Addr().String()
}
