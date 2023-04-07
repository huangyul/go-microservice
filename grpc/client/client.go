package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"micro-server/grpc/proto"
)

func main() {
	// 建立与服务端的链接
	conn, err := grpc.Dial("127.0.0.1:5555", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// 创建客户端的grpc实例，调用SayHello方法
	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "huang"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
