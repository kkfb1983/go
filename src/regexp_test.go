package src

import "regexp"

func Compile(expr string) (*Regexp, error) {

}

func MustCompile(expr string) *Regexp {
	re, err := Compile(expr)
	if err != nil {
		panic(err)
	}
	return re
}
