package main

import (
	"context"
	"flag"
	"log"

	"github.com/smallnest/rpcx/client"
)

var (
	addr2 = flag.String("addr", "localhost:8972", "server address")
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

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr2, "")

	xclinet := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclinet.Close()

	args := &Args{
		A: 10,
		B: 20,
	}

	reply := &Reply{}

	call, err := xclinet.Go(context.Background(), "Mul", args, reply, nil)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	replyCall := <-call.Done
	if replyCall.Error != nil {
		log.Fatalf("failed to call: %v ", replyCall.Error)
	} else {
		log.Fatalf("%d * %d = %d", args.A, args.B, reply.C)
	}
}
