package main

import (
	"fmt"

	"github.com/bverano/Go-Course/14-testing/03-exampleTests/coolmath"
)

func main() {
	fmt.Println(coolmath.Sum(2, 4, 5))
	fmt.Println(coolmath.Sum(2, 4, 4, 10))
}
