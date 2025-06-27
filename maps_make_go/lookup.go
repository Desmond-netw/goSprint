package main

import "fmt"

// looking up array
func getScore(name string, scores map[string]int) {
	// look
	fmt.Println("Looking up in Map")
	for keys, value := range scores {
		if keys == name {
			fmt.Printf("%v your score is := %v\n", keys, value)
			break
		} else {
			fmt.Println("Student Not found")
		}
	}

}
