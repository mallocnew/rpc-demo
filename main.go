package main

import (
	"fmt"
	"net"
	"net/rpc"
	"rpc-demo/rpc-protocol"
)

func main() {
	arith := new(protocol.Arith)
	rpc.Register(arith)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Listen error:", err)
		return
	}
	fmt.Println("RPC server listen on port 1234")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}