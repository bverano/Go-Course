package main

import "fmt"

type person struct {
	name string
	age  int
}

type human interface {
	talk()
}

func (p *person) talk() {
	fmt.Printf("Hello! I'm %s and this is my talk method!\n", p.name)
}

func saySomething(h human) {
	fmt.Println("From Say Something function.")
	h.talk()
}

func main() {
	p := person{
		name: "Bruno",
		age:  21,
	}
	// saySomething(p) Doesn't work
	saySomething(&p)
}
