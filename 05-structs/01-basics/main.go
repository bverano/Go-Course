package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

type secretAgent struct {
	person
	kl bool
}

func main() {
	p1 := person{
		firstName: "Bruno",
		lastName:  "Verano",
		age:       21,
	}

	p2 := person{
		firstName: "Luis",
		lastName:  "Castañeda",
		age:       19,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.firstName, p1.lastName, p1.age)
	fmt.Println(p2.firstName, p2.lastName, p2.age)

	sa1 := secretAgent{
		person: person{
			firstName: "James",
			lastName:  "Bond",
			age:       32,
		},
		kl: true,
	}

	fmt.Println(sa1)
	fmt.Println(sa1.age, sa1.person.firstName)

}
