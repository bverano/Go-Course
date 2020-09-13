package main

import "fmt"

func main() {
	m := map[string]int{
		"Bruno":  21,
		"Jacobo": 5,
	}

	fmt.Println(m)

	delete(m, "Bruno")
	fmt.Println(m)

	//Verifying if the value exists on the map before tryin to delete it
	if v, ok := m["Jacobo"]; ok {
		fmt.Println(v, "was deleted from the map")
		delete(m, "Bruno")
	}

}
