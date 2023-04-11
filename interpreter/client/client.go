package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"micro-server/interpreter/proto"
)

func main() {
	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		fmt.Println("发送前拦截")
		// 同一添加身份信息
		md := metadata.New(map[string]string{
			"appid":  "1010",
			"appkey": "huang",
		})
		ctx = metadata.NewOutgoingContext(context.Background(), md)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
	opt := grpc.WithUnaryInterceptor(interceptor)
	conn, _ := grpc.Dial("0.0.0.0:5555", grpc.WithInsecure(), opt)
	client := proto.NewGreeterClient(conn)
	r, err := client.Hello(context.Background(), &proto.ServerRequest{Name: "huang"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}
