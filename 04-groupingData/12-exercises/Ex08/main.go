package main

import "fmt"

func main() {

	m := map[string][]string{
		"eduar_tua":    []string{"computers", "mountains", "beach"},
		"carlos_ramon": []string{"reading", "shopping", "music"},
		"juan_bimba":   []string{"ice cream", "painting", "dancing"},
	}

	for k, v := range m {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Printf("\tIndex: %d\tValue: %s\n", i, v2)
		}
	}

}
