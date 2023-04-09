package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"micro-server/metadata/proto"
)

func main() {
	conn, err := grpc.Dial("tcp", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	r, err := c.Hello(context.Background(), &proto.HelloRequest{Name: "huang"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
