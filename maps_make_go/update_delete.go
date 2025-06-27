package main

import "fmt"

func UpdateDelete() {
	students := map[string]int{"John": 50, "Peter": 65, "Doe": 100}

	// update
	students["John"] = 70

	// print
	fmt.Println("Updating Array")
	for keys, val := range students {
		fmt.Println(keys, val)
	}

	fmt.Println() // new line

	// delete from map
	delete(students, "Doe")
	fmt.Println("Deleting From HashMap")
	for keys, val := range students {
		fmt.Println(keys, val)
	}

}
