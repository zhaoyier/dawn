package main

import (
	"zhao.com/dawn"
)

func main()  {
	t1 := dawn.NewServer()
	t1.Start("localhost:11000")
}
