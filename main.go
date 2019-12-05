package main

import (
	"./src"
	"fmt"
)

func main() {
	var c int
	// c = add(3, 8)
	param := src.Add(5, 15)
	fmt.Println("hello word", param)
	// go fmt.Println("hello word", param)
	c = int(src.CTof(12.13))
	fmt.Println(c)
}
