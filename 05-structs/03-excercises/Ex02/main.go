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

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	fmt.Println(m["Verano"])
	fmt.Println(m["Castañeda"])

	for k, v := range m {
		fmt.Println(k)
		for _, v2 := range v.favICFlavours {
			fmt.Printf("\t%s\n", v2)
		}
	}

}
