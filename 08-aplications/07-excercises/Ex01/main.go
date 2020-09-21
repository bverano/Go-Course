package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string
	Age  int
}

func main() {
	u1 := user{
		Name: "Eduar",
		Age:  32,
	}

	u2 := user{
		Name: "Juan",
		Age:  27,
	}

	u3 := user{
		Name: "Che",
		Age:  54,
	}

	users := []user{u1, u2, u3}

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%s\n", bs)

}
