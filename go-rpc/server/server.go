package main

import (
	"golang-code/go-rpc/proto"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

func main() {

	arith := new(proto.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	go http.Serve(l, nil)
	time.Sleep(60 * time.Second)
}
