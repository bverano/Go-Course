package main

import "fmt"

func main() {
	m := map[string][]string{
		"eduar_tua":    []string{"computers", "mountains", "beach"},
		"carlos_ramon": []string{"reading", "shopping", "music"},
		"juan_bimba":   []string{"ice cream", "painting", "dancing"},
	}

	m["bruno_verano"] = []string{"skate", "music", "programming"}

	for k, v := range m {
		fmt.Println(k)
		for _, v2 := range v {
			fmt.Printf("\t%s\n", v2)
		}
	}
}
