package main

import "fmt"

func main() {

	favSport := "Baseball"

	switch favSport {
	case "Baseball":
		fmt.Println("Home run!!!")
	case "Soccer":
		fmt.Println("Goal!!!")
	case "Basketball":
		fmt.Println("Point!!!")
	default:
		fmt.Println("Some cheer")
	}

}
