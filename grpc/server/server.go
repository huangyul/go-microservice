package main

import (
	"context"
	"net"

	"google.golang.org/grpc"

	"micro-server/grpc/proto"
)

type Server struct {
}

func (s Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "huang",
	}, nil
}

func main() {
	// 1.启动服务
	listen, _ := net.Listen("tcp", "0.0.0.0:5555")
	// 2.注册服务
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	g.Serve(listen)
}
