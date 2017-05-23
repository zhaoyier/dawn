package main

import (
	"fmt"
	"zhao.com/dawn"
	"zhao.com/examples/proto3"
	"github.com/golang/protobuf/proto"
)

func main()  {
	dawn.Register(1000, Call)
	t1 := dawn.NewServer()
	t1.Start("localhost:11000")

	m := make(map[string]int32)
	m["1"] = 1
	m["2"] = 2

	fmt.Println("=====>>:\t", len(m))
}

func Call(b []byte) (bs []byte, err error)  {
	temp := &proto3.Page{}
	proto.Unmarshal(b, temp)
	fmt.Println("======>>.9003:\t", temp)
	return proto.Marshal(&proto3.Page{
		PageNumber: 10,
		PageSize: 30,
	})
}