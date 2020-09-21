package main

import (
	"fmt"
	"sort"
)

type user struct {
	Name     string
	Lastname string
	Age      int
	Sayings  []string
}

type byAge []user

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type byLastname []user

func (a byLastname) Len() int           { return len(a) }
func (a byLastname) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byLastname) Less(i, j int) bool { return a[i].Lastname < a[j].Lastname }

func main() {
	u1 := user{
		Name:     "Eduar",
		Lastname: "Tua",
		Age:      32,
		Sayings: []string{
			"Cachicamo diciéndole a morrocoy conchudo",
			"La mona, aunque se vista de seda, mona se queda",
			"Poco a poco se anda lejos",
		},
	}

	u2 := user{
		Name:     "Carlos",
		Lastname: "Pérez",
		Age:      27,
		Sayings: []string{
			"Camarón que se duerme se lo lleva la corriente",
			"A ponerse las alpargatas que lo que viene es joropo",
			"No gastes pólvora en zamuro",
		},
	}

	u3 := user{
		Name:     "Che",
		Lastname: "López",
		Age:      54,
		Sayings: []string{
			"Ni lava ni presta la batea",
			"Hijo de gato, caza ratón",
			"Más vale pájaro en mano que cien volando",
		},
	}

	users := []user{u1, u2, u3}
	printUsers(users)

	fmt.Println("----------------------------------------------------------------")

	sort.Sort(byAge(users))
	printUsers(users)

	fmt.Println("----------------------------------------------------------------")

	sort.Sort(byLastname(users))
	printUsers(users)

}

func printUsers(u []user) {
	for _, v := range u {
		fmt.Println(v.Name, v.Lastname)
		fmt.Printf("\t%d years.\n", v.Age)
		fmt.Printf("\tSayings:\n")
		sort.Strings(v.Sayings)
		for _, v2 := range v.Sayings {
			fmt.Printf("\t\t%s\n", v2)
		}
	}
}
