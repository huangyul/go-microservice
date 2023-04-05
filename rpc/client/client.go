package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
)

type ResponseData struct {
	Data int `json:"data"`
}

func Add(a, b int) int {
	req := HttpRequest.NewRequest()
	res, _ := req.Get("http://127.0.0.1:3000/add?a=1&b=2")
	body, _ := res.Body()
	fmt.Println(string(body))
	resData := ResponseData{}
	_ = json.Unmarshal(body, &resData)
	return resData.Data
}

func main() {
	fmt.Println(Add(1, 2))
}
