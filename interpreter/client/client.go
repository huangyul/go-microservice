package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"micro-server/interpreter/proto"
)

func main() {
	conn, _ := grpc.Dial("0.0.0.0:5555", grpc.WithInsecure())
	client := proto.NewGreeterClient(conn)
	r, _ := client.Hello(context.Background(), &proto.ServerRequest{Name: "huang"})
	fmt.Println(r)
}
