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

	// Removing element of a map
	// using delete
	delete(cars, "year")
	fmt.Println()
	fmt.Printf("cars \t%v\n", cars)

	// checking if a val and key exists
	fruit, exists := myMapp[3]
	if exists {
		fmt.Println()
		fmt.Println("Fruit is :", fruit)
	} else {
		myMapp[4] = "Mango"
		fmt.Println("Not found")
	}

	fmt.Println() // new line
	// loop through a map
	for key, value := range myMapp {
		fmt.Println(key, value)
	}

	fmt.Println("---------- Execise Print out---------")
	GradeStudent()

	fmt.Println() // new line

	UpdateDelete()

	fmt.Println() // new line
	getScore("John", map[string]int{"John": 50, "Peter": 65, "Doe": 100})

	fmt.Println() // new line
	WordCount("go is fun and go is fast")

	// using Map and slice together
	GetSubjects("Alice", []string{"Math", "Physics"})

}
