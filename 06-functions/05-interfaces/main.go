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

func (s secretAgent) sayHi() {
	fmt.Println("Hello, my name is", s.name, "and I'm", s.age, "years old.")
	if s.killLicense {
		fmt.Println("I CAN KILL YOU")
	}

}
func (p person) sayHi() {
	fmt.Println("Hello, my name is", p.name, "and I'm", p.age, "years old. I'm not a S.A.")
}

type human interface {
	sayHi()
}

func bar(h human) {
	//type assertion
	switch h.(type) {
	case person:
		fmt.Println("I was passed to bar function. Person", h.(person).name)
	case secretAgent:
		fmt.Println("I was passed to bar function. Secret Agent", h.(secretAgent).name)
	}
}

type apple int

func main() {
	sa1 := secretAgent{
		person: person{
			name: "Bruno",
			age:  21,
		},
		killLicense: true,
	}

	p := person{
		name: "Luis",
		age:  19,
	}

	bar(sa1)
	bar(p)

	//convertion
	var x apple = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	sa1.sayHi()
}
