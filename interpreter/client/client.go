package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"micro-server/interpreter/proto"
)

func main() {
	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		fmt.Println("发送前拦截")
		return invoker(ctx, method, req, reply, cc, opts...)
	}
	opt := grpc.WithUnaryInterceptor(interceptor)
	conn, _ := grpc.Dial("0.0.0.0:5555", grpc.WithInsecure(), opt)
	client := proto.NewGreeterClient(conn)
	r, _ := client.Hello(context.Background(), &proto.ServerRequest{Name: "huang"})
	fmt.Println(r)
}
