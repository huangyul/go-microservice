package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 1. 建立链接
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic("链接失败")
	}
	// 2. 调用函数
	var reply = new(string)
	err = client.Call("HelloService.Hello", "huang", reply)
	if err != nil {
		fmt.Println(err)
		panic("调用失败")
	}
	fmt.Println(*reply)
}
