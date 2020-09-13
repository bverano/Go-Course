package main

import "fmt"

func main() {
	m := map[string]int{
		"Bruno":  21,
		"Jacobo": 5,
	}

	fmt.Println(m)

	//Adding element to map
	m["Luis"] = 19
	fmt.Println(m)

	for k, v := range m {
		fmt.Printf("%s's age is: %d\n", k, v)
	}

}
