package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name     string
	Lastname string
	Age      int
}

func main() {
	p1 := person{
		Name:     "Bruno",
		Lastname: "Verano",
		Age:      21,
	}
	p2 := person{
		Name:     "Luis",
		Lastname: "Castañeda",
		Age:      19,
	}
	persons := []person{p1, p2}
	fmt.Println(persons)

	bs, err := json.Marshal(persons)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%s", bs)

}
