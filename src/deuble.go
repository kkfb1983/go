package src

import "fmt"

func Deuble(x int) (result int) {
	defer func() { fmt.Printf("Deoble x(%d)->%D\n", x, result) }()
	return x + x
}
