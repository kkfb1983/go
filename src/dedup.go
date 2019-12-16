package src

import (
	"bufio"
	"fmt"
	"os"
)

func Dedup(){
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		line := input.Text()
		if !seen[line]{
			seen[line] = true
			fmt.Println("line->",line)
		}
	}
	if err := input.Err(); err != nil{
		fmt.Fprintf(os.Stderr, "dedup : %v\n", err)
	}
}

// equal(map[string]int{"A": 0}, map[string]int{"B": 42})
func equal(x,y map[string]int) bool{
	if len(x) != len(y){
		return false
	}
	for k,xv := range x{
		if yv, ok := y[k]; !ok || yv != xv{
			return false
		}
	}
	return  true
}
