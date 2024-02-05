package main

import (
	"context"
	"fmt"
	"rpc-demo/hello"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dial error:", err)
		return
	}
	defer conn.Close()

	client := hello.NewGreeterClient(conn)
	req := hello.HelloRequest{Name: "gRPC"}
	reply, err := client.SayHello(context.Background(), &req)
	if err != nil {
		fmt.Printf("failed to call sayhello: %v", err)
		return
	}
	fmt.Printf("reply from server: %s\n", reply.Message)
}
