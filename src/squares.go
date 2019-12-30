package src

import "fmt"

func anonymity_f() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func Squares() {
	f := anonymity_f()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
