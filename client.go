package main

import (
	"fmt"
	"net/rpc"
	"rpc-demo/rpc-protocol"
)
func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Dial error:", err)
		return
	}
	args := &protocol.Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		fmt.Println("Call error:", err)
		return
	}
	fmt.Printf("%d * %d = %d\n", args.A, args.B, reply)
}
