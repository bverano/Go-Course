package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheels bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	t := truck{
		vehicle: vehicle{
			color: "blue",
			doors: 2,
		},
		fourWheels: true,
	}

	s := sedan{
		vehicle: vehicle{
			color: "black",
			doors: 4,
		},
		luxury: false,
	}

	fmt.Println(t)
	fmt.Println(s)

	fmt.Println(t.color, t.doors, t.fourWheels)
	fmt.Println(s.color, s.doors, s.luxury)

}
