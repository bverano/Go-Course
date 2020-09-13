package main

import "fmt"

func main() {

	// map[keyType]ValuesType

	m := map[string]int{
		"Jacobo": 5,
		"Bruno":  21,
	}
	fmt.Println(m)

	fmt.Println(m["Bruno"])

	if v, ok := m["Jacobo"]; ok { //Use this form when you want to know if the element exists in the map
		fmt.Println("Printing from if statement", v)
	}

}
