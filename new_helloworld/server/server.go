package main

import (
	"micro-server/new_helloworld/handler"
	"micro-server/new_helloworld/server_proxy"
	"net"
	"net/rpc"
)

func main() {
	// 1. 新建服务
	listen, _ := net.Listen("tcp", ":5555")
	// 2. 注册服务
	server_proxy.RegisterHelloService(&handler.HelloService{})
	// 3. 获取链接得到的值，并给rpc处理
	for {
		conn, _ := listen.Accept()
		go rpc.ServeConn(conn)
	}
}
