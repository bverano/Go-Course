package main

import "fmt"

func main() {
	m := map[string][]string{
		"eduar_tua":    []string{"computers", "mountains", "beach"},
		"carlos_ramon": []string{"reading", "shopping", "music"},
		"juan_bimba":   []string{"ice cream", "painting", "dancing"},
	}

	delete(m, "juan_bimba")

	for k, v := range m {
		fmt.Println(k)
		for _, v2 := range v {
			fmt.Printf("\t%s\n", v2)
		}
	}
}
