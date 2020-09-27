package main

import (
	"fmt"
)

type errorPer struct {
	name string
	age  int
	err  error
}

func (ep errorPer) Error() string {
	return fmt.Sprintf("The person %s, of %d years old had an issue: %v", ep.name, ep.age, ep.err)
}

func main() {
	e1 := errorPer{
		name: "Bruno",
		age:  21,
		err:  fmt.Errorf("You are too old for playing with toys"),
	}
	foo(e1)
}

func foo(e error) {
	fmt.Println(e)
}
