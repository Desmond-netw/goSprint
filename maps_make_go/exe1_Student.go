package main

import "fmt"

func GradeStudent() {
	students := map[string]int{"John": 50, "Peter": 65, "Doe": 100}

	// print
	for keys, value := range students {
		if value > 80 {
			fmt.Println(keys, value, "(Execellent)")
		} else if value > 50 && value < 80 {
			fmt.Println(keys, value, "(pass)")
		} else {
			fmt.Println(keys, value)
		}

	}
}
