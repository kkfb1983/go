package main

import (
	"./src"
	"fmt"
)

func main() {
	//for _, arg := range os.Args[1:]{
	//
	//}
	/** add func */
	//var c int
	// c = add(3, 8)
	param := src.Add(5, 15)
	fmt.Println("hello word", param)
	// go fmt.Println("hello word", param)
	/** tempconv func */
	//var c src.Celsius
	//var f src.Fahrenheit
	//c = int(src.CTof(12.13))
	//fmt.Println(c)
	//fmt.Printf("%g\n",src.BoilingC-src.FreezingC)
	//boilingF := src.CTof(src.BoilingC)
	//fmt.Printf("%g\n",boilingF-src.CTof(src.FreezingC))
	//fmt.Printf("%g\n",boilingF-src.FreezingC)
	//fmt.Println(c==0)
	//fmt.Println(f>=0)
	////fmt.Println(c==f)
	//fmt.Println(c == src.Celsius(f))
	//c := src.FToC(212.0)
	//fmt.Println(c.String())
	//fmt.Printf("%v\n", c)
	//fmt.Printf("%s\n", c)
	//fmt.Println(c)
	//fmt.Printf("%g\n", c)
	//fmt.Println(float64(c))
	//c := src.CTok(1)
	//fmt.Println(c)
	//fmt.Println(float64(c))
	fmt.Println(src.PopCount(9))
}
