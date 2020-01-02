package src

import (
	"crypto/sha256"
	"fmt"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func Test() {
	// 数组
	//array()
	//sha()
	val := panic_test(2)
	fmt.Println(val)
}

func array() {
	//var a [3]int
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	//fmt.Println(a[0])
	//fmt.Println(a[len(a) - 1])
	//for i, v:=range a{
	//	fmt.Printf("%d %d\n",i,v)
	//}
	for _, v := range q {
		fmt.Printf("%d\n", v)
	}
	fmt.Println("\n##########################\n")
	for _, v := range r {
		fmt.Printf("%d\n", v)
	}
	fmt.Println("\n##########################\n")
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB])
	fmt.Println("\n##########################\n")
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"
	//d := [3]int{1, 2}
	//fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int
}

func sha() {
	fmt.Println("\n##########################\n")
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}
func results(d int) (r, f int, err error) {
	r += d
	f += d
	return
}

func panic_test(v int) int {
	var res int
	switch v {
	case 1:
		res = v
	}
	return res
}
