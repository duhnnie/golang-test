package main

import (
	"example/creational/singleton"
	"fmt"
)

func main() {
	s := singleton.GetInstance()
	s2 := singleton.GetInstance()

	fmt.Println(s.AddOne())
	fmt.Println(s2.AddOne())
	fmt.Println(s == s2)
}
