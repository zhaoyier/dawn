package main

import (
	"fmt"
	"zhao.com/dawn"
)

func main()  {
	t1 := dawn.NewServer()
	t1.Start("localhost:11000")

	m := make(map[string]int32)
	m["1"] = 1
	m["2"] = 2

	fmt.Println("=====>>:\t", len(m))
}
