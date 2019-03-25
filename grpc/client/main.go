package main

import (
	"context"
	"fmt"
	"github.com/yidane/gotest/grpc/message"
	"google.golang.org/grpc"
	"os"
	"time"
)

const (
	address     = "localhost:50022"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	client := message.NewGreeterClient(conn)
	request := message.HelloRequest{
		Name: "evan",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.SayHello(ctx, &request)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(response)
}
