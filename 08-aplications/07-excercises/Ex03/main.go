package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	Name     string
	Lastname string
	Age      int
	Sayings  []string
}

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

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("Error:", err)
	}

}
