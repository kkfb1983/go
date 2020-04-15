package src

import "fmt"

type PeopleDemo struct {
	name string
	age  int
}

type TestDemo interface {
	PrintDemo()
	SleepDemo()
}

type StudentDemo struct {
	name  string
	age   int
	score int
}

func (p StudentDemo) PrintDemo() {
	fmt.Println("name:", p.name)
	fmt.Println("age:", p.age)
	fmt.Println("score:", p.score)
}

func (p StudentDemo) SleepDemo() {
	fmt.Println("student sleep")
}

func (people PeopleDemo) PrintDemo() {
	fmt.Println("name:", people.name)
	fmt.Println("age:", people.age)
}

func (p PeopleDemo) SleepDemo() {
	fmt.Println("people sleep")
}

func InterfaceDemo() {

	var t TestDemo
	fmt.Println(t)
	//t.PrintDemo()

	var stu StudentDemo = StudentDemo{
		name:  "stu1",
		age:   20,
		score: 200,
	}

	t = stu
	t.PrintDemo()
	t.SleepDemo()

	var people PeopleDemo = PeopleDemo{
		name: "people",
		age:  100,
	}

	t = people
	t.PrintDemo()
	t.SleepDemo()

	fmt.Println("t:", t)
}
