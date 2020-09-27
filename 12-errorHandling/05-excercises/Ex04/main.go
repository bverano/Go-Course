package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("error matemático: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := raiz(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func raiz(f float64) (float64, error) {
	if f < 0 {
		err := fmt.Errorf("Math concept error")
		return 0, sqrtError{"50.2289 N", "99.4656 W", err}
	}
	return 42, nil
}
