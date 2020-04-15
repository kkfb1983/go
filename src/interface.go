package src

import (
	"fmt"
)

type Speaker interface {
	Hello()
}

type UserInterface struct {
	Name string
	Age  int
}

func (this *UserInterface) Hello() {
	fmt.Println("hello my name is", this.Name)
}

func Interface() {

	var s Speaker

	s = &UserInterface{"wss", 10}

	s.Hello()

}

