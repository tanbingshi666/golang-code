package main

import (
	"fmt"
	"golang-code/go-rpc/proto"
	"log"
	"net/rpc"
)

func main() {

	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &proto.Args{7, 8}

	// Synchronous call
	//var reply int
	//err = client.Call("Arith.Multiply", args, &reply)
	//if err != nil {
	//	log.Fatal("arith error:", err)
	//}
	//fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)

	// Asynchronous call
	quotient := new(proto.Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done // will be equal to divCall
	// check errors, print, etc.
	fmt.Println(replyCall.Reply)

}
