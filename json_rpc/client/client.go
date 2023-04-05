package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 这里可以理解为，使用net接收数据，然后交给rpc处理，之前直接使用rpc是因为没有自定义序列化格式
	conn, _ := net.Dial("tcp", ":5555")
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	var reply = new(string)
	_ = client.Call("HelloService.Hello", "huang", reply)
	fmt.Println(*reply) // hello,huang
}
