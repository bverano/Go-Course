package main

import "fmt"

type person struct {
	name string
	age  int
}

func changeMe(p *person) {
	p.name = "Jose"
	(*p).age = 50
}

func main() {
	p1 := person{
		name: "Bruno",
		age:  21,
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
