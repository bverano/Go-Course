package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name     string   `json:"Name"`
	Lastname string   `json:"Lastname"`
	Age      int      `json:"Age"`
	Sayings  []string `json:"Sayings"`
}

func main() {
	s := `[{"Name":"Eduar","Lastname":"Tua","Age":32,"Sayings":["Cachicamo diciéndole a morrocoy conchudo","La mona, aunque se vista de seda, mona se queda","Poco a poco se anda lejos"]},{"Name":"Carlos","Lastname":"Pérez","Age":27,"Sayings":["Camarón que se duerme se lo lleva la corriente","A ponerse las alpargatas que lo que viene es joropo","No gastes pólvora en zamuro"]},{"Name":"M","Lastname":"Hmmmm","Age":54,"Sayings":["Ni lava ni presta la batea","Hijo de gato, caza ratón","Más vale pájaro en mano que cien volando"]}]`

	var persons []person

	err := json.Unmarshal([]byte(s), &persons)
	if err != nil {
		fmt.Println("Error:", err)
	}

	for _, v := range persons {
		fmt.Println(v.Name, v.Lastname)
		fmt.Printf("\t%d years.\n", v.Age)
		fmt.Printf("\tSayings:\n")
		for _, v2 := range v.Sayings {
			fmt.Printf("\t\t%s\n", v2)
		}
	}

}
