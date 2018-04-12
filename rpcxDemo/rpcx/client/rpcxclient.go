package main

import (
	"context"
	"flag"
	"log"

	"github.com/smallnest/rpcx/client"
)

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

func main() {

	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclinet := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)

	defer xclinet.Close()

	args := &Args{
		A: 10,
		B: 20,
	}

	reply := &Reply{}
	err := xclinet.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call :%v ", err)
	}
	log.Printf("%d * %d = %d ", args.A, args.B, reply.C)

	err2 := xclinet.Call(context.Background(), "Add", args, reply)
	if err2 != nil {
		log.Fatalf("failed to call :%v ", err2)
	}
	log.Printf("%d + %d = %d ", args.A, args.B, reply.C)
}
