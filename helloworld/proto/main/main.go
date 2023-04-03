package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"micro-server/helloworld/proto"
)

func main() {
	// 定义数据
	req := __.HelloRequest{
		Name: "huang",
		Age:  18,
	}
	// 加密数据，变成二进制切片
	rsp, _ := proto.Marshal(&req)
	// 解密数据
	newReq := __.HelloRequest{}
	_ = proto.Unmarshal(rsp, &newReq)
	fmt.Println(newReq.Name, newReq.Age)
}
