package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (h *HelloService) Hello(s string, reply *string) error {
	*reply = "Hello, " + s
	return nil
}

func main() {
	listen, _ := net.Listen("tcp", ":5555")
	rpc.RegisterName("HelloService", &HelloService{})
	for {
		conn, _ := listen.Accept()
		// 改为使用json
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
