package main

import (
	"context"
	"fmt"
	"github.com/yidane/gotest/grpc/message"
	"google.golang.org/grpc"
	"net"
	"os"
)

const (
	port = ":50022"
)

type server struct {
}

func (server) SayHello(ctx context.Context, req *message.HelloRequest) (res *message.HelloResponse, err error) {
	fmt.Println(req.Name)

	return &message.HelloResponse{Message: "hello " + req.Name}, nil
}

func (server) SayBye(ctx context.Context, req *message.HelloRequest) (res *message.HelloResponse, err error) {
	fmt.Println(req.Name)

	return &message.HelloResponse{Message: "hello " + req.Name}, nil
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s := grpc.NewServer()

	message.RegisterGreeterServer(s, &server{})

	err = s.Serve(listener)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("exited")
}
