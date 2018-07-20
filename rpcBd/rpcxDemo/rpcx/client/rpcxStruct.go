package main

import "flag"

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}
