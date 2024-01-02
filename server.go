package main

import (
	"context"
	"fmt"
	"net"
	"rpc-demo/hello"

	"google.golang.org/grpc"
)

type server struct {
	hello.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloReply, error) {
	return &hello.HelloReply{Message: "Hello, " + req.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	srv := grpc.NewServer()
	hello.RegisterGreeterServer(srv, &server{})

	fmt.Println("Grpc server is running on 8080")
	if err := srv.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}

}
