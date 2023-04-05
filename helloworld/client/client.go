package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 1. 建立链接
	client, err := rpc.Dial("tcp", ":5555")
	if err != nil {
		panic("链接失败")
	}
	var reply = new(string)
	// 2. 调用远程方法
	err = client.Call("HelloService.Hello", "huang", reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(*reply) // hello huang
}
