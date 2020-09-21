package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name     string `json:"Name"`
	Lastname string `json:"Lastname"`
	Age      int    `json:"Age"`
}

func main() {
	bs := []byte(`[
		{"Name":"Bruno","Lastname":"Verano","Age":21},
		{"Name":"Luis","Lastname":"Castañeda","Age":19}
	]`)

	var persons []person
	err := json.Unmarshal(bs, &persons)
	if err != nil {
		fmt.Println("Error:", err)
	}

	for _, v := range persons {
		fmt.Println(v.Name, v.Lastname)
		fmt.Printf("\tAge: %d years old\n", v.Age)
	}
}
