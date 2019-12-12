package main

import (
	"fmt"
)

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Println(p.name, "says hi")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	person := person{"Kevin"}
	saySomething(&person)
	// Below doesn't work, person doesn't implement the human interface because the speak method takes a person pointer as an input. So passing in the var person, will not hit the speak() method
	// saySomething(person)
}
