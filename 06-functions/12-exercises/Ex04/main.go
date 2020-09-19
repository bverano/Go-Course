package main

import "fmt"

type person struct {
	name     string
	lastName string
	age      int
}

func (p person) introduceSelf() {
	fmt.Println("Hello! I'm", p.name, p.lastName, "and I'm", p.age, "years old.")
}

func main() {
	p1 := person{
		name:     "Bruno",
		lastName: "Verano",
		age:      21,
	}
	p1.introduceSelf()
}
