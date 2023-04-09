package main

import (
	"context"
	"google.golang.org/grpc"
	"micro-server/metadata/proto"
	"net"
)

type server struct {
}

func (s *server) Hello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "hello" + request.Name,
	}, nil
}

func main() {
	listen, _ := net.Listen("tcp", ":5555")

	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &server{})
	_ = g.Serve(listen)
}
