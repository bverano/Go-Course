package main

import "fmt"

type person struct {
	firstName     string
	lastName      string
	favICFlavours []string
}

func main() {

	p1 := person{
		firstName: "Bruno",
		lastName:  "Verano",
		favICFlavours: []string{
			"chocolate",
			"vanilla",
			"choco-chips",
		},
	}

	p2 := person{
		firstName: "Luis",
		lastName:  "Castañeda",
		favICFlavours: []string{
			"strawberry",
			"watermelon",
		},
	}

	persons := []person{p1, p2}

	for _, v := range persons {
		fmt.Println(v.firstName, v.lastName)
		for _, v2 := range v.favICFlavours {
			fmt.Printf("\t%s\n", v2)
		}
	}

}
