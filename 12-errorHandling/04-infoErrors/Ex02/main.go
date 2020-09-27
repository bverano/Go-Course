package main

import (
	"fmt"
	"log"
)

type ubicationError struct {
	lat  string
	long string
	err  error
}

func (n ubicationError) Error() string {
	return fmt.Sprintf("A math concept error has ocurred in: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10.45)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("Math Error: Square root of negative number: %v", f)
		return 0, ubicationError{"50.1293 N", "99.2019 W", nme}
	}
	return 42, nil
}
