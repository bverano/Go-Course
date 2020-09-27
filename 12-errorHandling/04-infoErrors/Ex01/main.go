package main

import (
	"errors"
	"fmt"
	"log"
)

var errSqrt = errors.New("There isn't natural square root of a negative number")

func main() {
	fmt.Printf("%T\n", errSqrt)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errSqrt
		// return 0, fmt.Errorf("There isn't natural square root of a negative number: %v", f)
	}
	return 42, nil
}
