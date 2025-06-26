package main

import "fmt"

func main() {

	// make maps
	myMapp := map[int]string{1: "Banana", 2: "Apple", 3: "Fish"}

	// using make

	cars := make(map[string]string)
	//inilize cars
	cars["Brand"] = "KIA"
	cars["model"] = "Mustibushi"
	cars["year"] = "1992"

	fmt.Println(len(myMapp))
	fmt.Println()
	// print cars with make
	fmt.Printf("cars \t%v\n", cars["Brand"])

	// Updating Value of Maps
	cars["Brand"] = "BMW"
	fmt.Printf("cars \t%v\n", cars["Brand"])
}
