package main

import "fmt"

type person struct {
	name string
	age  int
}

type secretAgent struct {
	person
	killLicense bool
}

func (s secretAgent) introduce() {
	fmt.Println("Hello, my name is", s.name, "and I'm", s.age, "years old.")
	if s.killLicense {
		fmt.Println("I CAN KILL YOU")
	}
}

func main() {
	sa1 := secretAgent{
		person: person{
			name: "Bruno",
			age:  21,
		},
		killLicense: true,
	}
	sa1.introduce()
}
