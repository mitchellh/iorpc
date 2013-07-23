// Package iorpc provides a way to serve and access io package interface
// implementations such as io.Reader over a net/rpc connection. The
// implementation strives for correctness over performance. For example,
// with the io.Reader implementation, if you request to Read one byte, it will
// perform a full RPC request/response to read only a single byte.
//
// Accessing io implementations across the network can be extremely useful
// in certain environments. And, if used properly, can be performant as well.
package iorpc

// This file only exists to set package documentation.
